package kvsctl

import (
	"os"

	"github.com/ishii1648/mem-kvsd/cmd/kvsctl/command"
	"github.com/spf13/cobra"
)

const (
	defaultHost       = "localhost"
	defaultListenPort = 10020
)

var (
	globalFlags = command.GlobalFlags{}
)

var (
	rootCmd = &cobra.Command{
		Use:        "kvsctl",
		Short:      "client for kvserver",
		SuggestFor: []string{"kvsctl"},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&globalFlags.Host, "host", defaultHost, "listen host")
	rootCmd.PersistentFlags().IntVar(&globalFlags.Port, "port", defaultListenPort, "listen port")
	rootCmd.PersistentFlags().BoolVar(&globalFlags.DebugMode, "debug", false, "debug mode")
}

func main() {
	rootCmd.AddCommand(
		command.NewPutCommand(),
	)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
