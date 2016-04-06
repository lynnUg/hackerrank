package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader :=bufio.NewReader(os.Stdin)
    text,_:= reader.ReadString('\n')
    text,_=reader.ReadString('\n')
    text = strings.TrimSpace(text)
    numbers :=strings.Split(text," ")
    misplaced :=numbers[len(numbers)-1]
    for i:=len(numbers)-2;i>=-1;i-- {
        if i==-1 {
            numbers[0]= misplaced
        } else if numbers[i] > misplaced {
            numbers[i+1]= numbers[i]
        }  else {
            numbers[i+1]= misplaced
            fmt.Println(strings.Join(numbers," "))
            break
        }
        fmt.Println(strings.Join(numbers," "))
    }
}