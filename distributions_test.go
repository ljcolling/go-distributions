package distributions

import (
	"github.com/google/go-cmp/cmp"
	"math"
	"testing"
)

func TestDistribution(t *testing.T) {

	// note: all test values are taken from R v4.0.2

	const tolerance = .0001
	opt := cmp.Comparer(func(x, y float64) bool {
		diff := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return (diff / mean) < tolerance
	})

	// non-central t distribution

	got := Dt(10, 10, 10)
	want := 0.1601017

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

	got = Dt(0.1, 12, .2)
	want = 0.3887798

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

	got = Dt(9, 12, .4)
	want = 2.318372e-06

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

	// cauchy distribution
	got = Dcauchy(0, 1, 1)
	want = 0.1591549

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

	got = Dcauchy(0, .707, 1)
	want = 0.212228

	if !cmp.Equal(got, want, opt) {
		t.Fatalf("got %v, wanted %v", got, want)
	}

}
