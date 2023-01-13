package spark

import (
	"encoding/json"
	"fmt"
	"go-yarn-spark-api/internal/yarn"
	"log"
	"net/http"
)

// func GetApplicationJobList(yarnApp yarn.YarnApplication) (*Job, error) {
// 	var job Job
// 	url := fmt.Sprintf(yarnApp.TrackingURL + fmt.Sprintf(JobListURL, yarnApp.ID))

// 	// TODO: delete
// 	fmt.Println(url)

// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	// TODO: change struct
// 	if err = json.NewDecoder(res.Body).Decode(&job); err != nil {
// 		return nil, err
// 	}

// 	// TODO: delete
// 	fmt.Println(job)

// 	return &job, nil
// }

func GetApplicationJobList(ch chan Summary, app yarn.YarnApplication) {

	var summary Summary

	var sj []Job

	url := fmt.Sprintf(app.TrackingURL + fmt.Sprintf(JobListURL, app.ID))

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		summary.IsError = true
		summary.LogError = err.Error()
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&sj); err != nil {
		log.Println(err)
		summary.IsError = true
		summary.LogError = err.Error()
	}

	summary.ID = app.ID
	summary.Status = app.State
	summary.TrackingURL = app.TrackingURL
	summary.TotalJob = len(sj)
	summary.SparkJobList = sj

	ch <- summary
}
