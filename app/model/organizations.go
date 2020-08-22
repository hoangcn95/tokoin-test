package model

// Organizations struct
type Organizations struct {
	ID            int      `json:"_id" bson:"_id"`
	URL           string   `json:"url" bson:"url"`
	ExternalID    string   `json:"external_id" bson:"external_id"`
	Name          string   `json:"name" bson:"name"`
	DomainNames   []string `json:"domain_names" bson:"domain_names"`
	CreatedAt     string   `json:"created_at" bson:"created_at"`
	Details       string   `json:"details" bson:"details"`
	SharedTickets bool     `json:"shared_tickets" bson:"shared_tickets"`
	Tags          []string `json:"tags" bson:"tags"`
}

// IsExists ..
func (m Organizations) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
