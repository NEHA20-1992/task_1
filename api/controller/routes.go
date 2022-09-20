package controller
import(
	"github.com/NEHA20-1992/task_1/api/middleware"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/users", middleware.JSONResponder(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middleware.JSONResponder(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middleware.JSONResponder(s.GetUser)).Methods("GET")
	 s.Router.HandleFunc("/users/{id}", middleware.JSONResponder(s.UpdateUser)).Methods("PUT")
	 s.Router.HandleFunc("/users/{id}", middleware.JSONResponder(s.DeleteUser)).Methods("DELETE")


}
