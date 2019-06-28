package main

/**
Knuth-Morris-Pratt
	 - An algorithm to search for substring in a given string.
Steps:
	 - Generate a suffic array which tells from where to begin the next search in pattern if a mismatch detected
	 - Iterate through text and while matching text
		 - If char is equal then increament pattern match count
		 - Else use suffix array to know where to start matching string again from next time
*/

func main() {
	var found_at = search("juaid", "d")
	if found_at >= 0 {
		println("Yes, found at", found_at)
	} else {
		println("Not found")
	}
}

/*
	Checks if `pattern` is substring of `text`, if yes the returns index where pattern begins inside text
	else returns -1
*/
func search(pattern string, text string) int {
	var suffix = make([]int, len(pattern))
	suffix_start := 0
	for suffix_end := 1; suffix_end < len(pattern); {
		ichar := pattern[suffix_end]
		jchar := pattern[suffix_start]
		if ichar == jchar {
			suffix[suffix_end] = suffix_start + 1
			suffix_start += 1
		} else {
			if suffix_start != 0 {
				suffix_start = suffix[suffix_start-1]
				continue
			}
		}
		suffix_end += 1

	}

	pattern_index := 0
	pattern_last_index := len(pattern) - 1
	for text_index := 0; text_index < len(text); text_index++ {
		schar := text[text_index]
		pchar := pattern[pattern_index]
		if schar == pchar {
			if pattern_index == pattern_last_index {
				return (text_index - pattern_last_index)
			}
			pattern_index += 1
		} else {
			if pattern_index > 0 {
				pattern_index = suffix[pattern_index]
				text_index -= 1
			}
		}
	}
	return -1
}
