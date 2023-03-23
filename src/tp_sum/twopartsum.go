package main

func sum(a []int, res chan<- int) {
	sum := 0
	arrlen := len(a)
	for i := 0; i < arrlen; i++ {
		sum += a[i]
	}

	res <- sum
}

func ConcurrentSum(a []int) int {
	n := len(a)
	ch := make(chan int)
	go sum(a[:n/2], ch)
	go sum(a[n/2:], ch)
	result := 0
	for i := 0; i < 2; i++ {
		// not sure how idiomatic this is
		select {
		case part_sum := <-ch:
			result += part_sum
		}
	}

	return result
}
