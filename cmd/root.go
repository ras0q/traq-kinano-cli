/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"fmt"

	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/util/traq"
	"github.com/spf13/cobra"
	traqbot "github.com/traPtitech/traq-bot"
)

type Cmds struct {
	h       handler.Handlers
	payload *traqbot.MessageCreatedPayload
}

func NewCmds(h handler.Handlers, pl *traqbot.MessageCreatedPayload) *Cmds {
	return &Cmds{
		h:       h,
		payload: pl,
	}
}

func (c *Cmds) rootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Use:   "Kinano",
		Short: "Hello, \"きなの\" World!!",
	}

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (c *Cmds) Execute(args []string) error {
	root := c.rootCmd()
	root.SetArgs(args)
	root.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		traq.MustPostMessage(c.payload.Message.ChannelID, fmt.Sprintf("```\n%s```", cmd.UsageString()))
	})

	// Add Subcommands
	root.AddCommand(
		c.pingCmd(),
	)

	return root.Execute() //nolint:wrapcheck
}
