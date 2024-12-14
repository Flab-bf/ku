## 说明
- 简介
通过Go语言实现留言板基础功能：用户注册、登录、发表留言、获取所有留言、删除留言. 

- 技术
使用 Go 语言开发，基于 Hertz 框架构建 Web 服务，采用 Gorm 对数据库 MySQL进行操作. 

- 项目结构
	- cmd/main.go ：项目入口.
	- api目录
		- router.go：配置路由，定义接口的路由规则.
		- user.go：存放用户相关接口（登录，注册）
		- comment.go:  存放留言相关接口（发表，逻辑删除，获取）.
	- service目录
		- user.go: 处理用户相关具体业务，与dao（数据层）交互.
		- comment.go：处理用户留言相关具体业务，与dao（数据层）交互.
	- dao目录
		- user.go：用户数据处理层，与数据库MySQL交互.
		- comment.go：用户留言数据处理层，与数据库MySQL交互.
	- model目录
		- 义用户结构体，包含用户的基本信息以及与数据库表的映射关系.
		- 留言结构体及相关方法，包括留言内容,用户关联信息,时间戳等,与数据库中   
		的留言表对应.
	- utiles/db.go:数据库连接工具文件，实现数据库的连接初始化.
- 注意事项
在修改数据库结构（如添加、修改表字段）后，可能需要使用数据库迁移工具（如 Gorm 的自动迁移功能）或者手动执行 SQL 脚本来更新数据库表结构，以确保项目与数据库的一致性.