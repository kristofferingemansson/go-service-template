package pkg

// QuoteRepository ..
type QuoteRepository interface {
	GetRandomNoun() (string, error)
	GetRandomVerb() (string, error)
	GetRandomParticle() (string, error)
	GetRandomArticle() (string, error)
	GetRandomPronoun() (string, error)
	GetRandomPreposition() (string, error)
	GetRandomAdverb() (string, error)
	GetRandomConjunction() (string, error)
}
