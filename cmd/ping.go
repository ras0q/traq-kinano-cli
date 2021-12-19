/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"github.com/Ras96/traq-kinano-cli/util/traq"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
func (c *Cmds) pingCmd() *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "Return pong",
		RunE: func(cmd *cobra.Command, args []string) error {
			traq.MustPostMessage(c.pl.Message.ChannelID, "pong!!!")

			return c.h.Ping()
		},
	}

	return pingCmd
}
