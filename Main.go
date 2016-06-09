package main

import (
    "fmt"
)


func main() {
    fmt.Println("main start")
    
    connstr := "server=localhost;user id=sa;password=Tjqj;database=TM_USER001;port=31051"
    
    query := &Query{}
    query.Open(1, connstr)
    defer query.Close()
    
    query.Query("select top 10 * from tb_user_dtl")
    
    fmt.Println("main end")
}