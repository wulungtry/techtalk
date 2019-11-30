package domain

import "time"

type Person struct {
	Name        string
	CityOfBirth string
	DateOfBirth time.Time
	Age         int32
	Weight      int32
	Height      int32
	IsMarried   bool
}
