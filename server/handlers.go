package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetChannels(w http.ResponseWriter, r *http.Request) {
	var includeCategories = r.URL.Query()["includeCategory"]

	channels := Filter(Channels, func(c Channel) bool {
		return All(includeCategories, func(category string) bool {
			return Any(c.Categories, func(channelCategory string) bool {
				return strings.Contains(strings.ToLower(channelCategory), strings.ToLower(category))
			})
		})
	})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(channels)
}
