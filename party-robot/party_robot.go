// package partyrobot
package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(`Happy birthday %s! You are now %v years old!`, name, age)
}

func aliasTable(table int) string {
	if table < 10 {
		return "00"
	}
	if table < 100 {
		return "0"
	}
	return ""
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	var nt string = aliasTable(table)

	var msm string = ""
	msm += fmt.Sprintf("Welcome to my party, %s!\n", name)
	msm += fmt.Sprintf("You have been assigned to table %s%v. Your table is %s, exactly %.1f meters from here.\n", nt, table, direction, distance)
	msm += fmt.Sprintf("You will be sitting next to %s.", neighbor)

	return msm
}

func main() {
	// fmt.Println(Welcome("Christiane"))
	// fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
