package main

//userinfo is a structure to represent user information
type Userinfo struct {
	uniqname  string
	userfirst string
	userlast  string
	password  string
	funds     float32
}

func newuser(uniqname string, userfirst string, userlast string, password string, funds float32) Userinfo {
	return Userinfo{
		uniqname:  uniqname,
		userfirst: userfirst,
		userlast:  userlast,
		password:  password,
		funds:     funds,
	}
}

//getuser is a function that collects username and password form new users

/* func getuser() {
	fmt.Println("Please enter a username below:")
	var collect userinfo
	collect.username = "cardib"
	fmt.Println(collect.username)

	fmt.Println("Please choose a password.")
	collect.password = "bodakyellow"
	fmt.Println(collect.password)

}
*/
