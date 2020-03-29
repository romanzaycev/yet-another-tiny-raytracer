package tinyraytracer

import (
	"fmt"
	"strconv"

	"github.com/romanzaycev/yet-another-tiny-raytracer/src/types"
)

//HexRgbToComponent Get component color from rgb hex string, eq: ff0000
func HexRgbToComponent(color string) (types.Vec3f, error) {
	intColor, err := strconv.ParseInt(fmt.Sprintf("ff%s", color), 16, 64)

	if err != nil {
		return types.Vec3f{V0: 0, V1: 0, V2: 0}, err
	}

	r := (intColor & 0x00ff0000) >> 16
	g := (intColor & 0x0000ff00) >> 8
	b := intColor & 0x000000ff

	return types.Vec3f{V0: float64(r) / 255., V1: float64(g) / 255., V2: float64(b) / 255.,}, nil
}
