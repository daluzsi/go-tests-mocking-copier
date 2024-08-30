package service_test

import (
	"errors"
	"example.com/tests/domain"
	"example.com/tests/mocks"
	"example.com/tests/service"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthorizationService_Authorize(t *testing.T) {
	scenarios := []struct {
		name          string
		firstCopyErr  error
		secondCopyErr error
		expectedErr   error
	}{
		{
			name:          "Both copies succeed",
			firstCopyErr:  nil,
			secondCopyErr: nil,
			expectedErr:   nil,
		},
		{
			name:          "First copy fails",
			firstCopyErr:  errors.New("first copy error"),
			secondCopyErr: nil,
			expectedErr:   errors.New("first copy error"),
		},
		{
			name:          "Second copy fails",
			firstCopyErr:  nil,
			secondCopyErr: errors.New("second copy error"),
			expectedErr:   errors.New("second copy error"),
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authReq := domain.AuthorizationReq{
				AuthorizationCode: "ASDQW1",
			}

			mockCopier := mocks.NewMockCopier(ctrl)

			mockCopier.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(scenario.firstCopyErr).Times(1)

			if scenario.firstCopyErr == nil {
				mockCopier.EXPECT().Copy(gomock.Any(), gomock.Any()).Return(scenario.secondCopyErr).Times(1)
			}

			authSvc := service.NewAuthorizationService(mockCopier)

			response, err := authSvc.Authorize(authReq)

			if scenario.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, scenario.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

			ctrl.Finish()
		})
	}
}

func TestAuthorizationService_AuthorizeT(t *testing.T) {
	scenarios := []struct {
		name               string
		firstCopyErr       error
		secondCopyErr      error
		expectedErr        error
		validateFirstCopy  bool
		validateSecondCopy bool
	}{
		{
			name:               "Both copies succeed",
			firstCopyErr:       nil,
			secondCopyErr:      nil,
			expectedErr:        nil,
			validateFirstCopy:  true,
			validateSecondCopy: true,
		},
		{
			name:               "First copy fails",
			firstCopyErr:       errors.New("first copy error"),
			secondCopyErr:      nil,
			expectedErr:        errors.New("first copy error"),
			validateFirstCopy:  false,
			validateSecondCopy: true,
		},
		{
			name:               "Second copy fails",
			firstCopyErr:       nil,
			secondCopyErr:      errors.New("second copy error"),
			expectedErr:        errors.New("second copy error"),
			validateFirstCopy:  true,
			validateSecondCopy: false,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			authReq := domain.AuthorizationReq{
				AuthorizationCode: "ASDQW1",
			}

			mockCopier := mocks.NewMockCopier(ctrl)

			mockCopier.EXPECT().Copy(gomock.Any(), gomock.Any()).DoAndReturn(func(dst, src interface{}) error {
				if scenario.validateFirstCopy {
					copier.Copy(dst, src)
					assert.Equal(t, dst.(*domain.Authorization).AuthorizationCode, src.(*domain.AuthorizationReq).AuthorizationCode)
				}
				return scenario.firstCopyErr
			}).Times(1)

			if scenario.firstCopyErr == nil {
				mockCopier.EXPECT().Copy(gomock.Any(), gomock.Any()).DoAndReturn(func(dst, src interface{}) error {
					if scenario.validateSecondCopy {
						copier.Copy(dst, src)
						assert.Equal(t, dst.(*domain.AuthorizationRes).AuthorizationCode, src.(*domain.ExternalRes).AuthorizationCode)
					}
					return scenario.secondCopyErr
				}).Times(1)
			}

			authSvc := service.NewAuthorizationService(mockCopier)

			response, err := authSvc.Authorize(authReq)

			if scenario.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, scenario.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, response)
			}

			ctrl.Finish()
		})
	}
}
