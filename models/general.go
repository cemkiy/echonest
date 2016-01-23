package models

// Status Struct
type Status struct {
	Version string `json:"id"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Response Struct
type Response struct {
	Status  Status   `json:"status"`
	Artists []Artist `json:"artists"`
	Songs   []Song   `json:"songs"`
}

// Result Struct
type Result struct {
	Response Response `json:"response"`
}
