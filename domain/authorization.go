package domain

type Authorization struct {
	AuthorizationCode string `json:"authorizationCode" copier:"AuthorizationCode"`
	LocalAmount       int64  `json:"localAmount" copier:"LocalAmount"`
	Mti               string `json:"mti" copier:"Mti"`
	AccountId         int64  `json:"accountId" copier:"AccountId"`
	CardId            int64  `json:"cardId" copier:"CardId"`
	ResponseCode      string `json:"responseCode" copier:"ResponseCode"`
}

type AuthorizationReq struct {
	AuthorizationCode string `json:"authorizationCode" copier:"AuthorizationCode"`
}

type AuthorizationRes struct {
	Mti               string `json:"mti" copier:"Mti"`
	AuthorizationCode string `json:"authorizationCode" copier:"AuthorizationCode"`
	ResponseCode      string `json:"responseCode" copier:"ResponseCode"`
}

type ExternalRes struct {
	AuthorizationCode string `json:"authorizationCode" copier:"AuthorizationCode"`
	LocalAmount       int64  `json:"localAmount" copier:"LocalAmount"`
	Mti               string `json:"mti" copier:"Mti"`
	AccountId         int64  `json:"accountId" copier:"AccountId"`
	CardId            int64  `json:"cardId" copier:"CardId"`
	ResponseCode      string `json:"responseCode" copier:"ResponseCode"`
}
