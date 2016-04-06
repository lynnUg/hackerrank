package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    
)

func hasPattern(grid ,pattern []string) string {
    has := "NO"
    for row,_:=range grid {
        found:= strings.Contains(grid[row],pattern[0])
        if  found {
            has = "YES"
     
            for i:=1;i<len(pattern);i++ {
                
                found2 := strings.Contains(grid[row+i],pattern[i])
                index  := strings.Index(grid[row+i],pattern[i])
                subString :=grid[(row+i)-1]
                if found2 {
                   subString = subString[index:index+len(pattern[i])]
                }
                
                if found2==false || subString!=pattern[i-1] {
                    
                    has ="NO"
                    break
                }
                 
            }
            
            if has =="YES" {
                break
            }
            
        }
    }
 return has
}
func readMatrix(reader *bufio.Reader,numberOfRows,numberOfColumns int)  []string {
 var matrix []string
 for i:=0 ; i < numberOfRows; i++ {
    line,_ := reader.ReadString('\n')
    matrix=append(matrix,strings.TrimSpace(line))
 }
 return matrix
}
func main() {
    var numberOfTestCases int
    reader :=bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')
    fmt.Sscanf(line, "%d", &numberOfTestCases)

    var testCases []string
    for i:=0; i< numberOfTestCases ; i++ {
        var numberOfRows ,numberOfColumns int

        line,_ := reader.ReadString('\n')
        fmt.Sscanf(line,"%d %d",&numberOfRows,&numberOfColumns)
        grid :=readMatrix(reader,numberOfRows,numberOfColumns)
        
        line,_ = reader.ReadString('\n')
        fmt.Sscanf(line,"%d %d",&numberOfRows,&numberOfColumns)
        pattern :=readMatrix(reader,numberOfRows,numberOfColumns)
        
        testCases=append(testCases,hasPattern(grid,pattern))
    }
    
    for _,i := range testCases {
        fmt.Println(i)
    }


}