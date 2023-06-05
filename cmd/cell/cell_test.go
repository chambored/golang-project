package main

import (
	"math"
	"reflect"
	"testing"
)

func TestAverageWeight(t *testing.T) {
	// Creating a test cell map
	cells := map[string]*Cell{
		"Google-Pixel 4 XL":  {bodyWeight: 193.0},
		"Google-Pixel 3":     {bodyWeight: 148.0},
		"Samsung-Galaxy S10": {bodyWeight: 157.0},
		"Empty-Weight":       {bodyWeight: 0}, // This one should be excluded from the average.
	}

	// Execute the function with the test cell map
	result := averageWeight(cells)

	// Expected average weight = (193.0 + 148.0 + 157.0) / 3 = 166.0
	var expected float32 = 166.0

	// If the result does not match the expected average, fail the test
	if result != expected {
		t.Errorf("averageWeight(cells) = %.2f; want %.2f", result, expected)
	}
}

func TestAverageDisplaySize(t *testing.T) {
	// Creating a test cell map
	cells := map[string]*Cell{
		"Google-Pixel 4 XL":  {displaySize: 6.3},
		"Google-Pixel 3":     {displaySize: 5.5},
		"Samsung-Galaxy S10": {displaySize: 6.1},
		"Zero-Display-Size":  {displaySize: 0}, // This one should be excluded from the average.
	}

	// Execute the function with the test cell map
	got := averageDisplaySize(cells)

	got = round(got, 2)

	want := 5.97

	if got != want {
		t.Errorf("averageDisplaySize(cells) = %v; want %v", got, want)
	}
}

func round(val float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return math.Round(val*p) / p
}

func TestCountPhonesByYear(t *testing.T) {
	// Create a map of cells
	cells := map[string]*Cell{
		"Google-Pixel 4 XL":  {launchAnnounced: 2019},
		"Google-Pixel 3":     {launchAnnounced: 2018},
		"Samsung-Galaxy S10": {launchAnnounced: 2019},
		"Samsung-Galaxy S9":  {launchAnnounced: 2018},
		"No-Announced-Year":  {launchAnnounced: 0}, // This one should be ignored.
	}

	// Expected result
	expected := YearCounts{
		Counts: map[uint]int{
			2018: 2,
			2019: 2,
		},
		Years: []uint{2018, 2019},
	}

	// Execute the function with the test data
	result := countPhonesByYear(cells)

	// If the result does not match the expected counts and years, fail the test
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("countPhonesByYear(cells) = %v; want %v", result, expected)
	}
}

func TestCountUniqueOS(t *testing.T) {
	cells := map[string]*Cell{
		"Phone1": {platformOS: "Android"},
		"Phone2": {platformOS: "Android"},
		"Phone3": {platformOS: "iOS"},
		"Phone4": {platformOS: "Windows"},
		"Phone5": {platformOS: "iOS"},
		"Phone6": {platformOS: ""},
		"Phone7": {platformOS: "Android"},
	}

	want := 3 // We expect 3 unique operating systems: Android, iOS, and Windows
	got := countUniqueOS(cells)

	if got != want {
		t.Errorf("countUniqueOS(cells) = %d; want %d", got, want)
	}
}

func TestCountPhonesByOEM(t *testing.T) {
	cells := map[string]*Cell{
		"Phone1": {oem: "Apple"},
		"Phone2": {oem: "Apple"},
		"Phone3": {oem: "Samsung"},
		"Phone4": {oem: "Samsung"},
		"Phone5": {oem: "Samsung"},
		"Phone6": {oem: "Google"},
	}

	want := map[string]int{"Apple": 2, "Samsung": 3, "Google": 1}
	got := countPhonesByOEM(cells)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("countPhonesByOEM(cells) = %v; want %v", got, want)
	}
}

