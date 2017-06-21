/**
 * @Author: Viacheslav Shynkarenko <Slavik>
 * @Date:   2017-06-22T01:20:59+03:00
 * @Email:  shinkarenko.vi@gmail.com
 * @Last modified by:   Slavik
 * @Last modified time: 2017-06-22T02:35:42+03:00
 * @Copyright: Viacheslav Shynkarenko. All Rights Reserved.
 */

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Server implements primary business logic of this app and handlers HTTP requests.
type Server struct {
	rqlog *RequestLog
}

type serverResponse struct {
	MinuteHitsCount int `json:"minute_hits_count"`
}

const defaultEndpoint = ":3030"

// NewServer create new instance of the server.
func NewServer(rqlog *RequestLog) *Server {
	server := &Server{}
	server.rqlog = rqlog
	server.mapHandlers()
	return server
}

// Start starts the server to listen for HTTP requests.
func (server *Server) Start() error {
	return http.ListenAndServe(defaultEndpoint, nil)
}

func (server *Server) mapHandlers() {
	http.HandleFunc("/", server.handleRequest)
}

func (server *Server) handleRequest(w http.ResponseWriter, req *http.Request) {
	server.rqlog.Hit()
	response := serverResponse{MinuteHitsCount: server.rqlog.MinuteHitsTotal()}
	server.respondJSON(response, w)
}

func (server *Server) respondJSON(v interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(json)
}
