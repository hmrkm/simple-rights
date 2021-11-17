package domain

type Rights interface {
	Verify(string, string) error
}

type rights struct {
	store Store
}

func NewRights(s Store) Rights {
	return rights{
		store: s,
	}
}

func (r rights) Verify(userID string, path string) error {
	userRoles := []UserRole{}
	if err := r.store.Load(&userRoles, &UserRole{UserID: userID}); err != nil {
		return err
	}
	roleIDs := []string{}
	for _, userRole := range userRoles {
		roleIDs = append(roleIDs, userRole.RoleID)
	}
	resource := Resource{}
	if err := r.store.Load(&resource, &Resource{Path: path}); err != nil {
		return err
	}
	permissions := []Permission{}
	if err := r.store.Load(&permissions, map[string]interface{}{"role_id": roleIDs, "resource_id": resource.ID}); err != nil {
		return err
	}
	if len(permissions) == 0 {
		return ErrNotAuthorized
	}

	return nil
}
