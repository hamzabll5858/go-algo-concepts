package dynamic_programming

func GetFibonacciNum(index int, mem map[int]int) int {
	if mem[index] != 0 {
		return mem[index]
	}
	if index == 1 || index == 0 {
		return 1
	}
	mem[index] = GetFibonacciNum(index-1, mem) + GetFibonacciNum(index-2, mem)
	return mem[index]
}
