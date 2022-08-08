package pg

type DomainEventDSO struct {
	Id            string `json:"id"`
	EventType     string `json:"event_type"`
	AggregateId   string `json:"aggregate_id"`
	AggregateType string `json:"aggregate_type"`
	EventData     string `json:"event_data"`
	Channel       string `json:"channel"`
	CreatedAt     string `json:"created_at"`
}
