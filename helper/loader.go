package helper

import (
	"embed"
	"image"
	_ "image/png"
	"io"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/tinne26/etxt"
)

const (
	SampleRate = 44100
)

func Image(fs embed.FS, Filename string) *ebiten.Image {
	file, err := fs.Open(Filename)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func Font(fs embed.FS, Filename string) *etxt.FontLibrary {
	file, err := fs.Open(Filename)
	if err != nil {
		panic(err)
	}

	bs, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fontLib := etxt.NewFontLibrary()
	fontLib.ParseFontBytes(bs)

	return fontLib
}

func LoadAllFilesAsAudio(fs embed.FS, Folder string, context *audio.Context) map[string]*audio.Player {
	streams := map[string]*audio.Player{}

	folder, err := fs.ReadDir(Folder)
	if err != nil {
		panic(err)
	}

	for _, v := range folder {
		name := Folder + "/" + v.Name()

		file, err := fs.Open(name)
		if err != nil {
			panic(err)
		}

		d, err := mp3.DecodeWithSampleRate(SampleRate, file)
		if err != nil {
			panic(err)
		}

		p, err := context.NewPlayer(d)
		if err != nil {
			panic("Cannot create player for: " + name)
		}
		streams[name] = p
	}

	return streams
}
