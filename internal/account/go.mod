module refound/internal/account

go 1.18

require (
	github.com/go-kit/kit v0.12.0
	github.com/google/uuid v1.3.0
	github.com/lib/pq v1.10.6
	github.com/nats-io/nats.go v1.16.0
	refound/internal/shared v0.0.0
)

replace refound/internal/shared => ../shared

require (
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
)
