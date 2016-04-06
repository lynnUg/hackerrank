package main
import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
    
)

func sweets(money int, costperc int, wrapperperc int ,numwrapper int) int {
    
    if money/costperc==0 && numwrapper/wrapperperc==0 {
        return 0
    }
    balanceMoney:=(money%costperc)
    balanceWrappers:=(numwrapper%wrapperperc)+(money/costperc)+(numwrapper/wrapperperc)
    return money/costperc + numwrapper/wrapperperc+ sweets(balanceMoney,costperc,wrapperperc, balanceWrappers)
}

func main() {
    input :=bufio.NewScanner(os.Stdin)
    i:=0
    for input.Scan() {
        
        if i>0{
            
            a:=strings.Split(input.Text()," ")
                 
                x,_:=strconv.Atoi(a[0])
                 y,_:=strconv.Atoi(a[1])
                 z,_:=strconv.Atoi(a[2])
                 fmt.Println(sweets(x,y,z,0))
          
          }
         i++
    }
}