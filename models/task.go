package models

import (
	"fmt"
	"time"
	"github.com/astaxie/beego"
)

var DefaultTaskList *TaskManager

type Task struct {
	Id    int64  // Unique identifier
	Title string
	Body  string // Description
	Date  time.Time
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	// t, err := time.Parse("2006-01-02", "2011-01-19")
	// t.String(); t.Format("2006-01-02 15:04:05")
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	return &Task{0, title, "", time.Now()}, nil
}

// TaskManager manages a list of tasks in memory.
type TaskManager struct {
	tasks  []*Task
	lastId int64
}

// NewTaskManager returns an empty TaskManager.
func NewTaskManager() *TaskManager {
	return &TaskManager{
		[]*Task{
			{1, "Title", "", time.Now()},
			{2, "Title 1", "", time.Now()},
		},
		2,
	}
}

// Save saves the given Task in the TaskManager.
func (m *TaskManager) Save(task *Task) error {
	if task.Id == 0 {
		m.lastId++
		task.Id = m.lastId
		m.tasks = append(m.tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.tasks {
		if t.Id == task.Id {
			m.tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

// All returns the list of all the Tasks in the TaskManager.
func (m *TaskManager) All() []*Task {
	return m.tasks
}

// Find returns the Task with the given id in the TaskManager and a boolean
// indicating if the id was found.
func (m *TaskManager) Find(Id int64) (*Task, bool) {
	for _, t := range m.tasks {
		if t.Id == Id {
			return t, true
		}
	}
	return nil, false
}

// Delete Task by Id.
func (m *TaskManager) Delete(Id int64) {
	index := -1
	tasks := m.tasks
	for i, t := range tasks {
		if t.Id == Id {
			beego.Info("Record to delete ", t)
			index = i
		}
	}
	if index == -1 {
		return
	}
	tasks[len(tasks)-1], tasks[index] = tasks[index], tasks[len(tasks)-1]
	m.tasks = tasks[:len(tasks)-1]
}

func init() {
	//UserList = make(map[string]*User)
	//u := User{"user_11111", "astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com"}}
	//UserList["user_11111"] = &u

	DefaultTaskList = NewTaskManager()
}
