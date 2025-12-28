package boardthemes

import (
	"encoding/json"
	"fmt"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
	uerr "github.com/mr-lasoma/chessgo/utils/errors"
)

const (
	fileName      = "boardthemes/themes.go"
	themeFilePath = "../../data/theme.json"
)

var (
	BoardThemes map[string]Theme
)

// Theme is the struct used for the chessboard themes
type Theme struct {
	Name       string    `json:"name"`
	LightColor colorJson `json:"light_color"`
	DarkColor  colorJson `json:"dark_color"`
}

type themesWrapper struct {
	Themes []Theme `json:"themes"`
}

// colorJson is the struct used to store the parsed file in a more secure way
// (i'm not sure if the rl.Color works with the go parser)
type colorJson struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

// ToRayLibColor returns the rl.Color based on the color chess.color
func (c *colorJson) ToRayLibColor() rl.Color {
	return rl.Color{
		R: c.R,
		G: c.G,
		B: c.B,
		A: c.A,
	}
}

// MustLoadThemes loads all the themes from a file, if it fails
// in any ways it panics
func MustLoadThemes() {
	data, err := os.ReadFile(themeFilePath)
	if err != nil {
		uerr.PanicStdMsg(fileName, fmt.Sprintf("unable to find %#v file", themeFilePath))
	}

	var file themesWrapper
	if err := json.Unmarshal(data, &file); err != nil {
		uerr.PanicStdMsg(fileName, fmt.Sprintf("unable to parse %#v jsonfile into themesWrapper struct", themeFilePath))
	}

	BoardThemes = make(map[string]Theme)

	for _, t := range file.Themes {
		BoardThemes[t.Name] = t
	}
}

func MustGetBoardTheme(name string) *Theme {
	t, ok := BoardThemes[name]
	if !ok {
		uerr.PanicStdMsg(fileName, fmt.Sprintf("unable to find %#v theme in %#v file", name, themeFilePath))
	}
	return &t
}
