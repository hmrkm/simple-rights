package usecase

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hmrkm/simple-rights/domain"
)

func TestVerify(t *testing.T) {
	err := errors.New("")
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
			[]domain.UserRole{
				{},
			},
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
		{
			"UserRole取得失敗の異常ケース",
			"userA",
			"path",
			[]domain.UserRole{},
			err,
			domain.Resource{},
			nil,
			[]domain.Permission{},
			nil,
			err,
		},
		{
			"Resource取得失敗の異常ケース",
			"userA",
			"path",
			[]domain.UserRole{},
			nil,
			domain.Resource{},
			err,
			[]domain.Permission{},
			nil,
			err,
		},
		{
			"Permission取得失敗の異常ケース",
			"userA",
			"path",
			[]domain.UserRole{},
			nil,
			domain.Resource{},
			nil,
			[]domain.Permission{},
			err,
			err,
		},
	}

	opts := []cmp.Option{cmpopts.EquateErrors()}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			sm := domain.NewMockStore(ctrl)

			expect := func() {
				sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
					func(target *[]domain.UserRole, cond *domain.UserRole) error {
						if tc.loadUserRoleErr == nil {
							*target = tc.loadUserRole
						}
						return tc.loadUserRoleErr
					},
				)
				if tc.loadUserRoleErr != nil {
					return
				}
				sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
					func(target *domain.Resource, cond *domain.Resource) error {
						if tc.loadResourceErr == nil {
							*target = tc.loadResource
						}
						return tc.loadResourceErr
					},
				)
				if tc.loadResourceErr != nil {
					return
				}
				sm.EXPECT().Load(gomock.Any(), gomock.Any()).DoAndReturn(
					func(target *[]domain.Permission, cond map[string]interface{}) error {
						if tc.loadPermissionErr == nil {
							*target = tc.loadPermission
						}
						return tc.loadPermissionErr
					},
				)
			}

			expect()

			ru := NewRights(sm)

			actualErr := ru.Verify(tc.userID, tc.path)

			if diff := cmp.Diff(actualErr, tc.expectedErr, opts...); diff != "" {
				t.Errorf("Verify() error miss match : %s", diff)
			}
		})
	}
}
