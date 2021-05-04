package ds

import (
	"errors"
	"github.com/buger/jsonparser"
	"time"
)

type Song struct {
	Id       int64
	Name     string
	Duration time.Duration
	Artists  []Artist
	Album
}

// NewSongFromDailySongsJson 从每日推荐歌曲获取数据
func NewSongFromDailySongsJson(json []byte) (Song, error) {
	var song Song
	if len(json) == 0 {
		return song, errors.New("json is empty")
	}

	id, err := jsonparser.GetInt(json, "id")
	if err != nil {
		return song, err
	}
	song.Id = id

	if name, err := jsonparser.GetString(json, "name"); err == nil {
		song.Name = name
	}
	if duration, err := jsonparser.GetInt(json, "dt"); err == nil {
		song.Duration = time.Millisecond * time.Duration(duration)
	}

	album, err := NewAlbumFromJson(json)
	if err == nil {
		song.Album = album
	}

	_, _ = jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		artist, err := NewArtist(value)

		if err == nil {
			song.Artists = append(song.Artists, artist)
		}
	}, "ar")

	return song, nil
}

// NewSongFromPlaylistSongsJson 从歌单获取数据
func NewSongFromPlaylistSongsJson(json []byte) (Song, error) {
	var song Song
	if len(json) == 0 {
		return song, errors.New("json is empty")
	}

	id, err := jsonparser.GetInt(json, "id")
	if err != nil {
		return song, err
	}
	song.Id = id

	if name, err := jsonparser.GetString(json, "name"); err == nil {
		song.Name = name
	}
	if duration, err := jsonparser.GetInt(json, "dt"); err == nil {
		song.Duration = time.Millisecond * time.Duration(duration)
	}
	if alId, err := jsonparser.GetInt(json, "al", "id"); err == nil {
		song.Album.Id = alId
	}
	if alName, err := jsonparser.GetString(json, "al", "name"); err == nil {
		song.Album.Name = alName
	}
	if alPic, err := jsonparser.GetString(json, "al", "picUrl"); err == nil {
		song.Album.PicUrl = alPic
	}

	_, _ = jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		artist, err := NewArtist(value)

		if err == nil {
			song.Artists = append(song.Artists, artist)
		}
	}, "ar")

	return song, nil
}

// NewSongFromFmJson 从私人FM获取数据
func NewSongFromFmJson(json []byte) (Song, error) {
	var song Song
	if len(json) == 0 {
		return song, errors.New("json is empty")
	}

	id, err := jsonparser.GetInt(json, "id")
	if err != nil {
		return song, err
	}
	song.Id = id

	if name, err := jsonparser.GetString(json, "name"); err == nil {
		song.Name = name
	}
	if duration, err := jsonparser.GetInt(json, "duration"); err == nil {
		song.Duration = time.Millisecond * time.Duration(duration)
	}
	if alId, err := jsonparser.GetInt(json, "album", "id"); err == nil {
		song.Album.Id = alId
	}
	if alName, err := jsonparser.GetString(json, "album", "name"); err == nil {
		song.Album.Name = alName
	}
	if alPic, err := jsonparser.GetString(json, "album", "picUrl"); err == nil {
		song.Album.PicUrl = alPic
	}

	_, _ = jsonparser.ArrayEach(json, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		artist, err := NewArtist(value)

		if err == nil {
			song.Artists = append(song.Artists, artist)
		}
	}, "artists")

	return song, nil
}