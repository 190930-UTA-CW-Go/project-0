package managers

import "fmt"

//Inventory is a slice that stores the items within the player's inventory.
var Inventory []string

//Backpack allows the player to check their inventory.
func Backpack() {
	if len(Inventory) >= 1 {
		fmt.Println("You search your backpack and find:")
		for i := 0; i < len(Inventory); i++ {
			fmt.Println(Inventory[i])
		}
	} else {
		fmt.Println("You search your backpack, but find nothing.")
	}
}

//InventoryAdd allows to the program to add items to the player's inventory.
func InventoryAdd(s string) {
	Inventory = append(Inventory, s)
}

//InventoryCheck checks the player's inventory for items necessary to complete certain actions.
func InventoryCheck(s string) bool {
	check := false
	for i := 0; i < len(Inventory); i++ {
		if Inventory[i] == s {
			check = true
		}
	}
	return check
}
