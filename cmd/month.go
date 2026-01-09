package cmd

import (
	"github.com/spf13/cobra"
)

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Visualize monthly stats",
	Long:  "Visualize monthly stats",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("month called %t \n", shouldGetLast)
		//period := domain.CreateMonth(shouldGetLast)
		//
		//activitiesSummary, _ := statService.ListActivities(period)
		//
		//fmt.Println(period.StartDay, "à", period.EndDay)
		//for _, activity := range activitiesSummary {
		//	fmt.Printf("Activité %s (%d): %d km, %d secondes, %dm de dénivelé positif\n",
		//		activity.SportType.String(),
		//		activity.Count,
		//		activity.TotalDistance,
		//		activity.TotalDuration,
		//		activity.TotalAscent,
		//	)
		//}
	},
}

func init() {
	//monthCmd.PersistentFlags().BoolVarP(&shouldGetLast, "last", "l", false, "Get last monthly stats")
	rootCmd.AddCommand(monthCmd)
}
