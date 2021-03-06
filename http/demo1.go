package main
import (
    "fmt"
    "net/http"
    "log"
)

func handleHello(w http.ResponseWriter,req *http.Request){
    //parse param
    req.ParseForm()
    fmt.Println(req)
    for k,v := range req.Form{
        fmt.Println("key:",k)
        fmt.Println("value:",v)
    }
    fmt.Fprintf(w,"Hello World")
}

func main(){

    //setting static dir
    http.Handle("/static/",http.FileServer(http.Dir("staticDir")))

    http.HandleFunc("/",handleHello)
    err := http.ListenAndServe(":12345", nil)
    if err!= nil{
        log.Fatal("ListenAndServen",err)
    }
}
