package dto

import (
	"errors"
	"strings"
	"tbwisk/public"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

//NewResourceTemplateInput 参数校验
type NewResourceTemplateInput struct {
	Name string   `form:"name" json:"name" validate:"required"`
	Keys []string `form:"keys" json:"keys" validate:"required"`
}

//BindingValidParams 检验参数
func (o *NewResourceTemplateInput) BindingValidParams(c *gin.Context) error {
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

//DelResourceTemplateInput 校验参数
type DelResourceTemplateInput struct {
	ID int64 `form:"id" json:"id" validate:"required"`
}

//BindingValidParams 检验参数
func (o *DelResourceTemplateInput) BindingValidParams(c *gin.Context) error {
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

//NewResourceInput 检验参数
type NewResourceInput struct {
	Name              string `form:"name" json:"name" validate:"required"`
	Section           string `form:"section" json:"section" validate:"required"`
	Env               string `form:"env" json:"env" validate:"required"`
	ResoureTemplateID uint   `form:"resid" json:"resid" validate:"required"`
	Data              string `form:"data" json:"data" validate:"required"`
}

//BindingValidParams 检验参数
func (o *NewResourceInput) BindingValidParams(c *gin.Context) error {
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

//DelResource 删除检验canshu
type DelResourceInput struct {
	ID int64 `form:"id" json:"id" validate:"required"`
}

//BindingValidParams 检验参数
func (o *DelResourceInput) BindingValidParams(c *gin.Context) error {
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
