ppackage main
import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
)
func findLargest(n int) string {
   threes := 0
   fives := 0
   digits := n
   for digits > 0 {
        if (digits % 3 == 0) {
            fives = digits
            break
        }
        digits -= 5;
    }
    threes = n-fives;
    if digits < 0  {
        return "-1"
    }
    return strings.Repeat("5", fives) + strings.Repeat("3", threes)
    
}

func main() {
    input :=bufio.NewScanner(os.Stdin)
    i:=0
    for input.Scan() {
        if i>0{
            size,_:=strconv.Atoi(input.Text())
            fmt.Println(findLargest(size))
        }
         i++
    }
    
}