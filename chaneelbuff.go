package main
import "fmt"

func main() {  

 ch := make (chan string,1)

 ch <- "hello buffered"
 fmt.Println("data is sent to channel")

 msg:= <-ch

 fmt.Println("received :" , msg)



}
