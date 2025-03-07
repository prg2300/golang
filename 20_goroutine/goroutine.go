package main

import (
	"fmt"
	"sync"
)

//lightweight threads for multithreading

func task(id int, w *sync.WaitGroup) {
	//defer:func complete hone ke baad run hoga
	defer w.Done()
	fmt.Println("doing task", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg) //go func first go to schedular than block

		// 	func() {
		// 		fmt.Println(i)
		// 	}()
		// }

		//time.Sleep(time.Second * 2) // runing concrently no sequence
	}

	wg.Wait()
}
