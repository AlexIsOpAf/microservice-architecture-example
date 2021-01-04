package internal

type Team struct {
	CreatedBy string   `json:"createdby"`
	Name      string   `json:"name"`
	Members   []string `json:"members"`
}

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}
