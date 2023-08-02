package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	employee "github.com/nithishgwd/employee_Module"
	office "github.com/nithishgwd/office_module"
	team "github.com/nithishgwd/team_module"
)

func main() {

	// Access employee details from the employee package
	emp1 := employee.NewEmployee("John", 30)
	emp2 := employee.NewEmployee("Jane", 28)

	// Create new teams
	team1 := team.NewTeam("Team A")
	team2 := team.NewTeam("Team B")

	// Add employees to the teams
	team1.AddEmployee(emp1)
	team2.AddEmployee(emp2)

	// Update employees today for the teams
	team1.UpdateEmployeesToday(10)
	team2.UpdateEmployeesToday(8)

	// Display team information
	fmt.Println("\nTeam Information:")
	fmt.Printf("Team: %s, Employees Today: %d\n", team1.GetName(), team1.GetEmployeesToday())
	fmt.Printf("Team: %s, Employees Today: %d\n", team2.GetName(), team2.GetEmployeesToday())

	// Access office details from the office package
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
