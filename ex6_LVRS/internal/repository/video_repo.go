package repository

import (
	"ex6_LVRS/internal/entity"
	"encoding/json"
	"os"
	"sync"
	"fmt"
)

var mutex sync.Mutex

const fileName = "store.json"

// this will process the video and change its status, return and save
func ProcessVideo(video entity.Video) (*entity.Video, error) {

}

// load videos
func LoadVideos() ([]*entity.Video, error) {
	var videos []*entity.Video
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []*entity.Video{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []*entity.Video{}, nil
	}

	err = json.Unmarshal(data, &videos)
	return videos, err
}