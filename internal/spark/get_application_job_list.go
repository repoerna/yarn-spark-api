package spark

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetApplicationJobList(serverURL, appID string) (*Job, error) {
	var job Job
	url := fmt.Sprintf(serverURL+JobListURL, appID)

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
