package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"shareddata/internal/config"
)

type Server struct {
	mux        *http.ServeMux
	handlerMap *map[string]http.HandlerFunc
	conf       *config.ServerConf
}

func NewServer(conf *config.ServerConf, handlerMap *map[string]http.HandlerFunc) (*Server, error) {
	if len(*handlerMap) < 1 {
		log.Fatal("Handler map is empty")
		return nil, errors.New("Empty handler map")
	}
	svr := &Server{
		mux:        http.NewServeMux(),
		handlerMap: handlerMap,
		conf:       conf,
	}
	svr.registerHandlers()
	return svr, nil
}

func (s *Server) registerHandlers() {
	for name, handlerFunc := range *s.handlerMap {
		s.mux.Handle(name, handlerFunc)
	}
}

func (s *Server) Start() error {
	if len(*s.handlerMap) < 1 {
		err := errors.New("Empty handler map")
		log.Fatal(err)
		return err
	}
	return http.ListenAndServe(fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port), s.mux)
}
