package http

import (
	// stdlib
	"log"
	"net"
	"net/http"
	// universe
	"github.com/universelabs/universe-server/universe"
)

// HTTP service
type Server struct {
	// net/http infrastructure
	ln net.Listener
	// handler to serve
	Handler *Handler
	// bind address to open
	Addr string
}

// Returns a new Server instantiated from the arguments 
func NewServer(addr string, ks universe.Keystore) *Server {
	srv := &Server {
		// set addr
		Addr: addr,
		// init handler
		Handler: NewHandler(ks),
	}
	return srv
}

// Listens and serves the server instance
func (srv *Server) Open() error {
	// open the socket
	ln, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		return err
	}
	srv.ln = ln

	// start HTTP server
	go func() { log.Fatal(http.Serve(srv.ln, srv.Handler)) }()
	// *** because http.Serve is called in a go routine, main() must hang
	// the process so that the server doesn't close!

	return nil
}

// Closes the server instance
func (srv *Server) Close() error {
	if srv.ln != nil {
		return srv.ln.Close()
	}
	return nil
}