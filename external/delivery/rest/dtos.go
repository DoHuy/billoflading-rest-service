package rest

type logDTO struct {
	LogID string   `json:"log_id"`
	Data  *logData `json:"data"`
}
type logData struct {
	Date        string `json:"date"`
	Type        string `json:"type"`
	ClientID    string `json:"client_id"`
	ClientName  string `json:"client_name"`
	IP          string `json:"ip"`
	Description string `json:"description"`
	UserAgent   string `json:"user_agent"`
	UserID      string `json:"user_id"`
	LogID       string `json:"log_id"`
}

type logsDTOReq struct {
	Logs []logDTO `json:"logs"`
}
