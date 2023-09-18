package model

type Role struct {
	ID       uint     `gorm:"primaryKey"`
	Name     string   `json:"name"`
	Accesses []Access `json:"access,omitempty"`
}
