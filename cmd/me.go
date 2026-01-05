package cmd

import "github.com/spf13/cobra"

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Visualize profile info",
	Long:  "Visualize profile info",
	Run: func(cmd *cobra.Command, args []string) {
		stravaClient.GetCurrentAthlete()
	},
}

func init() {
	rootCmd.AddCommand(meCmd)
}
