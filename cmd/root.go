/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/util/traq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	traqbot "github.com/traPtitech/traq-bot"
)

var cfgFile string

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
		Use:   "traq-kinano-cli",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	initRootCmd(rootCmd)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func (c *Cmds) Execute(args []string) error {
	root := c.rootCmd()
	root.SetArgs(args)
	root.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		traq.MustPostMessage(c.payload.Message.ChannelID, fmt.Sprintf("```\n%s\n```", cmd.HelpTemplate()))
	})

	// Add Subcommands
	root.AddCommand(
		c.pingCmd(),
	)

	return root.Execute() //nolint:wrapcheck
}

func initRootCmd(rootCmd *cobra.Command) {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.traq-kinano-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		traq.MustPostMessage(traq.GpsTimesRasBot, "pong!!!")
	})
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".traq-kinano-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".traq-kinano-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
