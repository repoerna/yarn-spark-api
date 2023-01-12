package cmd

import (
	"fmt"
	"go-yarn-spark-api/internal/spark"
	"go-yarn-spark-api/internal/yarn"
	"go-yarn-spark-api/pkg"

	"github.com/spf13/cobra"
)

var status string

func init() {
	SparkApplicationJobs.Flags().StringVar(&status, "status", "", "SPARK API - filter jobs in the specific state [running|succeeded|failed|unknown]")
}

var SparkApplicationJobs = &cobra.Command{
	Use:   "spark-application-job [yarn resource manager server]",
	Short: "Return Spark job/s detail managed by YARN",
	Long: `
	Command to get Spark job detail by YARN applications
	`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		getYarnClusterApps(cmd, args)

		yarnApps := cmd.Context().Value(YARN_APPS)

		for _, app := range yarnApps.(*yarn.YarnApplicationList).Apps.App {
			res, err := spark.GetApplicationJobList(app)
			fmt.Println(err)

			fmt.Println(pkg.PrettyResult(res))

		}
	},
}
