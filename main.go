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

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

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

	fmt.Println(Red + banner + Reset)
	fmt.Println("[ Discord Raid Panel ]")
	fmt.Println("Hello user, type" + Green + " 'help' " + Reset + "for view available commands.")

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
				commands.SpamWebhooks()

			case "4":
				commands.ChangeName(dg)
				commands.DeleteChannels(dg)

			case "5":
				cleanScreen()
				Menu()

			case "6":
				fmt.Println("Goodbye!")
				cleanScreen()
				return

			default:
				fmt.Println("Invalid Option...")
				cleanScreen()
				Menu()
			}
		}
	}
}

func help() {
	fmt.Println("Available Commands: ")
	fmt.Println("[1] Nuke Channels")
	fmt.Println("[2] Change server name")
	fmt.Println("[3] Webhook Spam")
	fmt.Println("[4] Nuke All")
	fmt.Println("[5] Clear Screen")
	fmt.Println("[6] Exit")
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
