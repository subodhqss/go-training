package main
import("net/http"
       "log"
)

func main(){
http.HandleFunc("/login", Login)
http.HandleFunc("/home", Home)
log.Fatal(http.ListenAndServe(":8500", nil))
}