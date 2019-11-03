package managers

import (
	"fmt"
	"strconv"
)

//TakeDamage is able to modify the hp value negatively. If you want to increase the player's hp, use a negative int.
func TakeDamage(i int, p *Player) {
	Hp -= i
	Playerhp = strconv.Itoa(Hp) + " hp"
	p.Hitpoints = Playerhp
	fmt.Println("You now have " + Playerhp + ".")
}

//GetMoney is able to modify the money value positively. If you want to take the player's money, use a negative int.
func GetMoney(i int, p *Player) {
	Money += i
	Playermoney = strconv.Itoa(Money) + " credits"
	p.Money = Playermoney
	fmt.Println("You now have " + Playermoney + ".")
}
