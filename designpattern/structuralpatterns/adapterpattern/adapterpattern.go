package main

import "fmt"

/*
Let's create an example where we have an interface for MediaPlayer
and another for AdvancedMediaPlayer. The MediaPlayer interface can play
MP3 files, while the AdvancedMediaPlayer interface can play MP4 and VLC
files. We'll use the Adapter pattern to make AdvancedMediaPlayer compatible
with MediaPlayer.

*/

// this interface defines the play method that our client will use
type MediaPlayer interface {
	Play(audioType, filename string)
}

// AdvancedMediaPlayer interface. methods for playing advanced media formats
type AdvancedMediaPlayer interface {
	PlayMP4(filename string)
	PlayVLC(filename string)
}

// implement concrete classes for advanced media player
type MP4Player struct { // implement the `Advanced Media Player` interface to play mp4 nd vlc files respectively
}

func (m MP4Player) PlayMP4(filename string) {
	fmt.Printf("Playing MP4 file. Name: %s\n", filename)
}

func (m MP4Player) PlayVLC(filename string) {
}

type VLCPlayer struct {
}

func (v VLCPlayer) PlayVLC(filename string) {
	fmt.Printf("Playing VLC file. Name: %s\n", filename)
}

func (v VLCPlayer) PlayMP4(filename string) {
}

// create the adapter class
type MediaAdapter struct { // Media adapter implements `MediaPlayer` interface and uses an
	//instance of `AdvancedMediaPlayer` to play appropriate media type
	advancedMusicPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	if audioType == "mp4" {
		return &MediaAdapter{advancedMusicPlayer: MP4Player{}}
	} else if audioType == "vlc" {
		return &MediaAdapter{advancedMusicPlayer: VLCPlayer{}}
	}
	return nil
}

// play method for MediaAdapter
func (adapter *MediaAdapter) Play(audioType, fileName string) {
	if audioType == "mp4" {
		adapter.advancedMusicPlayer.PlayMP4(fileName)
	} else if audioType == "vlc" {
		adapter.advancedMusicPlayer.PlayVLC(fileName)
	}
}

// implement the concrete class for Mediaplayer

type AudioPlayer struct {
	mediaplayer *MediaAdapter
}

func (player *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Printf("Playing MP3 files Name %s\n", fileName)
	} else if audioType == "mp4" || audioType == "vlc" {
		player.mediaplayer = NewMediaAdapter(audioType)
		player.mediaplayer.Play(audioType, fileName)
	} else {
		fmt.Printf("Invalid media. %s format not supported\n", audioType)
	}
}

func main() {

	audioPlayer := &AudioPlayer{}

	audioPlayer.Play("mp3", "song.mp3")
	audioPlayer.Play("mp4", "movie.mp4")
	audioPlayer.Play("vlc", "video.vlc")
	audioPlayer.Play("avi", "film.avi")
}
