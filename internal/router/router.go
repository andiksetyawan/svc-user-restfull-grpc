package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"svc-user/api/v1/userv1connect"
)

type router struct {
	port    string
	service userv1connect.UserServiceClient
	mux     *http.ServeMux
}

func NewRouter(port string, service userv1connect.UserServiceClient) *router {
	mux := http.NewServeMux()
	path, handler := userv1connect.NewUserServiceHandler(service)
	mux.Handle(path, handler)
	return &router{
		port:    port,
		service: service,
		mux:     mux,
	}
}

func (r *router) Listen() error {
	hostname, _ := os.Hostname()
	log.Info().Msgf("svc-user:%v is running on localhost:%v", hostname, r.port)

	if err := http.ListenAndServe(
		fmt.Sprintf(":%v", r.port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(r.mux, &http2.Server{}),
	); err != nil {
		log.Fatal().Err(err).Msg(err.Error())
		return err
	}
	return nil
}

func (r *router) Mux() *http.ServeMux {
	return r.mux
}
