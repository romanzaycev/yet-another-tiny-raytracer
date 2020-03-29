package tinyraytracer

import (
	"math"

	"github.com/romanzaycev/yet-another-tiny-raytracer/src/types"
)

type raytracer struct {
	cfg *Config
}

//NewRaytracer Get raytracer instance
func NewRaytracer(config *Config) *raytracer {
	r := &raytracer{
		cfg: config,
	}

	return r
}

//Render Render function
func (r *raytracer) Render(scene types.Scene) ([]types.Vec3f, error) {
	spheres := scene.Spheres
	lights := scene.Lights

	buffer := make([]types.Vec3f, r.cfg.OutImageWidth*r.cfg.OutImageHeight)

	fovFactor := 2.*math.Tan(r.cfg.FOV/2.)
	width := r.cfg.OutImageWidth
	fWidth := float64(width)
	height := r.cfg.OutImageHeight
	fHeight := float64(height)
	zeroPoint := types.Vec3f{V0: .0, V1: .0, V2: .0}

	fWidth05 := fWidth/2.
	fHeight05 := fHeight/2.
	dirZ := -fHeight/fovFactor

	for j := 0; j<height; j++ {
		fJ := float64(j)
		dirY := -(fJ + .5) + fHeight05

		for i := 0; i<width; i++ {
			fI := float64(i)

			dirX := (fI + .5) - fWidth05
			dir := types.Vec3f{V0: dirX, V1: dirY, V2: dirZ}
			buffer[i+j*width] = castRay(
				r.cfg,
				zeroPoint,
				dir.Normalize(),
				&spheres,
				&lights,
				0,
			)
		}
	}

	return buffer, nil
}

func castRay(cfg *Config, orig types.Vec3f, dir types.Vec3f, spheres *[]types.Sphere, lights *[]types.Light, depth int) types.Vec3f {
	var point types.Vec3f
	var N types.Vec3f
	material := types.DefaultMaterial()
	material.Albedo = []float64{.2, .1, .0, .0}

	if depth > 4 || !sceneIntersect(cfg, orig, dir, spheres, &point, &N, &material) {
		return cfg.EnvironmentColor
	}

	diffuseLightIntensity := 0.
	specularLightIntensity := 0.

	for i := 0; i < len(*lights); i++ {
		lightDir := (*lights)[i].Position.Sub(point); lightDir = lightDir.Normalize()
		lightDistanceVec := (*lights)[i].Position.Sub(point)
		lightDistance := lightDistanceVec.Norm()

		var shadowOrig types.Vec3f
		lightDirMulN := lightDir.Mul(N)

		if lightDirMulN.Dot() < 0 {
			shadowOrig = point.Sub(N.Scale(1e-3))
		} else {
			shadowOrig = point.Add(N.Scale(1e-3))
		}

		var shadowPt types.Vec3f
		var shadowN types.Vec3f
		var tmpMaterial types.Material
		shadowPtSub := shadowPt.Sub(shadowOrig)

		if sceneIntersect(cfg, shadowOrig, lightDir, spheres, &shadowPt, &shadowN, &tmpMaterial) && shadowPtSub.Norm() < lightDistance {
			continue
		}

		diffuseLightIntensity += (*lights)[i].Intensity * math.Max(0., lightDirMulN.Dot())
		negReflect := reflect(lightDir.Reverse(), N)
		negReflect = negReflect.Reverse()
		negReflect = negReflect.Mul(dir)
		specularLightIntensity += math.Pow(math.Max(0., negReflect.Dot()), material.SpecularExponent) * (*lights)[i].Intensity
	}

	vec111 := types.Vec3f{V0: 1., V1: 1., V2: 1.}

	return material.DiffuseColor.Scale(diffuseLightIntensity * material.Albedo[0] + vec111.Dot() * specularLightIntensity * material.Albedo[1])
}

func sceneIntersect(cfg *Config, orig types.Vec3f, dir types.Vec3f, spheres *[]types.Sphere, hit *types.Vec3f, N *types.Vec3f, material *types.Material) bool {
	spheresDist := math.MaxFloat64

	for i := 0; i < len(*spheres); i++ {
		var distI float64

		if (*spheres)[i].RayIntersect(orig, dir, &distI) && distI < spheresDist {
			spheresDist = distI

			*hit = orig.Add(dir.Scale(distI))
			hitSub := hit.Sub((*spheres)[i].Center)
			*N = hitSub.Normalize()
			*material = (*spheres)[i].Material
		}
	}

	checkerboardDist := math.MaxFloat64

	if math.Abs(dir.V1) > 1e-3 {
		d := -(orig.V1 + 4) / dir.V1
		pt := orig.Add(dir.Scale(d))

		if d > 0 && math.Abs(pt.V0) < 10 && pt.V2 < -10 && pt.V2 > -30 && d < spheresDist {
			checkerboardDist = d
			*hit = pt
			*N = types.Vec3f{V0: 0., V1: 1., V2: 0.}

			checkerCheck := (int(.5*hit.V0+1000) + int(.5*hit.V2)) & 1

			if checkerCheck > 0 {
				material.DiffuseColor = cfg.CheckerboardColorA
			} else {
				material.DiffuseColor = cfg.CheckerboardColorB
			}
		}
	}

	return math.Min(spheresDist, checkerboardDist) < 1000
}

func reflect(i types.Vec3f, n types.Vec3f) types.Vec3f {
	ni := i.Mul(n)
	n2f := i.Scale(2.)

	return i.Sub(n2f.Mul(ni))
}
