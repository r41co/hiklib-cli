package cmd

import (
	"log"

	"github.com/r41co/hiklib"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hikvision SDK",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("hiklib-cli [%s]\n", hiklib.HikVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
