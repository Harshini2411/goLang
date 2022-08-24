//Arrays-->size is fixed to overcome that we use lists in other lang 
package main
 import "fmt"
 func main(){
    movieName:="Queen"
    var firstName string
    var lastName string
    var cost int 
    var bookings [50]string
    // bookings[0]=firstName+" "+lastName-->this must be at last after defining and declaring 
    fmt.Println("Enter your first name:")
    fmt.Scan(&firstName)
     fmt.Println("Enter your last name:")
    fmt.Scan(&lastName)
     fmt.Println("Enter your cost of ticket:")
    fmt.Scan(&cost)
    bookings[0]=firstName+" "+lastName
    fmt.Printf("whole array:%v\n",bookings)
    fmt.Printf("the type of the array is:%T\n",bookings)
    fmt.Printf("length of array:%v\n",len(bookings))
    fmt.Printf("Thank you %v for booking tickets for %v Movie of Rs.%v",bookings[0],movieName,cost)
    
 }


///where as in GO we use slices-->which is dynamic and abstraction of array i.e. resizable array,variable length,can also gets sub-array.It is also index-based

package main
 import "fmt"
 func main(){
    movieName:="Queen"
    var firstName string
    var lastName string
    var cost int 
    //var bookings []string
    bookings:=[]string{}
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
