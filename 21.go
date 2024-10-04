package main

import (
	"fmt"
	"strings"
)

// Реализовать паттерн «адаптер» на любом примере.

// MediaPlayer - интерфейс, который умеет проигрывать MP3 файлы.
type MediaPlayer interface {
	PlayAudio(filename string)
}

// AdvancedMediaPlayer - интерфейс, который умеет проигрывать MP4 и VLC файлы.
type AdvancedMediaPlayer interface {
	PlayMP4(filename string)
	PlayVLC(filename string)
}

// AdvancedPlayer - структура, реализующая интерфейс AdvancedMediaPlayer.
type AdvancedPlayer struct{}

func (a *AdvancedPlayer) PlayMP4(filename string) {
	fmt.Printf("Playing MP4 file: %s\n", filename)
}

func (a *AdvancedPlayer) PlayVLC(filename string) {
	fmt.Printf("Playing VLC file: %s\n", filename)
}

// MediaAdapter - адаптер, который делает AdvancedMediaPlayer совместимым с MediaPlayer.
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

// PlayAudio - метод адаптера, который вызывает PlayMP4 или PlayVLC для AdvancedPlayer.
func (m *MediaAdapter) PlayAudio(filename string) {
	fmt.Printf("Using adapter to play audio...\n")

	// Определяем формат файла и вызываем соответствующий метод проигрывания.
	if strings.HasSuffix(filename, ".mp4") {
		m.advancedPlayer.PlayMP4(filename)
	} else if strings.HasSuffix(filename, ".vlc") {
		m.advancedPlayer.PlayVLC(filename)
	} else {
		fmt.Println("Unsupported format")
	}
}

// AudioPlayer - структура, реализующая интерфейс MediaPlayer.
type AudioPlayer struct{}

func (a *AudioPlayer) PlayAudio(filename string) {
	// Если формат MP3, то проигрываем напрямую
	if strings.HasSuffix(filename, ".mp3") {
		fmt.Printf("Playing MP3 file: %s\n", filename)
	} else {
		// Используем адаптер для воспроизведения других форматов.
		adapter := &MediaAdapter{advancedPlayer: &AdvancedPlayer{}}
		adapter.PlayAudio(filename)
	}
}

func main() {
	// Создаем экземпляр AudioPlayer
	player := &AudioPlayer{}

	// Проигрываем MP3 файл напрямую
	player.PlayAudio("song.mp3")

	// Проигрываем MP4 файл через адаптер
	player.PlayAudio("video.mp4")

	// Проигрываем VLC файл через адаптер
	player.PlayAudio("movie.vlc")

	// Попробуем проиграть неподдерживаемый формат
	player.PlayAudio("document.pdf")
}
