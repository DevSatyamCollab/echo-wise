package internal

// quote struct
type Quote struct {
	Quote  string
	Author string
	Id     int
}

func NewQuote(id int, q, a string) *Quote {
	return &Quote{
		Id:     id,
		Quote:  q,
		Author: a,
	}
}
