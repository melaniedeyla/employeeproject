package person_go

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	PackagePay  int
	TotalLeave  int
	LeaveTaken  int
	Onboarding  float32
	Termination float32
}

// Function
func about(e Employee) string {
	return fmt.Sprintf("The Employee %s %s has the salary package of $%d, total leave of %d, was onboarded on %d and termination date %d", e.FirstName, e.LastName, e.PackagePay, e.TotalLeave, e.Onboarding, e.Termination)
}

// Method
func (e *Employee) about() string {
	return fmt.Sprintf("The Employee %s %s has the salary package of $%d, total leave of %d, was onboarded on %d and termination date %d", e.FirstName, e.LastName, e.PackagePay, e.TotalLeave, e.Onboarding, e.Termination)
}
