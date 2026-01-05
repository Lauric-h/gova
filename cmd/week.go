package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Visualize weekly stats",
	Long:  "Visualize weekly stats",
	Run: func(cmd *cobra.Command, args []string) {
		activitiesSummary, _ := statService.ListActivities(shouldGetLast)

		for _, activity := range activitiesSummary {
			fmt.Printf("Activité %s (%d): %d km, %d secondes, %dm de dénivelé positif\n",
				activity.SportType.String(),
				activity.Count,
				activity.TotalDistance,
				activity.TotalDuration,
				activity.TotalAscent,
			)
		}
	},
}

func init() {
	weekCmd.PersistentFlags().BoolVarP(&shouldGetLast, "last", "l", false, "Get last weekly stats")
	rootCmd.AddCommand(weekCmd)
}
