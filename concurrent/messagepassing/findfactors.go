package messagepassing

import "fmt"

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func Run5() {
	rCh := make(chan []int)
	go func() {
		rCh <- findFactors(3419110721)
	}()
	fmt.Println(findFactors(4033836233))
	fmt.Println(<-rCh)
}
