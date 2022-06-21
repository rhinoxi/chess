package internal

import (
	"embed"
	"os"
	"strings"
)

//go:embed pieces/*.svg
var svgPieces embed.FS

type asset struct {
	bytes []byte
	info  os.FileInfo
}

func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	bs, err := svgPieces.ReadFile(cannonicalName)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}
