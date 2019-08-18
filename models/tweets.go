package models

// Tweet represents a single, public user tweet
type Tweet struct {
	Username string `json:"username,omitempy"`
	Text     string `json:"text,omitempty"`
	Time     int64  `json:"-"`
}

type Follower struct {
	ID string `json:id,omitempty`
}
