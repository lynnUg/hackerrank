package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)
func insert(numbers []string, index int) []string {
    misplaced :=numbers[index]
    
    for i:=index-1;i>=-1;i-- {
         if i==-1 {
            numbers[0]= misplaced
            return numbers
         }
         z,_:=strconv.Atoi(string(numbers[i]))
         y,_:=strconv.Atoi(misplaced)
       
        
        
        if z>y {
            numbers[i+1]= numbers[i]
        }  else {
            
            numbers[i+1]= misplaced
            return numbers
            
            
        }
        
        
    }
    return numbers
}
func main() {
    reader :=bufio.NewReader(os.Stdin)
    text,_:= reader.ReadString('\n')
    text,_=reader.ReadString('\n')
    text = strings.TrimSpace(text)
    numbers :=strings.Split(text," ")
    
    for i:=1;i<len(numbers);i++{
        if numbers[i] < numbers[i-1] {
            numbers = insert(numbers,i)
        }
        fmt.Println(strings.Join(numbers," "))
    }
    
}