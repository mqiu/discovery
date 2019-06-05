package streams

import (
	"net/http"
	"encoding/json"
)

func fetchAds(idStr string) (*Ads, error) {
	resp, err := http.Get("https://coding-challenge.dsc.tv/v1/ads/"+idStr)
    if err != nil {
        panic(err)
    }
	defer resp.Body.Close()
	var record Ads
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		return nil, err
	}
	
	return &record, nil
}