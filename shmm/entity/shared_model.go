package entity

type SharedModel struct {
	Name     string `json:"name" gorm:"primaryKey"`
	Tag      string `json:"tag" gorm:"primaryKey"`
	Shmname  string `json:"shmname" gorm:"unique"`
	Shmsize  int64  `json:"shmsize"`
	Status   int    `json:"status" gorm:"default:0"` // Pending:0, Activated:1
	Metadata string `json:"metadata"`
}
