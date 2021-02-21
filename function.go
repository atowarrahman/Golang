package main

import "fmt"
/*
func add(x int,y int) int{
r :=x+y
return r
}
*/
//method-2
/*
func add(x, y int) int{
r :=x+y
return r
}
*/
//method-3
/*
func add(x, y int)(r int){
r=x+y
return r
}
*/
/*
func add(x, y int)(r int){
r=x+y
return
}
*/
func rectangle(l int,b int)(area int,parameter int){
	parameter=2*(l+b)
	area=l*b
	return
}

func update(a *int,t *string){
	*a=*a+5;
	*t=*t+"Doe";
	return
}
func main(){

//add(30,40)
a,p :=rectangle(10,30)
fmt.Println(a,p)
a:=10
t:="atowar"
update(&a,&t) 
//update(a *int,t *string)
}
