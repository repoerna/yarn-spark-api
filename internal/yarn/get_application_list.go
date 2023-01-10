package yarn

import (
	"encoding/json"
	"fmt"
	"go-yarn-spark-api/pkg"
	"net/http"
)

// QueryParams struct have fields to Cluster Application API query Parameter
// states - applications matching the given application states, specified as a comma-separated list.
// finalStatus - the final status of the application - reported by the application itself
// user - user name
// queue - unfinished applications that are currently in this queue
// limit - total number of app objects to be returned
// startedTimeBegin - applications with start time beginning with this time, specified in ms since epoch
// startedTimeEnd - applications with start time ending with this time, specified in ms since epoch
// finishedTimeBegin - applications with finish time beginning with this time, specified in ms since epoch
// finishedTimeEnd - applications with finish time ending with this time, specified in ms since epoch
// applicationTypes - applications matching the given application types, specified as a comma-separated list.
// applicationTags - applications matching any of the given application tags, specified as a comma-separated list.
// name - name of the application
// deSelects - a generic fields which will be skipped in the result.
type GetApplicationListQueryParams struct {
	User              string `json:"user"`
	FinalStatus       string `json:"finalStatus,omitempty"`
	Queue             string `json:"queue,omitempty"`
	Limit             string `json:"limit,omitempty"`
	StartedTimeBegin  string `json:"startedTimeBegin,omitempty"`
	FinishedTimeBegin string `json:"finishedTimeBegin,omitempty"`
	FinishedTimeEnd   string `json:"finishedTimeEnd,omitempty"`
	Name              string `json:"name,omitempty"`
	DeSelect          string `json:"deSelect,omitempty"`
	States            string `json:"states,omitempty"`
	ApplicationTypes  string `json:"applicationTypes,omitempty"`
	ApplicationTags   string `json:"applicationTags,omitempty"`
}

func (g *GetApplicationListQueryParams) GetApplicationList(server string) (*YarnApplicationList, error) {
	var y YarnApplicationList
	qp, err := pkg.BuildQueryParams(*g)
	fmt.Println(qp)
	if err != nil {
		return nil, err
	}
	var url = fmt.Sprintf(server + ApplicationURL + "?" + qp)
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&y); err != nil {
		return nil, err
	}
	fmt.Println(y)
	return &y, nil
}
