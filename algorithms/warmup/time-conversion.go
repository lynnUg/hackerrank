package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)
func convertTime (input string) string {
    start:=input[:2]
    end:=input[len(input)-2:]
    if end == "PM" {
            b,_:=strconv.Atoi(start)
            b = ( b%12 )+12
            start=fmt.Sprintf("%02d",b)
            
            
    } else {
            b,_:=strconv.Atoi(start)
            b = b%12 
            start=fmt.Sprintf("%02d",b)
            
    }
    return start+input[2:len(input)-2]
    
}

func main() {
 
    input := bufio.NewScanner(os.Stdin)
    
    for input.Scan() {
        a:=input.Text()
        
        fmt.Println( convertTime(a))
        
    }
}