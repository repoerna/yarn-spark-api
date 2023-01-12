package cmd

import (
	"context"
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

func init() {
	YarnClusterApps.PersistentFlags().StringVarP(&user, "user", "u", "", "YARN API - applications matching the given application states, specified as a comma-separated list")
	YarnClusterApps.PersistentFlags().StringVar(&finalStatus, "finalStatus", "", "YARN API - the final status of the application - reported by the application itself")
	YarnClusterApps.PersistentFlags().StringVarP(&queue, "queue", "q", "", "YARN API - unfinished applications that are currently in this queue")
	YarnClusterApps.PersistentFlags().StringVarP(&limit, "limit", "l", "", "YARN API - total number of app objects to be returned")
	YarnClusterApps.PersistentFlags().StringVar(&startedTimeBegin, "startedTimeBegin", "", "YARN API - applications with start time beginning with this time, specified in ms since epoch")
	YarnClusterApps.PersistentFlags().StringVar(&finishedTimeBegin, "finishedTimeBegin", "", "YARN API - applications with finish time beginning with this time, specified in ms since epoch")
	YarnClusterApps.PersistentFlags().StringVar(&finishedTimeEnd, "finishedTimeEnd", "", "YARN API - applications with finish time ending with this time, specified in ms since epoch")
	YarnClusterApps.PersistentFlags().StringVarP(&name, "name", "n", "", "YARN API - name of the application")
	YarnClusterApps.PersistentFlags().StringVar(&deSelect, "deSelect", "", "YARN API - a generic fields which will be skipped in the result.")
	YarnClusterApps.PersistentFlags().StringVarP(&states, "states", "s", "", "YARN API - applications matching the given application states, specified as a comma-separated list")
	YarnClusterApps.PersistentFlags().StringVar(&applicationTypes, "applicationTypes", "", "YARN API - applications matching the given application types, specified as a comma-separated list")
	YarnClusterApps.PersistentFlags().StringVar(&applicationTags, "applicationTags", "", "YARN API - applications matching any of the given application tags, specified as a comma-separated list")
}

var YarnClusterApps = &cobra.Command{
	Use:   "yarn-cluster-apps [yarn resource manager server]",
	Short: "Obtain a collection of resources, each of which represents an application",
	Long: `
	With the Applications API, you can obtain a collection of resources, each of which represents an application. 
	When you run a GET operation on this resource, you obtain a collection of Application Objects.
	`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		getYarnClusterApps(cmd, args)

		ctx := cmd.Context()

		fmt.Println(pkg.PrettyResult(ctx.Value(YARN_APPS)))
	},
}

func getYarnClusterApps(cmd *cobra.Command, args []string) {

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

	res, err := qp.GetApplicationList(args[0])
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, YARN_APPS, res)

	cmd.SetContext(ctx)
}
