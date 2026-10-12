# Railway 部署说明（后端）

前端仓库见 `vbcodeweb-web`。

## 1. 准备数据库

在 Railway 项目里添加一个 **MySQL** 服务。

## 2. 后端服务变量

在 api 服务的 Variables 中添加（`MySQL` 换成你的数据库服务名）：

| 变量 | 示例值 | 说明 |
| --- | --- | --- |
| `MYSQLHOST` | `${{MySQL.MYSQLHOST}}` | 数据库地址 |
| `MYSQLPORT` | `${{MySQL.MYSQLPORT}}` | 数据库端口 |
| `MYSQLUSER` | `${{MySQL.MYSQLUSER}}` | 数据库用户 |
| `MYSQLPASSWORD` | `${{MySQL.MYSQLPASSWORD}}` | 数据库密码 |
| `MYSQLDATABASE` | `${{MySQL.MYSQLDATABASE}}` | 数据库名 |
| `JWT_SECRET` | 一段足够长的随机字符串 | JWT 签名密钥 |
| `JWT_EXPIRE` | `168h` | token 有效期，可选 |

> 也可用 `DB_HOST / DB_PORT / DB_USER / DB_PASSWORD / DB_NAME` 等价变量。
> 两者都设置时优先级：`DB_*` > `MYSQL*`。
> 程序不再强制要求 `prod.yaml`，配置文件缺失时会完全使用环境变量。

如果数据库用户没有建库权限，`CREATE DATABASE` 会报权限错误；此时把 `MYSQLDATABASE`
设为已存在的库名即可。

## 3. 监听端口

程序优先读取环境变量 `PORT`（Railway 自动注入），无需手动设置。

## 4. 初始化管理员

表结构会自动迁移，但管理员账号需要跑一次 `seed`：

```bash
railway run --service <api 服务名> ./vbcodeweb seed
```

默认可登录账号：`admin / admin123`（登录后请尽快修改）。

## 5. 公网域名

在 api 服务 **Settings → Networking → Generate Domain** 生成公网地址，
把它填到前端服务的 `NEXT_PUBLIC_API_BASE`。

## 6. 关于 railway.json

仓库根目录的 `railway.json` 固化了构建方式（Dockerfile）与健康检查路径。
Railway 已将 config-as-code 标记为 deprecated（现有文件支持到 2026-12-01），
若后续失效，在控制台按同样内容设置即可，不影响部署。
