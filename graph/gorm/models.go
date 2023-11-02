package gormModel

import (
	"gorm.io/gorm"
)

type NewCompany struct {
	gorm.Model
	CompanyID string    `json:"companyId"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Jobs      []*NewJob `json:"jobs" gorm:"foreignKey:CompanyId"`
}

type NewUser struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewJob struct {
	gorm.Model
	Title     string `json:"title"`
	CompanyId uint64 `json:"companyId"`
}
