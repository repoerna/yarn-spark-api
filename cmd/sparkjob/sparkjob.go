package sparkjob

import (
	"fmt"
	yarn "go-yarn-spark-api/cmd/yarn/clusterapps"
	"go-yarn-spark-api/internal/spark"

	"github.com/spf13/cobra"
)

// var (
// 	resourceManagerServer,
// 	user,
// 	finalStatus,
// 	queue,
// 	limit,
// 	startedTimeBegin,
// 	finishedTimeBegin,
// 	finishedTimeEnd,
// 	name,
// 	deSelect,
// 	states,
// 	applicationTypes,
// 	applicationTags string
// )

// func Cmd() *cobra.Command {
// 	CmdSparkJobList.AddCommand(All)

// 	return CmdSparkJobList
// }

func init() {
	// CmdSparkJobList.PersistentFlags().StringVar(&states, "state", "", "filter jobs in the specific state [running|succeeded|failed|unknown]")
	// All.Flags().StringVarP(&resourceManagerServer, "rmserver", "r", "localhost:8088", "YARN Resource Manager server")
}

var CmdSparkJobList = &cobra.Command{
	Use:   "sparkjob [resource manager server]",
	Short: "Return Spark job/s detail managed by YARN",
	Long: `
	Command to get Spark job detail by YARN applications
	`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if yarn.Apps != nil {
			apps := yarn.Apps.Apps.App
			for _, app := range apps {
				job, _ := spark.GetApplicationJobList(app)
				fmt.Println(job)
			}
		}

	},
}

// var All = &cobra.Command{
// 	Use:   "all-yarn-app",
// 	Short: "Return all of spak job detail in all YARN applications",
// 	Long: `
// 	Will get all YARN Application with "SPARK" application type
// 	and get all of its SPARK job detail with given filter
// 	`,
// 	// Args: cobra.MinimumNArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println()
// 	},
// }
