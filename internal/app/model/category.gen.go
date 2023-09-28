package model

const TableNameCategory = "category"

// Category mapped from table <category>
type Category struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Model         string `gorm:"column:model;not null" json:"model"`
	Type          bool   `gorm:"column:type;not null" json:"type"`
	ModelID       int32  `gorm:"column:model_id;not null" json:"model_id"`
	ParentID      int32  `gorm:"column:parent_id;not null" json:"parent_id"`
	CatName       string `gorm:"column:cat_name;not null" json:"cat_name"`
	CatDir        string `gorm:"column:cat_dir;not null" json:"cat_dir"`
	Thumb         string `gorm:"column:thumb" json:"thumb"`
	URL           string `gorm:"column:url;not null" json:"url"`
	Sort          int32  `gorm:"column:sort;not null" json:"sort"`
	IsShow        bool   `gorm:"column:is_show;not null" json:"is_show"`
	IsMeshow      bool   `gorm:"column:is_meshow" json:"is_meshow"`
	IsTarget      bool   `gorm:"column:is_target;not null" json:"is_target"`
	IsHTML        bool   `gorm:"column:is_html;not null" json:"is_html"`
	IsLink        bool   `gorm:"column:is_link;not null" json:"is_link"`
	TemplateCate  string `gorm:"column:template_cate;not null" json:"template_cate"`
	TemplateList  string `gorm:"column:template_list;not null" json:"template_list"`
	TemplateShow  string `gorm:"column:template_show;not null" json:"template_show"`
	SeoTitle      string `gorm:"column:seo_title;not null" json:"seo_title"`
	SeoKey        string `gorm:"column:seo_key;not null" json:"seo_key"`
	SeoDes        string `gorm:"column:seo_des;not null" json:"seo_des"`
	Power         string `gorm:"column:power" json:"power"`
	URLList       int32  `gorm:"column:url_list;not null" json:"url_list"`
	URLShow       int32  `gorm:"column:url_show;not null" json:"url_show"`
	CategoryModel ModelM `gorm:"foreignKey:id;references:model_id" json:"category_model"`
}

// TableName Category's table name
func (*Category) TableName() string {
	return TableNameCategory
}
