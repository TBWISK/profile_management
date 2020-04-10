package dao

import (
	"encoding/json"
	"fmt"
	"tbwisk/public"

	"github.com/jinzhu/gorm"
)

//Resoure 资源落地 资源固定资产
type Resoure struct {
	gorm.Model
	Name              string            // 对外的资源命名
	Section           string            //section 名字; 不唯一或者唯一
	ResoureTemplateID uint              //资源模板的id
	Data              string            `gorm:"type:TEXT;default:'{}'"` //资源落地实现 {} map 格式
	Env               string            //选择环境
	DataMap           map[string]string `gorm:"-"`
}

func (f *Resoure) decodeData() {
	if err := json.Unmarshal([]byte(f.Data), &f.DataMap); err != nil {
		fmt.Println("decodeData err", err)
	}
}

//Find f
func (f *Resoure) Find(id int64) error {
	// var user Resoure
	err := public.GormPool.Where("id = ?", id).First(&f).Error

	if err != nil {
		return err
	}
	f.decodeData()
	return nil
}

//FindMany 寻找多数
func (f *Resoure) FindMany(ids []string) (*[]Resoure, error) {
	var users []Resoure
	err := public.GormPool.Where("id IN (?)", ids).Find(&users).Error

	if err != nil {
		return nil, err
	}
	for idx := 0; idx < len(users); idx++ {
		users[idx].decodeData()
	}
	return &users, nil
}

//Save s
func (f *Resoure) Save() error {
	if err := public.GormPool.Save(f).Error; err != nil {
		return err
	}
	return nil
}

//Del 删除
func (f *Resoure) Del() error {
	if err := public.GormPool.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}

//ResoureTemplate 资源类型
type ResoureTemplate struct {
	gorm.Model
	Name         string   //资源类型命名
	Template     string   `gorm:"type:TEXT;default:'[]'"` //资源模板 [格式]
	TemplateList []string `gorm:"-"`
}

func (f *ResoureTemplate) decodeTemplate() {
	if err := json.Unmarshal([]byte(f.Template), &f.TemplateList); err != nil {
		fmt.Println("decodeTemplate err", err)
	}
}

func (f *ResoureTemplate) encodeResoure() error {
	b, err := json.Marshal(f.TemplateList)
	if err != nil {
		fmt.Println("encodeResoure err", err)
		return err
	}
	f.Template = string(b)
	return nil
}

//Save 保存
func (f *ResoureTemplate) Save() error {
	if err := f.encodeResoure(); err != nil {
		return err
	}
	if err := public.GormPool.Save(f).Error; err != nil {
		return err
	}
	return nil
}

//Del 删除
func (f *ResoureTemplate) Del() error {
	if err := public.GormPool.Delete(&f).Error; err != nil {
		return err
	}
	return nil
}

//Find findone
func (f *ResoureTemplate) Find(id int64) (*ResoureTemplate, error) {
	var user ResoureTemplate
	err := public.GormPool.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.decodeTemplate()
	return &user, nil
}

//FindALL all
func (f *ResoureTemplate) FindALL() ([]*ResoureTemplate, error) {
	var ff []*ResoureTemplate
	query := public.GormPool
	err := query.Find(&ff).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return ff, nil
}
