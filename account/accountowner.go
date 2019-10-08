package account

import (
	"strconv"
)

//struct of the account owner

type accountowner struct {
	custfname, custlname     string
	street, city, state, zip string
	accountno                int
	avalbal                  float64
}

//function of account information

func (a accountowner) string() string {
	var acctinfo string
	acctinfo := a.custfname + "\n" + a.custlname + "\n" + a.street + "\n" + a.city + "\n" + a.state +
	a.accountno + strconv.Itoa(a.accountno) + a.avalbal + strconv.FormatFloat(a.avalbal, 'f', 2, 64)
	return acctinfo
}



/* deposit to availbalr*/

func (a *accountowner) deposit (amount float64) {
	if  amount > 0 {
		a.avalbal = a.avalbal + amount
		fmt.Println(avaibal)
	}
else
	
{
	fmt.Println("You have the same amount of balance")
}
}
/*withdrwal from avalbal*/
func (a *accountowner) withdrwal (amount float64) {
	if a.avalbal > amount{
		a.avalbal = a.avalbal - amount
		fmt.Println(avaibal)
	}
else
	
{
	fmt.Println("You don't have sufficent money")
}
