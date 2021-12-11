package usecase

import "github.com/hmrkm/simple-rights/domain"

type Rights interface {
	Verify(string, string) error
}

type rights struct {
	store domain.Store
}

func NewRights(s domain.Store) Rights {
	return rights{
		store: s,
	}
}

func (r rights) Verify(userID string, path string) error {
	userRoles := []domain.UserRole{}
	if err := r.store.Load(&userRoles, &domain.UserRole{UserID: userID}); err != nil {
		return err
	}
	roleIDs := []string{}
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}
	resource := domain.Resource{}
	if err := r.store.Load(&resource, &domain.Resource{Path: path}); err != nil {
		return err
	}
	permissions := []domain.Permission{}
	if err := r.store.Load(&permissions, map[string]interface{}{"role_id": roleIDs, "resource_id": resource.ID}); err != nil {
		return err
	}
	if len(permissions) == 0 {
		return domain.ErrNotAuthorized
	}

	return nil
}
