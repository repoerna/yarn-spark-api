package cluster

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
	FinalStatus,
	Queue,
	Limit,
	StartedTimeBegin,
	FinishedTimeBegin,
	FinishedTimeEnd,
	Name,
	DeSelect string
	States,
	ApplicationTypes,
	ApplicationTags []string
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

func (a *Application) BuildQueryParams() map[string]string {
	queryParams := make(map[string]string)

	if a.QueryParams != nil {
		if len(a.QueryParams.States) > 0 {
			var statesStr string
			for _, state := range a.QueryParams.States {
				statesStr += state + ","
			}
			queryParams["states"] = statesStr[:len(statesStr)-1]
		}
	}

	return queryParams
}
