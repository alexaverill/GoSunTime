package main
import (
	"fmt"
	"net/http"
)
func test(w http.ResponseWriter, req *http.Request){
	fmt.Fprintf(w,"Test\n");
}
func main(){
	fmt.Println("Test");
	http.HandleFunc("/test",test);
	http.ListenAndServe(":8090",nil);
}