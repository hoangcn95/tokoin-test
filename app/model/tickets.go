package model

// Tickets struct
type Tickets struct {
	ID             string   `json:"_id" bson:"_id"`
	URL            string   `json:"url" bson:"url"`
	ExternalID     string   `json:"external_id" bson:"external_id"`
	CreatedAt      string   `json:"created_at" bson:"created_at"`
	Type           string   `json:"type" bson:"type"`
	Subject        string   `json:"subject" bson:"subject"`
	Description    string   `json:"description" bson:"description"`
	Priority       string   `json:"priority" bson:"priority"`
	Status         string   `json:"status" bson:"status"`
	SubmitterID    int      `json:"submitter_id" bson:"submitter_id"`
	AssigneeID     int      `json:"assignee_id" bson:"assignee_id"`
	OrganizationID int      `json:"organization_id" bson:"organization_id"`
	Tags           []string `json:"tags" bson:"tags"`
	HasIncidents   bool     `json:"has_incidents" bson:"has_incidents"`
	DueAt          string   `json:"due_at" bson:"due_at"`
	Via            string   `json:"via" bson:"via"`
}

// IsExists ..
func (m Tickets) IsExists() (ok bool) {
	if m.ID != "" {
		ok = true
	}
	return
}
