package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nats-io/nats.go"

	persistence "refound/internal/account/infrastructure/persistence"

	repo "refound/internal/account/repo"
	accountService "refound/internal/account/service/account"
)

const SERVICE_SUBJECT = "account_service"

func main() {
	// CONNECT TO EVENT BUS
	// note: you'll need to run "docker compose up" in project root to launch NATS server
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

	// INIT DB

	repositories, repoErr := persistence.NewRepositories(os.Getenv("DB_PG_USER"), os.Getenv("DB_PG_PASSWORD"), os.Getenv("DB_PG_HOST"), os.Getenv("DB_PG_PORT"), os.Getenv("DB_PG_DBNAME"))
	if repoErr != nil {
		panic(repoErr)
	}
	defer repositories.Close()

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
