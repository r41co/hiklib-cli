package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/r41co/hiklib"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get a list of saved videos from the camera today.",
	Run: func(cmd *cobra.Command, args []string) {
		loglevel := -1
		if Verbose {
			loglevel = 0
		}
		layout := "2006-01-02T15:04:05"
		st, s_error := time.Parse(layout, s_date)
		et, e_error := time.Parse(layout, e_date)
		u, dev := hiklib.HikLoginLog(camera, port, username, password, loglevel)
		err := 0
		lst := hiklib.MotionVideos{}
		if (s_error != nil) || (e_error != nil) {
			err, lst = hiklib.HikListVideo(u, dev.ByStartChan)
		} else {
			err, lst = hiklib.HikListVideo(u, dev.ByStartChan, st.Year(), int(st.Month()), st.Day(), st.Hour(), st.Minute(), st.Second(), et.Year(), int(et.Month()), et.Day(), et.Hour(), et.Minute(), et.Second())
		}
		log.Printf("Found %d video:\n", lst.Count)
		if err == 0 {
			for i, v := range lst.Videos {
				fmt.Printf("[%d-%.02d-%.02d %.02d:%.02d:%.02d - %d-%.02d-%.02d %.02d:%.02d:%.02d] #%d %s %s \n", v.From_year, v.From_month, v.From_day, v.From_hour, v.From_min, v.From_sec, v.To_year, v.To_month, v.To_day, v.To_hour, v.To_min, v.To_sec, i, v.Filename, humanize.Bytes(uint64(v.Size)))
			}
		}
		hiklib.HikLogout(u)
	},
}

func init() {
	listCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Username (required if password is set)")
	listCmd.Flags().StringVarP(&password, "password", "p", "", "Password (required if username is set)")
	listCmd.Flags().StringVarP(&camera, "cameraip", "c", "", "Camera ip address")
	listCmd.Flags().StringVarP(&s_date, "start", "s", "", "Starting date/time (YYYY-MM-DDTHH:II:SS)")
	listCmd.Flags().StringVarP(&e_date, "end", "e", "", "Ending date/time (YYYY-MM-DDTHH:II:SS)")
	listCmd.MarkFlagsRequiredTogether("username", "password", "cameraip")

	listCmd.MarkFlagRequired("username")
	listCmd.MarkFlagRequired("password")
	listCmd.MarkFlagRequired("cameraip")

	listCmd.Flags().IntVarP(&port, "port", "z", 8000, "Camera port")

	rootCmd.AddCommand(listCmd)
}
