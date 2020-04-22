package config

/*

 框架的配置目录

*/
/*----------------------------------APPLICATION CONFIG------------------------------------------ */

//应用端口号
const GIN_PORT string = "8081"

//应用模板目录 ， 注意 */*
const GIN_TEMPLATE_DIR string = "templates/*"

/*---------------------------------- MYSQL CONFIG -----------------------------------------------*/

// mysql 主机
const MYSQL_HOST string  = "127.0.0.1"

//mysql  数据库
const MYSQL_DATABASE string = "blog"

// mysql 端口号
const MYSQL_PORT string = "3306"

// mysql 用户名
const MYSQL_USER string  = "root"

//mysql  密码
const MYSQL_PASS string  = "root"

//链接池
const MYSQL_MAX_IDES int = 10
const MYSQL_MAX_CONNECTIONS int= 100




