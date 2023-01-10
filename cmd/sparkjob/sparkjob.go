package sparkjob

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	resourceManagerServer string
	states                []string
)

func init() {
	CmdSparkJobList.Flags().StringArrayVarP(&states, "state", "s", nil, "filter jobs in the specific state")
	AllSparkJobList.Flags().StringVarP(&resourceManagerServer, "rmserver", "r", "localhost:8088", "YARN Resource Manager server")
}

var CmdSparkJobList = &cobra.Command{
	Use:   "sparkjob",
	Short: "Return all of spak job detail",
	Long: `
	use all command to get all YARN Application with "SPARK" application type 
	and get all of its SPARK job detail with given filter
	`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(resourceManagerServer)
	},
}

var AllSparkJobList = &cobra.Command{
	Use:   "all",
	Short: "Return all of spak job detail in all YARN applications",
	Long: `
	Will get all YARN Application with "SPARK" application type 
	and get all of its SPARK job detail with given filter
	`,
	// Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(resourceManagerServer)
	},
}
