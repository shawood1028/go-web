package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-superset/common"
	"go-superset/model"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Login(ctx *gin.Context) {
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	// 判断手机号是否存在
	DB := common.GetDB()
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	// 判断密码是否正确
	fmt.Println(user.Password, password)
	fmt.Println([]byte(user.Password), []byte(password))
	bts, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err := bcrypt.CompareHashAndPassword(bts, []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v", err)
		return
	}

	// 返回结果
	ctx.JSON(200, gin.H{
		"code":    200,
		"data":    gin.H{"token": token},
		"message": "注册成功",
	})

}
