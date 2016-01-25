package main
import (
	"log"
	"time"
	"runtime"
)

func main() {

	log.Println(runtime.NumGoroutine())
	go func() {
		log.Println("end")
	}()
	go func() {
		log.Println("return")
		return
	}()
	go func() {
		log.Println("exit")
		runtime.Goexit()
	}()

	log.Println(runtime.NumGoroutine())
	time.Sleep(time.Second)
}