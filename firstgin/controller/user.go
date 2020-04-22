package controller

import (
	"github.com/gin-gonic/gin"
	"firstgin/model"
	"net/http"
)

func GetUserList(this *gin.Context){

	 var userlist []model.User;

	 userlist =  model.UserList();

	 this.JSON(200,gin.H{
		"msg":"this is user list test",
		"data":userlist,
	}) 

}

func GetUserListToHtml(this *gin.Context){

	var userlist []model.User;

	 userlist =  model.UserList();

	this.HTML(http.StatusOK,"userlist.html",gin.H{
		"msg":"this is user data list",
		"data":userlist,
	});

}