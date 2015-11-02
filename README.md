# mercator
Mercator coordinate system conversions

Golang port of the code from http://www.maptiler.org/google-maps-coordinates-tile-bounds-projection/

## Install
```
go get github.com/davvo/mercator
```

## Usage
```
package main

import (
	"github.com/davvo/mercator"
)

var x, y, z = 1569604.8201851572, 8930630.669201756, 10
var lat, lon = mercator.MetersToLatLon(x, y)
var px, py = mercator.MetersToPixels(x, y, z)
var tx, ty = mercator.LatLonToTile(lat, lon, z)

fmt.Printf("Meters: %f, %f\n", x, y)
fmt.Printf("Lat Lon: %f, %f\n", lat, lon)
fmt.Printf("Pixels: %f, %f\n", px, py)
fmt.Printf("Tile: %d, %d\n", tx, ty)
```

