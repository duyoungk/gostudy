package main

import (
    "fmt"
)


func main() {
    fmt.Println("main start")
    
    connstr := "server=175.207.11.31;user id=sa;password=Tlemskdls4096;database=SW_USER001;port=1433"
    
    query := NewQuery()
    query.Open(1, connstr)
    defer query.Close()
    
    // results := query.Query("select top 1 * from tb_user with(nolock)")
    
    // if results != nil {
    //     begin := results.Front()
    //     for begin != nil {
    //         r := begin.Value.(map[string]interface{})
    //         for k, v := range r {
    //             fmt.Printf("%s:%v, ", k, v)
    //         }
    //         fmt.Println()
    //         begin = begin.Next()
    //     }
    // }
    
    results := query.Proc("p_test", 123, 456)
    if results != nil {
        begin := results.Front()
        for begin != nil {
            row := begin.Value.(map[string]interface{})
            for k, v := range row {
                fmt.Printf("%s:%v\t", k, v)
            }
            fmt.Println()
            begin = begin.Next()
        }
    }
    

    fmt.Println("main end")
}