package model

type CommentLabel struct {
	Id           int64  `gorm:"id" json:"id"`                       // 自增ID
	CommentLabel string `gorm:"comment_label" json:"comment_label"` // 评价标签名称
	CreateTime   int64  `gorm:"create_time" json:"create_time"`     // 创建时间
	UpdateTime   int64  `gorm:"update_time" json:"update_time"`     // 修改时间
	StarRating   int64  `gorm:"star_rating" json:"star_rating"`     // 标签所属星级
}

type NewComment struct {
	OrderID        int64                 `json:"OrderId"`
	DoctorUin      int64                 `json:"DoctorUin"`
	PatientUin     int64                 `json:"PatientUin"`
	IsShow         int64                 `json:"IsShow"`
	CommentContent string                `json:"CommentContent"`
	CommentLabels  []*NewCommentLabel    `json:"CommentLabels"`
	CommentStars   int64                 `json:"CommentStars"`
}

type NewCommentLabel struct {
	ID           int    `json:"Id"`
	CommentLabel string `json:"CommentLabel"`
	StarRating   int    `json:"StarRating"`
}

type Comment struct {
	Id             int64  `gorm:"id" json:"id"`                           // 自增ID
	DoctorUin      int64  `gorm:"doctor_uin" json:"doctor_uin"`           // 医生ID(t_doctor_info表的uin)
	PatientUin     int64  `gorm:"patient_uin" json:"patient_uin"`         // 患者ID(t_patient_info表的uin)
	OrderId        int64  `gorm:"order_id" json:"order_id"`               // 订单ID(t_order表的id)
	OperatorUin    string `gorm:"operator_uuid" json:"operator_uuid"`     // 操作人员uuid(t_operator_info表的uuid)
	IsShow         int64  `gorm:"is_show" json:"is_show"`                 // 是否展示(1-展示,2-隐藏)
	CommentStars   int64  `gorm:"comment_stars" json:"comment_stars"`     // 评价星级(取值范围1-5)
	CommentContent string `gorm:"comment_content" json:"comment_content"` // 评价内容(长度1000)
	CreateTime     int64  `gorm:"create_time" json:"create_time"`         // 创建时间
	UpdateTime     int64  `gorm:"update_time" json:"update_time"`         // 修改时间
	DbFrom         string `gorm:"db_from" json:"db_from"`                 // 来源数据库（新，旧）
}

type CommentQueryParam struct {
	Id  	 int64 `json:"id"`
	OrderId  int64 `json:"order_id"`
}
