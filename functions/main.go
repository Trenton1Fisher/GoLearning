package main

import "fmt"

func sayGreeting(name string){
	fmt.Printf("Good Morning %s\n", name)
}

func sayBye(name string){
	fmt.Printf("Goodbye %s\n", name)
}

func cycleNames(names []string, f func(string)){
	for _, value := range names {
		f(value)
	}
}

func main(){
	sayGreeting("Trenton")
	sayBye("Trenton")	

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
}