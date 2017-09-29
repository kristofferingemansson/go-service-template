package quote

import (
	"github.com/kristofferingemansson/go-service-template/pkg"
	"strings"
)

// Service ..
type Service interface {
	GenerateQuote() (string, error)
}

type service struct {
	repository pkg.QuoteRepository
}

// NewService ..
func NewService(repository pkg.QuoteRepository) Service {
	return &service{
		repository: repository,
	}
}

type wordgenerator func() (string, error)

func (s *service) GenerateQuote() (string, error) {
	sent := []wordgenerator{
		s.repository.GetRandomPronoun,
		s.repository.GetRandomParticle,
		s.repository.GetRandomVerb,
		s.repository.GetRandomArticle,
		s.repository.GetRandomNoun,
		s.repository.GetRandomAdverb,
	}

	ret := make([]string, len(sent))
	for i, gen := range sent {
		var err error
		ret[i], err = gen()
		if err != nil {
			return "", err
		}
	}

	return strings.Join(ret, " "), nil
}
