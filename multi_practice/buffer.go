package main
import (
	"log"
	"runtime"
	"fmt"
)

func worker2(msg string) <-chan string {
	limit := make(chan int, 5)
	receiver := make(chan string)

	go func() {
		for i := 0; i< 100; i++ {
			log.Println(runtime.NumGoroutine())
			limit <- 1
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				<-limit
			}(i)
		}
	}()
	return receiver
}

func main() {
	receiver := worker2("job")
	for i := 0; i < 100; i++ {
		log.Println(<-receiver)
	}
}