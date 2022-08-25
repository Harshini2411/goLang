package main
 import "fmt"
 func main(){
    movieName:="Queen"
    var firstName string
    var lastName string
    var cost int 
   bookings:=[]string{}//bookings=[]string{}
    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)
     fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)
     fmt.Println("Enter your cost of ticket:")
    fmt.Scan(&cost)
    bookings=append(bookings,firstName+" "+lastName)
    fmt.Printf("whole slice:%v\n",bookings)
    fmt.Printf("the type of the slice is:%T\n",bookings)
    fmt.Printf("length of slice:%v\n",len(bookings))
    fmt.Printf("Thank you %v for booking tickets for %v Movie of Rs.%v",bookings[0],movieName,cost)
    
 }
