package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// На вход подаются два неупорядоченных слайса любой длины. Надо написать функцию, которая возвращает их пересечение
func twoSlices(slc1, slc2 []int) (map[int]int, error) {
	var mp map[int]int = make(map[int]int)

	var ptr1, ptr2 *[]int = nil, nil
	if len(slc1) > len(slc2) {
		ptr1 = &slc2
		ptr2 = &slc1
	} else {
		ptr1 = &slc1
		ptr2 = &slc2
	}

	for _, val := range *ptr1 {
		mp[val] = 0
	}

	for _, val := range *ptr2 {
		if _, ok := mp[val]; ok {
			mp[val]++
		}
	}

	for i := range mp {
		if mp[i] == 0 {
			delete(mp, i)
		}

	}

	return mp, nil
}

//Написать генератор случайных чисел
func randNum(n int) int {
	rn := rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- rn.Intn(n)
		}

	}()

	x := <-ch

	return x
}

// Слить N каналов в один

func clueChan(chs ...<-chan int) {
	mergedChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(len(chs))

	for _, ch := range chs {
		go func(ch <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for id := range ch {
				mergedChan <- id
			}
		}(ch, &wg)
	}
}

// Написать WorkerPool с заданной функцией
func WorkerPoolBiz(f func(int) int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- f(job)
	}

}

func WorkerPool() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	tasker := func(x int) int {
		return x * 10
	}

	for w := 0; w <= 5; w++ {
		go WorkerPoolBiz(tasker, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}

func Tasks() {
	a := []int{23, 3, 1, 2}
	b := []int{6, 2, 4, 23, 23, 7}
	fmt.Println(twoSlices(a, b))

	fmt.Println(randNum(10))

	WorkerPool()
}
