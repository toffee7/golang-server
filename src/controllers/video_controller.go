package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../data"
)

// GetVideos specifies handler function to fetch videos
func GetVideos(w http.ResponseWriter, req *http.Request) {
	videos := make([]data.Video, 0)
	// List<data.Video> videos = new ArrayList<>();
	video := data.Video{}

	counts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	for _, num := range counts {
		video = data.Video{
			Name: "video" + strconv.Itoa(num),
			Link: "https://www.youtube.com/embed/BnYq7JapeDA",
		}
		videos = append(videos, video)
	}

	fmt.Println(videos)

	bytes, _ := json.Marshal(videos)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}
