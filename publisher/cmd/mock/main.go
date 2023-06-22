package main

//go:generate mockgen publisher/internal/person Male

// 👇 an interface acting as API Client
type ApiClient interface {
	GetData() string
}

// 👇 a function using the ApiClient interface
func Process(client ApiClient) int {
	data := client.GetData()
	return len(data)
}

func main() {

}
