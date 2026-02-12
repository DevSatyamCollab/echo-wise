package suffle

import (
	"math/rand"

	core "github.com/DevSatyamCollab/echo-wise/internal/core"
)

func Suffle(quotes []core.Quote) core.Quote {
	r := rand.Intn(len(quotes))
	q := quotes[r]
	return q
	//return q.Id, q.Quote, q.Author
}
