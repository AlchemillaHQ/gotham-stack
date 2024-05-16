package handlers

import (
	"fmt"
	"github.com/AlchemillaHQ/gotham-stack/internal"
	"github.com/AlchemillaHQ/gotham-stack/internal/i18n"
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"github.com/AlchemillaHQ/gotham-stack/internal/services"
	"github.com/AlchemillaHQ/gotham-stack/internal/utils"
	"net/http"
	"strconv"
)

func createTaskHTML(r *http.Request, id int, task string) string {
	html := fmt.Sprintf(`
		<li class="flex justify-between items-center p-2 bg-gray-100 rounded mt-2 shadow">
			<span class="text-gray-800"><b>%d.</b> %s</span>
			<button class="text-red-500 hover:text-red-700" hx-delete="/task/%d" hx-target="#todo-list" hx-swap="innerHTML">%s</button>
		</li>
	`, id, task, id, i18n.GetTranslation("dashboard", "delete", middleware.GetSessionLanguage(r)))

	return html
}

func TaskHandler(w http.ResponseWriter, r *http.Request, taskService *services.TaskService) {
	userId := internal.GetUserID(r)

	if userId == 0 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	switch r.Method {
	case "GET":
		handleGetTasks(w, r, taskService, userId)
	case "POST":
		handleCreateTask(w, r, taskService, userId)
	case "DELETE":
		handleDeleteTask(w, r, taskService, userId)
	}
}

func handleGetTasks(w http.ResponseWriter, r *http.Request, taskService *services.TaskService, userId int) {
	tasks, err := taskService.GetTasks(userId)
	if err != nil {
		http.Error(w, i18n.GetTranslation("dashboard", "get_all_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
		return
	}

	var html string

	for _, task := range tasks {
		html += createTaskHTML(r, task.ID, task.Task)
	}

	SendHTMLResponse(w, html)
}

func handleCreateTask(w http.ResponseWriter, r *http.Request, taskService *services.TaskService, userId int) {
	task := r.FormValue("task")

	if task == "" {
		http.Error(w, i18n.GetTranslation("dashboard", "no_empty_task", middleware.GetSessionLanguage(r)), http.StatusBadRequest)
		return
	}

	id, err := taskService.CreateTask(userId, task)

	if err != nil {
		http.Error(w, i18n.GetTranslation("dashboard", "task_create_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
		return
	}

	createdTask, err := taskService.GetTask(userId, id)

	if err != nil {
		http.Error(w, i18n.GetTranslation("dashboard", "get_one_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
		return
	}

	html := createTaskHTML(r, createdTask.ID, createdTask.Task)

	SendHTMLResponse(w, html)
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request, taskService *services.TaskService, userId int) {
	taskId := r.URL.Path[len("/task/"):]

	if taskId == "all" {
		count, err := taskService.GetTaskCount(userId)

		if err != nil {
			http.Error(w, i18n.GetTranslation("dashboard", "get_all_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
			return
		}

		if count == 0 {
			http.Error(w, i18n.GetTranslation("dashboard", "no_tasks_to_delete", middleware.GetSessionLanguage(r)), http.StatusBadRequest)
			return
		}

		err = taskService.DeleteTasks(userId)
		if err != nil {
			http.Error(w, i18n.GetTranslation("dashboard", "get_all_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
			return
		}

		w.Header().Set("HX-Trigger", utils.UIAlert("success", `Deleted `+strconv.Itoa(int(count))+` tasks`))

		SendHTMLResponse(w, "")
		return
	} else {
		id, err := strconv.Atoi(taskId)
		if err != nil {
			http.Error(w, i18n.GetTranslation("dashboard", "invalid_task_id", middleware.GetSessionLanguage(r)), http.StatusBadRequest)
			return
		}

		err = taskService.DeleteTask(userId, id)
		if err != nil {
			http.Error(w, i18n.GetTranslation("dashboard", "delete_one_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
			return
		}

		tasks, err := taskService.GetTasks(userId)

		if err != nil {
			http.Error(w, i18n.GetTranslation("dashboard", "get_all_failed", middleware.GetSessionLanguage(r)), http.StatusInternalServerError)
			return
		}

		var html string

		for _, task := range tasks {
			html += createTaskHTML(r, task.ID, task.Task)
		}

		w.Header().Set("HX-Trigger", utils.UIAlert("success", fmt.Sprintf(`Deleted Task %d`, id)))

		SendHTMLResponse(w, html)
	}
}
