package models

type Profile struct {
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Email     string   `json:"email,omitempty"`
	Following []string `json:"following"`
}
