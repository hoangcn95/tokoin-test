package model

// Users struct
type Users struct {
	ID             int      `json:"_id" bson:"_id"`
	URL            string   `json:"url" bson:"url"`
	ExternalID     string   `json:"external_id" bson:"external_id"`
	Name           string   `json:"name" bson:"name"`
	Alias          string   `json:"alias" bson:"alias"`
	CreatedAt      string   `json:"created_at" bson:"created_at"`
	Active         bool     `json:"active" bson:"active"`
	Verified       bool     `json:"verified" bson:"verified"`
	Shared         bool     `json:"shared" bson:"shared"`
	Locale         string   `json:"locale" bson:"locale"`
	Timezone       string   `json:"timezone" bson:"timezone"`
	LastLoginAt    string   `json:"last_login_at" bson:"last_login_at"`
	Email          string   `json:"email" bson:"email"`
	Phone          string   `json:"phone" bson:"phone"`
	Signature      string   `json:"signature" bson:"signature"`
	OrganizationID int      `json:"organization_id" bson:"organization_id"`
	Tags           []string `json:"tags" bson:"tags"`
	Suspended      bool     `json:"suspended" bson:"suspended"`
	Role           string   `json:"role" bson:"role"`
}

// IsExists ..
func (m Users) IsExists() (ok bool) {
	if m.ID != 0 {
		ok = true
	}
	return
}
