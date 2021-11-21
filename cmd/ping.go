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
		Short: "Return pong",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.h.Ping(c.payload.Message.ChannelID)
		},
	}

	return pingCmd
}
