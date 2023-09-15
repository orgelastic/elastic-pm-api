package core

import "net/http"

type AccountCreatedEvent struct {
	ID         string `json:"id"`
	TimeJoined uint64 `json:"timeJoined"`
	Email      string `json:"email"`
	ThirdParty *struct {
		ID     string `json:"id"`
		UserID string `json:"userId"`
	} `json:"thirdParty"`
	TenantIds []string `json:"tenantIds"`
}

type UnauthorisedAccessEvent struct {
	Message string
	Req     *http.Request
	Res     http.ResponseWriter
}
