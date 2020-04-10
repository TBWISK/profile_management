package controller

import (
	"tbwisk/dao"
	"tbwisk/dto"
	"tbwisk/middleware"

	"github.com/gin-gonic/gin"
)

//RegisterResource 注册资源
func RegisterResource(router *gin.RouterGroup) {
	api := &apiResource{}
	router.GET("/index", api.Index)
	router.POST("/edit_update_resource", api.EditUpdate)
	router.POST("/new_resource", api.NewResource)
	router.POST("/del_resource", api.DelResource)
	router.POST("/new_resource_template", api.NewResourceTemplate)
	router.POST("/del_resource_template", api.DelResourceTemplate)
}

type apiResource struct {
}

func (a *apiResource) Index(c *gin.Context) {
	args := dto.ProductConfArgs{}
	if err := args.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	productConf := dao.ProductConf{}
	resList, err := productConf.FindResoureList(args.ID)
	if err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	resoure := dao.Resoure{}
	resoures, err := resoure.FindMany(resList)
	if err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	//TODO 数据需要预处理
	middleware.ResponseSuccess(c, resoures)
}

//EditUpdate 资源更改
func (a *apiResource) EditUpdate(c *gin.Context) {
	args := dto.ProductConfEditUpdate{}
	if err := args.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	// 新增或者删除配置项
	productconf := dao.ProductConf{}
	productcon, err := productconf.Find(args.ID)
	if err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	if args.Type == "update" {
		productcon.ResoureList = append(productcon.ResoureList, args.Resource)
	} else if args.Type == "del" {
		for i := 0; i < len(productcon.ResoureList); i++ {
			if args.Resource == productcon.ResoureList[i] {
				productcon.ResoureList = append(productcon.ResoureList[0:i], productcon.ResoureList[i+1:]...)
				break
			}
		}
	} else {
		middleware.ResponseError(c, 2005, err)
		return
	}
	if err := productcon.Save(); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}

//NewResource 新资源
func (a *apiResource) NewResource(c *gin.Context) {
	input := dto.NewResourceInput{}
	if err := input.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	res := dao.Resoure{Name: input.Name, Section: input.Section, ResoureTemplateID: input.ResoureTemplateID, Env: input.Env, Data: input.Data}
	if err := res.Save(); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}

//DelResource 删除资源
func (a *apiResource) DelResource(c *gin.Context) {
	input := dto.DelResourceInput{}
	if err := input.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	res := dao.Resoure{}
	if err := res.Find(input.ID); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	if err := res.Del(); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}
func (a *apiResource) ResourceTemplateList(c *gin.Context) {
	temp := dao.ResoureTemplate{}
	// temp.Find(id)
	templates, err := temp.FindALL()
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, templates)
}

//NewResourceTemplate 新模板
func (a *apiResource) NewResourceTemplate(c *gin.Context) {
	input := dto.NewResourceTemplateInput{}
	if err := input.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	template := dao.ResoureTemplate{Name: input.Name, TemplateList: input.Keys}
	if err := template.Save(); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}

//DelResourceTemplate 移除模板
func (a *apiResource) DelResourceTemplate(c *gin.Context) {
	input := dto.DelResourceTemplateInput{}
	if err := input.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	temp := dao.ResoureTemplate{}
	template, err := temp.Find(input.ID)
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	if err := template.Del(); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, nil)
}
