package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../data"
)

// GetVideos specifies handler function to fetch videos
func GetVideos(w http.ResponseWriter, req *http.Request) {
	videos := make([]data.Video, 0)

	video := data.Video{
		Name: "video1",
		Link: "https://www.youtube.com/embed/BnYq7JapeDA",
	}

	videos = append(videos, video)
	fmt.Println(videos)

	bytes, _ := json.Marshal(videos)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
	return
}
