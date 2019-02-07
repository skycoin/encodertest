package encodertest

import (
	mathrand "math/rand"
	"testing"
	"time"
)

type Foo struct {
	A string
	B []int64 `enc:",maxlen=2"`
	C int64
}

func TestPopulate(t *testing.T) {
	rand := mathrand.New(mathrand.NewSource(time.Now().Unix()))

	var f Foo
	err := PopulateRandom(&f, rand, PopulateRandomOptions{
		MaxRandLen: 5,
		MinRandLen: 3,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(f.A) < 3 || len(f.A) > 5 {
		t.Fatal("String length should be between 3 and 5, inclusive")
	}

	if len(f.B) != 2 {
		t.Fatal("Slice length should have been capped at 2")
	}
}
