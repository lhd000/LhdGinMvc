package controller

import (
	 "github.com/gin-gonic/gin"
)


func Test(this *gin.Context)  {

	this.JSON(200,gin.H{
		"msg":"this is controller test",
	})
	
}

func Test1(this *gin.Context){
	this.JSON(200,gin.H{
		"msg":"this is another test",
	})
}