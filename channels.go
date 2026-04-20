package main
import "fmt"

func main() {

	ch:= make (chan string)

	go func() {

		ch <- "hello Devops Engineer" 
	}	()
        

	msg := <-ch  
	fmt.Println(msg)

}



