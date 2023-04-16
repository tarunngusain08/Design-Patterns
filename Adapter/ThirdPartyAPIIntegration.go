package Adapter

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// Third-party API response struct
type ThirdPartyResponse struct {
	Data string `json:"data"`
}

// Target interface
type DataFetcher interface {
	FetchData() string
}

// Adaptee struct
type ThirdPartyAPI struct{}

// Adaptee implementation
func (api *ThirdPartyAPI) request() (*ThirdPartyResponse, error) {
	resp, err := http.Get("https://api.example.com/data")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data ThirdPartyResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Adapter
type ThirdPartyAdapter struct {
	thirdPartyAPI *ThirdPartyAPI
}

func (a *ThirdPartyAdapter) FetchData() string {
	resp, err := a.thirdPartyAPI.request()
	if err != nil {
		return ""
	}

	return resp.Data
}

// Client code
func GetData(f DataFetcher) string {
	return f.FetchData()
}

func main() {
	// Using the ThirdPartyAPI directly
	thirdPartyAPI := &ThirdPartyAPI{}
	response, err := thirdPartyAPI.request()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Data from ThirdPartyAPI:", response.Data)

	// Using the DataFetcher interface with the ThirdPartyAdapter
	adapter := &ThirdPartyAdapter{thirdPartyAPI: thirdPartyAPI}
	data := GetData(adapter)
	fmt.Println("Data from ThirdPartyAdapter:", data)
}
