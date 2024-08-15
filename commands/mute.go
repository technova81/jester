package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type MuteCmd struct {
	UserID string
}

func (cmd *MuteCmd) FromArgs(args []string) error {
	if len(args) == 1 {
		cmd.UserID = args[0][2 : len(args[0])-1]
		return nil
	}

	return errors.New("invalid command \n Usage: !mute @user")
}

func (cmd *MuteCmd) Name() string { return "mute" }
func (cmd *MuteCmd) Help() string { return "Mute a given user" }
func (cmd *MuteCmd) LongHelp() LongHelp {
	// return `This command mutes a user on the server, depending on the context.
	// 	**Usage:**` + "`!mute @user`" +
	// 	`**Arguments:**
	// 	* @user (Required): Mention the user to unmute using the "@" symbol followed by their username.
	// 	**Note:User must be connect to a voice channel**`

	return LongHelp{
		About: "This command mutes a user on the server",
		Usage: "`!mute @user`",
		Arguments: []Argument{
			{
				Name:     "user",
				Required: true,
				Help:     `Mention the user to mute using the "@" symbol followed by their username(User must be connected to a voice channel).`,
			},
		},
		Subcommands: nil,
	}
}

func (cmd *MuteCmd) Run(sess *discordgo.Session, msg *discordgo.Message) error {
	return sess.GuildMemberMute(msg.GuildID, cmd.UserID, true)
}

func (cmd *MuteCmd) Subcommands() map[string]Command {
	return nil
}
