package utils

import (
	"fmt"
	"github.com/kkdai/youtube/v2"
	"io"
	"os"
)

func downloadManager(content interface{}, decision string, bulk bool) {
	convertToAudio := false
	if decision == "audio" {
		convertToAudio = true
	}
	if bulk {
		myContent := content.(*youtube.Playlist)
		bulkDownloader(myContent, convertToAudio)
	} else {
		myContent := content.(*youtube.PlaylistEntry)
		downloader(myContent, convertToAudio)
	}

}

func bulkDownloader(playlist *youtube.Playlist, convertToAudio bool) {
	for _, videoFromPlaylist := range playlist.Videos {
		downloader(videoFromPlaylist, convertToAudio)
	}
}

func downloader(videoFromPlaylist *youtube.PlaylistEntry, convertToAudio bool) {
	video, err := youtube.Client.VideoFromPlaylistEntry(&videoFromPlaylist)
	// Now it's fully loaded.

	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := youtube.Client.GetStream(&video, &video.Formats[0])
	if err != nil {
		panic(err)
	}
	file, err := os.Create(video.Title + ".mp4")

	//if convertToAudio {
	//	file, err := os.Create(video.Title + ".mp3")
	//} else {
	//}

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = io.Copy(file, stream)

	if err != nil {
		panic(err)
	}
}
