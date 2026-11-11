package main

import ( "time" 
 "fmt" )

func Pingpongpattern(){

	first := make(chan struct{})
	second := make(chan struct{})

	
go  printFirstName(first,second)
go  printLastName(first,second)
	

	first <- struct{}{}

time.Sleep(10 * time.Millisecond)

fmt.Println("go routine successfully executed")

}



func printFirstName(first chan struct{},second chan struct{}){
for{
	<- first
	fmt.Println("Abhishek")
	second <- struct{}{}
}
}


func printLastName(first chan struct{},second chan struct{}){

	for{
	<- second
	fmt.Println("Rajput")
	first <- struct{}{}
	}
}