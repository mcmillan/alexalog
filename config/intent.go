package config

import "github.com/mcmillan/alexalog/domain"

type IntentOptions struct {
	Name string
}

type Intent struct {
	Options    IntentOptions
	OnStart    func(*domain.Intent) error
	OnComplete func(*domain.Intent) (*domain.Response, error)
}

type Intents []*Intent
