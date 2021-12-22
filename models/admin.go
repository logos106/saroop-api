package model

type Admin struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Domain int `json:"domain,omitempty"`
	Password string `json:"password,omitempty"`
	Role string `json:"role,omitempty"`
	Status string `json:"status,omitempty"`
}

// Address description
type Domain struct {
	ID  int `json:"id,omitempty"`
	Name string `json:"state,omitempty"`
}
