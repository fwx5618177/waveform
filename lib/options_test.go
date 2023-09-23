package lib

import (
	"fmt"
	"image/color"
	"testing"
)

// TestOptionsError verifies that the format of OptionsError.Error does
// not change.
func TestOptionsError(t *testing.T) {
	var tests = []struct {
		option string
		reason string
	}{
		{"foo", "bar"},
		{"baz", "qux"},
		{"one", "two"},
	}

	for _, test := range tests {
		// Generate options error
		opErr := &OptionsError{
			Option: test.option,
			Reason: test.reason,
		}

		// Verify correct format
		if opErr.Error() != fmt.Sprintf("%s: %s", test.option, test.reason) {
			t.Fatalf("unexpected Error string: %v", opErr.Error())
		}
	}
}

// TestOptionBGColorFunctionOK verifies that BGColorFunction returns no error
// with acceptable input.
func TestOptionBGColorFunctionOK(t *testing.T) {
	testWaveformOptionFunc(t, BGColorFunction(SolidColor(color.Black)), nil)
}

// TestOptionBGColorFunctionNil verifies that BGColorFunction does not accept
// a nil ColorReduceFunc.
func TestOptionBGColorFunctionNil(t *testing.T) {
	testWaveformOptionFunc(t, BGColorFunction(nil), errBGColorFunctionNil)
}

// TestOptionFGColorFunctionOK verifies that FGColorFunction returns no error
// with acceptable input.
func TestOptionFGColorFunctionOK(t *testing.T) {
	testWaveformOptionFunc(t, FGColorFunction(SolidColor(color.Black)), nil)
}

// TestOptionFGColorFunctionNil verifies that FGColorFunction does not accept
// a nil ColorReduceFunc.
func TestOptionFGColorFunctionNil(t *testing.T) {
	testWaveformOptionFunc(t, FGColorFunction(nil), errFGColorFunctionNil)
}

// TestOptionSampleFunctionOK verifies that SampleFunction returns no error
// with acceptable input.
func TestOptionSampleFunctionOK(t *testing.T) {
	testWaveformOptionFunc(t, SampleFunction(RMSF64Samples), nil)
}

// TestOptionSampleFunctionNil verifies that SampleFunction does not accept
// a nil SampleReduceFunc.
func TestOptionSampleFunctionNil(t *testing.T) {
	testWaveformOptionFunc(t, SampleFunction(nil), errSampleFunctionNil)
}

// TestOptionResolutionOK verifies that Resolution returns no error with acceptable input.
func TestOptionResolutionOK(t *testing.T) {
	testWaveformOptionFunc(t, Resolution(1), nil)
}

// TestOptionResolutionZero verifies that Resolution does not accept integer 0.
func TestOptionResolutionZero(t *testing.T) {
	testWaveformOptionFunc(t, Resolution(0), errResolutionZero)
}

// TestOptionScaleOK verifies that Scale returns no error with acceptable input.
func TestOptionScaleOK(t *testing.T) {
	testWaveformOptionFunc(t, Scale(1, 1), nil)
}

// TestOptionScaleXZero verifies that Scale does not accept an X value integer 0.
func TestOptionScaleXZero(t *testing.T) {
	testWaveformOptionFunc(t, Scale(0, 1), errScaleXZero)
}

// TestOptionScaleYZero verifies that Scale does not accept an Y value integer 0.
func TestOptionScaleYZero(t *testing.T) {
	testWaveformOptionFunc(t, Scale(1, 0), errScaleYZero)
}

// TestOptionScaleClippingOK verifies that ScaleClipping returns no error.
func TestOptionScaleClippingOK(t *testing.T) {
	testWaveformOptionFunc(t, ScaleClipping(), nil)
}

// TestOptionSharpnessOK verifies that Sharpness returns no error.
func TestOptionSharpnessOK(t *testing.T) {
	testWaveformOptionFunc(t, Sharpness(0), nil)
}

// TestlibetOptionsNil verifies that Waveform.SetOptions ignores any
// nil OptionsFunc arguments.
func TestlibetOptionsNil(t *testing.T) {
	testWaveformOptionFunc(t, nil, nil)
}

// TestlibetBGColorFunction verifies that the Waveform.SetBGColorFunction
// method properly modifies struct members.
func TestlibetBGColorFunction(t *testing.T) {
	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetBGColorFunction(SolidColor(color.Black)); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.bgColorFn == nil {
		t.Fatalf("SetBGColorFunction failed, nil function member")
	}
}

// TestlibetFGColorFunction verifies that the Waveform.SetFGColorFunction
// method properly modifies struct members.
func TestlibetFGColorFunction(t *testing.T) {
	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetFGColorFunction(SolidColor(color.Black)); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.fgColorFn == nil {
		t.Fatalf("SetFGColorFunction failed, nil function member")
	}
}

// TestlibetSampleFunction verifies that the Waveform.SetSampleFunction
// method properly modifies struct members.
func TestlibetSampleFunction(t *testing.T) {
	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetSampleFunction(RMSF64Samples); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.sampleFn == nil {
		t.Fatalf("SetSampleFunction failed, nil function member")
	}
}

// TestlibetResolution verifies that the Waveform.SetResolution method properly
// modifies struct members.
func TestlibetResolution(t *testing.T) {
	// Predefined test values
	res := uint(1)

	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetResolution(res); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.resolution != res {
		t.Fatalf("unexpected resolution: %v != %v", w.resolution, res)
	}
}

// TestlibetScale verifies that the Waveform.SetScale method properly
// modifies struct members.
func TestlibetScale(t *testing.T) {
	// Predefined test values
	x := uint(1)
	y := uint(1)

	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetScale(x, y); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.scaleX != x {
		t.Fatalf("unexpected scale X: %v != %v", w.scaleX, x)
	}
	if w.scaleY != y {
		t.Fatalf("unexpected scale Y: %v != %v", w.scaleY, y)
	}
}

// TestlibetScaleClipping verifies that the Waveform.SetScaleClipping method properly
// modifies struct members.
func TestlibetScaleClipping(t *testing.T) {
	// Generate empty Waveform, apply function
	w := &Waveform{}
	if err := w.SetScaleClipping(); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if !w.scaleClipping {
		t.Fatalf("SetScaleClipping failed, false scaleClipping member")
	}
}

// TestlibetSharpness verifies that the Waveform.SetSharpness method properly
// modifies struct members.
func TestlibetSharpness(t *testing.T) {
	// Predefined test values
	sharpness := uint(1)

	// Generate empty Waveform, apply parameters
	w := &Waveform{}
	if err := w.SetSharpness(sharpness); err != nil {
		t.Fatal(err)
	}

	// Validate that struct members are set properly
	if w.sharpness != sharpness {
		t.Fatalf("unexpected sharpness: %v != %v", w.sharpness, sharpness)
	}
}

// testWaveformOptionFunc is a test helper which verifies that applying the
// input OptionsFunc to a new Waveform struct generates the appropriate
// error output.
func testWaveformOptionFunc(t *testing.T, fn OptionsFunc, err error) {
	if _, wErr := New(nil, fn); wErr != err {
		t.Fatalf("unexpected error: %v != %v", wErr, err)
	}
}
