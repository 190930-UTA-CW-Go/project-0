package app

/*
NewTech documentation
*/
func NewTech(acc string, pass string, comp string, first string, last string) *Tech {
	return &Tech{
		account:   acc,
		password:  pass,
		company:   comp,
		firstname: first,
		lastname:  last,
	}
}
