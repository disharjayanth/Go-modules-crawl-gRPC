package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want = %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want = %q", got, want)
	}
}

func TestWorldTravelers(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := worldTraverlers(); got != want {
		t.Errorf("worldTravelers = %q, want = %q", got, want)
	}
}

func TestGlassV3(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := glassV3(); got != want {
		t.Errorf("glassv3 = %q, want = %q", got, want)
	}
}

func TestOptv2(t *testing.T) {
	want := "If a program is too slow, it must have a loop."
	if got := optV2(); got != want {
		t.Errorf("optv2 = %q, want = %q", got, want)
	}
}
