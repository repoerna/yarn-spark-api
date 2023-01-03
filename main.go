package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	resourceManagerAddress string
	yarnStates,
	yarnApplicationType []string
)

func init() {
	flag.StringVar(&resourceManagerAddress, "rma", "", "YARN Resource Manager Address")
	flag.StringArrayVar(&yarnStates, "ys", []string{}, "ys - YARN States. To filter YARN API result using states")
	flag.StringArrayVar(&yarnApplicationType, "yat", []string{}, "yat - YARN Application Type. To filter YARN API result using application type")
}

func main() {
	flag.Args()

	flag.Parse()
	if resourceManagerAddress == "" {
		fmt.Fprintln(os.Stderr, "please input YARN Resource Manager Address using flag --rma")
		flag.PrintDefaults()
		os.Exit(2)
	}

	fmt.Println(yarnStates)
	fmt.Println(yarnApplicationType)
	fmt.Println(resourceManagerAddress)

	// app := cluster.Application{
	// 	QueryParams: &cluster.QueryParams{
	// 		States: []string{
	// 			"RUNNING", "STOPPED",
	// 		},
	// 	},
	// }

	// qp := app.BuildQueryParams()
	// fmt.Println(qp)
	// fmt.Println(app.BuildURL(qp))
}
