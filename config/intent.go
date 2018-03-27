package config

import "github.com/mcmillan/alexalog/domain"

type Intent struct {
	Name       string
	OnStart    func(*domain.Intent) error
	OnComplete func(*domain.Intent) (*domain.Response, error)
}

type Intents []*Intent
