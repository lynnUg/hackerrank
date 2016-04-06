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
    sum, _ := strconv.ParseFloat(a[0], 64)
    
    pos:=0
    neg:=0
    zero:=0
    for _,i:=range b {
        z,_:=strconv.Atoi(i)
        if z==0 {
            zero+=1
        } else if z <0{
            neg+=1
        }else{
            pos+=1
        }
        
    }
    fmt.Printf("%.6f \n",float64(pos)/sum)
    fmt.Printf("%.6f \n",float64(neg)/sum)
    fmt.Printf("%.6f \n",float64(zero)/sum)
    
    
}