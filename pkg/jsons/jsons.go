package jsons

//easyjson:json
//go:generate easyjson -all jsons.go
type TestStruct struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Tags    []string `json:"tags"`
	Enabled bool     `json:"enabled"`
}
