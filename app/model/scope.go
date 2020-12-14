package model

// Scope ...
type Scope struct {
	Order         string       `json:"order"`
	Tag           string       `json:"tag"`
	Text          string       `json:"text"`
	FriendsOnly   bool         `json:"friendsOnly"`
}