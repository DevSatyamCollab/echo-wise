package suffle

import (
	"math/rand"

	internal "github.com/DevSatyamCollab/echo-wise/internal/core"
)

func Suffle(quotes []internal.Quote) internal.Quote {
	r := rand.Intn(len(quotes))
	q := quotes[r]
	return q
	//return q.Id, q.Quote, q.Author
}
