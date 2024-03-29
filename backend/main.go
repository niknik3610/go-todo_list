package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type TodoHandler struct {
    todoItems map[string]TodoItem;
    finishedTodoItems map[string]TodoItem;
}

func main() {
    todoHandler := TodoHandler {
        todoItems: make(map[string]TodoItem),
        finishedTodoItems: make(map[string]TodoItem),
    };

    mux := http.NewServeMux();
    mux.HandleFunc("/", serveFile);
    mux.HandleFunc("/api/getTodoTasks", todoHandler.getTodoTasks);
    mux.HandleFunc("/api/newTodoTask", todoHandler.newTodoTask);
    mux.HandleFunc("/api/finishTodoTask", todoHandler.finishTodoTask);
    mux.HandleFunc("/api/getFinishedTodoTasks", todoHandler.getFinishedTodoTasks);
    err := http.ListenAndServe(":6969", mux);

    if err != nil {
        fmt.Printf("There was an error while running the server:\n%s", err.Error());
        return;
    }
}

const FRONTEND_ROOT  = "../frontend";
const INDEX_HTML_PATH = "src/index.html";

func serveFile(w http.ResponseWriter, req *http.Request) {
    path := FRONTEND_ROOT + req.URL.Path;

    if req.URL.Path == "/" {
        path += INDEX_HTML_PATH;
    }

    fmt.Println("Received a request for path " + path);
    http.ServeFile(w, req, path);
}

type TodoItem struct {
    TaskName string;
}
const TODO_ITEM_TEMPLATE = "templates/todo_item.html"
func (handle *TodoHandler) getTodoTasks(w http.ResponseWriter, req *http.Request) {
    tmpl, err := template.ParseFiles(TODO_ITEM_TEMPLATE);

    if err != nil {
        fmt.Printf("Error while parsing html template: %s", err.Error());
    }
    tmpl.Execute(w, handle.todoItems);
}

const FINISHED_TODO_ITEM_TEMPLATE = "templates/finished_todo_item.html"
func (handle *TodoHandler) getFinishedTodoTasks(w http.ResponseWriter, req *http.Request) {
    tmpl, err := template.ParseFiles(FINISHED_TODO_ITEM_TEMPLATE);

    if err != nil {
        fmt.Printf("Error while parsing html template: %s", err.Error());
    }
    tmpl.Execute(w, handle.finishedTodoItems);
}

func (handle *TodoHandler) newTodoTask(w http.ResponseWriter, req *http.Request) {
    todoItem := TodoItem {
        TaskName: req.FormValue("taskName"),
    }

    handle.todoItems[todoItem.TaskName] = todoItem;
    w.WriteHeader(http.StatusOK)
}

func (handle *TodoHandler) finishTodoTask(w http.ResponseWriter, req *http.Request) {
    itemName :=  req.FormValue("taskName");
    item, ok := handle.todoItems[itemName];
    if ok == false {
        w.WriteHeader(http.StatusBadRequest);
        io.WriteString(w, "Todo Item not Found");
        return;
    }

    delete(handle.todoItems, itemName);
    handle.finishedTodoItems[itemName] = item;
    w.WriteHeader(http.StatusOK);
}
