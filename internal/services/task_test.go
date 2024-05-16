package services

import (
	"testing"
)

func TestGetTask(t *testing.T) {
	id, err := taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	task, err := taskService.GetTask(TestUserId, id)

	if err != nil {
		t.Fatalf("Failed to get task: %v", err)
	}

	if task.Task != "Test task" {
		t.Fatalf("Expected task to be 'Test task', got '%s'", task.Task)
	}

	err = taskService.DeleteTask(TestUserId, id)

	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}
}

func TestGetTasks(t *testing.T) {
	_, err := taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	tasks, err := taskService.GetTasks(TestUserId)

	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task to be returned, got %d", len(tasks))
	}

	if tasks[0].Task != "Test task" {
		t.Fatalf("Expected task to be 'Test task', got '%s'", tasks[0].Task)
	}

	err = taskService.DeleteTask(TestUserId, tasks[0].ID)

	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	tasks, err := taskService.GetTasks(TestUserId)

	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Task != "Test task" {
		t.Fatalf("Expected task to be 'Test task', got '%s'", tasks[0].Task)
	}

	err = taskService.DeleteTask(TestUserId, tasks[0].ID)

	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}
}

func TestDeleteTask(t *testing.T) {
	id, err := taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	err = taskService.DeleteTask(TestUserId, id)

	if err != nil {
		t.Fatalf("Failed to delete task: %v", err)
	}
}

func TestDeleteTasks(t *testing.T) {
	_, err := taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	err = taskService.DeleteTasks(TestUserId)

	if err != nil {
		t.Fatalf("Failed to delete tasks: %v", err)
	}

	tasks, err := taskService.GetTasks(TestUserId)

	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	if len(tasks) != 0 {
		t.Fatalf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestGetTaskCount(t *testing.T) {
	count, err := taskService.GetTaskCount(TestUserId)

	if err != nil {
		t.Fatalf("Failed to get task count: %v", err)
	}

	if count != 0 {
		t.Fatalf("Expected 0 tasks, got %d", count)
	}

	_, err = taskService.CreateTask(TestUserId, "Test task")

	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	count, err = taskService.GetTaskCount(TestUserId)

	if err != nil {
		t.Fatalf("Failed to get task count: %v", err)
	}

	if count != 1 {
		t.Fatalf("Expected 1 task, got %d", count)
	}

	err = taskService.DeleteTasks(TestUserId)

	if err != nil {
		t.Fatalf("Failed to delete tasks: %v", err)
	}
}
