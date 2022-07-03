package base

type ExampleConfig struct {}

func (ExampleConfig) BaseUrl() string {
	return "https://jsonplaceholder.typicode.com"
}