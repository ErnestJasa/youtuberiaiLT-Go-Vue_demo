package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func GetChannels(w http.ResponseWriter, r *http.Request) {
	var includeCategories = RemoveEmptyStrings(r.URL.Query()["includeCategory"])
	var excludeCategories = RemoveEmptyStrings(r.URL.Query()["excludeCategory"])

	var channelHandle = strings.ToLower(r.URL.Query().Get("channelHandle"))
	result := Channels
	if channelHandle != "" {

		result = Filter(result, func(c Channel) bool {
			return strings.Contains(strings.ToLower(c.Title), channelHandle) ||
				strings.Contains(strings.ToLower(c.CustomUrl), channelHandle)
		})
	}
	if len(includeCategories) > 0 {
		result = Filter(result, func(c Channel) bool {
			return All(includeCategories, func(category string) bool {
				return Any(c.Categories, func(channelCategory string) bool {
					return strings.Contains(strings.ToLower(channelCategory), strings.ToLower(category))
				})
			})
		})
	}

	if len(excludeCategories) > 0 {
		result = Filter(result, func(c Channel) bool {
			return !Any(excludeCategories, func(category string) bool {
				return Any(c.Categories, func(channelCategory string) bool {
					return strings.Contains(strings.ToLower(channelCategory), strings.ToLower(category))
				})
			})
		})
	}
	w.Header().Set("Content-Type", "application/json")
	if result == nil {
		result = []Channel{}
	}
	json.NewEncoder(w).Encode(result)
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Categories)
}
