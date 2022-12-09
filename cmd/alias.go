/*
Copyright © 2021 ras0q <asymptote.k.k@gmail.com>
*/
package cmd

import (
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

func (c *Cmds) aliasCmd() *cobra.Command {
	// aliasCmd represents the alias command
	aliasCmd := &cobra.Command{
		Use:     "alias",
		Short:   "Call aliases",
		Aliases: []string{"a"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias, err := c.h.CallAlias(c.ctx, args[0])
			if err != nil {
				return err
			}

			if err := c.q.PostMessage(
				uuid.FromStringOrNil(c.pl.Message.ChannelID),
				alias.Long,
				true,
			); err != nil {
				return err
			}

			return nil
		},
	}

	aliasCmd.AddCommand(c.addAliasCmd())

	return aliasCmd
}

func (c *Cmds) addAliasCmd() *cobra.Command {
	// addAliasCmd represents the addAlias command
	addAliasCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an new alias",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			userID, err := uuid.FromString(c.pl.Message.User.ID)
			if err != nil {
				return err
			}

			if err := c.h.AddAlias(c.ctx, userID, args[0], args[1]); err != nil {
				return err
			}

			if err := c.q.PostMessage(
				uuid.FromStringOrNil(c.pl.Message.ChannelID),
				":attoteki_seicho:",
				true,
			); err != nil {
				return err
			}

			return nil
		},
	}

	return addAliasCmd
}
