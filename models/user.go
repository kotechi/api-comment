package model

type Users struct {
	ID           int    `json:"id" gorm:"type:int; primaryKey; autoIncrement"`
	Username     string `json:"username" gorm:"not null; type:varchar(50); unique_index;"`
	PasswordHash string `json:"password_hash" gorm:"not null; type:varchar(255); unique_index;"`
	Email        string `json:"email" gorm:"not null; type:varchar(255); unique_index;"`
	RoleId       int    `json:"role_id" gorm:"type:int; unique_index"`
}
