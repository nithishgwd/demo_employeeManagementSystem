package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	office "github.com/nithishgwd/office_module"
)

func main() {
	off := office.NewOffice()

	off.AddTeam("Team A")
	off.AddTeam("Team B")

	for {
		fmt.Println("\nAvailable Teams:")
		teamNames := off.GetTeamNames()
		for i, name := range teamNames {
			fmt.Printf("%d. %s\n", i+1, name)
		}

		fmt.Print("Enter the number corresponding to the team to update employees today (0 to exit): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		if choice == 0 {
			break
		}

		if choice <= 0 || choice > len(teamNames) {
			fmt.Println("Invalid choice. Please select a valid team.")
			continue
		}

		selectedTeam := teamNames[choice-1]
		fmt.Printf("Updating employees today for Team %s\n", selectedTeam)
		fmt.Print("Enter the number of employees present today: ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		numEmployees, err := strconv.Atoi(input)
		if err != nil || numEmployees < 0 {
			fmt.Println("Invalid number of employees. Please enter a valid number.")
			continue
		}

		off.UpdateEmployeesTodayForTeam(selectedTeam, numEmployees)
		fmt.Printf("Employees today for Team %s has been updated to %d\n", selectedTeam, numEmployees)

		// Display all office information
		off.DisplayOfficeInformation()
	}
}
