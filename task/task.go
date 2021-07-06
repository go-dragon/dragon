package task

import (
	"fmt"
	"time"
)

func Run() {
	TestTask()
}

func TestTask() {
	ticker := time.Tick(time.Second)
	for tk := range ticker {
		fmt.Println(tk.Unix())
	}
}
