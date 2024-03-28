package main

import (
    "html/template"
	"fmt"
	"net/http"
)

type TodoHandler struct {
    todoItems []TodoItem;
}

func main() {
    mux := http.NewServeMux();
    mux.HandleFunc("/", serveFile);
    mux.HandleFunc("/api/getTodoTasks", getTodoTasks);
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
func getTodoTasks(w http.ResponseWriter, req *http.Request) {
    tmpl, err := template.ParseFiles(TODO_ITEM_TEMPLATE);

    todoItem1 := TodoItem {
        TaskName: "Test Task",
    };

    todoItem2 := TodoItem {
        TaskName: "Test Task2",
    };

    items := []TodoItem {
        todoItem1,
        todoItem2,
    }
    if err != nil {
        fmt.Printf("Error while parsing html template: %s", err.Error());
    }

    tmpl.Execute(w, items)
}

