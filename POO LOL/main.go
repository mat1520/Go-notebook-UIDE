package main

import "fmt"

type Stats struct {
	HP	 int
	MANA int
	Attack int
	Defense int
}
type Champion struct {
	Name  string
	Role  string
	Stats Stats
}

func (c Champion) Show() {
	fmt.Printf("Champion: %s\n", c.Name)
	fmt.Printf("Role: %s\n", c.Role)
	fmt.Printf("Stats: HP=%d, MANA=%d, Attack=%d, Defense=%d\n", c.Stats.HP, c.Stats.MANA, c.Stats.Attack, c.Stats.Defense)
}

func (c Champion) Attack(target *Champion) {
	damage := c.Stats.Attack - target.Stats.Defense
	if damage > 0 {
		target.Stats.HP -= damage
		fmt.Printf("%s attacked %s for %d damage!\n", c.Name, target.Name, damage)
	} else {
		fmt.Printf("%s attacked %s but it was not very effective...\n", c.Name, target.Name)
	}
}
