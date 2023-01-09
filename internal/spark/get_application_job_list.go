package spark

import (
	"encoding/json"
	"fmt"
	"go-yarn-spark-api/internal/yarn"
	"net/http"
)

func GetApplicationJobList(yarnApp yarn.YarnApplication) (*Job, error) {
	var job Job
	url := fmt.Sprintf(yarnApp.TrackingURL + fmt.Sprintf(JobListURL, yarnApp.ID))

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&job); err != nil {
		return nil, err
	}

	return &job, nil
}
