package main
import ("fmt"
        "strings"
        "bufio"
        "os"
        "strconv"
       )
func spaces (n int) (h []string){
for i:=0 ;i< n ; i++ {
  h=append(h, " ")
}
return
}
func main() {
reader := bufio.NewReader(os.Stdin)
//text, _ := reader.ReadString('\n')
var a string
for input.Scan() {
        a=input.Text()
        i++
    }
n,_:=strconv.Atoi(a)
h:=spaces (n)
for i:=n-1 ; i>=0 ;i-- {
  h[i]="#"
  fmt.Println(strings.Join(h,""))
 }
}