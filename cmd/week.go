package cmd

import (
	"fmt"
	"gova/internal/domain"

	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Visualize weekly stats",
	Long:  "Visualize weekly stats",
	RunE: func(cmd *cobra.Command, args []string) error {
		if appCtx == nil || appCtx.StatService == nil {
			return fmt.Errorf("application not initialized, try to login again")
		}

		shouldGetLast, err := cmd.Flags().GetBool("last")
		if err != nil {
			return err
		}

		period := domain.CreateWeek(shouldGetLast)
		activitiesSummary, err := appCtx.StatService.ListActivities(period)
		if err != nil {
			return err
		}

		fmt.Println(period.StartDay, "à", period.EndDay)
		for _, activity := range activitiesSummary {
			fmt.Printf("Activité %s (%d): %d km, %d secondes, %dm de dénivelé positif\n",
				activity.SportType.String(),
				activity.Count,
				activity.TotalDistance,
				activity.TotalDuration,
				activity.TotalAscent,
			)
		}

		return nil
	},
}

func init() {
	weekCmd.Flags().BoolP("last", "l", false, "Get last week's stats")
	rootCmd.AddCommand(weekCmd)
}
