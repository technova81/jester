package commands

import (
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TicketCloseCmd struct {
}

func (cmd *TicketCloseCmd) FromArgs(args []string) error {
	return nil
}

func (cmd *TicketCloseCmd) Name() string { return "ticket:close" }
func (cmd *TicketCloseCmd) Help() string {
	return "Deletes a specified number of messages in the channel."
}
func (cmd *TicketCloseCmd) LongHelp() string {
	return `This command deletes a specified number of messages from the current text channel.
		**Usage:**` + "`!purge <limit>`" +
		`**Arguments:**
		* <limit> (Required): An integer specifying the number of messages to delete (up to a maximum limit 100).`
}
func (cmd *TicketCloseCmd) Run(sess *discordgo.Session, msg *discordgo.Message) error {
	channel, err := sess.Channel(msg.ChannelID)
	if err != nil {
		return errors.New("unable to find the channel")
	}

	if strings.HasPrefix(channel.Name, "ticket-") {

		_, err := sess.ChannelDelete(msg.ChannelID)
		return err
	}

	return nil
}

func (cmd *TicketCloseCmd) Subcommands() map[string]Command {
	return nil
}
