package database

type Buku struct {
	Id        int64  `gorm:"primaryKey;autoIncrement:true"`
	Judul     string `json:"judul"`
	Pengarang string `json:"pengarang"`
}
