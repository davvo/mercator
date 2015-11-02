package mercator

import (
	"math"
	"testing"
)

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < 0.00000001
}

func TestResolution(t *testing.T) {
	zoom := 10
	expected := 152.8740565703525
	res := Resolution(zoom)
	if !floatEquals(res, expected) {
		t.Errorf("Resolution(%d) == %f, want %f", zoom, res, expected)
	}
}

func TestZoom(t *testing.T) {
	res := 152.8740565703525
	expected := 10
	zoom := Zoom(res)
	if zoom != expected {
		t.Errorf("Zoom(%f) == %d, want %d", res, zoom, expected)
	}
}

func TestLatLonToMeters(t *testing.T) {
	lat, lon := 62.3, 14.1
	expectedX, expectedY := 1569604.8201851572, 8930630.669201756
	x, y := LatLonToMeters(lat, lon)
	if !floatEquals(x, expectedX) || !floatEquals(y, expectedY) {
		t.Errorf("LatLonToMeters(%f, %f) == %f, %f, want %f, %f", lat, lon, x, y, expectedX, expectedY)
	}
}

func TestMetersToLatLon(t *testing.T) {
	x, y := 1569604.8201851572, 8930630.669201756
	expectedLat, expectedLon := 62.3, 14.1
	lat, lon := MetersToLatLon(x, y)
	if !floatEquals(lat, expectedLat) || !floatEquals(lon, expectedLon) {
		t.Errorf("MetersToLatLon(%f, %f) == %f, %f, want %f, %f", x, y, lat, lon, expectedLat, expectedLon)
	}
}

func TestPixelsToMeters(t *testing.T) {
	px, py, zoom := 123456789.0, 123456789.0, 15
	expectedX, expectedY := 569754371.206588, 569754371.206588
	x, y := PixelsToMeters(px, py, zoom)
	if !floatEquals(x, expectedX) || !floatEquals(y, expectedY) {
		t.Errorf("PixelsToMeters(%d, %d, %d) == %f, %f, want %f, %f", px, py, zoom, x, y, expectedX, expectedY)
	}
}

func TestMetersToPixels(t *testing.T) {
	x, y, zoom := 569754371.206588, 569754371.206588, 15
	expectedPx, expectedPy := 123456789.0, 123456789.0
	px, py := MetersToPixels(x, y, zoom)
	if !floatEquals(px, expectedPx) || !floatEquals(py, expectedPy) {
		t.Errorf("MetersToPixels(%f, %f, %d) == %d, %d, want %d, %d", x, y, zoom, px, py, expectedPx, expectedPy)
	}
}

func TestLatLonToPixels(t *testing.T) {
	lat, lon, zoom := 62.3, 14.1, 15
	expectedPx, expectedPy := 4522857.8133333335, 6063687.123767246
	px, py := LatLonToPixels(lat, lon, zoom)
	if !floatEquals(px, expectedPx) || !floatEquals(py, expectedPy) {
		t.Errorf("LatLonToPixels(%f, %f, %d) == %f, %f, want %f, %f", lat, lon, zoom, px, py, expectedPx, expectedPy)
	}
}

func TestPixelsToLatLon(t *testing.T) {
	px, py, zoom := 4522857.8133333335, 6063687.123767246, 15
	expectedLat, expectedLon := 62.3, 14.1
	lat, lon := PixelsToLatLon(px, py, zoom)
	if !floatEquals(lat, expectedLat) || !floatEquals(lon, expectedLon) {
		t.Errorf("PixelsToLatLon(%f, %f, %d) == %f, %f, want %f, %f", px, py, zoom, lat, lon, expectedLat, expectedLon)
	}
}

func TestPixelsToTile(t *testing.T) {
	px, py := 123456789.0, 123456789.0
	expectedTileX, expectedTileY := 482253, 482253
	tileX, tileY := PixelsToTile(px, py)
	if tileX != expectedTileX || tileY != expectedTileY {
		t.Errorf("PixelsToTile(%f, %f) == %d, %d, want %d, %d", px, py, tileX, tileY, expectedTileX, expectedTileY)
	}
}

func TestMetersToTile(t *testing.T) {
	x, y, zoom := 569754371.206588, 569754371.206588, 15
	expectedTileX, expectedTileY := 482253, 482253
	tileX, tileY := MetersToTile(x, y, zoom)
	if tileX != expectedTileX || tileY != expectedTileY {
		t.Errorf("MetersToTile(%f, %f, %d) == %d, %d, want %d, %d", x, y, zoom, tileX, tileY, expectedTileX, expectedTileY)
	}
}
