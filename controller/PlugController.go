package controller

import (
	"ShowWeb/common"
	"ShowWeb/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"path"
	"strconv"
	"strings"
)

func GetPlugByMid(c *gin.Context)  {
	var data string
	var res string
	mid:=c.Query("mid")
	mids:=strings.Split(mid,",")
	//rmid,_:=strconv.Atoi(mid)
	if len(mids)>1 {
		for i:=0;i<len(mids);i++   {
			if i<len(mids)-1 {
				data=mids[i]+" or"
			}else {
				res =data+" "+mids[i]
			}
		}
	}else {
		res = mids[0]
	}
	page:=c.Query("page")
	//rmid,_:=strconv.Atoi(mid)
	rpage,_:=strconv.Atoi(page)
	rpage=(rpage-1)*20
	lpage:=rpage+20
	plist,perr:=common.GetPlugByMidLimit(res,rpage,lpage)
	allist,_:=common.GetPlugByMid(res)
	if perr!=nil {
		fmt.Print("s%",perr)
	}
	count:=float64(len(allist))
	allpage:=float64(count/20)
	allpage=math.Ceil(allpage)
	c.JSON(200,gin.H{"data":plist,"allpage":allpage})
}

func GetPlugsByMidHot(c *gin.Context)  {
	res:=[]models.Plug{}
	mid:=c.Query("mid")
	rmid,_:=strconv.Atoi(mid)

	plist,perr:=common.GetPlugsByMidHot(rmid)
	if perr!=nil {
		fmt.Print("s%",perr)
	}
	if len(plist)>4 {
		for i:=0 ;i<4 ;i++  {
			res=append(res,plist[i])
		}
	}else {
		res=plist
	}



	c.JSON(200,gin.H{"data":res})
}

func DownloadAdd(c *gin.Context )  {
	pid:=c.Query("pid")
	rpid,_:=strconv.Atoi(pid)
	getp,_:=common.GetPlugsByPid(rpid)
	new:=getp.P_downum+1
	(&getp).P_downum=new
	//if err!=nil {
	//	fmt.Printf("s%",err)
	//}
	_,err:=common.UpdateDByPid(&getp)
	if err!=nil {
		fmt.Printf("s%",err)
		c.JSON(200,gin.H{"data":"fail"})
	}else {
		c.JSON(200,gin.H{"data":"success"})
	}
	fmt.Printf("s%",getp)
}

func GetPlugByKeyWord(c *gin.Context)  {
	keyword:=c.Query("keyword")
	keywords:="%"+keyword+"%"
	res,err:=common.GetPlugByKeyWord(keywords)
	if err!=nil {
		c.JSON(400,gin.H{"data":"false"})

	}else {
		c.JSON(200,gin.H{"data":res})
	}
}

func test()  {
	
}

func CreateOnePlug(c *gin.Context)  {
	title, okTitle := c.GetPostForm("title")
	content, okContent := c.GetPostForm("content")
	menuId, okMenuId := c.GetPostForm("menu_id")
	var plug models.Plug

	videoFile, okVideo := c.FormFile("video_file")
	imgFile, okImg := c.FormFile("img_file")
	plugFile, okPlug := c.FormFile("plug_file")

	if !okTitle {
		c.JSON(200,gin.H{"code":1,"msg":"标题参数错误！"})
		return
	}

	if !okContent {
		c.JSON(200,gin.H{"code":1,"msg":"简介参数错误！"})
		return
	}

	if !okMenuId {
		c.JSON(200,gin.H{"code":1,"msg":"菜单参数错误！"})
		return
	}

	if okPlug != nil{
		c.JSON(200,gin.H{"code":1,"msg":"上传插件文件错误！"})
		return
	}

	if okImg != nil{
		c.JSON(200,gin.H{"code":1,"msg":"上传图片文件错误！"})
		return
	}

	if okVideo != nil{
		c.JSON(200,gin.H{"code":1,"msg":"上传视频文件错误！"})
		return
	}

	videoUrl := path.Join("./static_file/video/", videoFile.Filename)
	imgUrl := path.Join("./static_file/img/", imgFile.Filename)
	plugUrl := path.Join("./static_file/plug/", plugFile.Filename)

	errVideo := c.SaveUploadedFile(videoFile, videoUrl)
	if errVideo != nil{
		c.JSON(200,gin.H{"code":1,"msg":"视频文件保存错误！"})
		return
	}

	errImg :=c.SaveUploadedFile(imgFile, imgUrl)
	if errImg != nil{
		c.JSON(200,gin.H{"code":1,"msg":"图片文件保存错误！"})
		return
	}

	errPlug := c.SaveUploadedFile(plugFile, plugUrl)
	if errPlug != nil{
		c.JSON(200,gin.H{"code":1,"msg":"插件文件保存错误！"})
		return
	}

	plug.M_id,_ = strconv.Atoi(menuId)
	plug.P_title = title
	plug.P_content = content
	plug.P_v_dress = fmt.Sprintf("/static_file/video/%s", videoFile.Filename)
	plug.P_img_adress = fmt.Sprintf("/static_file/img/%s", imgFile.Filename)
	plug.P_d_adress = fmt.Sprintf("/static_file/plug/%s", plugFile.Filename)

	err := common.CreatePlug(&plug)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"数据记录创建失败！"})
		return
	}
	c.JSON(200,gin.H{"code":0,"msg":"创建成功！"})
	return
}

