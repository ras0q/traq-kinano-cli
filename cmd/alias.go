/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

func (c *Cmds) aliasCmd() *cobra.Command {
	// aliasCmd represents the alias command
	aliasCmd := &cobra.Command{
		Use:   "alias",
		Short: "A brief description of your command",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.h.CallAlias(c.payload.Message.ChannelID, args[0])
		},
	}

	aliasCmd.AddCommand(c.addAliasCmd())

	return aliasCmd
}

func (c *Cmds) addAliasCmd() *cobra.Command {
	// addAliasCmd represents the addAlias command
	addAliasCmd := &cobra.Command{
		Use:   "add",
		Short: "A brief description of your command",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			userID, err := uuid.FromString(c.payload.Message.User.ID)
			if err != nil {
				return err
			}

			return c.h.AddAlias(c.payload.Message.ChannelID, userID, args[0], args[1])
		},
	}

	return addAliasCmd
}
