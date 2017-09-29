package inmem

import (
	"github.com/kristofferingemansson/go-service-template/pkg"
	"math/rand"
)

type dict []string

func (d dict) getRandom() string {
	i := rand.Intn(len(d))
	return d[i]
}

type quoteRepository struct {
	nouns        dict
	verbs        dict
	particles    dict
	articles     dict
	pronouns     dict
	prepositions dict
	adverbs      dict
	conjunctions dict
}

// NewQuoteRepository ..
func NewQuoteRepository() pkg.QuoteRepository {
	return &quoteRepository{
		nouns:        dict{"house", "car", "cake", "dog", "hose", "bridge", "bride", "bribe"},
		verbs:        dict{"build", "drive", "cook", "sell", "buy", "rent", "paint"},
		particles:    dict{"to"},
		articles:     dict{"a", "an"},
		pronouns:     dict{"I", "me", "mine", "myself", "she", "her", "hers", "herself", "we", "us", "ours", "ourselves"},
		prepositions: dict{"in", "on", "at", "around", "above", "near", "underneath", "alongside", "of", "for"},
		adverbs:      dict{"good", "well"},
		conjunctions: dict{"for", "and", "nor", "but", "or", "yet", "so"},
	}
}

func (r *quoteRepository) GetRandomNoun() (string, error) {
	return r.nouns.getRandom(), nil
}

func (r *quoteRepository) GetRandomVerb() (string, error) {
	return r.verbs.getRandom(), nil
}

func (r *quoteRepository) GetRandomParticle() (string, error) {
	return r.particles.getRandom(), nil
}

func (r *quoteRepository) GetRandomArticle() (string, error) {
	return r.articles.getRandom(), nil
}

func (r *quoteRepository) GetRandomPronoun() (string, error) {
	return r.pronouns.getRandom(), nil
}

func (r *quoteRepository) GetRandomPreposition() (string, error) {
	return r.prepositions.getRandom(), nil
}

func (r *quoteRepository) GetRandomAdverb() (string, error) {
	return r.adverbs.getRandom(), nil
}

func (r *quoteRepository) GetRandomConjunction() (string, error) {
	return r.conjunctions.getRandom(), nil
}
