package youtube_api_go

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/kkdai/youtube/v2"
)

type Youtube struct {
	Name string
}

func (yt *Youtube) Download_Video(url string, output_name string) {
	videoID := url
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels() // only get videos with audio
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(output_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
func (yt *Youtube) Download_Playlist(url string, output_name string) {
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

	/* ----- Downloading the 1st video ----- */
	entry := playlist.Videos[0]
	video, err := client.VideoFromPlaylistEntry(entry)
	if err != nil {
		panic(err)
	}
	// Now it's fully loaded.

	fmt.Printf("Downloading %s by '%s'!\n", video.Title, video.Author)

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create(output_name)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	_, err = io.Copy(file, stream)

	if err != nil {
		panic(err)
	}

	println("Downloaded /video.mp4")
}
