package main

import "go-yarn-spark-api/cmd"

func main() {

	// server := "http://bdgbnbcldnn01.intra.excelcom.co.id:8088"

	// yarnQP := yarn.GetApplicationListQueryParams{}
	// yarnQP.User = "ampid"
	// yarnQP.States = "RUNNING"
	// yarnQP.ApplicationTypes = "SPARK"

	// yarnAppList, err := yarnQP.GetApplicationList(server)
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(yarnAppList)

	// for _, app := range yarnAppList.Apps.App {
	// 	fmt.Println(app)
	// 	// go func(app yarn.YarnApplication) {
	// 	res, _ := spark.GetApplicationJobList(app)
	// 	fmt.Println("dfsd")
	// 	fmt.Println(res)
	// 	// }(app)

	// }

	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// wg.Wait()
	cmd.Execute()
}
