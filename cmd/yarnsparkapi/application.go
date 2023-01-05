package yarnsparkapi

var URL = "/ws/v1/cluster/apps"

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
type QueryParams struct {
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

type Application struct {
	QueryParams *QueryParams
}

func (a *Application) BuildURL(queryParams map[string]string) string {
	url := URL + "?"

	if len(queryParams) > 0 {
		for k, v := range queryParams {
			url += k + "=" + v + "&"
		}
	}

	return url[:len(url)-1]

}
