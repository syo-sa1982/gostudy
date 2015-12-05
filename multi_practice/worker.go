package main
import (
	"log"
	"fmt"
)

func worker1(msg string) <-chan string {
	receiver := make(chan string)
	for i := 0; i < 4; i++ {
		go func(i int) {
			msg := fmt.Sprintf("%d %s done", i, msg)
			receiver <- msg
		}(i)
	}
	return  receiver
}

func main() {
	receiver := worker1("job")
	for i := 0; i < 4; i++ {
		log.Println(<-receiver)
	}
}