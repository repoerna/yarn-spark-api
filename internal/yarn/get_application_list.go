package yarn

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetApplicationList(server string) (*YarnApplicationList, error) {
	var y YarnApplicationList
	var url = fmt.Sprintf(server + ApplicationURL)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&y); err != nil {
		return nil, err
	}
	return &y, nil
}
