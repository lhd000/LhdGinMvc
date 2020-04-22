package app

import (
	"github.com/gin-gonic/gin"
	"firstgin/router"
	"firstgin/config"
) 

/**
    应用启动函数
*/
func Run(){

	r := gin.Default()
   
	//设置模板目录
	r.LoadHTMLGlob(config.GIN_TEMPLATE_DIR)


	router.RegistRouter(r)

    //设置应用启动 端口
    r.Run(":"+config.GIN_PORT) // listen and serve on 0.0.0.0:8080

}