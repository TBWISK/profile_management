package dto

import (
	"errors"
	"strings"
	"tbwisk/public"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

//ProjectList 页数
type ProjectList struct {
	Page string `form:"page" json:"page" validate:"required"`
}

//BindingValidParams 检验参数
func (o *ProjectList) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

//ProjectInput 项目输入参数
type ProjectInput struct {
	Name      string `form:"name" json:"name" validate:"required"`
	GitLabURL string `form:"gitlab" json:"gitlab" validate:"required"`
}

//BindingValidParams 检验参数
func (o *ProjectInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

//ProductConfArgs 校验参数
type ProductConfArgs struct {
	ID string `form:"id" json:"id" validate:"required"`
}

//BindingValidParams 检验参数
func (o *ProductConfArgs) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

//ProductConfEditUpdate 产品update
type ProductConfEditUpdate struct {
	ID       string `form:"id" json:"id" validate:"required"`
	Resource string `form:"resourceid" json:"resourceid" validate:"required"`
	Type     string `form:"type" json:"type" validate:"required"` // 更改类型 update是新增 del是删除
}

//BindingValidParams 检验参数
func (o *ProductConfEditUpdate) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}

//ProjectConfigInput 检验用户层面
type ProjectConfigInput struct {
	Env         string `form:"env" json:"env" validate:"required"`
	APPID       string `form:"appid" json:"appid" validate:"required"`
	AppSecret   string `form:"appsecret" json:"appsecret" validate:"required"`
	ProjectName string `form:"project" json:"project" validate:"required"`
}

//BindingValidParams 检验参数
func (o *ProjectConfigInput) BindingValidParams(c *gin.Context) error {
	if err := c.ShouldBind(o); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("zh")
	}
	err := public.Validate.Struct(o)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
