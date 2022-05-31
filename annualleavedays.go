package person

type AnnualLeaveDays struct{}

func (al AnnualLeaveDays) Leave(e *person.Employee) bool {
	//Annual Leave left for employee
	//Remember to do the subtraction from annualleave days to leavetaken
	return true
}

