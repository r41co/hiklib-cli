package cmd

import (
	"github.com/r41co/hiklib"
	"github.com/spf13/cobra"
)

var (
	filename string
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Get snapshot from camera.",
	Run: func(cmd *cobra.Command, args []string) {
		loglevel := -1
		if Verbose {
			loglevel = 0
		}
		u, dev := hiklib.HikLoginLog(camera, port, username, password, loglevel)
		hiklib.HikCaptureImage(u, dev.ByStartChan, filename)
		hiklib.HikLogout(u)
	},
}

func init() {
	snapshotCmd.Flags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	snapshotCmd.Flags().StringVarP(&password, "password", "p", "", "Password (required if username is set)")
	snapshotCmd.Flags().StringVarP(&camera, "cameraip", "c", "", "Camera ip address")
	snapshotCmd.MarkFlagsRequiredTogether("username", "password", "cameraip")

	snapshotCmd.MarkFlagRequired("username")
	snapshotCmd.MarkFlagRequired("password")
	snapshotCmd.MarkFlagRequired("cameraip")

	snapshotCmd.Flags().IntVarP(&port, "port", "z", 8000, "Camera port")
	snapshotCmd.LocalFlags().StringVarP(&filename, "filename", "f", "./cam.jpeg", "The filename where the snapshot will be saved")

	rootCmd.AddCommand(snapshotCmd)
}
