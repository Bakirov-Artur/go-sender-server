package main

import (
    "fmt"
    "net"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Message struct {
	gorm.Model
	Content string
	RemoteAddr string
}

func main() {
	flog, logger := LogInit("main")
	defer flog.Close()
    listener, err := net.Listen("tcp", ":9876")
    if err != nil {
        fmt.Println(err)
        return
    }

    defer listener.Close()

	//Connect to data storage
	db, err := gorm.Open("sqlite3", "all-msg.db")
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}
	defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&Message{})
	// Start shell
	go ShellRun(db)
    logger.Println("Сервер ждет подключений...")
    for {
        conn, err := listener.Accept()
		if err != nil {
            logger.Panicln(err)
			logger.Printf("%s -> Отключен\n", conn.RemoteAddr().String())
            conn.Close()
			continue
        }
		logger.Printf("%s <- Подключен\n", conn.RemoteAddr().String())
		go Communicates(conn, db)
    }
}

func Communicates(conn net.Conn, db *gorm.DB) {
	defer conn.Close()
	flog, logger := LogInit("Communicates")
	defer flog.Close()
	remoteAddr := conn.RemoteAddr().String()
	for {
		// Считываем полученные в запросе данные
        inputBuff := make([]byte, (1024 * 4))
		conn.SetDeadline(time.Now().Add(time.Second*30))
		req, err := conn.Read(inputBuff)
        if req == 0 || err != nil {
            logger.Println("Read error:", err)
			logger.Printf("%s <- Соединение закрыта\n", remoteAddr)
            break
		}else {
			logger.Printf("%s <- Данные получены\n", remoteAddr)
			//Создаю новую запись в базе данных
			db.Create(&Message{Content: string(inputBuff), RemoteAddr: remoteAddr})
			logger.Printf("%s <- Данные сохранен\n", remoteAddr)
		}
		fmt.Println()
	}
}
