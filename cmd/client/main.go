package main

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
	userv1 "svc-user/api/v1"
	"svc-user/api/v1/userv1connect"
)

const addressClient = "http://localhost:9998"

func main() {
	cGrpc := userv1connect.NewUserServiceClient(
		http.DefaultClient,
		addressClient,
		connect.WithGRPC(),
	)

	r, err := cGrpc.Create(context.TODO(), connect.NewRequest(&userv1.CreateRequest{Name: "John Client"}))
	if err != nil {
		log.Fatal().Err(err).Msg("fail creating new user")
	}

	log.Debug().Msg(r.Msg.String())
}
