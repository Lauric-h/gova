package cmd

import (
	"fmt"
	"gova/internals/client"

	"github.com/spf13/cobra"
)

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Visualize monthly stats",
	Long:  "Visualize monthly stats",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("month called %t \n", shouldGetLast)
		client.GetActivityList()
	},
}

func init() {
	monthCmd.PersistentFlags().BoolVarP(&shouldGetLast, "last", "l", false, "Get last monthly stats")
	rootCmd.AddCommand(monthCmd)
}
