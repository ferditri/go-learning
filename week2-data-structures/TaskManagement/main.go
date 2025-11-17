package main

import "fmt"

type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     string
	Priority    int    // 1-5 (1=tertinggi)
	Status      string // "todo", "in progress", "done"
	Category    string
}

// TODO: Buat methods untuk Task:
// 1. func (t *Task) MarkInProgress()
// 2. func (t *Task) MarkDone()
// 3. func (t *Task) UpdatePriority(newPriority int) error
// 4. func (t Task) IsOverdue(currentDate string) bool
// 5. func (t Task) GetTaskInfo() string

func (t *Task) MarkInProgress() {
	if t.Status != "in progress" {
		t.Status = "in progress"
	} else {
		fmt.Printf("Task '%s' is already in progress\n", t.Title)
	}
}

func (t *Task) MarkDone() {
	if t.Status != "done" {
		t.Status = "done"
	} else {
		fmt.Printf("Task '%s' is already done\n", t.Title)
	}
}
func (t *Task) UpdatePriority(newPriority int) error {
	if newPriority < 1 || newPriority > 5 {
		return fmt.Errorf("priority must be between 1 and 5")
	} else {
		t.Priority = newPriority
		return nil
	}
}

func (t Task) IsOverdue(currentDate string) bool {
	return t.DueDate < currentDate && t.Status != "done"
}

func (t Task) GetTaskInfo() string {
	return fmt.Sprintf("ID: %d\nTitle: %s\nDescription: %s\nDue Date: %s\nPriority: %d\nStatus: %s\nCategory: %s",
		t.ID, t.Title, t.Description, t.DueDate, t.Priority, t.Status, t.Category)
}

type TaskManager struct {
	Tasks []Task
}

// TODO: Buat methods untuk TaskManager:
// 1. func (tm *TaskManager) AddTask(task Task)
// 2. func (tm *TaskManager) RemoveTask(id int) error
// 3. func (tm TaskManager) FindByStatus(status string) []Task
// 4. func (tm TaskManager) FindByCategory(category string) []Task
// 5. func (tm TaskManager) GetHighPriorityTasks() []Task
// 6. func (tm TaskManager) GetTasksDueToday(today string) []Task
// 7. func (tm *TaskManager) UpdateTaskStatus(id int, status string) error

func (tm *TaskManager) AddTask(task Task) {
	tm.Tasks = append(tm.Tasks, task)
	fmt.Printf("Task '%s' added successfully\n", task.Title)
}

func (tm *TaskManager) RemoveTask(id int) error {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			fmt.Printf("Task with ID %d removed successfully\n", id)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func (tm TaskManager) FindByStatus(status string) []Task {
	var result []Task
	for _, task := range tm.Tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}
	return result
}

func (tm TaskManager) FindByCategory(category string) []Task {
	var result []Task
	for _, task := range tm.Tasks {
		if task.Category == category {
			result = append(result, task)
		}
	}
	return result
}

func (tm TaskManager) GetHighPriorityTasks() []Task {
	var result []Task
	for _, task := range tm.Tasks {
		if task.Priority == 1 {
			result = append(result, task)
		}
	}
	return result
}

func (tm TaskManager) GetTasksDueToday(today string) []Task {
	var result []Task
	for _, task := range tm.Tasks {
		if task.DueDate == today {
			result = append(result, task)
		}
	}
	return result
}

func (tm *TaskManager) UpdateTaskStatus(id int, status string) error {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i].Status = status
			fmt.Printf("Task with ID %d status updated to '%s'\n", id, status)
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}

func main() {
	// TODO: Test task management system
	taskManager := TaskManager{}

	task1 := Task{ID: 1, Title: "Complete project", Description: "Finish the Go project", DueDate: "2024-06-10", Priority: 1, Status: "todo", Category: "Work"}
	task2 := Task{ID: 2, Title: "Grocery shopping", Description: "Buy groceries for the week", DueDate: "2024-06-05", Priority: 3, Status: "todo", Category: "Personal"}
	taskManager.AddTask(task1)
	taskManager.AddTask(task2)

	fmt.Println(taskManager.FindByStatus("todo"))
	taskManager.UpdateTaskStatus(1, "in progress")
	fmt.Println(taskManager.FindByStatus("in progress"))

}
