package main
import "fmt"
import "reflect"
func main (){


var students[3]string
students[0]="Atowar"
students[1]="Rahman"
students[2]="Emon"
fmt.Println(students)
fmt.Println(len(students))
//data pull,data retrive,data get
fmt.Println(students[2])
//short hand way,string literals
students :=[...]string{"Atowar","Rahman","Emon"}
fmt.Println(students)
x :=students[0:2]
x :=make([]string,0)
var fruits []string
append(fruits,"apple","banana","mango")
fmt.Println(fruits,len(fruits))

fmt.Println(x)
fmt.Println(fruits)
fmt.Println("%T\n",fruits)
fmt.Println("%T",students)
a :=reflect.TypeOf(students).kind().string()
fmt.Println(a)


}