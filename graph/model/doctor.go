package model

type Doctor struct {
	Uin                          int64  `gorm:"uin" json:"uin"`             // 医生uin
	Name                         string `gorm:"name" json:"name"`           // 医生姓名
	Department                   string `gorm:"department" json:"department"`
	Title                        string `gorm:"title" json:"title"`
	Supplier                     string `gorm:"supplier" json:"supplier"`   // 供应商
}

type DoctorQueryParam struct {
	Uin int64 `json:"Uin"`
}
