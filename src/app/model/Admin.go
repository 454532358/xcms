package model

//CREATE TABLE `sf_admin` (
//  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//  `username` varchar(50) NOT NULL,
//  `password` varchar(50) NOT NULL,
//  `encrypt` char(20) NOT NULL,
//  `last_time` int(10) NOT NULL DEFAULT '0',
//  `creat_time` int(10) NOT NULL,
//  `is_lock` smallint(2) NOT NULL DEFAULT '0',
//  `last_ip` varchar(50) NOT NULL,
//  `logins` int(10) NOT NULL DEFAULT '0',
//  `groupid` smallint(4) unsigned NOT NULL DEFAULT '0',
//  PRIMARY KEY (`id`) USING BTREE
//) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;
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
	Groupid   int    `json:"groupid"`
}
