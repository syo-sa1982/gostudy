package main
import (
	"log"
)


func main() {
	fin := make(chan bool)
	go func(){
		log.Println("worker working.. ")
		fin <- false
	}()
	<-fin
}