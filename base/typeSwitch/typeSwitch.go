package main

import "fmt"

func main(){

	whoami:= func(i interface{}){
switch t := i.(type){
case int :
	 fmt.Println("Integer:",t)
case string:
	fmt.Println("String:",t)
}
	}
	whoami("uoshhs")

}
