package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rtnhnd255/payment_service_emulator/app/storage"
)

func (s *Server) healthcheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "healthcheck")
	}
}

func (s *Server) createTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t *storage.Transaction
		err := decoder.Decode(t)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		err = s.Storage.WriteTransaction(t)
		if err != nil {
			log.Println("Error inserting transaction")
			log.Println(err)
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
	}
}

func (s *Server) updateTransactionStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t *storage.Transaction
		err := decoder.Decode(t)
		if err != nil {
			w.WriteHeader(400)
			return
		}

		tSaved := s.Storage.ReadTransactionByID(t.ID)
		if tSaved.UserEmail != t.UserEmail {
			log.Println("Unauthorized attempt to update transaction status")
			w.WriteHeader(403)
			return
		}

		err = s.Storage.UpdateTransactionStatus(t.ID, t.Status)
		if err != nil {
			log.Println("Unable to update status")
			log.Println(err)
			w.WriteHeader(500)
		}
		w.WriteHeader(200)
	}
}

func (s *Server) checkTransactionStatusHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 0 {
			w.WriteHeader(400)
			return
		}
		_id := uint64(id)
		t := s.Storage.ReadTransactionByID(_id)
		w.Header().Add("status", t.Status)
		w.WriteHeader(200)
	}
}

func (s *Server) readAllUsersTransactionsByIDHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userid := vars["userid"]

		ts := s.Storage.ReadAllUsersTransactionsByID(userid)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(ts)
		w.WriteHeader(200)
	}
}

func (s *Server) readAllUsersTransactionsByEmailHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email := vars["email"]

		ts := s.Storage.ReadAllUsersTransactionsByEmail(email)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(ts)
		w.WriteHeader(200)
	}
}

func (s *Server) undoTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 0 {
			w.WriteHeader(400)
			return
		}
		email := vars["email"]
		_id := uint64(id)
		t := s.Storage.ReadTransactionByID(_id)

		if t.UserEmail != email {
			w.WriteHeader(403)
			return
		}
		if t.Status == "SUCCESS" || t.Status == "FAILURE" {
			w.WriteHeader(400)
			w.Header().Add("reason", fmt.Sprintf("Cannot undo transaction because it is in terminal status - %s", t.Status))
		}
	}
}
