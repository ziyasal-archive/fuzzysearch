package main

import (
	"fmt"
)

func Fuzzysearch(needle string, haystack string) bool {
	var _needle = []rune(needle)
	var _haystack = []rune(haystack)

	var hlen = len(haystack)
	var nlen = len(needle)
	
	if nlen > hlen {
		return false
	}
	
	if nlen == hlen {
		return needle == haystack
	}
L:
	for i, j := 0, 0; i < nlen; i++ {
		var nch = _needle[i]
		for j < hlen {
			j := j + 1
			if _haystack[j] == nch {
				continue L
			}
		}
		return false
	}
	return true
}
