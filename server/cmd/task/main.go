package main

import (
	"log"
	task "summer/server/shared/kitex_gen/task/taskservice"
)

func main() {
	svr := task.NewServer(new(TaskServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
