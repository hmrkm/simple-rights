package usecase

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/hmrkm/simple-rights/domain"
)

func TestVerify(t *testing.T) {
	testCases := []struct {
		name              string
		userID            string
		path              string
		loadUserRole      []domain.UserRole
		loadUserRoleErr   error
		loadResource      domain.Resource
		loadResourceErr   error
		loadPermission    []domain.Permission
		loadPermissionErr error
		expectedErr       error
	}{
		{
			"正常ケース",
			"userA",
			"path",
			[]domain.UserRole{},
			nil,
			domain.Resource{},
			nil,
			[]domain.Permission{{}},
			nil,
			nil,
		},
		{
			"権限が無い異常ケース",
			"userA",
			"path",
			[]domain.UserRole{},
			nil,
			domain.Resource{},
			nil,
			[]domain.Permission{},
			nil,
			domain.ErrNotAuthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := domain.NewMockStore(ctrl)
			sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
				func(target *[]domain.UserRole, cond *domain.UserRole) error {
					if tc.loadUserRoleErr == nil {
						*target = tc.loadUserRole
					}
					return tc.loadUserRoleErr
				},
			)
			sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
				func(target *domain.Resource, cond *domain.Resource) error {
					if tc.loadResourceErr == nil {
						*target = tc.loadResource
					}
					return tc.loadResourceErr
				},
			)
			sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
				func(target *[]domain.Permission, cond map[string]interface{}) error {
					if tc.loadPermissionErr == nil {
						*target = tc.loadPermission
					}
					return tc.loadPermissionErr
				},
			)
			ru := NewRights(sm)

			actualErr := ru.Verify(tc.userID, tc.path)

			if !errors.Is(actualErr, tc.expectedErr) {
				t.Errorf("Verify() actualErr: %v, ecpectedErr: %v", actualErr, tc.expectedErr)
			}
		})
	}
}
