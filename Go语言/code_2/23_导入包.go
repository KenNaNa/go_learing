package main //必须

//忽略此包
import _ "fmt"

func main() {

}

/*
//给包名起别名
import io "fmt"

func main() {
	io.Println("this is a test")
}
*/

/*
//.操作
import . "fmt" //调用函数，无需通过包名
import . "os"

func main() {
	Println("this is a test")
	Println("os.Args = ", Args)
}
*/

/*
//方式1
//import "fmt" //导入包，必须使用，否则编译不过
//import "os"

//方法2，常用
import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("this is a test")
	fmt.Println("os.Args = ", os.Args)
}
*/
