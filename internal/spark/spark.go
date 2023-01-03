package spark

import (
	"errors"
	"strings"
)

const JobListURL = "/api/v1/applications/%s/jobs"

type Job struct {
	JobID               int    `json:"jobId"`
	Name                string `json:"name"`
	SubmissionTime      string `json:"submissionTime"`
	CompletionTime      string `json:"completionTime"`
	StageIds            []int  `json:"stageIds"`
	Status              Status `json:"status"`
	NumTasks            int    `json:"numTasks"`
	NumActiveTasks      int    `json:"numActiveTasks"`
	NumCompletedTasks   int    `json:"numCompletedTasks"`
	NumSkippedTasks     int    `json:"numSkippedTasks"`
	NumFailedTasks      int    `json:"numFailedTasks"`
	NumKilledTasks      int    `json:"numKilledTasks"`
	NumCompletedIndices int    `json:"numCompletedIndices"`
	NumActiveStages     int    `json:"numActiveStages"`
	NumCompletedStages  int    `json:"numCompletedStages"`
	NumSkippedStages    int    `json:"numSkippedStages"`
	NumFailedStages     int    `json:"numFailedStages"`
	KilledTasksSummary  struct {
	} `json:"killedTasksSummary"`
}

// running|succeeded|failed|unknown
type Status int

const (
	Running Status = iota + 1
	Succeeded
	Failed
	Unknown
)

func (s Status) String() string {
	return [...]string{"running", "succeeded", "failed", "unknown"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Status) EnumIndex() int {
	return int(s)
}

func ParseSparkStatus(status string) (Status, error) {

	status = strings.ToLower(status)

	switch status {
	case "running":
		return Running, nil
	case "succeeded":
		return Succeeded, nil
	case "failed":
		return Failed, nil
	case "unknown":
		return Unknown, nil
	default:
		return 0, errors.New("invalid status value")
	}
}
