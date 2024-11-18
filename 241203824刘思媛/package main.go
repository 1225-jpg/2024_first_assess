package main
import "fmt"
func main (){
    var computerchoice string
    var mychoice string
    Zhaoshi := map[int]string{0:"石头",1:"剪刀",2:"布"}
    fmt.Println("请输入剪刀/石头/布")
    fmt.Scan(&my)
    computerchoice = Zhaoshi[n]
    mychoice = Zhaoshi[my]
    if(computerchoice == "石头"&&mychoice=="布")
    fmt.Println("你赢了")
    if(computerchoice == "石头"&&mychoice=="剪刀")
    fmt.Println("你输了")
    if(computerchoice == "石头"&&mychoice=="石头")
    fmt.Println("平局")
}