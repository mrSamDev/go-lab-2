package main

import "fmt"

func main() {
	chartype := map[string][]string{
		"bond_james":       {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_jenny": {"Smart, witty", "Loyal", "Resourceful"},
		"no_dr":            {"Always ready", "Resourceful", "Clever"},
	}

	chartype["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(chartype, "no_dr")

	fmt.Println("Characters and their IDs:", chartype)

}
