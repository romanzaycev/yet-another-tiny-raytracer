package scenes

import (
	"github.com/romanzaycev/yet-another-tiny-raytracer/src/types"
)

func SpheresScene() types.Scene {
	ivory := types.Material{
		RefractionFactor: 1.,
		Albedo: []float64{.6, .3, .1, .0},
		DiffuseColor: types.Vec3f{V0: .4, V1: .4, V2: .3},
		SpecularExponent: 50.,
	}
	redRubber := types.Material{
		RefractionFactor: 1.0,
		Albedo: []float64{.9, .1, .0, .0},
		DiffuseColor: types.Vec3f{V0: .3, V1: .1, V2: .1},
		SpecularExponent: 10.,
	}
	grassRubber := types.Material{
		RefractionFactor: 1.0,
		Albedo: []float64{.2, .1, .0, .0},
		DiffuseColor: types.Vec3f{V0: .15, V1: .68, V2: .37},
		SpecularExponent: 3.,
	}

	return types.Scene{
		Spheres: []types.Sphere{
			types.Sphere{
				Center: types.Vec3f{V0: -3., V1: 0., V2: -16.},
				Radius: 2.,
				Material: ivory,
			},
			types.Sphere{
				Center: types.Vec3f{V0: -1., V1: -1.5, V2: -12.},
				Radius: 2.,
				Material: grassRubber,
			},
			types.Sphere{
				Center: types.Vec3f{V0: 1.5, V1: -.5, V2: -18.},
				Radius: 3.,
				Material: redRubber,
			},
			types.Sphere{
				Center: types.Vec3f{V0: 7., V1: 6., V2: -14.},
				Radius: 4.,
				Material: types.DefaultMaterial(),
			},
		},
		Lights: []types.Light{
			types.Light{
				Position: types.Vec3f{V0: -20., V1: 20.,  V2: 20.},
				Intensity: 1.5,
			},
			types.Light{
				Position: types.Vec3f{V0: 30., V1: 50.,  V2: -25.},
				Intensity: 1.8,
			},
			types.Light{
				Position: types.Vec3f{V0: 30., V1: 20.,  V2: 30.},
				Intensity: 1.7,
			},
		},
	}
}
