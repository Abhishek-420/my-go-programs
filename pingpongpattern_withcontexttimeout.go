package main

import ( "time" 
 "fmt" 
"sync"
"context")

func PingpongpatternWithContext(){

	ctx, cancel := context.WithTimeout(context.Background(),10 * time.Millisecond)

	defer cancel()

	 var wg *sync.WaitGroup
	first := make(chan struct{})
	second := make(chan struct{})

	wg.Add(2)
go  printFirstNameWithContext(first,second, ctx, wg)
go  printLastNameWithContext(first,second,ctx,  wg)
	
	first <- struct{}{}

	wg.Wait()


fmt.Println("go routine successfully executed")

}



func printFirstNameWithContext(first chan struct{},second chan struct{}, ctx context.Context, wg *sync.WaitGroup){


	defer wg.Done()
for{

	if ctx.Err()!= nil{
		return
	}

	<- first
	fmt.Println("Abhishek")
	second <- struct{}{}
}
}


func printLastNameWithContext(first chan struct{},second chan struct{}, ctx context.Context, wg *sync.WaitGroup){


	defer wg.Done()
	for{

				if ctx.Err()!= nil{
		return
	}

	<- second
	fmt.Println("Rajput")
	first <- struct{}{}
	}
}