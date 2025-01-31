package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"
)

func GetChannels(w http.ResponseWriter, r *http.Request) {
	var includeCategories = RemoveEmptyStrings(r.URL.Query()["includeCategory"])
	var excludeCategories = RemoveEmptyStrings(r.URL.Query()["excludeCategory"])
	var sortOrder = r.URL.Query().Get("sortOrder")

	var channelHandle = strings.ToLower(r.URL.Query().Get("channelHandle"))
	result := make([]Channel, len(Channels))
	copy(result, Channels)
	if channelHandle != "" {

		result = Filter(result, func(c Channel) bool {
			return strings.Contains(strings.ToLower(c.Title), channelHandle) ||
				strings.Contains(strings.ToLower(c.CustomUrl), channelHandle)
		})
	}
	if len(includeCategories) > 0 {
		result = Filter(result, func(c Channel) bool {
			return AllInclude(includeCategories, func(category string) bool {
				return AnyOneIncludes(c.Categories, func(channelCategory string) bool {
					return strings.Contains(strings.ToLower(channelCategory), strings.ToLower(category))
				})
			})
		})
	}

	if len(excludeCategories) > 0 {
		result = Filter(result, func(c Channel) bool {
			return !AnyOneIncludes(excludeCategories, func(category string) bool {
				return AnyOneIncludes(c.Categories, func(channelCategory string) bool {
					return strings.Contains(strings.ToLower(channelCategory), strings.ToLower(category))
				})
			})
		})
	}
	switch sortOrder {
	case "ByHandleAscending", "1":
		sort.Slice(result, func(i, j int) bool {
			return AscendingString(result[i].Title, result[j].Title)
		})
	case "ByHandleDescending", "2":
		sort.Slice(result, func(i, j int) bool {
			return DescendingString(result[i].Title, result[j].Title)
		})
	case "BySubCountAscending", "3":
		sort.Slice(result, func(i, j int) bool {
			return AscendingNumber(int64(result[i].SubscriberCount), int64(result[j].SubscriberCount))
		})
	case "BySubCountDescending", "4":
		sort.Slice(result, func(i, j int) bool {
			return DescendingNumber(int64(result[i].SubscriberCount), int64(result[j].SubscriberCount))
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
