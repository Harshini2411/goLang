//infinite for-loop->to break/interrupt this app use ctrl+C
package main
 import "fmt"
 func main(){
    movieName:="Queen"
    for{
    var firstName string
    var lastName string
    var cost int 
     bookings :=[]string{}
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
    fmt.Printf("Thank you %v for booking tickets for %v Movie of Rs.%v\n",bookings,movieName,cost)
    }
    
 }
//FOR-EACH LOOP--loop through a specific list of elements
