package main
import "fmt"
   type Book struct{
    	Title string;
    	Author string;
    	ISBN string;
    	Price float32;
    	Pages int;
    	

    }
  
func main(){
   

  var b1 Book;
  b1.Title="Golang Book";
  b1.Author="Caleb doxy";
  b1.ISBN="12345";
  b1.Price=50.56;
  b1.Pages=166;

  fmt.Println(b1);
   fmt.Println(b1.Title);
    fmt.Println(b1.Author);
     fmt.Println(b1.ISBN);
      fmt.Println(b1.Price);
      fmt.Println(b1.Pages);

        b2 := struct{
    	color string
    	Name string
    	Age int
    	Roll string
        }{
        color:"red"
    	Name:"atowar"
    	Age:20
    	Roll:"19-41012-2"
        }
    fmt.Println(b2);
   fmt.Println(b2.color);
    fmt.Println(b2.Name);
     fmt.Println(b2.Age);
      fmt.Println(b2.Roll);
      //in computer science ,a literal is a notation for repesenting a fixed value in source code.
      //struct literals
      //var name string="atowar"
      var w1 weight
      w1=30.31
      fmt.Println(w1);
      

}