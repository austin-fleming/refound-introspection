package main

import (
	"fmt"
	"net/http"

	"github.com/nats-io/nats.go"

	repo "refound/internal/account/repo"
	accountService "refound/internal/account/service/account"
)

const SERVICE_SUBJECT = "account_service"

func main() {
	// CONNECT TO NATS
	wait := make(chan bool)

	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}

	natsConnection.Subscribe(SERVICE_SUBJECT, func(msg *nats.Msg) {
		fmt.Printf("Received: %s", string(msg.Data))
		natsConnection.Publish(msg.Reply, []byte("Hello"))
	})

	fmt.Println("Subscribed to", SERVICE_SUBJECT)

	<-wait

	// INIT REPOS
	accountRepo := repo.MakeAccountRepo()
	accountRelationRepo := repo.MakeAccountRelationRepo()

	// INIT SERVICES
	service := accountService.NewService(accountRepo, accountRelationRepo)

	// INIT HANDLERS
	accountExistsHandler := accountService.MakeAccountExistsHandler(service)

	// INIT ROUTES
	http.Handle("/exists", accountExistsHandler)
}
