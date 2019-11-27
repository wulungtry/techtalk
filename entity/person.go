package entity

import "time"

type Person struct {
	PersonalId  string    `gorm:"primary_key" gorm:"column:PersonalId"`
	Name        string    `gorm:"column:Name"`
	CityOfBirth string    `gorm:"column:CityOfBirth"`
	DateOfBirth time.Time `gorm:"column:DateOfBirth"`
	Height      int       `gorm:"column:Height"`
	Weight      int       `gorm:"column:Weight"`
	IsMarried   bool      `gorm:"column:IsMarried"`
}
