package main

import "github.com/uber/h3-go/v4"
import "fmt"

func main() {
 latLng := h3.NewLatLng(37.775938728915946, -122.41795063018799)
 resolution := 9

 cell := h3.LatLngToCell(latLng, resolution)

 fmt.Printf("%s\n", cell)
}
