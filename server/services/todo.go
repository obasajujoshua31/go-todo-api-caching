package services

type Todo struct {
	ID          int    `gorm:"AUTO_INCREMENT"`
	Title       string `gorm:"size:255"`
	IsCompleted bool
}
