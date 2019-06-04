package db

import "github.com/emikohmann/go-structuring/domain-driven/domain/ban"

type repo struct {
}

func New() ban.Repository {
    return &repo{}
}

func (repo *repo) Create(b ban.Ban) error {
    // DB implementation for create
    return nil
}

func (repo *repo) Update(b ban.Ban) error {
    // DB implementation for update
    return nil
}

func (repo *repo) Delete(b ban.Ban) error {
    // DB implementation for delete
    return nil
}
