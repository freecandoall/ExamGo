// 참고
// https://brownbears.tistory.com/313	고루틴의 이해
// http://golang.site/	고 언어 기본 문법들
// https://subicura.com/2016/06/13/start-go-shipment-tracking-opensource.html 배송조회 서비스 프로젝트 소개글

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func add(a int, b int, c chan int) {
	c <- a + b
}

func add2(a int, b int, c *int) {
	*c = a + b
}

func main() {

	println("Hello World")

	var c = make(chan int)
	go add(1, 3, c)
	println(<-c)

	var cc int
	add2(2, 4, &cc)
	println(cc)

	runtime.GOMAXPROCS(runtime.NumCPU())
	// desc := "cpu : " + runtime.GOMAXPROCS(0)
	desc := fmt.Sprint("cpu : ", runtime.GOMAXPROCS(0))
	fmt.Println(desc)

	// 고루틴 끝날때까지 기다리기 예제
	TestGoroutine()
	fmt.Println("Finished call TestFor()")

	// 작은수 반환 예제 (고루틴 사용 전/후)
	fmt.Println(TestMin([]int{83, 46, 49, 23, 92, 48, 39, 91, 44, 99, 25, 42, 35, 56, 23}))
	fmt.Println(TestParallelMin([]int{83, 46, 49, 23, 92, 48, 39, 91, 44, 99, 25, 42, 35, 56, 23}, 4))
	fmt.Println("Finished call TestMin()")
}

func TestGoroutine() {
	var wg sync.WaitGroup

	s := "Hello, world!"
	anomyFuncCallCount := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			anomyFuncCallCount++
			fmt.Println(s, n, "anomy : ", anomyFuncCallCount)
			defer wg.Done()
		}(i)
		fmt.Println(s+" for", i)
	}

	wg.Wait()
}

func TestMin(a []int) int {

	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}

	return min
}

func TestParallelMin(a []int, n int) int {

	if len(a) < n {
		return TestMin(a)
	}

	mins := make([]int, n)
	size := (len(a) + n - 1) / n

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}

			mins[i] = TestMin(a[begin:end])
		}(i)
	}

	wg.Wait()
	return TestMin(mins)
}
