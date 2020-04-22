package router

import (
	"firstgin/controller"
	"github.com/gin-gonic/gin"
)

func RegistRouter(this *gin.Engine){

	this.GET("/test",controller.Test);

	this.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
	})
	
    

	this.GET("/",func(c *gin.Context){
		c.JSON(200,gin.H{
			"msg":"this is lhd's firt gin framework programe",
			"test":"this is hot compile test",
		})
	})

	this.GET("/test1",controller.Test1);

	this.GET("/another",controller.Another);

	this.GET("/html",controller.Html);

	this.GET("/userlist",controller.GetUserList);

	this.GET("/ulisthtml",controller.GetUserListToHtml);
	
}