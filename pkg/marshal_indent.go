package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrettyResult(v any) string {
	resJSON, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	return fmt.Sprintf("%s\n", string(resJSON))
}
