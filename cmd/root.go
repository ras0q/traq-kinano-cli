/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"context"
	"fmt"

	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/spf13/cobra"
	"github.com/traPtitech/traq-ws-bot/payload"
)

var CmdNames = map[string]struct{}{
	"a":     {},
	"alias": {},
	"help":  {},
	"ping":  {},
}

type Cmds struct {
	ctx context.Context
	h   handler.Handlers
	pl  *payload.MessageCreated
	w   Writer
}

func NewCmds(ctx context.Context, h handler.Handlers, pl *payload.MessageCreated, w Writer) *Cmds {
	return &Cmds{ctx, h, pl, w}
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
func (c *Cmds) Execute(args []string) {
	root := c.rootCmd()
	c.addSubCmds(root)
	root.SetArgs(args)
	root.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		_, _ = c.w.
			SetChannelID(c.pl.Message.ChannelID).
			Write([]byte(fmt.Sprintf("```\n%s```", cmd.UsageString())))
	})

	if err := root.Execute(); err != nil {
		_, _ = c.w.
			SetChannelID(c.pl.Message.ChannelID).
			Write([]byte(fmt.Sprintf("```\n%s```", err.Error())))
	}
}

func (c *Cmds) addSubCmds(cmd *cobra.Command) {
	cmd.AddCommand(
		c.aliasCmd(),
		c.pingCmd(),
	)
}
