package main

import (
	"bot-discord/commands"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const version = "v1.0.0"

func Menu() {
	cleanScreen()

	banner := `
 ██▀███   ▄▄▄       ██▓▓█████▄     ██▓███   ▄▄▄       ███▄    █ ▓█████  ██▓    
▓██ ▒ ██▒▒████▄    ▓██▒▒██▀ ██▌   ▓██░  ██▒▒████▄     ██ ▀█   █ ▓█   ▀ ▓██▒    
▓██ ░▄█ ▒▒██  ▀█▄  ▒██▒░██   █▌   ▓██░ ██▓▒▒██  ▀█▄  ▓██  ▀█ ██▒▒███   ▒██░    
▒██▀▀█▄  ░██▄▄▄▄██ ░██░░▓█▄   ▌   ▒██▄█▓▒ ▒░██▄▄▄▄██ ▓██▒  ▐▌██▒▒▓█  ▄ ▒██░    
░██▓ ▒██▒ ▓█   ▓██▒░██░░▒████▓    ▒██▒ ░  ░ ▓█   ▓██▒▒██░   ▓██░░▒████▒░██████▒
░ ▒▓ ░▒▓░ ▒▒   ▓▒█░░▓   ▒▒▓  ▒    ▒▓▒░ ░  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ ░░ ▒░ ░░ ▒░▓  ░
  ░▒ ░ ▒░  ▒   ▒▒ ░ ▒ ░ ░ ▒  ▒    ░▒ ░       ▒   ▒▒ ░░ ░░   ░ ▒░ ░ ░  ░░ ░ ▒  ░
  ░░   ░   ░   ▒    ▒ ░ ░ ░  ░    ░░         ░   ▒      ░   ░ ░    ░     ░ ░   
   ░           ░  ░ ░     ░                      ░  ░         ░    ░  ░    ░  ░
                        ░                                                      `

	fmt.Println(banner)
	fmt.Printf(" [ Discord Raid Panel - %s ]\n", version)
	fmt.Println("Hello user, type 'help' for view available commands.")

}

func main() {

	opcion := bufio.NewScanner(os.Stdin)

	fmt.Print("Insert Bot Token: ")
	var token string
	if opcion.Scan() {
		token = strings.TrimSpace(opcion.Text())
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Sesion Error.", err)
		return
	}

	err = dg.Open()
	if err != nil {
		fmt.Println("Auth Error.", err)
		return
	}
	defer dg.Close()
	Menu()

	for {
		fmt.Print("panel@user:~$ ")

		if opcion.Scan() {
			entrada := opcion.Text()

			switch entrada {

			case "help":
				cleanScreen()
				Menu()
				help()

			case "1":
				fmt.Println("Starting...")
				commands.DeleteChannels(dg)

			case "2":
				fmt.Println("Changing server name...")
				commands.ChangeName(dg)

			case "3":
				cleanScreen()
				Menu()

			case "4":
				fmt.Println("Goodbye!")
				cleanScreen()
				return

			default:
				fmt.Println("Invalid Option...")
			}
		}
	}
}

func help() {
	fmt.Println("Available Commands: ")
	fmt.Println("[1] Delete & Create channels")
	fmt.Println("[2] Change server name")
	fmt.Println("[3] Clear Screen")
	fmt.Println("[4] Exit")
}

func cleanScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		fmt.Print("\033[H\033[2J")
	}
}
