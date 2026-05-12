package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGOMAXPROCS(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(20) //cara menambah thread yang digunakan untuk menjalankan goroutine, defaultnya adalah jumlah cpu yang tersedia
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
	group.Wait()
}
