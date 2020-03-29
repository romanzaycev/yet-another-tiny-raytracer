package types

//Material Material structure
type Material struct {
	RefractionFactor float64
	Albedo []float64
	DiffuseColor Vec3f
	SpecularExponent float64
}

//DefaultMaterial Get default gray material
func DefaultMaterial() Material  {
	return Material{
		RefractionFactor: 0.,
		Albedo:           []float64{1., 0., 0, 0.,},
		DiffuseColor:     Vec3f{.333, .333,.333},
		SpecularExponent: 0.,
	}
}