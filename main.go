package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"projects/task-cli/helper"
	"projects/task-cli/types"
	"strconv"
	"strings"
)

func AddTask(task types.Task) {
    tasks := helper.JsonToTasks()
    tasks.Tasks = append(tasks.Tasks, task)
    helper.TasksToJson(tasks)
    fmt.Println("Taks added successfully")
}

func DeleteTask (taskId int) {
    tasks := helper.JsonToTasks()
    var newTasks types.Tasks
    for i := 0; i < len(tasks.Tasks); i++{
        if tasks.Tasks[i].Id == taskId {
            newTasks.Tasks = append(tasks.Tasks[:i], tasks.Tasks[i + 1:]...)
            helper.TasksToJson(newTasks)
            fmt.Println("Task deleted successfully")
            return
        }
    }
    fmt.Printf("Couldn't find task with Id: %d", taskId)
}

func DeleteAllTasks (){
    var tasks types.Tasks
    helper.TasksToJson(tasks)
}

func ListTasks() {
    tasks := helper.JsonToTasks() 
    if len(tasks.Tasks) == 0 {
        println("No tasks registered")
    }else {
        fmt.Print("\n")
        for i := 0; i < len(tasks.Tasks); i++ {
            curr := tasks.Tasks[i]
            fmt.Printf("(%d) Task description: \"%s\", Status: %s\n", curr.Id, curr.Description, curr.Status)
        }
    }

}

func UpdateTask(taskId int, description string){
    tasks := helper.JsonToTasks()
    if taskId < 0 || taskId >= len(tasks.Tasks) {
        error := errors.New("Invalid task id")
        log.Fatal(error) 
    }

    for i := 0; i < len(tasks.Tasks); i++{
        if tasks.Tasks[i].Id == taskId {
            tasks.Tasks[i].Description =  description
            helper.TasksToJson(tasks)
            fmt.Println("Task description updated successfully")
            return
        } 
    }

    println("Could not update task")
}

func MarkTask(taskId int, status string) {
    tasks := helper.JsonToTasks()
    for i:=0; i < len(tasks.Tasks); i++{
        if tasks.Tasks[i].Id == taskId {
            tasks.Tasks[i].Status = status 
            helper.TasksToJson(tasks)
            fmt.Println("Task marked successfully")
            return
        } 
    }

    println("Could not mark task, check if task id is valid")
}

func main(){
    args := os.Args
    if len(args) < 2 {
        fmt.Println("Welcome to task-cli")
    }else {
        actionArr := strings.Split(args[1], "-")
        switch actionArr[0] {
        case "add":
            if len(args) != 3 {
                error := errors.New("Invalid command")
                log.Fatal(error)
            }
            newTask := types.Task{Id: len(helper.JsonToTasks().Tasks), Description: args[2], Status: "todo"} 
            AddTask(newTask)
        case "list":
            if len(args) != 2 {
                error := errors.New("Invalid command")
                log.Fatal(error)
            }
            ListTasks()
        case "delete":
            if len(args) != 3 {
                error := errors.New("Invalid command")
                log.Fatal(error)
            } 

            taskId, err := strconv.Atoi(args[2])
            if err != nil {
                log.Fatal(err)
            }

            DeleteTask(taskId)        
        case "update":
            if len(args) != 4 {
                error := errors.New("Invalid command")
                log.Fatal(error)
            } 

            taskId,err := strconv.Atoi(args[2])
            if err != nil{
                log.Fatal(err)
            }
            
            UpdateTask(taskId, args[3])

        case "mark":

            if len(args) != 3 {
                error := errors.New("Invalid command")
                log.Fatal(error)
            } 
            var status string
            for i:=1; i < len(actionArr); i++{ 
                status = status + " " + actionArr[i]
            }
            
            taskId,err := strconv.Atoi(args[2])
            if err != nil{
                log.Fatal(err)
            }

            MarkTask(taskId, status)
        }
    }
}
