package main

func tp(millimeters float64) float64 {
	// 1 inch == 72 typographic points; 1 inch == 25.4 millimeters.
	return millimeters / 25.4 * 72
}
