package types

import (
	"testing"
)

func TestVec3f_Normalize(t *testing.T) {
	vec := Vec3f{1, 0, 0}
	vec = vec.Normalize()

	if vec.V0 != 1. {
		t.Fatalf("Expected %v, but got %v", 1., vec.V0)
	}
}

func TestVec3f_Add(t *testing.T) {
	vec := Vec3f{1, 2, 3}
	vec = vec.Add(Vec3f{2, 2, 2})

	if vec.V0 != 3 {
		t.Fatalf("Expected %v, but got %v", 3, vec.V0)
	}

	if vec.V1 != 4 {
		t.Fatalf("Expected %v, but got %v", 4, vec.V1)
	}

	if vec.V2 != 5 {
		t.Fatalf("Expected %v, but got %v", 5, vec.V2)
	}
}

func TestVec3f_Sub(t *testing.T) {
	vec := Vec3f{1, 2, 3}
	vec = vec.Sub(Vec3f{1, 1, 1})

	if vec.V0 != 0 {
		t.Fatalf("Expected %v, but got %v", 1, vec.V0)
	}

	if vec.V1 != 1 {
		t.Fatalf("Expected %v, but got %v", 2, vec.V1)
	}

	if vec.V2 != 2 {
		t.Fatalf("Expected %v, but got %v", 3, vec.V2)
	}
}

func TestVec3f_Mul(t *testing.T) {
	vec := Vec3f{1, 2, 3}
	vec = vec.Mul(Vec3f{2, 2, 2})

	if vec.V0 != 2 {
		t.Fatalf("Expected %v, but got %v", 2, vec.V0)
	}

	if vec.V1 != 4 {
		t.Fatalf("Expected %v, but got %v", 4, vec.V1)
	}

	if vec.V2 != 6 {
		t.Fatalf("Expected %v, but got %v", 6, vec.V2)
	}
}

func TestVec3f_Scale(t *testing.T) {
	vec := Vec3f{1, 2, 3}
	vec = vec.Scale(3)

	if vec.V0 != 3 {
		t.Fatalf("Expected %v, but got %v", 3, vec.V0)
	}

	if vec.V1 != 6 {
		t.Fatalf("Expected %v, but got %v", 6, vec.V1)
	}

	if vec.V2 != 9 {
		t.Fatalf("Expected %v, but got %v", 9, vec.V2)
	}
}

func TestVec3f_Reverse(t *testing.T) {
	vec := Vec3f{1, 2, 3}
	vec = vec.Reverse()

	if vec.V0 != -1 {
		t.Fatalf("Expected %v, but got %v", -1, vec.V0)
	}

	if vec.V1 != -2 {
		t.Fatalf("Expected %v, but got %v", -2, vec.V1)
	}

	if vec.V2 != -3 {
		t.Fatalf("Expected %v, but got %v", -3, vec.V2)
	}
}

