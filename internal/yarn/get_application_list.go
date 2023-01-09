package yarn

import (
	"encoding/json"
	"fmt"
	"go-yarn-spark-api/pkg"
	"net/http"
)

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
	return &y, nil
}
