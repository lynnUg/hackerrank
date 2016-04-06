package main
import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "sort"
    "strings"
    
)
func removeZeros(trees []int) []int {
    i:=0
    for {
        
        if i>len(trees)-1 {
            break
        }
        if trees[i]==0{
            
            trees = append(trees[:i], trees[i+1:]...)
            
        } else {
            i++
        }
    }
    return trees
}


func cut(trees []int, z *int) []int{
    
    *z = len(trees)
     x := trees[0]
    for i:=0 ; i<len(trees) ; i++{
        trees[i]=trees[i]-x
    }
    return trees
}

func cutprocess(trees []int) {
    i:=0
    
    for len(trees)>=1 {
        
        trees=cut(trees,&i)
        fmt.Println(i)
        trees=removeZeros(trees)
        
        
    }
    
}

func main() {
    input :=bufio.NewScanner(os.Stdin)
    i:=0
    trees:=[]int{}
    for input.Scan() {
        
        if i>0{
            
            a:=strings.Split(input.Text()," ")
            for _,i:=range a{
                z,_:=strconv.Atoi(i)
                trees=append(trees,z)
            }
            sort.Ints(trees)
            cutprocess(trees)
          
           
            
        } 
         i++
    }
   
    
    
 
}