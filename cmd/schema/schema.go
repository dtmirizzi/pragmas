package schema

type Foo struct {
	ID          string   `json:"$id"`
	Schema      string   `json:"$schema"`
	Definitions struct{} `json:"definitions"`
	Properties  struct {
		Age struct {
			ID       string  `json:"$id"`
			Default  int64   `json:"default"`
			Examples []int64 `json:"examples"`
			Title    string  `json:"title"`
			Type     string  `json:"type"`
		} `json:"age"`
		Name struct {
			ID       string   `json:"$id"`
			Default  string   `json:"default"`
			Examples []string `json:"examples"`
			Pattern  string   `json:"pattern"`
			Title    string   `json:"title"`
			Type     string   `json:"type"`
		} `json:"name"`
	} `json:"properties"`
	Required []string `json:"required"`
	Title    string   `json:"title"`
	Type     string   `json:"type"`
}
