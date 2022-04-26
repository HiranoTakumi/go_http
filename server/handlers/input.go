package handlers

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type Persons struct {
	Persons []Person `json:"persons"`
}
