package app

import (
	"goWeb/todo/model"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var rd *render.Render = render.New()

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func getSessionID(r *http.Request) string {
	session, err := store.Get(r, "session")
	if err != nil {
		return ""
	}

	// Set some session values.
	val := session.Values["id"]
	if val == nil {
		return ""
	}
	return val.(string)
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := a.db.GetTodos()
	err := rd.JSON(w, http.StatusOK, list)
	if err != nil {
		return
	}
}

func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	todo := a.db.AddTodo(name)
	err := rd.JSON(w, http.StatusCreated, todo)
	if err != nil {
		return
	}
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	ok := a.db.RemoveTodo(id)
	if ok {
		err := rd.JSON(w, http.StatusOK, Success{true})
		if err != nil {
			return
		}
	} else {
		err := rd.JSON(w, http.StatusOK, Success{false})
		if err != nil {
			return
		}
	}
}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		err := rd.JSON(w, http.StatusOK, Success{true})
		if err != nil {
			return
		}
	} else {
		err := rd.JSON(w, http.StatusOK, Success{false})
		if err != nil {
			return
		}
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func CheckSignin(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if strings.Contains(r.URL.Path, "/signin.html") ||
		strings.Contains(r.URL.Path, "/auth") {
		next(w, r)
		return
	}

	sessionID := getSessionID(r)
	if sessionID != "" {
		next(w, r)
		return
	}
	http.Redirect(w, r, "/signin.html", http.StatusTemporaryRedirect)
}

func MakeHandler(filepath string) *AppHandler {
	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.HandlerFunc(CheckSignin),
		negroni.NewStatic(http.Dir("public")))
	n.UseHandler(r)

	a := &AppHandler{
		Handler: n,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/auth/google/login", googleLoginHandler)
	r.HandleFunc("/auth/google/callback", googleAuthCallback)
	r.HandleFunc("/", a.indexHandler)

	return a
}
