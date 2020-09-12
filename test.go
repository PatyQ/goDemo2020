package main
import (
	"fmt"
	"log"
	"net/http"
)

var h int = 9
func main(){
	// fmt.Println("Hello Go!")
	// fmt.Println("Hello \rGO!")

	// var s = "la"
	// var ss = "na";
	// fmt.Println(s + ss)

	// num := "嘎嘎"
	// fmt.Printf("num数据类型:%v哒哒",num)

	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}
//处理程序回显请求的URL r的路径部分
func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"URL.Path = %q\n",r.URL.Path)
	fmt.Printf("执行第二十六行%v\n",h)
}