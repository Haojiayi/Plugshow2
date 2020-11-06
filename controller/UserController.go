package controller

import (
	"ShowWeb/common"
	"ShowWeb/models"
	_ "fmt"
	"github.com/gin-gonic/gin"
)

//func GetAllUserAction(c *gin.Context) {
//	reslist, err := common.GetAllInfo()
//	if err != nil {
//		c.JSON(200, gin.H{"code": 1})
//		return
//	}
//	c.JSON(200, gin.H{"code": 0, "msg": "查询成功！", "count": len(reslist), "data": reslist})
//}

func LoginHandler(c *gin.Context) {
	account,oka := c.GetPostForm("account")
	password,okp := c.GetPostForm("password")
	if !oka || !okp {
		c.JSON(200,gin.H{"code":1,"msg":"参数错误"})
		return
	}
	if account == "" || password == "" {
		c.JSON(200,gin.H{"code":1,"msg":"账号或密码为空！"})
		return
	}
	_,status,token,msg := common.LoginFunc(account,password)
	if status == 0 {
		c.JSON(200,gin.H{"code":1,"msg":msg})
		return
	} else if status == 1 {
		c.JSON(200,gin.H{"code":1,"msg":msg})
		return
	} else if status == 2 {
		c.JSON(200,gin.H{"code":1,"msg":msg})
		return
	} else if status == 3 {
		c.JSON(200, gin.H{"code": 0, "msg": "登录成功！", "token": token})
		return
	}else {
		c.JSON(200, gin.H{"code": 1, "msg": "登录失败！"})
		return
	}

}

func RegisterHandler(c *gin.Context)  {
	account,oka := c.GetPostForm("account")
	password,okp := c.GetPostForm("password")
	if !oka || !okp {
		c.JSON(200,gin.H{"code":1,"msg":"参数错误"})
		return
	}
	err,msg := common.CreateUser(account,password)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":msg})
		return
	}
	c.JSON(200,gin.H{"code":0,"msg":msg})
	return
}


type result struct {
	Mid int `json:"mid"`
	Name string `json:name`
	Children []models.Menu `json:"children"`
	Url string `json:"url"`
	M_l_id int `m_l_id`
}




