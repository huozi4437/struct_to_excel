package object

type TestScores struct {
	Id        string `json:"id" gorm:"id" bson:"id" desc:"id"`
	StudentId string `json:"studentId" gorm:"student_id" bson:"student_id" desc:"学生id"`
	Subject   string `json:"subject" gorm:"subject" bson:"subject" desc:"科目"`
	Scores    int    `json:"scores" gorm:"scores" bson:"scores" desc:"分数"`
}
