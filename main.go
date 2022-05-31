package main

import (
	"fmt"
	person "secondProject"
)

// NineSalary interface
type NineSalary interface {
	Salary(e *person.Employee) bool
}

//AnnualLeaveDays interface
type AnnualLeaveDays interface {
	Leave(e *person.Employee) bool
}

func main() {
	e1 := &person.Employee{
		FirstName:   "Mel",
		LastName:    "Deyla",
		PackagePay:  1000,
		TotalLeave:  3,
		LeaveTaken:  1,
		Onboarding:  01.22,
		Termination: 02.22,
	}

	e2 := &person.Employee{
		FirstName:   "blah",
		LastName:    "blah",
		PackagePay:  2000,
		TotalLeave:  3,
		LeaveTaken:  1,
		Onboarding:  02.22,
		Termination: 03.22,
	}

	e3 := &person.Employee{
		FirstName:   "blah",
		LastName:    "blah",
		PackagePay:  3000,
		TotalLeave:  3,
		LeaveTaken:  1,
		Onboarding:  03.22,
		Termination: 04.22,
	}

	fmt.Println(melsapp.salary.Salary(e1))
	fmt.Println(melsapp.salary.Salary(e2))
}
