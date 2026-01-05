package cmd

import (
	"fmt"
	"gova/internals/client"

	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Visualize weekly stats",
	Long:  "Visualize weekly stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("week called %t \n", shouldGetLast)
		client.GetCurrentAthlete()
	},
}

func init() {
	weekCmd.PersistentFlags().BoolVarP(&shouldGetLast, "last", "l", false, "Get last weekly stats")
	rootCmd.AddCommand(weekCmd)
}
