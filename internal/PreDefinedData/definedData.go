package predefineddata

import core "github.com/DevSatyamCollab/echo-wise/internal/core"

func GetPreData() []core.Quote {
	return []core.Quote{
		*core.NewQuote(0,
			"Don't listen to what people say, watch what they do.",
			"Churchill",
		),

		*core.NewQuote(1,
			"The only way to do great work is to love what you do.",
			"Steve Jobs",
		),

		*core.NewQuote(2,
			"It is not the mountain we conquer, but ourselves.",
			"Sir Edmund Hilary",
		),

		*core.NewQuote(3,
			"Success is not final, failure is not fatal: it is the courage to continue that counts.",
			"Winston Churchill",
		),

		*core.NewQuote(4,
			"The trouble with having an open mind, of course, is that people will insist on coming along and trying to put things in it.",
			"Terry Pratchett",
		),

		*core.NewQuote(5,
			"Life is what happens when you're busy making other plans.",
			"John Lennon",
		),

		*core.NewQuote(6,
			"Everything is funny, as long as it's happening to somebody else",
			"Will Rogers",
		),

		*core.NewQuote(7,
			"No act of kindness, no matter how small, is ever wasted.",
			"Aesop",
		),

		*core.NewQuote(8,
			"In the middle or every difficult lies opportunity.",
			"Albert Einstein",
		),

		*core.NewQuote(9,
			"Do what you can, with what you have, where you are.",
			"Theodore Roosevelt",
		),

		*core.NewQuote(10,
			"Yesterday is history, tomorrow is a mystery, but today is a gift. That is why it is called the present",
			"Alice Morse Earle",
		),

		*core.NewQuote(11,
			"I’m not in this world to live up to your expectations and you’re not in this world to live up to mine",
			"Bruce Lee",
		),
	}
}
