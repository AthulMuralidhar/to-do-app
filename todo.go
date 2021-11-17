package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func todoPost(ctx *gin.Context) {
	user:= new(User)

	if err := ctx.Bind(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	Users = append(Users, user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "note added successfully",
		"jwt": "tester-123",
	})
}