package main

import (
    "fmt"
    "net"
)

func main() {
    message := "Привет, я сервер сообщение"   // отправляемое сообщение
    listener, err := net.Listen("tcp", ":9876")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer listener.Close()

    fmt.Println("Сервер ждет подключений...")
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }
		conn.Write([]byte(message))
		fmt.Println("Cообщение отправлено: %s", message)
        conn.Close()
    }
}
