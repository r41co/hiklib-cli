package cmd

import (
	"log"

	"github.com/r41co/hiklib"
	"github.com/spf13/cobra"
)

var (
	videoname      string
	filename_video string
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download video from camera.",
	Run: func(cmd *cobra.Command, args []string) {
		loglevel := -1
		if Verbose {
			loglevel = 0
		}
		u, _ := hiklib.HikLoginLog(camera, port, username, password, loglevel)
		log.Println(filename)
		if filename_video == "" {
			filename_video = videoname + ".ts"
		}

		hiklib.HikSaveFile(u, videoname, filename_video)
		hiklib.HikLogout(u)
	},
}

func init() {
	downloadCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	downloadCmd.Flags().StringVarP(&password, "password", "p", "", "Password (required if username is set)")
	downloadCmd.Flags().StringVarP(&camera, "cameraip", "c", "", "Camera ip address")
	downloadCmd.MarkFlagsRequiredTogether("username", "password", "cameraip")

	downloadCmd.MarkFlagRequired("username")
	downloadCmd.MarkFlagRequired("password")
	downloadCmd.MarkFlagRequired("cameraip")

	downloadCmd.Flags().IntVarP(&port, "port", "z", 8000, "Camera port")

	downloadCmd.Flags().StringVarP(&videoname, "name", "n", "", "Name of video on camera")
	downloadCmd.LocalFlags().StringVarP(&filename_video, "filename", "f", "", "The filename where video will be saved")
	downloadCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(downloadCmd)
}
