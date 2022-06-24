package domain

type Book struct {
	Id          int    `json:"id" gorm:"primaryKey" gorm:"AUTO_INCREMENT" faker:"oneof: 1, 2,3,4,5,6,7,8,9,10"`
	ISBN        string `json:"isbn"  faker:"uuid_digit"`
	Title       string `json:"title" faker:"word"`
	Author      string `json:"author" faker:"name"`
	Description string `json:"descr" faker:"sentence"`
	CreatedAt   string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP" faker:"timestamp"`
}

type Pagination struct {
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
	Sort  string `json:"sort"`
}

func (b *Book) TableName() string {
	return "book"
}
