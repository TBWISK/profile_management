package dao

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"tbwisk/public"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// `gorm:"type:TEXT"`

//Product 产品
type Product struct {
	gorm.Model
	Name              string
	Del               int
	GitLabURL         string
	ProductConfProd   int
	ProductConfProdok bool
	ProductConfTest   int
	ProductConfTestok bool
	ProductConfDev    int  //对应实现的配置
	ProductConfDevok  bool //是否有 开发
}

//NewUUID uuid
func NewUUID() string {
	ui := uuid.Must(uuid.NewV4(), nil)
	fmt.Println(ui)
	return ui.String()
}

//PageList 分页查询
func (f *Product) PageList(pageNo int, pageSize int) ([]*Product, int64, error) {
	var user []*Product
	var userCount int64
	offset := (pageNo - 1) * pageSize
	query := public.GormPool
	// if name != "" {
	// query = query.Where("name = ?", name)
	// }
	err := query.Limit(pageSize).Offset(offset).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	errCount := query.Table("product").Count(&userCount).Error
	if errCount != nil {
		return nil, 0, err
	}
	return user, userCount, nil
}

//Find find
func (f *Product) Find(id int64) (*Product, error) {
	var user Product
	err := public.GormPool.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//Save 存储
func (f *Product) Save() error {
	if err := public.GormPool.Save(f).Error; err != nil {
		return err
	}
	return nil
}

// 单个最后的落地情况
// http:{
//		boss_url:,
//		dsp_uro:,
//		host:,
//	}

//ProductConf 产品实现
type ProductConf struct {
	gorm.Model
	Del         int //1是shanchu
	APPID       string
	AppSecret   string
	ProductID   uint
	Env         string
	Resoure     string   `gorm:"type:TEXT;default:'[]'"` // []实现了对应的资源
	ResoureList []string `gorm:"-"`
}

func (f *ProductConf) decodeResoure() {
	if err := json.Unmarshal([]byte(f.Resoure), &f.ResoureList); err != nil {
		fmt.Println("decodeResoure err", err)
	}
}

func (f *ProductConf) encodeResoure() error {
	b, err := json.Marshal(f.ResoureList)
	if err != nil {
		fmt.Println("encodeResoure err", err)
		return err
	}
	f.Resoure = string(b)
	return nil
}

//MD5 md5
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

//NewProductConf 实现配置
func NewProductConf(Env string) ProductConf {
	return ProductConf{Env: Env, APPID: ""}
}

//Save s
func (f *ProductConf) Save() error {
	if err := f.encodeResoure(); err != nil {
		return err
	}
	if err := public.GormPool.Save(f).Error; err != nil {
		return err
	}
	return nil
}

//Find f
func (f *ProductConf) Find(id string) (*ProductConf, error) {
	var user ProductConf
	err := public.GormPool.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	user.decodeResoure()
	return &user, nil
}

//FindResoureList 寻找资源列表 输入参数产品id
func (f *ProductConf) FindResoureList(id string) ([]string, error) {
	var user ProductConf
	err := public.GormPool.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	var resoureList []string
	json.Unmarshal([]byte(user.Resoure), &resoureList)
	return resoureList, nil
}

//RandMd5 随机md5
func RandMd5() string {
	return MD5(NewUUID())
}
