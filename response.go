package cloudflare

import (
	"time"

	"github.com/google/uuid"
)

type ZeroTrustUser struct {
	ID        uuid.UUID `json:"id"`
	UID       uuid.UUID `json:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"` // TODO: parse the time
	UpdatedAt time.Time `json:"updated_at"`
}

type FetchZeroTrustUsersResponse struct {
	Success    bool            `json:"success"`
	Result     []ZeroTrustUser `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}
