package domain

import "time"

type Person struct {
	Name        string
	CityOfBirth string
	DateOfBirth time.Time
	Age         int
	IsMarried   bool
}
