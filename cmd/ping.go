/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
func (c *Cmds) pingCmd() *cobra.Command {
	pingCmd := &cobra.Command{
		Use:   "ping",
		Short: "Return pong",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := c.q.PostMessage(
				uuid.FromStringOrNil(c.pl.Message.ChannelID),
				"pong",
				true,
			); err != nil {
				return err
			}

			return c.h.Ping()
		},
	}

	return pingCmd
}
