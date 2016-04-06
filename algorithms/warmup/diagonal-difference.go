package main
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
    "math"
)

func main() {
 
    a:= []string{}
    input := bufio.NewScanner(os.Stdin)
   
    for input.Scan() {
        a=append(a,input.Text())
        
    }
    b,_:=strconv.Atoi(a[0])
    dia_a:=0
    for i,v:= range a[1:] {
        z:=strings.Split(v," ")
        e,_:=strconv.Atoi(z[i])
        dia_a+= e
        
    }
    dia_b:=0
    for i,v:= range a[1:] {
        z:=strings.Split(v," ")
        e,_:=strconv.Atoi(z[(b-1)-i])
        dia_b+= e
        
    }
    fmt.Println(math.Abs(float64(dia_a-dia_b)))
  
    
}