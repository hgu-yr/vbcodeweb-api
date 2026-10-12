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

## 3. 监听端口（最容易踩坑）

程序监听环境变量 `PORT`，未设置时回退到 **8899**。

> ⚠️ Railway **不一定**会替你注入 `PORT`。如果你在 **Settings → Networking** 的
> 域名里把 **target port** 设成了某个值（例如 `8888`），就必须在 Variables 里
> 显式设置 `PORT=8888`，**两者必须一致**，否则请求无法到达容器。

不一致时的表现：

```json
HTTP 502 {"status":"error","code":502,"message":"Application failed to respond"}
```

排查要点：

- 部署状态是 `SUCCESS`、日志也有「数据库初始化成功！！！」，却仍是 502 → 基本就是端口不一致；
- `GIN_MODE=release` 下 gin **不会**打印 `Listening and serving HTTP on :xxxx`，
  别据此以为进程没起来（这也是本项目曾误判的原因）；
- 用 Railway API 查看域名实际 target port：
  `serviceInstance(serviceId, environmentId) { domains { serviceDomains { domain targetPort } } }`。

不想管理端口的话：**删掉自定义 target port**，或把 `PORT` 设成与 target port 一致的值。

## 4. 初始化管理员

表结构会自动迁移，但管理员账号需要跑一次 `seed`：

```bash
railway run --service <api 服务名> ./vbcodeweb seed
```

默认可登录账号：`admin / admin123`（登录后请尽快修改）。

## 5. 公网域名

在 api 服务 **Settings → Networking → Generate Domain** 生成公网地址。

前端服务 `vbcodeweb-web` 需要把该地址填到：

```
NEXT_PUBLIC_API_BASE = https://vbcodeweb-api-production.up.railway.app
```

> ⚠️ 两个坑：
> 1. **必须带 `https://`**。前端用服务端 `fetch`，Node 对非绝对 URL 会直接抛错，
>    错误被 `catch` 吞掉后页面就显示成「共 0 个工具 / 暂无数据」——看起来像没数据，
>    其实是请求根本没发出。
> 2. 该变量是**构建期**注入（Dockerfile 里的 `ARG NEXT_PUBLIC_API_BASE`），改完必须
>    **Redeploy 重新构建**，只重启无效。

## 6. 关于 railway.json

仓库根目录的 `railway.json` 固化了构建方式（Dockerfile）与健康检查路径。
Railway 已将 config-as-code 标记为 deprecated（现有文件支持到 2026-12-01），
若后续失效，在控制台按同样内容设置即可，不影响部署。
