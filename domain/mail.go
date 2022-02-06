package domain

// Mail - strut mail
type Mail struct {
    Name    string `json:"name"`
    Subject string `json:"subject"`
    Text    string `json:"text"`
    Email   string `json:"email"`
}
