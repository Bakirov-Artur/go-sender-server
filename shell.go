package main

import (
    "fmt"
	"github.com/jinzhu/gorm"
)

func ShellRun(db *gorm.DB) {
	fmt.Println( "Командная оболочка сервера" )
	fmt.Println( "Для получения справки введите: help или ?" )
	fmt.Println()
	cmds:= make(chan string, 1)
	for {
		inCmd(cmds)
		runCmd(cmds, db)
	}
}

func inCmd(cmd chan<- string) {
	var bufCmds string
	fmt.Print("Введите команду: ")
	fmt.Scanln(&bufCmds)
	fmt.Printf("\n")
	cmd <- bufCmds
}

func runCmd(cmd <-chan string, db *gorm.DB) {
	bufCmds := <-cmd

	if bufCmds == "print"{
		var msgReconds []Message
		db.Select("created_at, remote_addr").Find(&msgReconds)
		//Вывод полученных данных
		for _, msg := range msgReconds {
			fmt.Printf("--------------------------------------\n", )
			fmt.Printf("ip-адрес:        %s\n", msg.RemoteAddr)
			fmt.Printf("время получения: %s\n", msg.CreatedAt)
		}
		fmt.Printf("\n")
	} else if bufCmds == "help" || bufCmds == "?"{
		fmt.Printf("print - ввывод всех записей\n")
		fmt.Printf("help - получить справку команд\n")
	} else {
		fmt.Printf("команда не найдена: %s\n", bufCmds)
	}
}

