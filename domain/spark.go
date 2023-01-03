package domain

type SparkJob struct {
	JobID               int    `json:"jobId"`
	Name                string `json:"name"`
	SubmissionTime      string `json:"submissionTime"`
	CompletionTime      string `json:"completionTime"`
	StageIds            []int  `json:"stageIds"`
	Status              string `json:"status"`
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
