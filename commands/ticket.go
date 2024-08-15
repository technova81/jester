package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

var TicketSubCommands = map[string]Command{
	// "create": &TicketCreateCmd{},
	// "close":  &TicketCloseCmd{},
}

type TicketCmd struct {
	Runner func(sess *discordgo.Session, msg *discordgo.Message)
}

func (cmd *TicketCmd) FromArgs(args []string) error {
	if subCmd, ok := TicketSubCommands[args[0]]; ok {
		subCmd.FromArgs(args[1:])
		cmd.Runner = func(sess *discordgo.Session, msg *discordgo.Message) {
			subCmd.Run(sess, msg)
		}
	}
	return nil
}

func (cmd *TicketCmd) Name() string { return "ticket" }
func (cmd *TicketCmd) Help() string { return "Create or Closes a ticket for user" }
func (cmd *TicketCmd) LongHelp() LongHelp {

	// arguments := `* **@user** (Required): Mention the user to ban using the "@" symbol followed by their username.` + "\n" +
	// 	`* **[days]** (Required): Specify the duration of the ban in days. If omitted, the ban will be permanent.` + "\n" +
	// 	`* **[reason]** (Required): Provide a reason for the ban. The reason will be DM to user.`

	return LongHelp{
		About:       "This command bans a given user for specific days from the server.",
		Usage:       "`!ban @user [days] [reason]`",
		Arguments:   nil,
		Subcommands: nil,
	}

}
func (cmd *TicketCmd) Run(sess *discordgo.Session, msg *discordgo.Message) error {
	if cmd.Runner != nil {
		cmd.Runner(sess, msg)
	} else {
		return errors.New("invalid command format \n Usage: !ticket create <reason> or !ticket close")
	}
	cmd.Runner = nil
	return nil
}

func (cmd *TicketCmd) Subcommands() map[string]Command {
	return TicketSubCommands
}
