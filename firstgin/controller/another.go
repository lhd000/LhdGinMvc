package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Another(this *gin.Context){
	this.JSON(200,gin.H{
      "msg":"this is another test",
	})
}

func Html(this *gin.Context){

	str := []string {"dfdfdf","sdfsdfsdf","sdfsdfsf"} 

    this.HTML(http.StatusOK,"index.html",gin.H{
		"msg":"this is gin template test",
		"data":str,
	});
}