package leave_go

import (
	person "secondProject/person"
)

type AnnualLeaveDays struct{}

func (al AnnualLeaveDays) LeavesTaken(e *person.Employee) bool {
	//Annual Leave left for employee
	//Remember to do the subtraction from annualleave days to leavetaken
	return true
}
