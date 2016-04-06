package main
import (
    "fmt"
    "math"
    "bufio"
    "os"
    "strings"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    text,_:= reader.ReadString('\n')
    var length int
    var newText string 
    for i,_:=range text {
        if string(text[i])!= " " {
          newText += string(text[i])
          length +=1
        }
    }
   
    ceil :=math.Ceil(math.Sqrt(float64(length)))
    finalArray := []string{}
    
    for i,_:=range newText {
         
        if i<int(ceil){
           
            finalArray =append(finalArray,string(newText[i]))
        } else {
            finalArray[i%int(ceil)] += string(newText[i])
            
        }
    }
    
    fmt.Println(strings.Join(finalArray," "))
}