package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
 
    a:= []string{}
    input := bufio.NewScanner(os.Stdin)
    
    for input.Scan() {
        a=append(a,input.Text())
    }
    for i:=1 ;i<len(a)-1;i=i+2{
        b:= strings.Split(a[i]," ")
        thres,_:=strconv.Atoi(b[1])
        c:=strings.Split(a[i+1]," ")
        sum:=0
        for _,d:=range c{
            if d <="0"{
                sum+=1
            }
            
        }
        
        if sum>=thres{
            fmt.Println("NO")
        } else {
            fmt.Println("YES")
        }
        
        
    }
    
    
}