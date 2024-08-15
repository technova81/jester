package commands

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type HelpCmd struct {
	args []string
}

func (cmd *HelpCmd) FromArgs(args []string) error {
	cmd.args = args
	return nil
}

func (cmd *HelpCmd) Name() string { return "help" }
func (cmd *HelpCmd) Help() string { return "Displays command usage details." }
func (cmd *HelpCmd) LongHelp() LongHelp {
	return LongHelp{
		About:       "Provides detailed information on how to use the available commands.",
		Usage:       "`!help <commandName>` or `!help`",
		Arguments:   nil,
		Subcommands: nil,
	}
}
func (cmd *HelpCmd) Run(sess *discordgo.Session, msg *discordgo.Message) error {

	if len(cmd.args) == 0 {
		embed := &discordgo.MessageEmbed{
			Title:       "Available Commands:",
			Description: "Here's a list of available commands:",
			Color:       0x00FFFF,
		}
		for k, v := range Commands {
			field := &discordgo.MessageEmbedField{
				Name:  k,
				Value: v.Help(),
			}
			embed.Fields = append(embed.Fields, field)
		}

		_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
		return err
	} else {
		// case 1:
		// 	{
		// 		embed := &discordgo.MessageEmbed{
		// 			Title:       args[0],
		// 			Description: Commands[args[0]].Help(),
		// 			Fields: []*discordgo.MessageEmbedField{
		// 				{
		// 					Name:  "About",
		// 					Value: Commands[args[0]].LongHelp(),
		// 				},
		// 			},
		// 		}
		// 		_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
		// 		return err
		// 	}

		// default:
		// 	return errors.New("invalid command format \n Usage: !help or !help <commandName>")

		var currCmd Command = Commands[cmd.args[0]]
		for i := 1; i < len(cmd.args); i++ {
			if subCmds := currCmd.Subcommands(); subCmds != nil {
				currCmd = subCmds[cmd.args[i]]
			} else {
				return errors.New("invalid subcommand")
			}

		}
		longHelp := currCmd.LongHelp()
		fields := []*discordgo.MessageEmbedField{
			{
				Name:  "About",
				Value: longHelp.About,
			},
			{
				Name:  "Usage",
				Value: longHelp.Usage,
			},
		}

		if longHelp.Arguments != nil {
			argumentDescriptions := ""
			for _, arg := range longHelp.Arguments {
				description := fmt.Sprintf("* `%s` (Required): %s\n", arg.Name, arg.Help)
				argumentDescriptions += description
			}
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  "Arguments",
				Value: argumentDescriptions,
			})
		}

		// if longHelp.Subcommands != nil {
		// 	fields = append(fields, &discordgo.MessageEmbedField{
		// 		Name:  "Subcommands",
		// 		Value: *longHelp.Subcommands,
		// 	})
		// }

		embed := &discordgo.MessageEmbed{
			Title:       currCmd.Name(),
			Description: currCmd.Help(),
			Fields:      fields,
		}
		_, err := sess.ChannelMessageSendEmbed(msg.ChannelID, embed)
		return err
	}
}

func (cmd *HelpCmd) Subcommands() map[string]Command {
	return nil
}
