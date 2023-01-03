package spark

import (
	"errors"
	"strings"
)

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
