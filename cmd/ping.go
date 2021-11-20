/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
func (c *Cmds) pingCmd() *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.h.Ping(c.payload.Message.ChannelID)
		},
	}

	return pingCmd
}
