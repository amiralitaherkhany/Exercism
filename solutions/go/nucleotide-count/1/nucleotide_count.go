package dna

import "errors"

type Histogram map[rune]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h Histogram = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, v := range d {
		switch v {
		case 'A':
			h[v]++
		case 'C':
			h[v]++
		case 'G':
			h[v]++
		case 'T':
			h[v]++
		default:
			return h, errors.New("")
		}
	}
	return h, nil
}
