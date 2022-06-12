package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtnhnd255/payment_service_emulator/app/storage"
)

type Server struct {
	Config  *Config
	Router  *mux.Router
	Storage *storage.Storage
}

func NewServer(config *Config) *Server {
	return &Server{
		Config:  config,
		Router:  mux.NewRouter(),
		Storage: storage.NewStorage(config.StorageConfig),
	}
}

func (s *Server) ConfigureRouter() {
	s.Router.Use(s.logMiddleware())

	s.Router.HandleFunc("/", s.healthcheck())
	s.Router.HandleFunc("/createTransaction", s.createTransactionHandler())
	s.Router.HandleFunc("/updateTransactionStatus", s.updateTransactionStatusHandler())
	s.Router.HandleFunc("/chechTransactionStatus/{id}", s.checkTransactionStatusHandler())
	s.Router.HandleFunc("/readAllUsersTransactions/userid/{userid}", s.readAllUsersTransactionsByIDHandler())
	s.Router.HandleFunc("/readAllUsersTransactions/email/{email}", s.readAllUsersTransactionsByEmailHandler())
	s.Router.HandleFunc("/undoTransaction/{id}/{email}", s.undoTransactionHandler())
}

func (s *Server) Run() error {
	s.ConfigureRouter()
	log.Println("Opening db pool")
	err := s.Storage.Open()
	if err != nil {
		log.Println("Error opening storage")
		return err
	}
	log.Println("Starting server")
	return http.ListenAndServe(s.Config.Addr, s.Router)
}

func (s *Server) Shutdown() {
	s.Storage.Close()
}
