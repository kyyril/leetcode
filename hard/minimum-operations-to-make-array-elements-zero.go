package main

import (
	"fmt"
)

func minOperations(queries [][]int) int64 {
	prefix := func(x int) int64 {
		if x <= 0 {
			return 0
		}
		var res int64
		pow := int64(1)
		cost := int64(1)
		xx := int64(x)
		for {
			next := pow * 4
			if xx < next {
				res += (xx - pow + 1) * cost
				break
			}
			res += (next - pow) * cost
			pow = next
			cost++
		}
		return res
	}

	var ans int64
	for _, q := range queries {
		l, r := q[0], q[1]
		totalDigits := prefix(r) - prefix(l-1)
		ops := (totalDigits + 1) / 2 // ceil(totalDigits/2)
		ans += ops
	}
	return ans
}

func main() {
	fmt.Println(minOperations([][]int{{1, 2}, {2, 4}})) // 3
	fmt.Println(minOperations([][]int{{2, 6}}))         // 4
}
