package main
import "fmt"

type Person struct {

      Name string
      Age int 
      }


      func Person123(p Person){

	      fmt.Println("Name" ,p.Name , "Age:", p.Age)
              }

	 func main () {


		 p1 := Person{

			 Name: "HItesh",
			 Age : 25,
		 }

		 p2 := Person{}
		 p2.Name = "Rahul"
		 p2.Age =30

		 fmt.Println(p1.Name)
		 fmt.Println(p1.Age)

                Person123(p1)
		Person123(p2)

	}


