package models

import (
	"time"
	_ "github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"reflect"
	"errors"
)

type Tasks map[uuid.UUID]*Task

type Task struct {
	Id    uuid.UUID
	Title string
	Body  string
	Date  time.Time
}

// Check if tasks equal, Id is NOT checked.
func (t Task) Equal(task Task) bool {
	if (t.Title == task.Title && t.Body == task.Body && t.Date == task.Date) {
		return true
	}
	return false
}

// NewTaskManager returns an empty Tasks.
func NewTaskManager() Tasks {
	return Tasks{}
}

// Validate tasks.
func validateTask(t Task) error {
	if (t.Id == uuid.Nil) {
		return errors.New("Nil Id.")
	}
	if (t.Title == "") {
		return errors.New("Empty title.")
	}
	return nil
}

// NewTask creates a new task given a title, that can't be empty.
func NewTask(title string, body string) (Task, error) {
	// t, err := time.Parse("2006-01-02", "2011-01-19")
	// t.String(); t.Format("2006-01-02 15:04:05")
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	task := Task{uuid.NewV4(), title, body, time.Now()}

	err := validateTask(task);
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// Save saves the given Task in the TaskManager.
func (t Tasks) Save(task Task) (uuid.UUID, error) {
	err := validateTask(task);
	if err != nil {
		return uuid.Nil, err
	}

	task.Id = uuid.NewV4();
	t[task.Id] = &task
	return task.Id, nil
}

// Add a task with existing Id.
func (t Tasks) Add(task Task) error {
	err := validateTask(task);
	if err != nil {
		return err
	}

	t[task.Id] = &task

	return nil
}

// All returns the list of all the Tasks.
func (t Tasks) All() []*Task {
	v := []*Task{}
	keys := reflect.ValueOf(t).MapKeys()

	for _, id := range keys {
		uid, _ := id.Interface().(uuid.UUID)
		v = append(v, t[uid])
	}
	return v
}

// Find returns the Task with the given id in the TaskManager.
func (t Tasks) Find(Id uuid.UUID) (*Task, bool) {
	if t, ok := t[Id]; ok {
		return t, true
	}
	return nil, false
}

// Update task.
func (t Tasks) Update(uid uuid.UUID, task Task) (*Task, error) {
	if err := validateTask(task); err != nil {
		return nil, err
	}

	if _, ok := t[uid]; ok {
		// Save the id.
		task.Id = uid
		t[uid] = &task
		return t[uid], nil
	}
	return nil, errors.New("Task not exist.")
}

// Delete Task by Id.
func (t Tasks) Delete(Id uuid.UUID) {
	delete(t, Id)
}