func TestFindLatestPhoneByOEM(t *testing.T) {
	cells := map[string]*Cell{
		"Phone1": {oem: "Apple", model: "iPhone 13", launchAnnounced: 2023},
		"Phone2": {oem: "Apple", model: "iPhone 12", launchAnnounced: 2022},
		"Phone3": {oem: "Samsung", model: "Galaxy S21", launchAnnounced: 2022},
		"Phone4": {oem: "Samsung", model: "Galaxy S22", launchAnnounced: 2023},
		"Phone5": {oem: "Google", model: "Pixel 6", launchAnnounced: 2023},
	}

	want := map[string]*Cell{
		"Apple":   {oem: "Apple", model: "iPhone 13", launchAnnounced: 2023},
		"Samsung": {oem: "Samsung", model: "Galaxy S22", launchAnnounced: 2023},
		"Google":  {oem: "Google", model: "Pixel 6", launchAnnounced: 2023},
	}

	got := findLatestPhoneByOEM(cells)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("findLatestPhoneByOEM(cells) = %v; want %v", got, want)
	}
}

func TestFindHeaviestAndLightestPhones(t *testing.T) {
	cells := map[string]*Cell{
		"Phone1": {model: "iPhone 13", bodyWeight: 174},
		"Phone2": {model: "iPhone 12 Mini", bodyWeight: 135},
		"Phone3": {model: "Galaxy S21", bodyWeight: 0}, // This should be ignored
		"Phone4": {model: "Galaxy S22", bodyWeight: 200},
		"Phone5": {model: "Pixel 6", bodyWeight: 143},
	}

	wantHeaviest := &Cell{model: "Galaxy S22", bodyWeight: 200}
	wantLightest := &Cell{model: "iPhone 12 Mini", bodyWeight: 135}

	gotHeaviest, gotLightest := findHeaviestAndLightestPhones(cells)

	if !reflect.DeepEqual(gotHeaviest, wantHeaviest) {
		t.Errorf("Heaviest phone incorrect, got: %v, want: %v.", gotHeaviest, wantHeaviest)
	}

	if !reflect.DeepEqual(gotLightest, wantLightest) {
		t.Errorf("Lightest phone incorrect, got: %v, want: %v.", gotLightest, wantLightest)
	}
}

func TestParseYear(t *testing.T) {
	tests := []struct {
		input string
		want  *uint
	}{
		{"Announced 2023", uintPtr(2023)},
		{"Invalid year", nil},
		{"2000 2023", uintPtr(2000)}, // If there are multiple 4-digit numbers, the first one is taken
		{"", nil},
	}

	for _, test := range tests {
		got := parseYear(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseYear(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestParseWeight(t *testing.T) {
	tests := []struct {
		input string
		want  *float32
	}{
		{"174 g", float32Ptr(174)},
		{"Invalid weight", nil},
		{"", nil},
	}

	for _, test := range tests {
		got := parseWeight(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseWeight(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestParseSim(t *testing.T) {
	tests := []struct {
		input string
		want  *string
	}{
		{"yes", nil},
		{"no", nil},
		{"Nano", strPtr("Nano")},
		{"", strPtr("")},
	}

	for _, test := range tests {
		got := parseSim(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseSim(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestParseSize(t *testing.T) {
	sizeStr := "6.1 inches"
	got := parseSize(sizeStr)

	if got == nil {
		t.Errorf("parseSize(%q) returned nil, want non-nil", sizeStr)
	}

	const tolerance = 1e-6
	want := 6.1

	if !almostEqual(*got, want, tolerance) {
		t.Errorf("parseSize(%q) = %v, want %v", sizeStr, *got, want)
	}
}

func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestParseSensors(t *testing.T) {
	tests := []struct {
		input string
		want  *string
	}{
		{"12.2", nil},
		{"accelerometer", strPtr("accelerometer")},
		{"", strPtr("")},
	}

	for _, test := range tests {
		got := parseSensors(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parseSensors(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestParsePlatformOS(t *testing.T) {
	tests := []struct {
		input string
		want  *string
	}{
		{"10.0", nil},
		{"Android, 10.0", strPtr("Android")},
		{"", strPtr("")},
	}

	for _, test := range tests {
		got := parsePlatformOS(test.input)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("parsePlatformOS(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

// Helper functions to create pointers to string and numerical values
func strPtr(s string) *string       { return &s }
func uintPtr(u uint) *uint          { return &u }
func float32Ptr(f float32) *float32 { return &f }
func float64Ptr(f float64) *float64 { return &f }
