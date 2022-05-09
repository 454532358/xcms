package model

type Category struct {
	Catid        int    `json:"catid"`
	Model        string `json:"model"`
	Type         int    `json:"type"`
	ModelId      int    `json:"model_id"`
	ParentId     int    `json:"parent_id"`
	CatName      string `json:"cat_name"`
	CatDir       string `json:"cat_dir"`
	Thumb        string `json:"thumb"`
	Url          string `json:"url"`
	Sort         int    `json:"sort"`
	IsShow       int    `json:"is_show"`
	IsMeshow     int    `json:"is_meshow"`
	IsTarget     int    `json:"is_target"`
	IsHtml       int    `json:"is_html"`
	IsLink       int    `json:"is_link"`
	TemplateCate string `json:"template_cate"`
	TemplateList string `json:"template_list"`
	TemplateShow string `json:"template_show"`
	SeoTitle     string `json:"seo_title"`
	SeoKey       string `json:"seo_key"`
	SeoDes       string `json:"seo_des"`
	Power        string `json:"power"`
	UrlList      int    `json:"url_list"`
	UrlShow      int    `json:"url_show"`
}

func (c Category) GetList(pageSize int) {
	//TODO implement me
	panic("implement me")
}
