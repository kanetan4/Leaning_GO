package main

import (
	"context"
	videos "ex6_LVRS/kitex_gen/videos"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// AddVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) AddVideo(ctx context.Context, request *videos.AddVideoRequest) (resp *videos.Video, err error) {
	// TODO: Your code here...
	return
}

// ManualOverride implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) ManualOverride(ctx context.Context, request *videos.ManualOverrideRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}

// GetVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideo(ctx context.Context, request *videos.GetVideoRequest) (resp *videos.Video, err error) {
	// TODO: Your code here...
	return
}

// GetVideos implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideos(ctx context.Context, request *videos.GetVideoListRequest) (resp *videos.VideoList, err error) {
	// TODO: Your code here...
	return
}

// DeleteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteVideo(ctx context.Context, request *videos.DeleteVideoRequest) (resp bool, err error) {
	// TODO: Your code here...
	return
}
