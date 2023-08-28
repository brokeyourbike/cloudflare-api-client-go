package cloudflare

import "github.com/google/uuid"

type FetchZeroTrustUsersResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		ID        uuid.UUID `json:"id"`
		UID       uuid.UUID `json:"uid"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt string    `json:"created_at"`
		UpdatedAt string    `json:"updated_at"`
	} `json:"result"`
	ResultInfo struct {
		Page       int `json:"page"`
		PerPage    int `json:"per_page"`
		Count      int `json:"count"`
		TotalCount int `json:"total_count"`
	} `json:"result_info"`
}
