package yarn

import (
	"fmt"
	"go-yarn-spark-api/internal/yarn"
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

func init() {
	ClusterApps.Flags().StringVarP(&user, "user", "u", "", "applications matching the given application states, specified as a comma-separated list")
	ClusterApps.Flags().StringVar(&finalStatus, "finalStatus", "", "the final status of the application - reported by the application itself")
	ClusterApps.Flags().StringVarP(&queue, "queue", "q", "", "user name")
	ClusterApps.Flags().StringVarP(&limit, "limit", "l", "", "total number of app objects to be returned")
	ClusterApps.Flags().StringVar(&startedTimeBegin, "startedTimeBegin", "", "applications with start time beginning with this time, specified in ms since epoch")
	ClusterApps.Flags().StringVar(&finishedTimeBegin, "finishedTimeBegin", "", "applications with finish time beginning with this time, specified in ms since epoch")
	ClusterApps.Flags().StringVar(&finishedTimeEnd, "finishedTimeEnd", "", "applications with finish time ending with this time, specified in ms since epoch")
	ClusterApps.Flags().StringVarP(&name, "name", "n", "", "name of the application")
	ClusterApps.Flags().StringVar(&deSelect, "deSelect", "", "a generic fields which will be skipped in the result.")
	ClusterApps.Flags().StringVarP(&states, "states", "s", "", "applications matching the given application states, specified as a comma-separated list")
	ClusterApps.Flags().StringVar(&applicationTypes, "applicationTypes", "", "applications matching the given application types, specified as a comma-separated list")
	ClusterApps.Flags().StringVar(&applicationTags, "applicationTags", "", "applications matching any of the given application tags, specified as a comma-separated list")
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
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(args)

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

		apps, err := qp.GetApplicationList(args[0])
		if err != nil {
			log.Println(err)
		}

		fmt.Println(apps)

		// for _, app := range apps.Apps.App {

		// }
	},
}
