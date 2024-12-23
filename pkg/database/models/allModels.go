package models

import (
	"time"
)

type UserRole string

const (
	StaffRole     = "StaffRole"
	AdminRole     = "AdminRole"
	PartnerRole   = "PartnerRole"
	VendorRole    = "VendorRole"
	SuperUserRole = "SuperUserRole"
	PersonRole    = "PersonRole"
)

type Donations struct {
	ID        uint64    `json:"id,omitempty"`
	Title     string    `gorm:"not null" json:"title,omitempty"`
	Amount    float64   `gorm:"not null" json:"amount,omitempty"`
	Ward      *Ward     `gorm:"foreignKey:ward_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ward,omitempty"`
	WardID    uint64    `json:"wardId,omitempty"`
	User      *User     `json:"user,omitempty"`
	UserID    uint64    `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type Ward struct {
	ID        uint64       `json:"id,omitempty"`
	Title     string       `gorm:"not null" json:"title,omitempty"`
	FullName  string       `gorm:"not null" json:"fullName,omitempty"`
	Address   string       `gorm:"not null" json:"address,omitempty"`
	Want      string       `gorm:"not null" json:"want,omitempty"`
	Collected float64      `gorm:"default:0" json:"collected"`
	Necessary float64      `gorm:"not null" json:"necessary,omitempty"`
	Donations []*Donations `gorm:"foreignKey:ward_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donations,omitempty"`
	CreatedAt time.Time    `json:"createdAt,omitempty"`
	UpdatedAt time.Time    `json:"updatedAt,omitempty"`
}

type CardCompany struct {
	ID        uint64    `json:"id,omitempty"`
	FullName  string    `gorm:"not null" json:"fullName,omitempty"`
	Number    string    `gorm:"index,unique" json:"number,omitempty"`
	Date      string    `gorm:"not null" json:"date,omitempty"`
	Cvv       uint64    `gorm:"not null" json:"cvv,omitempty"`
	CompanyID uint64    `json:"companyId,omitempty"`
	Company   *Company  `json:"omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type Company struct {
	ID        uint64       `json:"id,omitempty"`
	Title     string       `json:"title,omitempty"`
	Phone     string       `gorm:"index,unique" json:"phone,omitempty"`
	Address   string       `gorm:"not null" json:"address,omitempty"`
	Site      string       `json:"site,omitempty"`
	INN       string       `gorm:"not null" json:"inn,omitempty"`
	KPP       string       `gorm:"not null" json:"kpp,omitempty"`
	OKPO      string       `gorm:"not null" json:"okpo,omitempty"`
	Card      *CardCompany `gorm:"foreignKey:company_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"card,omitempty"`
	UserID    uint64       `json:"userId,omitempty"`
	User      *User        `json:"user,omitempty"`
	CreatedAt time.Time    `json:"createdAt,omitempty"`
	UpdatedAt time.Time    `json:"updatedAt,omitempty"`
}

type Card struct {
	ID        uint64    `json:"id,omitempty"`
	FullName  string    `gorm:"not null" json:"fullName,omitempty"`
	Number    string    `gorm:"index,unique" json:"number,omitempty"`
	Date      string    `gorm:"not null" json:"date,omitempty"`
	Cvv       uint64    `gorm:"not null" json:"cvv,omitempty"`
	UserID    uint64    `json:"userId,omitempty"`
	User      *User     `json:"omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// User - модель пользователей
type User struct {
	ID        uint64       `json:"id,omitempty"`
	Email     string       `gorm:"index:,unique" json:"email,omitempty"`
	Username  string       `gorm:"not null" json:"username,omitempty"`
	Password  string       `gorm:"not null" json:"password,omitempty"`
	Phone     string       `gorm:"index:,unique,not null" json:"phone,omitempty"`
	Card      []*Card      `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"card,omitempty"`
	Role      string       `gorm:"not null" json:"role,omitempty"`
	Company   *Company     `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"company,omitempty"`
	Type      uint64       `gorm:"default:0" json:"type,omitempty"`
	Donations []*Donations `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"donations,omitempty"`
	CreatedAt time.Time    `json:"createdAt,omitempty"`
	UpdatedAt time.Time    `json:"updatedAt,omitempty"`
}

// Session - сессия пользователей
type Session struct {
	ID           uint64    `json:"id,omitempty"`
	UserID       uint64    `gorm:"index:,unique" json:"userId,omitempty"`
	RefreshToken string    `json:"refreshToken,omitempty"`
	CreatedAt    time.Time `json:"createdAt,omitempty"`
	UpdatedAt    time.Time `json:"updatedAt,omitempty"`
}
