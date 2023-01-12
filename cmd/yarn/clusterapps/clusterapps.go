package yarn

import (
	"fmt"
	"go-yarn-spark-api/internal/yarn"
	"go-yarn-spark-api/pkg"
	"log"

	"github.com/spf13/cobra"
)

var (
	user,
	finalStatus,
	queue,
	limit,
	startedTimeBegin,
	finishedTimeBegin,
	finishedTimeEnd,
	name,
	deSelect,
	states,
	applicationTypes,
	applicationTags string
)

var err error
var Apps *yarn.YarnApplicationList

func init() {
	ClusterApps.PersistentFlags().StringVarP(&user, "user", "u", "", "YARN API - applications matching the given application states, specified as a comma-separated list")
	ClusterApps.PersistentFlags().StringVar(&finalStatus, "finalStatus", "", "YARN API - the final status of the application - reported by the application itself")
	ClusterApps.PersistentFlags().StringVarP(&queue, "queue", "q", "", "YARN API - unfinished applications that are currently in this queue")
	ClusterApps.PersistentFlags().StringVarP(&limit, "limit", "l", "", "YARN API - total number of app objects to be returned")
	ClusterApps.PersistentFlags().StringVar(&startedTimeBegin, "startedTimeBegin", "", "YARN API - applications with start time beginning with this time, specified in ms since epoch")
	ClusterApps.PersistentFlags().StringVar(&finishedTimeBegin, "finishedTimeBegin", "", "YARN API - applications with finish time beginning with this time, specified in ms since epoch")
	ClusterApps.PersistentFlags().StringVar(&finishedTimeEnd, "finishedTimeEnd", "", "YARN API - applications with finish time ending with this time, specified in ms since epoch")
	ClusterApps.PersistentFlags().StringVarP(&name, "name", "n", "", "YARN API - name of the application")
	ClusterApps.PersistentFlags().StringVar(&deSelect, "deSelect", "", "YARN API - a generic fields which will be skipped in the result.")
	ClusterApps.PersistentFlags().StringVarP(&states, "states", "s", "", "YARN API - applications matching the given application states, specified as a comma-separated list")
	ClusterApps.PersistentFlags().StringVar(&applicationTypes, "applicationTypes", "", "YARN API - applications matching the given application types, specified as a comma-separated list")
	ClusterApps.PersistentFlags().StringVar(&applicationTags, "applicationTags", "", "YARN API - applications matching any of the given application tags, specified as a comma-separated list")

}

// func Cmd() *cobra.Command {
// 	ClusterApps.AddCommand()
// }

var ClusterApps = &cobra.Command{
	Use:   "clusterapps [resource manager server]",
	Short: "Obtain a collection of resources, each of which represents an application",
	Long: `
	With the Applications API, you can obtain a collection of resources, each of which represents an application. 
	When you run a GET operation on this resource, you obtain a collection of Application Objects.
	`,
	Args: cobra.MinimumNArgs(1),
	// TraverseChildren: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		qp := yarn.GetApplicationListQueryParams{
			User:              user,
			FinalStatus:       finalStatus,
			Queue:             queue,
			Limit:             limit,
			StartedTimeBegin:  startedTimeBegin,
			FinishedTimeBegin: finishedTimeBegin,
			FinishedTimeEnd:   finishedTimeEnd,
			Name:              name,
			DeSelect:          deSelect,
			States:            states,
			ApplicationTypes:  applicationTypes,
			ApplicationTags:   applicationTags,
		}

		Apps, err = qp.GetApplicationList(args[0])
		if err != nil {
			log.Fatal(err)
		}

	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(pkg.PrettyResult(Apps))
	},
}
