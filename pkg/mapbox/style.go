package mapbox

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ListStyle struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Owner      string    `json:"owner"`
	Version    int32     `json:"version"`
	Center     []float64 `json:"center"`
	Zoom       float64   `json:"zoom"`
	Bearing    float64   `json:"bearing"`
	Pitch      float64   `json:"pitch"`
	Created    time.Time `json:"created"`
	Modified   time.Time `json:"modified"`
	Visibility string    `json:"visibility"`
}

func GetTableHeader() string {
	return "name, version"
}

func (row *ListStyle) AsTableRow() string {
	return "abc123, 2"
}

func GetStyles(accessToken string, username string) ([]ListStyle, error) {
	if username == "" {
		return nil, fmt.Errorf("Username missing, please provide with --username/-u")
	}

	if accessToken == "" {
		return nil, fmt.Errorf("Access token is missing, please provide with --access-token or env: MAPBOX_ACCESS_TOKEN")
	}

	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v", username)

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("failed to fetch styles for user: %v", err)
	}

	if res.StatusCode > 200 {
		return nil, fmt.Errorf("failed to fetch styles: %v", GetErrorMessage(res.StatusCode, res.Payload))
	}

	var styles []ListStyle
	if err := json.Unmarshal(res.Payload, &styles); err != nil {
		log.Fatalf("failed to parse styles: %v", err)
	}

	return styles, nil
}

func GetStyle(accessToken string, username string, styleId string) ([]ListStyle, error) {
	if username == "" {
		return nil, fmt.Errorf("Username missing, please provide with --username/-u")
	}

	if accessToken == "" {
		return nil, fmt.Errorf("Access token is missing, please provide with --access-token or env: MAPBOX_ACCESS_TOKEN")
	}

	client := GetDefaultClient(accessToken)

	endpoint := fmt.Sprintf("/styles/v1/%v", username)

	res, err := client.Get(endpoint, nil, nil)
	if err != nil {
		log.Fatalf("Failed to fetch styles for user: %v", err)
	}

	var styles []ListStyle
	if err := json.Unmarshal(res.Payload, &styles); err != nil {
		log.Fatalf("Failed to parse styles: %v", err)
	}

	out, err := json.Marshal(styles)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))

	return styles, nil
}
