package tinyraytracer

import (
	"bufio"
	"fmt"
	"math"

	"github.com/romanzaycev/yet-another-tiny-raytracer/src/types"
)

//WritePpmFile Writes PPM file from bitmap data
func WritePpmFile(width int, height int, in []types.Vec3f, out *bufio.Writer) error {
	_, err := out.WriteString(fmt.Sprintf("P6\n%d %d \n255\n", width, height))

	if err != nil {
		return err
	}

	length := width * height

	for i := 0; i < length; i++ {
		c := in[i]
		max := math.Max(c.V0, math.Min(c.V1, c.V2))

		if max > 1 {
			c = c.Scale(1./max)
		}

		_ = out.WriteByte(uint8(255 * math.Max(0., math.Min(1., c.V0))))
		_ = out.WriteByte(uint8(255 * math.Max(0., math.Min(1., c.V1))))
		_ = out.WriteByte(uint8(255 * math.Max(0., math.Min(1., c.V2))))
	}

	err = out.Flush()

	if err != nil {
		return err
	}

	return nil
}
