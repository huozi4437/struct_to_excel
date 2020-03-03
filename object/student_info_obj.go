package object

type StudentInfo struct {
	Id      string `json:"id" gorm:"id" bson:"id" desc:"id"`
	Name    string `json:"name" gorm:"name" bson:"name" desc:"姓名"`
	Age     int    `json:"age" gorm:"age" bson:"age" desc:"年龄"`
	Class   string `json:"class" gorm:"class" bson:"class" desc:"班级"`
	Address string `json:"address" gorm:"address" bson:"address" desc:"住址"`
}
