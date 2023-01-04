package yarn

import (
	"errors"
	"strings"
)

type Entry struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type PreemptedResourceSecondsMap struct {
	Entry Entry `json:"entry"`
}

type ResourceSecondsMap struct {
	Entry Entry `json:"entry"`
}

type Timeouts struct {
	Timeout []TimeoutData `json:"timeout"`
}

type TimeoutData struct {
	Type                   string `json:"type"`
	ExpiryTime             string `json:"expiryTime"`
	RemainingTimeInSeconds int    `json:"remainingTimeInSeconds"`
}

type YarnApplication struct {
	ID                          string                      `json:"id"`
	User                        string                      `json:"user"`
	Name                        string                      `json:"name"`
	Queue                       string                      `json:"queue"`
	State                       string                      `json:"state"`
	FinalStatus                 string                      `json:"finalStatus"`
	Progress                    float64                     `json:"progress"`
	TrackingUI                  string                      `json:"trackingUI"`
	TrackingURL                 string                      `json:"trackingUrl"`
	Diagnostics                 string                      `json:"diagnostics"`
	ClusterID                   int64                       `json:"clusterId"`
	ApplicationType             string                      `json:"applicationType"`
	ApplicationTags             string                      `json:"applicationTags"`
	Priority                    int                         `json:"priority"`
	StartedTime                 int                         `json:"startedTime"`
	LaunchTime                  int                         `json:"launchTime"`
	FinishedTime                int                         `json:"finishedTime"`
	ElapsedTime                 int                         `json:"elapsedTime"`
	AllocatedMB                 int                         `json:"allocatedMB"`
	AllocatedVCores             int                         `json:"allocatedVCores"`
	ReservedMB                  int                         `json:"reservedMB"`
	ReservedVCores              int                         `json:"reservedVCores"`
	RunningContainers           int                         `json:"runningContainers"`
	MemorySeconds               int64                       `json:"memorySeconds"`
	VcoreSeconds                int                         `json:"vcoreSeconds"`
	QueueUsagePercentage        float64                     `json:"queueUsagePercentage"`
	ClusterUsagePercentage      float64                     `json:"clusterUsagePercentage"`
	ResourceSecondsMap          ResourceSecondsMap          `json:"resourceSecondsMap"`
	PreemptedResourceMB         int                         `json:"preemptedResourceMB"`
	PreemptedResourceVCores     int                         `json:"preemptedResourceVCores"`
	NumNonAMContainerPreempted  int                         `json:"numNonAMContainerPreempted"`
	NumAMContainerPreempted     int                         `json:"numAMContainerPreempted"`
	PreemptedMemorySeconds      int                         `json:"preemptedMemorySeconds"`
	PreemptedVcoreSeconds       int                         `json:"preemptedVcoreSeconds"`
	PreemptedResourceSecondsMap PreemptedResourceSecondsMap `json:"preemptedResourceSecondsMap"`
	UnmanagedApplication        bool                        `json:"unmanagedApplication"`
	AmNodeLabelExpression       string                      `json:"amNodeLabelExpression"`
	Timeouts                    Timeouts                    `json:"timeouts"`
}

type YarnApplicationList struct {
	Apps struct {
		App []YarnApplication `json:"app"`
	} `json:"apps"`
}

// to filter applications based on input comma-separated list of application states.
// The valid application state can be one of the following:
//  ALL, NEW, NEW_SAVING, SUBMITTED, ACCEPTED, RUNNING, FINISHED, FAILED, KILLED
type States int

const (
	ALL States = iota + 1
	NEW
	NEW_SAVING
	SUBMITTED
	ACCEPTED
	RUNNING
	FINISHED
	FAILED
	KILLED
)

func (s States) String() string {
	return [...]string{"ALL", "NEW", "NEW_SAVING", "SUBMITTED", "ACCEPTED", "RUNNING", "FINISHED", "FAILED", "KILLED"}[s-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s States) EnumIndex() int {
	return int(s)
}

func ParseSparkStatus(states string) (States, error) {

	states = strings.ToUpper(states)

	switch states {
	case "ALL":
		return ALL, nil
	case "NEW":
		return NEW, nil
	case "NEW_SAVING":
		return NEW_SAVING, nil
	case "SUBMITTED":
		return SUBMITTED, nil
	case "ACCEPTED":
		return ACCEPTED, nil
	case "RUNNING":
		return RUNNING, nil
	case "FINISHED":
		return FINISHED, nil
	case "FAILED":
		return FAILED, nil
	case "KILLED":
		return KILLED, nil
	default:
		return 0, errors.New("invalid state value")
	}
}
