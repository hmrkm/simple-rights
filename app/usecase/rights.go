package usecase

import "github.com/hmrkm/simple-rights/domain"

type Rights interface {
	Verify(string, string) error
}

type rights struct {
	rights domain.Rights
}

func NewRights(r domain.Rights) Rights {
	return rights{
		rights: r,
	}
}

func (r rights) Verify(userID string, resource string) error {
	return r.rights.Verify(userID, resource)
}
