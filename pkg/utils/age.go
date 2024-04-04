package utils

import "time"

func CalculateAge(dob string) int {
	birthDate, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return -1
	}
	age := time.Since(birthDate).Hours() / 24 / 365
	return int(age)
}
