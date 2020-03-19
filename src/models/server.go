package models

// Server info
type Server struct {
	ID     uint   `gorm:"primary_key" json:"id" form:"id"`
	Domain string `gorm:"not null" json:"domain" form:"domain"`
	Port   uint   `gorm:"not null;unique" json:"port" form:"port"`
	Active int    `gorm:"default:0" json:"active" form"active"`
}
