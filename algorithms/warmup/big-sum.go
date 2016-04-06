package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
 
    a:= [2]string{}
    input := bufio.NewScanner(os.Stdin)
    i:=0
    for input.Scan() {
        a[i]=input.Text()
        i++
    }
    b:= strings.Split(a[1]," ")
    var sum uint64
    
    for _,i:=range b {
        z,_:=strconv.ParseUint(i,10,64)
        sum+=z
    }
    fmt.Println(sum)
    
}