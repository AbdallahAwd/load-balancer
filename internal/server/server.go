package server

import (
	"fmt"
	"net/http"
)

func Start(server *http.Server) error {
	fmt.Print("server is listening on ", server.Addr, " ...")
	if err := server.ListenAndServe(); err != nil {
		return err
	}
	return nil

}
