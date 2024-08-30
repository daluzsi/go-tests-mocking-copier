package service

import (
	"example.com/tests/copier"
	"example.com/tests/domain"
)

type AuthorizationService struct {
	copier copier.Copier
}

func NewAuthorizationService(copier copier.Copier) *AuthorizationService {
	return &AuthorizationService{
		copier: copier,
	}
}

func (as *AuthorizationService) Authorize(authReq domain.AuthorizationReq) (domain.AuthorizationRes, error) {
	authRes := domain.AuthorizationRes{}
	authD := domain.Authorization{}

	if err := as.copier.Copy(&authD, &authReq); err != nil {
		return authRes, err
	}

	externalRes, err := as.ExternalCall(authD)
	if err != nil {
		return authRes, err
	}

	if err := as.copier.Copy(&authRes, &externalRes); err != nil {
		return authRes, err
	}

	return authRes, nil
}

func (as *AuthorizationService) ExternalCall(authD domain.Authorization) (domain.ExternalRes, error) {
	return domain.ExternalRes{AuthorizationCode: authD.AuthorizationCode}, nil
}
