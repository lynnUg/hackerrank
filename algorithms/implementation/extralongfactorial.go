package main
import (
    "fmt"
    "math/big"
    "bufio"
    "os"
    "strconv"
    "strings"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    input,_:= reader.ReadString('\n')
    value,_:=strconv.Atoi(strings.TrimSpace(input))
    bi :=big.NewInt(1)
    for i:=1;i<=value;i++ {
        temp := new(big.Int)
        bi = temp.Mul(bi,big.NewInt(int64 (i)))
    }
    fmt.Println(bi)
}