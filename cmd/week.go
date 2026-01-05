package cmd

import (
	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Visualize weekly stats",
	Long:  "Visualize weekly stats",
	Run: func(cmd *cobra.Command, args []string) {
		stravaClient.ListActivities(1767552000, 1766956800)
	},
}

func init() {
	weekCmd.PersistentFlags().BoolVarP(&shouldGetLast, "last", "l", false, "Get last weekly stats")
	rootCmd.AddCommand(weekCmd)
}
