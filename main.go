package main

import (
	"fmt"
	leave "secondProject/leave"
	person "secondProject/person"
	"secondProject/salary"
)

// NineSalary interface
type NineSalary interface {
	Salary(e *person.Employee) bool
}

// AnnualLeaveDays interface
type AnnualLeaveDays interface {
	LeavesTaken(e *person.Employee) bool
}

// EmployeeDatabase interface
type EmployeeDatabase struct {
	Salary        NineSalary
	EmployeeLeave AnnualLeaveDays
}

func main() {
	// Pointing to the package person, where employee struct is found
	// person still red as external package not imported
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

	// Declaring 'es' as the variable of employee salary, pointing to salary package.
	es := &salary.NineSalary{}
	// Declaring 'el' as the variable of employee leave, pointing to leave package.
	el := &leave.AnnualLeaveDays{}

	// Declaring melsapp as variable of EmployeeDatabase struct which contains NineSalary and AnnualLeaveDays.
	melsapp := EmployeeDatabase{
		Salary:        es,
		EmployeeLeave: el,
	}

	// melsapp variable of EmployeeDatabase which contains the properties of salary and leave, lowercase salary as its from different package.
	fmt.Println(melsapp.Salary.Salary(e1))
	fmt.Println(melsapp.Salary.Salary(e2))
	// melsapp variable of EmployeeDatabase which contains the properties of salary and leave, lowercase leave as its from different package.
	fmt.Println(melsapp.EmployeeLeave.LeavesTaken(e1))
	fmt.Println(melsapp.EmployeeLeave.LeavesTaken(e2))
}
