package entity

import "time"

type Person struct {
	Personalid  string    `gorm:"primary_key" gorm:"column:PersonalId"`
	Name        string    `gorm:"column:Name"`
	Cityofbirth string    `gorm:"column:CityOfBirth"`
	Dateofbirth time.Time `gorm:"column:DateOfBirth"`
	Height      int32     `gorm:"column:Height"`
	Weight      int32     `gorm:"column:Weight"`
	Ismarried   bool      `gorm:"column:IsMarried"`
}
