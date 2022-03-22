package service

import entitiy "github.com/meharifitih/go-gin/entity"

type VideoService interface {
	FindAll() []entitiy.Video
	Save(entitiy.Video) entitiy.Video
}

type videoService struct {
	videos []entitiy.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) FindAll() []entitiy.Video {
	return service.videos
}
func (service *videoService) Save(video entitiy.Video) entitiy.Video {
	service.videos = append(service.videos, video)
	return video
}
