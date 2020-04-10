package controller

import (
	"strconv"
	"tbwisk/dao"
	"tbwisk/dto"
	"tbwisk/middleware"

	"github.com/gin-gonic/gin"
)

//APIProject api/project
type APIProject struct {
}

//GetLists 获取列表
func (api *APIProject) GetLists(c *gin.Context) {
	projectList := dto.ProjectList{}
	if err := projectList.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	pageInt, err := strconv.ParseInt(projectList.Page, 10, 64)
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	project := dao.Product{}
	// project.PageList(name, pageNo, pageSize)
	if userList, total, err := project.PageList(int(pageInt), 20); err != nil {
		middleware.ResponseError(c, 2005, err)
	} else {
		m := map[string]interface{}{
			"list":  userList,
			"total": total,
		}
		middleware.ResponseSuccess(c, m)
	}
}

//GetProjectConfig 获取配置方式
func (api *APIProject) GetProjectConfig(c *gin.Context) {

}
