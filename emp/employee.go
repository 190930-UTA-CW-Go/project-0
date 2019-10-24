package emp

//Employee structure holds the values needed to edit accounts
type Employee struct {
	EmpID       string
	EmpUsername string
	EmpPass     string
	EmpFName    string
	EmpLName    string
}

//CreateNewEmployee constructor for the Employee struct
func CreateNewEmployee(id string, username string, pass string, fName string, lName string) Employee {
	return Employee{id, username, pass, fName, lName}
}

//GetEmployeeID returns the employee ID
func (a *Employee) GetEmployeeID() string {
	return a.EmpID
}

//GetEmployeeUsername returns the employee Username
func (a *Employee) GetEmployeeUsername() string {
	return a.EmpUsername
}

//GetEmployeePass returns the employee Password
func (a *Employee) GetEmployeePass() string {
	return a.EmpPass
}

//GetEmployeeFName returns the employee First Name
func (a *Employee) GetEmployeeFName() string {
	return a.EmpFName
}

//GetEmployeeLName returns the employee Last Name
func (a *Employee) GetEmployeeLName() string {
	return a.EmpLName
}
