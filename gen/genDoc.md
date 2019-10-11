# gen
Welcome to the technical documentation of the gen package.
The purpose of the gen package is to generate a new vending machine.
The information pertaining to the vending machine is stored in the database table **machine**.
Each time the main function is run, gen should be called to create a new machine.
To get this to work properly, gen should first erase all data in the machine table.
Then, gen generates new data and writes it to the machine table.

>package gen // import "github.com/Tony-Moon/project-0/gen"


## Functions in gen

>func Generate(rows int, columns int, max int) ([]string, []string, []int)
    Generate is the main function inside the gen package. Really, Generate should
    be the only function in the gen package to be called outside of the gen and the
    gen_test package. Additionally, it should only be called once, at the start of
    main.go inside the main package. Generate should be called to generate an index,
    beverage type and a stock amount for each row in the vending machine. First the
    table needs to be cleared of all data. Then, the "make" family of methods should
    be called and return a slice. Generate should send those slices to be written
    into the database table. Finally, Generate has a catch if it is sent a zero or
    less. It does not make sense to have a vending machine with zero rows or zero
    stock capacity.

>func MakeIndex(rows int, columns int) []string
    MakeIndex generates an index key for the database. The key is a letter
    followed by a number (such as A3). A vending machine can have multiple rows
    (denoted by the letter) and multiple columns (denoted by the number).
    main.go in the main package specifies how many rows and columns the vending
    machine will have. The index will always start at "A1."

    >func ToCharStr(i int) string
        ToCharStr increments the letter for each new row.

>func MakeBeverage(rows int, columns int) []string
    MakeBeverage generates a slice that wil tell the database which type of
    drink goes where.

>func MakeStock(rows int, columns int, max int) []int
    MakeStock generates a slice that wil tell the database how many of each
    drink is in the row. It does by creating a random integer between 0 and a
    maximum specified by main.go in the main package.