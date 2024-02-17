package loops

import "log"

func CountTo(countTo int) []int {
	result := make([]int, 0, countTo)
	for i := range countTo {
		result = append(result, i+1)
	}
	return result
}

func CountLetters(items []string) int {
	var result int

	for _, item := range items {
		result += len(item)
	}

	return result
}

func CountLettersWithConcurrency(done <-chan struct{}, items []string) int {
	workers := make([]chan int, 0, len(items))
	for _, word := range items {
		countStream := make(chan int)
		go func() {
			defer close(countStream)
			select {
			case <-done:
				return
			case countStream <- len(word):
				log.Println("word", word)
				return
			}
		}()
		workers = append(workers, countStream)
	}

	var result int

	for _, words := range workers {
		result += <-words
	}

	return result
}

func countLetters(done <-chan struct{}, word string) chan int {
	countStream := make(chan int)
	go func() {
		defer close(countStream)
		var count int
		for range word {
			count++
		}

		select {
		case <-done:
			return
		case countStream <- len(word):
			return
		}
	}()
	return countStream
}
