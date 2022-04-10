package techpalace

import (
	"fmt"
	"strings"
)

const MSM = `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars string = ""
	for i := 0; i < numStarsPerLine; i++ {
		stars += "*"
	}
	msm := stars + "\n" + welcomeMsg + "\n" + stars
	return msm
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	fmt.Println(AddBorder("Welcome!", 10))

	msm := CleanupMessage(MSM)
	fmt.Println(msm)
}
