package ban

type Ban struct {
    Type  string `json:"type"`
    Value string `json:"value"`
}

type Repository interface {
    Create(b Ban) error
    Update(b Ban) error
    Delete(b Ban) error
}
