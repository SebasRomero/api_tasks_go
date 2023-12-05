package tasks

import (
	"errors"

	"github.com/graphql-go/graphql"
)

func getTaskById(id string) (*Task, error) {
	for i, task := range Tasks {
		if task.ID == id {
			return &Tasks[i], nil
		}
	}

	return nil, errors.New("Task not found")
}

func removeById(tasks []Task, id string) []Task {
	index := -1
	for i, element := range tasks {
		if element.ID == id {
			index = i
			break
		}
	}

	if index != -1 {
		tasks = append(tasks[:index], tasks[index+1:]...)
	}

	return tasks
}

func GetTasks(params graphql.ResolveParams) (interface{}, error) {
	return Tasks, nil
}

func GetTask(params graphql.ResolveParams) (interface{}, error) {
	idQuery := params.Args["id"].(string)

	task, err := getTaskById(idQuery)

	if err != nil {
		return nil, nil
	}

	return task, nil

}

func AddTask(params graphql.ResolveParams) (interface{}, error) {
	task := Task{
		ID:          params.Args["id"].(string),
		Title:       params.Args["title"].(string),
		Description: params.Args["description"].(string),
		Completed:   params.Args["completed"].(bool),
	}

	Tasks = append(Tasks, task)
	return task, nil
}

func DeleteTask(params graphql.ResolveParams) (interface{}, error) {
	task, err := getTaskById(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	Tasks = removeById(Tasks, task.ID)
	return Tasks, nil
}

func UpdateTask(params graphql.ResolveParams) (interface{}, error) {
	task, err := getTaskById(params.Args["id"].(string))

	if err != nil {
		return nil, err
	}

	task.Completed = !task.Completed
	return task, nil
}
