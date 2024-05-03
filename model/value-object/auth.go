package valueobject

import "github.com/google/uuid"

type Auth struct {
	BusinessID uint64    `json:"business_id" valid:"Required"`
	UserID     uuid.UUID `json:"user_id" valid:"Required"`
	RequestID  string    `json:"request_id" valid:"Required"`
	OrgID      uint64    `json:"org_id" valid:"Required"`
}
