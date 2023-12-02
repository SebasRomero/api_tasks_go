package tasks

type Task struct {
	ID          string `json: "id"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Completed   bool   `json: "completed"`
}

var Tasks = []Task{
	{ID: "1", Title: "Clean Room", Description: "I have to clean my room", Completed: false},
	{ID: "2", Title: "Clean Bathroom", Description: "I have to clean the bathroom", Completed: false},
	{ID: "3", Title: "Study Go", Description: "I have to study", Completed: false},
}
