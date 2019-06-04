package ds

import "github.com/emikohmann/go-structuring/domain-driven/domain/ban"

type repo struct {
}

func New() ban.Repository {
    return &repo{}
}

func (repo *repo) Create(b ban.Ban) error {
    // DS implementation for create
    return nil
}

func (repo *repo) Update(b ban.Ban) error {
    // DS implementation for update
    return nil
}

func (repo *repo) Delete(b ban.Ban) error {
    // DS implementation for delete
    return nil
}
