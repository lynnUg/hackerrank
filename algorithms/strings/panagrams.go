package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    input :=bufio.NewScanner(os.Stdin)
    for input.Scan() {
        test:=strings.ToLower(input.Text())
        start :="a"
        end :="z"
        alphabet:= map[string]int{}
        count:=0
        alpha:=0
        for i:=start[0];i<=end[0];i++ {
            if strings.Contains(test,string(i)) {
                if _,ok:=alphabet[string(i)];ok {
                    alphabet[string(i)]+=1
                } else {
                    alphabet[string(i)]=1
                    count+=1
                }
            }
            alpha+=1
        }
        
        if count==alpha{
            fmt.Println("pangram")
        } else {
             fmt.Println("not pangram")
        }
        
    }
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}