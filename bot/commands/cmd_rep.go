package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/Bios-Marcel/discordemojimap"
	"github.com/bwmarrin/discordgo"

	"github.com/Southclaws/cj/discord"
	"github.com/Southclaws/cj/storage"
	"github.com/Southclaws/cj/types"
)

func (cm *CommandManager) commandRep(
	interaction *discordgo.InteractionCreate,
	args map[string]*discordgo.ApplicationCommandInteractionDataOption,
	settings types.CommandSettings,
) (
	context bool,
	err error,
) {
	user := cm.Storage.GetUserOrCreate(interaction.Member.User.ID)
	sort.Slice(user.ReceivedReactions, func(i, j int) bool {
		return user.ReceivedReactions[i].Counter > user.ReceivedReactions[j].Counter
	})
	embed, err := FormatUserReactions(&user.ReceivedReactions, interaction.Member.User, cm.Discord)
	if err != nil {
		cm.replyDirectly(interaction, fmt.Sprintf("Failed to format reactions with error %s", err.Error()))
		return
	}

	cm.replyDirectlyEmbed(interaction, "", embed)

	return
}

// Only post emoji that we have access to
func validateEmoji(input string, serverEmoji []*discordgo.Emoji) bool {
	isValid := discordemojimap.ContainsEmoji(input)
	if isValid == false {
		for _, v := range serverEmoji {
			if v.MessageFormat() == input {
				return true
			}
		}
	}
	return isValid
}

func FormatUserReactions(reactions *[]storage.ReactionCounter, author *discordgo.User, session *discord.Session) (embed *discordgo.MessageEmbed, err error) {
	statsMessage := strings.Builder{}
	statsMessage.WriteString(fmt.Sprintf("**%s's Reactions**\n\n", author.Username)) //nolint:errcheck

	embed = &discordgo.MessageEmbed{Color: 0x3498DB}
	if len(*reactions) == 0 {
		statsMessage.WriteString("There are no reactions to display here!")
	}

	var serverEmoji []*discordgo.Emoji
	// Colllect emoji from every server CJ is in (since it can display them from those)
	for _, guild := range session.S.State.Guilds {
		serverEmoji = append(serverEmoji, guild.Emojis...)
	}

	for _, reaction := range *reactions {
		if validateEmoji(reaction.Reaction, serverEmoji) {
			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:   fmt.Sprintf("%s​", reaction.Reaction),
				Value:  fmt.Sprintf("%dx", reaction.Counter),
				Inline: true,
			})
		}
	}

	// Add one more 'filling' embed field in case there are 2
	// fields (with 1 and 3 fields the formatting is fine)
	if len(*reactions)%3 == 2 {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   "", // Note: the content of this seemingly empty string is u0080
			Value:  "", // Note: the content of this seemingly empty string is u0080
			Inline: true,
		})
	}

	embed.Description = statsMessage.String()

	return embed, nil
}
