//go:build linux || darwin
package linux

import (
	"os/exec"

	"github.com/elias-gill/walldo-in-go/globals"
)

func setFehBackground(file string, mode Mode) error {
	return exec.Command("feh", mode.setFehMode(), file).Run()
}

func (mode Mode) setFehMode() (string) {
	switch mode {
	case globals.FILL_CENTER:
		return "--bg-center"
	case globals.FILL_SCALE:
		return "--bg-scale"
	case globals.FILL_ZOOM:
		return "--bg-fill"
	case globals.FILL_ORIGINAL:
		return "--bg-max"
	case globals.FILL_TILE:
		return "--bg-tile"
	default:
		panic("invalid wallpaper mode")
	}
}
