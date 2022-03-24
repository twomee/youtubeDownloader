package utils

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
	"strings"
)

func requestManager(content string, decision string) {
	if decision == "playlist" {
		getBulkVideo(content)
	} else if decision == "song" {
		getSingleVideo(content)
	}
}

func getBulkVideo(url string) *youtube.Playlist {
	playlistID := url
	client := youtube.Client{}

	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		panic(err)
	}

	/* ----- Enumerating playlist videos ----- */
	header := fmt.Sprintf("Playlist %s by %s", playlist.Title, playlist.Author)
	println(header)
	println(strings.Repeat("=", len(header)) + "\n")

	for k, v := range playlist.Videos {
		fmt.Printf("(%d) %s - '%s'\n", k+1, v.Author, v.Title)
	}

	return playlist
}
func getSingleVideo(url string) {
	client := youtube.Client{Debug: true}
	video, err := client.GetVideo(url)
	if err != nil {
		panic(err)
	}

	// Typically youtube only provides separate streams for video and audio.
	// If you want audio and video combined, take a look a the downloader package.
	format := video.Formats.FindByQuality("medium")
	reader, _, err := client.GetStream(video, format)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		panic(err)
	}

	reader.Close()
}
