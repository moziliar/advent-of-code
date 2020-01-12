package main

func Permute(seq []int) (out [][]int) {
	// base case
	if len(seq) < 2 {
		return [][]int{seq}
	}

	first := seq[0]
	seq = seq[1:]
	for i := 0; i < len(seq)+1; i++ {
		for _, v := range Permute(seq) {
			out = append(out, insert(first, i, v))
		}
	}
	return out
}

func insert(num int, i int, seq []int) []int {
	return append(seq[:i], append([]int{num}, seq[i:]...)...)
}
