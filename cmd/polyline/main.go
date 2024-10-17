package main

import (
	"fmt"

	"github.com/twpayne/go-polyline"
)

func runDecode() error {
	originalEncoded := "qeymAgpbyMGRUx@IXIXi@bD[hBe@zAQ`@{CtFU\\i@l@q@f@i@Z_@Po@Je@NSVIBS\\[x@[t@i@pAKV_@x@INIVGNw@pBEJKZELu@|CCHk@hCMj@AFMb@Qh@{AjDyAjDc@lASl@c@dBEPEVK^CNJBH@|Bb@VqAuBY"
	coords, _, err := polyline.DecodeCoords([]byte(originalEncoded))
	if err != nil {
		return err
	}

	encoded := polyline.EncodeCoords(coords)

	fmt.Println("Matched:", originalEncoded == string(encoded))
	return nil
}

func runConcat() {
	fmt.Println("--- Original ---")
	coords := [][]float64{
		{38.5, -120.2},
		{40.7, -120.95},
		{43.252, -126.453},
	}
	poly := string(polyline.EncodeCoords(coords))
	fmt.Println("Coordinates:", coords)
	fmt.Println("Polyline:", poly)
	fmt.Println()

	// 1st segment
	fmt.Println("--- Segment 1 ---")
	coords1 := coords[0:2]
	poly1 := string(polyline.EncodeCoords(coords1))
	fmt.Println("Coordinates 1:", coords1)
	fmt.Println("Polyline 1:", poly1)
	fmt.Println()

	// 2nd segment
	fmt.Println("--- Segment 2 ---")
	coords2 := coords[1:3]
	poly2 := string(polyline.EncodeCoords(coords2))
	fmt.Println("Coordinates 2:", coords2)
	fmt.Println("Polyline 2:", poly2)
	fmt.Println()

	fmt.Println("--- Merged ---")
	mergedCoords := append(coords1, coords2[1:]...) // assume always valid that the coords1[lastIndex] == coords2[firstIndex]
	mergedPoly := string(polyline.EncodeCoords(mergedCoords))
	fmt.Println("Merged coordinates:", mergedCoords)
	fmt.Println("Merged polyline:", mergedPoly)
	fmt.Println()

	fmt.Println("--- Comparison ---")
	fmt.Println("Matched:", poly == mergedPoly)
}

func main() {
	runConcat()
}
