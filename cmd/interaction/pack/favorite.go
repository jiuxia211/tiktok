package pack

import (
	"strconv"

	"github.com/ozline/tiktok/cmd/video/kitex_gen/video"
	"github.com/ozline/tiktok/kitex_gen/interaction"
	"github.com/ozline/tiktok/kitex_gen/user"
)

// build *user.User
func BuildUser(data *video.User) *user.User {
	if data == nil {
		return nil
	}

	totalFavorited, _ := strconv.ParseInt(data.TotalFavorited, 10, 64)
	return &user.User{
		Id:              data.Id,
		Name:            data.Name,
		Avatar:          data.Avatar,
		FollowCount:     data.FollowCount,
		FollowerCount:   data.FollowerCount,
		IsFollow:        data.IsFollow,
		BackgroundImage: data.BackgroundImage,
		Signature:       data.Signature,
		WorkCount:       data.WorkCount,
		FavoritedCount:  data.FavoriteCount,
		TotalFavorited:  totalFavorited,
	}
}

// convert video.Video To interaction.Video
func BuildVideo(data *video.Video) *interaction.Video {
	if data == nil {
		return nil
	}

	return &interaction.Video{
		Id:            data.Id,
		Author:        BuildUser(data.Anthor),
		PlayUrl:       data.PlayUrl,
		CoverUrl:      data.CoverUrl,
		FavoriteCount: data.FavoriteCount,
		CommentCount:  data.CommentCount,
		IsFavorite:    data.IsFavorite,
		Title:         data.Title,
	}
}

// build []*interaction.Video
func BuildVideos(datas []*video.Video) []*interaction.Video {
	if datas == nil {
		return nil
	}

	videos := make([]*interaction.Video, 5, 10)

	for _, item := range datas {
		videos = append(videos, BuildVideo(item))
	}
	return videos
}