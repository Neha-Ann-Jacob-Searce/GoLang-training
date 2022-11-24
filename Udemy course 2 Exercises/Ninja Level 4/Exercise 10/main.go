package main

import "fmt"

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"Shaken, not stirred", "Martinis", "Women"}
	m["moneypenny_miss"] = []string{"James Bond", "Literature", "Computer Science"}
	m["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}
	m["dr"] = []string{"evil", "Ice", "Sun"}
	delete(m, "dr")
	for key, value := range m {
		fmt.Println(key, value)
	}
}
