package model

// Models 返回需要 AutoMigrate 的模型列表。
func Models() []any {
	return []any{
		&User{},
		&Category{},
		&Product{},
		&MiniProgram{},
		&Article{},
		&Submission{},
		&Upload{},
	}
}
