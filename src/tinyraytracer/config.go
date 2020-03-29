package tinyraytracer

import (
	"math"

	"github.com/romanzaycev/yet-another-tiny-raytracer/src/types"
)

//Config
type Config struct {
	OutImageWidth      int
	OutImageHeight     int
	OutImageName       string
	FOV                float64
	EnvironmentColor   types.Vec3f
	CheckerboardColorA types.Vec3f
	CheckerboardColorB types.Vec3f
}

//NewConfig Get config instance
func NewConfig() *Config {
	return &Config{
		OutImageWidth:  1000,
		OutImageHeight: 100,
		OutImageName:   "out.ppm",
		FOV:            math.Pi / 3,
	}
}
