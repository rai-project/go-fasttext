package fasttext

import (
	"math"
	"sort"
)

type Vector struct {
	Element float32 `json:"probability"`
}

type Vectors []Vector

// Len is the number of elements in the collection.
func (p Vectors) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p Vectors) Less(i, j int) bool {
	pi := p[i].Element
	pj := p[j].Element
	return !(pi < pj || math.IsNaN(float64(pi)) && !math.IsNaN(float64(pj)))
}

// Swap swaps the elements with indexes i and j.
func (p Vectors) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Vectors) Sort() {
	sort.Sort(p)
}
