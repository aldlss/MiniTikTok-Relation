# MiniTikTok-Relation

由于文档中描述不清，这里对于好友的定义为互相关注。

由于要获取 Message 信息，因此要连接到 Message 数据库获取消息，因此需要 PostgreSQL 的相关环境变量。  
表内的字段相关可参考 Message 服务中的定义

## 监听

监听设置是 `[::]:19198`

## 环境变量

### NEO4J_URL

neo4j 的地址 

### NEO4J_USERNAME

用于登录 neo4j 的用户名

### NEO4J_PASSWORD

用于登录 neo4j 的密码

### PGSQL_HOST

要连接的 PostgreSQL 的地址

### PGSQL_PORT （可选）

连接的 PostgreSQL 地址的端口，不填为默认端口

### PGSQL_USER

登录 PostgreSQL 的用户名

### PGSQL_PASSWORD

登录 PostgreSQL 的密码

### PGSQL_DBNAME

使用的数据库的名字

### TABLE_NAME

要查的表的名字，注意是全名
