package controllers

import (
	"fmt"
	"struck/configs"
	"struck/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db, err := configs.DBConn()
	if err != nil {
		return
	} else {
		var body models.Body_tblAccount
		var tblAccounts models.TBLAccount
		c.ShouldBind(&body)
		var count int64
		userName := body.UserName
		password := body.Password
		fmt.Println(userName)
		fmt.Println(password)
		db.Model(tblAccounts).Count(&count)

	}

}