func UpdateOnePlug(c *gin.Context)  {
	mId,okMId := c.GetPostForm("id")
	title, okTitle := c.GetPostForm("title")
	content, okContent := c.GetPostForm("content")
	menuId, okMenuId := c.GetPostForm("menu_id")

	videoFile, okVideo := c.FormFile("video_file")
	imgFile, okImg := c.FormFile("img_file")
	plugFile, okPlug := c.FormFile("plug_file")

	if !okMId {
		c.JSON(200,gin.H{"code":1,"msg":"关键参数错误！"})
		return
	}

	plug,err := common.GetPlugById(mId)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"数据异常！"})
		return
	}

	if okTitle {
		plug.P_title = title
	}

	if okContent {
		plug.P_content = content
	}

	if okMenuId {
		plug.M_id,_ = strconv.Atoi(menuId)
	}

	if plugFile != nil{
		if okPlug != nil{
			c.JSON(200,gin.H{"code":1,"msg":"上传插件文件错误！"})
			return
		}
		plugUrl := path.Join("./static_file/plug/", plugFile.Filename)
		errPlug := c.SaveUploadedFile(plugFile, plugUrl)
		if errPlug != nil{
			c.JSON(200,gin.H{"code":1,"msg":"插件文件保存错误！"})
			return
		}
		plug.P_d_adress = fmt.Sprintf("/static_file/plug/%s", plugFile.Filename)
	}

	if imgFile != nil{
		if okImg != nil{
			c.JSON(200,gin.H{"code":1,"msg":"上传图片文件错误！"})
			return
		}
		imgUrl := path.Join("./static_file/img/", imgFile.Filename)
		errImg :=c.SaveUploadedFile(imgFile, imgUrl)
		if errImg != nil{
			c.JSON(200,gin.H{"code":1,"msg":"图片文件保存错误！"})
			return
		}
		plug.P_img_adress = fmt.Sprintf("/static_file/img/%s", imgFile.Filename)
	}

	if videoFile != nil{
		if okVideo != nil{
			c.JSON(200,gin.H{"code":1,"msg":"上传视频文件错误！"})
			return
		}
		videoUrl := path.Join("./static_file/video/", videoFile.Filename)
		errVideo := c.SaveUploadedFile(videoFile, videoUrl)
		if errVideo != nil{
			c.JSON(200,gin.H{"code":1,"msg":"视频文件保存错误！"})
			return
		}
		plug.P_v_dress = fmt.Sprintf("/static_file/video/%s", videoFile.Filename)
	}

	err = common.UpdatePlug(&plug)
	if err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"修改失败！"})
		return
	}
	c.JSON(200,gin.H{"code":0,"msg":"修改成功！"})
	return
}

func DeleteOnePlug(c *gin.Context)  {
	mId, okMid := c.GetPostForm("id")
	if !okMid {
		c.JSON(200,gin.H{"code":1,"msg":"未传入关键参数！"})
		return
	}
	if err := common.DeletePlug(mId); err != nil {
		c.JSON(200,gin.H{"code":1,"msg":"删除失败！"})
		return
	}else {
		c.JSON(200,gin.H{"code":0,"msg":"删除成功！"})
		return
	}
}