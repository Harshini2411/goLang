package main
import "fmt"
func main(){
  var availableTickets=50
  const leftTickets=30
  var conferenceName="Go Conference"
  fmt.Println("Welcome to ",conferenceName, "there are",availableTickets,"availableTickets and ",leftTickets,"Tickets are left.")
  fmt.Printf("Welcome to %v there are %v available tickets and  %v tickets are left.\n",conferenceName,availableTickets,leftTickets)
}
//-->Datatypes<---\\
package main
import "fmt"
func main(){
  var num int=20
  var name string="Harshini"
  //also for var-->name:="Harshini"
  //its not for const-->num:=20[doesn't work]
  fmt.Println(&num)//prints the address of that variable
  fmt.Scan(&num)//takes input from user
  fmt.Printf("My name is %v and number of chocolates in my hand are %v",name,num)
}
///--------->Simple program using Print and scan stmts<------------\\
package main
import "fmt"
func main(){
    var FirstName string
    var LastName string
    var Email string
    var Tickets int
    fmt.Scan(&Tickets)
    fmt.Println("Enter your firstName:")
    fmt.Scan(&FirstName)
    fmt.Println("Enter your lastName:")
    fmt.Scan(&LastName)
    fmt.Println("Enter your email:")
    fmt.Scan(&Email)
    fmt.Printf("Thank you %v %v for booking %v at once.You will get an confirmation email to %v.\n",FirstName,LastName,Tickets,Email)
}
