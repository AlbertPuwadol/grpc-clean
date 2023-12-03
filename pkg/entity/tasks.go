package entity

type Entity struct {
	Start  int
	End    int
	Entity string
	Label  string
}

type Task1 struct {
	Text     string
	Entities []Entity
}

type TaskResult struct {
	Task1 Task1
	Task2 string
	Task3 int
}
