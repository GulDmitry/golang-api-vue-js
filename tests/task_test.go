package tests

import (
	"testing"
	. "github.com/guldmitry/go-api-vue-js/models"
)

func newTaskOrFatal(t *testing.T, title string, body string) Task {
	task, err := NewTask(title, body)
	if err != nil {
		t.Fatalf("New task: %v.", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	title := "Learn Go"
	body := "Body"
	task := newTaskOrFatal(t, title, body)
	if task.Title != title || task.Body != body {
		t.Errorf("Expected title %s, body %s, got %s, %s.", title, body, task.Title, task.Body)
	}
}

func TestNewTaskEmptyTitle(t *testing.T) {
	_, err := NewTask("", "")
	if err == nil {
		t.Error("Expected 'empty title' error, got nil.")
	}
}

func TestSaveTaskAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go", "Body")

	m := NewTaskManager()
	_, err := m.Save(task)
	if err != nil {
		t.Error("Unable to save a task.")
	}

	all := m.All()
	if len(all) != 1 {
		t.Errorf("Expected 1 task, got %v.", len(all))
	}
	if all[0].Equal(task) == false {
		t.Errorf("Expected %v, got %v.", task, all[0])
	}
}

func TestAddTaskAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go", "Body")

	m := NewTaskManager()
	m.Add(task)

	all := m.All()
	if len(all) != 1 {
		t.Errorf("Expected 1 task, got %v.", len(all))
	}
	if all[0].Equal(task) == false || all[0].Id != task.Id {
		t.Errorf("Expected %v, got %v.", task, all[0])
	}
	if _, ok := m[task.Id]; !ok {
		t.Errorf("Expected to find task: %v.", task)
	}
}

func TestSaveAndRetrieveTwoTasks(t *testing.T) {
	learnGo := newTaskOrFatal(t, "Learn Go", "")
	learnTDD := newTaskOrFatal(t, "Learn TDD", "")

	m := NewTaskManager()
	learnGoId, _ := m.Save(learnGo)
	learnTDDId, _ := m.Save(learnTDD)

	all := m.All()
	if len(all) != 2 {
		t.Errorf("Expected 2 tasks, got %v.", len(all))
	}
	if _, ok := m[learnGoId]; !ok {
		t.Errorf("Expected to find task: %v.", learnGo)
	}
	if _, ok := m[learnTDDId]; !ok {
		t.Errorf("Expected to find task: %v.", learnTDD)
	}
}

func TestSaveModifyAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "Learn Go", "")
	updTask := newTaskOrFatal(t, "Learn Go upd", "Body")

	m := NewTaskManager()
	id, _ := m.Save(task)

	m.Update(id, updTask)

	actualTask := m[id]
	if actualTask.Title != updTask.Title || actualTask.Body != updTask.Body {
		t.Error("Task wasn't updated.", actualTask, updTask)
	}
}

func TestDelete(t *testing.T) {
	m := NewTaskManager()
	id1, _ := m.Save(newTaskOrFatal(t, "Title 1", ""))
	id2, _ := m.Save(newTaskOrFatal(t, "Title 2", ""))

	m.Delete(id1)

	if len(m.All()) != 1 {
		t.Error("Expected 1 tasks.")
	}

	m.Delete(id2)

	if len(m.All()) != 0 {
		t.Error("Expected 0 tasks.")
	}
}

func TestSaveAndFind(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go", "")
	m := NewTaskManager()
	id, _ := m.Save(task)

	actualTask, ok := m.Find(id)
	if !ok {
		t.Error("didn't find task")
	}
	if !actualTask.Equal(task) {
		t.Errorf("expected %v, got %v", task, actualTask)
	}
}
