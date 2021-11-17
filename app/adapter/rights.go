package adapter

import "github.com/hmrkm/simple-rights/usecase"

type Rights interface {
	Verify(RequestRights) error
}

type rights struct {
	rights usecase.Rights
}

func NetRights(r usecase.Rights) Rights {
	return rights{
		rights: r,
	}
}

func (r rights) Verify(req RequestRights) error {
	return r.rights.Verify(req.UserId, req.Resource)
}
