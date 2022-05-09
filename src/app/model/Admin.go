package model

type Admin struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Encrypt   string `json:"encrypt"`
	LastTime  int    `json:"last_time"`
	CreatTime int    `json:"creat_time"`
	IsLock    int    `json:"is_lock"`
	LastIp    string `json:"last_ip"`
	Logins    int    `json:"logins"`
	GroupId   int    `json:"groupid"`
}

func (a Admin) TableName() string {
	return "sf_admin"
}
