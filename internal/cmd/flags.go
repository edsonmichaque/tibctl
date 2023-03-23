package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func withGlobalFlags() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.PersistentFlags().String(optSecret, "", "Access token")
		cmd.PersistentFlags().String(optBaseURL, "", "Base URL")
		cmd.PersistentFlags().StringVar(&profile, optProfile, defaultProfile, "Profile")
		cmd.PersistentFlags().StringVarP(&configFile, optConfigFile, "c", "", "Configuration file")

		viper.SetEnvPrefix(envPrefix)
	}
}

// withOutputFlag adds output flag to command
func withOutputFlag(value string) cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP(optOutput, "o", value, "Output format")
	}
}

// withQueryFlag adds query flag to command
func withQueryFlag() cmdOption {
	return func(cmd *cobra.Command) {
		cmd.Flags().StringP(optQuery, "q", "", "Query")
	}
}

func withOptions(opts *Options) cmdOption {
	return func(cmd *cobra.Command) {
		cmd.SetOutput(opts.Stdout)
		cmd.SetIn(opts.Stdin)
		cmd.SetErr(opts.Stderr)
	}
}

// withSubcommand adds subcommand to command
func withSubcommand(children ...*Cmd) cmdOption {
	return func(cmd *cobra.Command) {
		for _, c := range children {
			cmd.AddCommand(c.Command)
		}
	}
}
