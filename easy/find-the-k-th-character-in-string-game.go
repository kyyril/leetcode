package main

func kthCharacter(k int) byte {
	wordLengths := []int{1} // panjang word setelah tiap iterasi
	for wordLengths[len(wordLengths)-1] < k {
		wordLengths = append(wordLengths, wordLengths[len(wordLengths)-1]*2)
	}

	var getChar func(k int, iter int) byte
	getChar = func(k int, iter int) byte {
		if iter == 0 {
			return 'a'
		}
		prevLen := wordLengths[iter-1]
		if k <= prevLen {
			return getChar(k, iter-1)
		} else {
			// next character dari posisi (k - prevLen)
			ch := getChar(k-prevLen, iter-1)
			if ch == 'z' {
				return 'a'
			}
			return ch + 1
		}
	}

	return getChar(k, len(wordLengths)-1)
}
