package entity

//Person object for REST(CRUD)
type Person struct {
	ID        int    `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	City      string `json:"city"`
}

type Tabler interface {
	TableName() string
}

func (Person) TableName() string {
	return "person"
}
