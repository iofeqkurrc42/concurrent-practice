package messagepassing

import "sync"

func FanIn[K any](quit <-chan int, allChannels ...<-chan K) chan K {
	wg := sync.WaitGroup{}
	wg.Add(len(allChannels))
	output := make(chan K)
	for _, c := range allChannels {
		go func(channel <-chan K) {
			defer wg.Done()
			for i := range channel {
				select {
				case output <- i:
				case <-quit:
					return
				}
			}
		}(c)
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}

func CreateAll[K any](n int) []chan K {
	channels := make([]chan K, n)
	for i := range channels {
		channels[i] = make(chan K)
	}
	return channels
}

func CloseAll[K any](channels ...chan K) {
	for _, o := range channels {
		close(o)
	}
}

func Broadcast[K any](quit <-chan int, input <-chan K, n int) []chan K {
	outputs := CreateAll[K](n)
	go func() {
		defer CloseAll(outputs...)
		var msg K
		moreData := true
		for moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					for _, output := range outputs {
						output <- msg
					}
				}
			case <-quit:
				return
			}
		}
	}()
	return outputs
}

func Take[K any](quit chan int, n int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		moreData := true
		var msg K
		for n > 0 && moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					output <- msg
					n--
				}
			case <-quit:
				return
			}
			if n == 0 {
				close(quit)
			}
		}
	}()
	return output
}
