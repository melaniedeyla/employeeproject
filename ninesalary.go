package salary

import (
	person "secondProject/person"
)

type NineSalary struct{}

func (s *NineSalary) Salary(e *person.Employee) bool {
	// salary employee
	return true
}
