package utils

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"vbCodeWeb/model"
)

// charset 返回配置的字符集，为空时使用 utf8mb4。
func (c DatabaseConfig) charset() string {
	if c.Charset == "" {
		return "utf8mb4"
	}
	return c.Charset
}

// DSN 生成连接到指定库的 DSN。
func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=5s&readTimeout=5s&writeTimeout=5s",
		c.User, c.Password, c.Host, c.Port, c.Name, c.charset())
}

// serverDSN 生成不指定库的 DSN，用于创建数据库。
func (c DatabaseConfig) serverDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=%s&parseTime=True&loc=Local&timeout=5s&readTimeout=5s&writeTimeout=5s",
		c.User, c.Password, c.Host, c.Port, c.charset())
}

var (
	conf *Config
	db   *gorm.DB
)

// LoadConfig 读取环境对应的 yaml 配置文件，默认使用 dev.yaml。
func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	path := fmt.Sprintf("%s.yaml", env)

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config %s: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("decode config %s: %w", path, err)
	}
	if cfg.Database.Host == "" || cfg.Database.User == "" || cfg.Database.Name == "" {
		return nil, fmt.Errorf("database host, user and name are required")
	}
	return &cfg, nil
}

// InitDB 加载配置、创建数据库、初始化连接池并迁移模型。
func InitDB() *gorm.DB {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	conf = cfg
	if err := ensureDatabase(cfg.Database); err != nil {
		panic(fmt.Errorf("ensure database: %w", err))
	}

	gdb, err := gorm.Open(mysql.Open(cfg.Database.DSN()), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("connect mysql: %w", err))
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		panic(fmt.Errorf("get sql.DB: %w", err))
	}
	sqlDB.SetMaxOpenConns(cfg.Database.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.Database.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(cfg.Database.ConnMaxLifetime)

	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("ping mysql: %w", err))
	}

	models := model.Models()
	for _, m := range models {
		if err := migrate(gdb, m); err != nil {
			panic(fmt.Errorf("auto migrate: %w", err))
		}
	}
	db = gdb
	fmt.Println("数据库初始化成功！！！")
	return db
}

// tableCommenter 由模型实现以声明表注释。
type tableCommenter interface {
	TableComment() string
}

// migrate 迁移单张表：建表/改列，并同步表注释到数据库。
func migrate(db *gorm.DB, m any) error {
	autoMigrateDB := db
	if tc, ok := m.(tableCommenter); ok && tc.TableComment() != "" {
		autoMigrateDB = db.Set("gorm:table_options", fmt.Sprintf("COMMENT='%s'", tc.TableComment()))
	}
	if err := autoMigrateDB.AutoMigrate(m); err != nil {
		return err
	}
	if err := syncTableComment(db, m); err != nil {
		return err
	}
	return syncColumnComments(db, m)
}

// syncColumnComments 同步所有列注释到数据库。
// GORM 的 AutoMigrate 会跳过主键列的注释更新，这里统一补齐。
func syncColumnComments(db *gorm.DB, m any) error {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(m); err != nil {
		return err
	}
	if stmt.Schema == nil {
		return nil
	}

	rows, err := db.Raw(
		"SELECT column_name, column_comment FROM information_schema.COLUMNS WHERE table_schema = DATABASE() AND table_name = ?",
		stmt.Schema.Table,
	).Rows()
	if err != nil {
		return err
	}
	current := make(map[string]string)
	for rows.Next() {
		var name, comment string
		if err := rows.Scan(&name, &comment); err != nil {
			rows.Close()
			return err
		}
		current[name] = comment
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return err
	}

	for _, field := range stmt.Schema.Fields {
		if field.DBName == "" || field.IgnoreMigration {
			continue
		}
		if comment, ok := current[field.DBName]; ok && comment == field.Comment {
			continue
		}
		if err := db.Migrator().AlterColumn(m, field.DBName); err != nil {
			return err
		}
	}
	return nil
}

// syncTableComment 对比模型与数据库的表注释，有变动时更新数据库。
func syncTableComment(db *gorm.DB, m any) error {
	tc, ok := m.(tableCommenter)
	if !ok || tc.TableComment() == "" {
		return nil
	}

	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(m); err != nil {
		return err
	}
	if stmt.Schema == nil {
		return nil
	}
	table := stmt.Schema.Table

	var current sql.NullString
	if err := db.Raw(
		"SELECT TABLE_COMMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = DATABASE() AND TABLE_NAME = ?",
		table,
	).Row().Scan(&current); err != nil {
		return err
	}
	if current.Valid && current.String == tc.TableComment() {
		return nil
	}

	comment := db.Dialector.Explain("?", tc.TableComment())
	return db.Exec(fmt.Sprintf("ALTER TABLE `%s` COMMENT = %s", table, comment)).Error
}

// ensureDatabase 在目标库不存在时创建它。
func ensureDatabase(cfg DatabaseConfig) error {
	conn, err := sql.Open("mysql", cfg.serverDSN())
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET %s", cfg.Name, cfg.charset())
	if _, err := conn.Exec(stmt); err != nil {
		return err
	}
	return nil
}

// DB 返回已初始化的默认数据库连接。
func DB() *gorm.DB {
	return db
}

// GetConfig 返回当前加载的配置。
func GetConfig() *Config {
	return conf
}
