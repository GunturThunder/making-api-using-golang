package models

import "time"

type Address struct {
	ID         string `gorm:"size:36;nol null;uniqueIndex;primary_key"`
	User       User
	UserID     string `grom:"size:36;index"`
	Name       string `gorm:size:100`
	isPrimary  bool
	CityID     string `gorm:size:100`
	ProvinceID string `gorm:size:100`
	Address1   string `gorm:size:255`
	Address2   string `gorm:size:255`
	Phone      string `gorm:size:100`
	Email      string `gorm:size:100`
	PostCode   string `gorm:size:100`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}