package ban

type Service interface {
    Apply(b Ban) error
    Remove(b Ban) error
}

type service struct {
    repo Repository
}

func NewService(r Repository) Service {
    return &service{
        repo: r,
    }
}

func (s *service) Apply(b Ban) error {
    if err := s.repo.Create(b); err != nil {
        return err
    }
    if err := s.repo.Update(b); err != nil {
        return err
    }
    return nil
}

func (s *service) Remove(b Ban) error {
    if err := s.repo.Delete(b); err != nil {
        return err
    }
    return nil
}
