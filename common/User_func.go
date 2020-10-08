package common

import (
	myjwt "ShowWeb/middleware"
	"ShowWeb/models"
	"ShowWeb/mysql"
	"crypto/md5"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"time"
)

var db = mysql.GetDb()

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	UserID string `json:"user_id"`
}

// 生成令牌
func generateToken(user models.User) (err error,msg string) {
	nowTime:=time.Now()
	expireTime:=nowTime.Add(3*time.Hour)
	j := &myjwt.JWT{
		[]byte("newtrek"),
	}
	claims := myjwt.CustomClaims{
		fmt.Sprintf("%d",user.ID),
		user.Account,
		jwtgo.StandardClaims{
			NotBefore: nowTime.Unix(), // 签名生效时间
			ExpiresAt: expireTime.Unix(), // 过期时间 一小时
			Issuer:    "newtrek",                   	//签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		//c.JSON(http.StatusOK, gin.H{
		//	"status": -1,
		//	"msg":    err.Error(),
		//})
		return err,""
	}

	//log.Println(token)
	//
	//data := LoginResult{
	//	UserID:  fmt.Sprintf("%d",user.ID),
	//	Token: token,
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"status": 0,
	//	"msg":    "登录成功！",
	//	"data":   data,
	//})
	return nil,token
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}

func MD5Secret(data string) string {

	t := md5.New()

	io.WriteString(t, data)

	return fmt.Sprintf("%x", t.Sum(nil))

}

func CreateUser(account string,password string) (err error,msg string) {
	var user models.User
	var new_str string
	count := 0
	err = db.Where("account=?",account).Find(&user).Count(&count).Error
	if count > 0 {
		return err,"账号已存在"
	}
	rand.Seed(time.Now().UnixNano())
	num4 := rand.Intn(1000000-900000) + 900000 //[0,100)
	user.Account = account
	user.Salt = num4
	new_str = fmt.Sprintf("%d%s", num4,password)
	user.Password = MD5Secret(new_str)
	err = db.Create(&user).Error
	if err != nil {
		return err,"注册失败"
	}
	return nil,"注册成功"
}

func LoginFunc(account string,password string) (err error,status int,token string,msg string)  {

	var user models.User
	num := 0
	err = db.Where("account=?",account).Find(&user).Count(&num).Error
	if num <= 0 {
		return err,0,"","账号错误"
	}
	str := fmt.Sprintf("%d%s",user.Salt,password)
	if user.Password != MD5Secret(str) {
		return err,1,"","密码错误"
	}

	err,token = generateToken(user)
	if err != nil {
		return err, 2,"", "生成token错误"
	}else {
		return err, 3, token,"生成token成功"
	}

}