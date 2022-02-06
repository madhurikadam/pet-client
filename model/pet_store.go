package model

type Pet struct {
	ID    int     `json:"id"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type PetPostReq struct {
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

type PetPostResp struct {
	Pet     PetPostReq `json:"pet"`
	Message string     `json:"message"`
}
