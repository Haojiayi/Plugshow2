package controller

import (
	"ShowWeb/common"
	"ShowWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReturnMenu(c *gin.Context){
	resultlist:=[]result{}
	flist,ferr:=common.GetAllFatherMenu()
	if ferr!=nil {
		fmt.Printf("v%",ferr)
	}
	for _,y :=range flist{
		clist,_:=common.GetChildrenMenuByFid(int(y.ID))
		resultlist=append(resultlist, result{Mid:int(y.ID),Name:y.Name,Children:clist,Url:y.M_url})
	}
	//
	fmt.Printf("s%","s%","s%",flist)
	//res,_:=json.Marshal(resultlist)
	//fmt.Print(res)
	c.JSON(200,gin.H{"data":resultlist})
}

func GetMenuByMid(c *gin.Context)  {
	mid:=c.Query("mid")
	rmid,_:=strconv.Atoi(mid)

	resultlist,_:=common.GetMenuByMid(rmid)


	//
	fmt.Printf("s%","s%","s%",resultlist)
	//res,_:=json.Marshal(resultlist)
	//fmt.Print(res)
	c.JSON(200,gin.H{"data":resultlist})
}

func CreateOneMenu(c *gin.Context)  {
	name,oka := c.GetPostForm("name")
	url,okb := c.GetPostForm("url")
	parent_id,okc := c.GetPostForm("parent_id")
	var menu models.Menu
	if !oka || !okb {
		c.JSON(200,gin.H{"code":1,"msg":"参数错误！"})
		return
	}
	if !okc {
		menu.M_url = url
		menu.Name = name
	}else {
		menu.M_url = url
		menu.Name = name
		menu.M_l_id,_ = strconv.Atoi(parent_id)
	}
	err := common.CreateMenu(&menu)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"创建失败！"})
		return
	}
	c.JSON(200,gin.H{"code":0,"msg":"创建成功！"})
	return
}

func UpdateOneMenu(c *gin.Context)  {
	mId, okMid := c.GetPostForm("m_id")
	name, okName := c.GetPostForm("name")
	mLId, okMLId := c.GetPostForm("m_l_id")
	mUrl, okMUrl := c.GetPostForm("m_url")
	if !okMid || !okMLId || !okMUrl || !okName {
		c.JSON(200,gin.H{"code":1,"msg":"未传入关键参数！"})
		return
	}

	menu,err := common.GetMenuById(mId)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"数据查询异常！"})
		return
	}
	menu.M_l_id,_ = strconv.Atoi(mLId)
	menu.Name = name
	menu.M_url = mUrl
	if err = common.UpdateMenu(menu); err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"数据修改失败！"})
		return
	}else {
		c.JSON(200,gin.H{"code":0,"msg":"修改成功！"})
		return
	}
}

func DeleteOneMenu(c *gin.Context)  {
	mId, okMid := c.GetPostForm("m_id")
	if !okMid {
		c.JSON(200,gin.H{"code":1,"msg":"未传入关键参数！"})
		return
	}
	if err := common.DeleteMenu(mId); err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"删除失败！"})
		return
	}else {
		c.JSON(200,gin.H{"code":0,"msg":"删除成功！"})
		return
	}
}
