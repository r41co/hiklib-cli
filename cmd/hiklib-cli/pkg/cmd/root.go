package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Verbose  bool
	username string
	password string
	camera   string
	s_date   string
	e_date   string
	port     int
)

var (
	rootCmd = &cobra.Command{
		Use:   "hiklib-cli",
		Short: "HIKLIB client app",
		Long:  `Console app for connect to hikvision camera.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

}

func Execute() error {
	return rootCmd.Execute()
}
