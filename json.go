package main

type Font struct {
	BrSize float64 `json:"br_size"`
	Name   string  `json:"name"`
	Size   float64 `json:"size"`
}
type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type RichText struct {
	Font     Font     `json:"font"`
	Position Position `json:"pos"`
	Text     []string `json:"text"`
}
type Header struct {
	Left  RichText `json:"left"`
	Right RichText `json:"right"`
}
