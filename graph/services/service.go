package services

import (
	"context"

	"graphql/graph/model"
)

type Service interface {
	CreateUser(ctx context.Context, nu model.NewUser) (*model.User, error)
	ViewUser(ctx context.Context) ([]*model.User, error)

	CreateCompany(newComp model.NewCompany) (*model.Company, error)
	ViewCompanies(ctx context.Context) ([]*model.Company, error)
	FetchCompanyByID(ctx context.Context, companyId string) (*model.Company, error)

	CreateByCompanyId(jobs []*model.NewJob, compId string) ([]*model.Job, error)
	FetchJobByCompanyId(ctx context.Context, companyId string) ([]*model.Job, error)
	GetJobById(ctx context.Context, jobId string) (*model.Job, error)
	GetAllJobs(ctx context.Context) ([]*model.Job, error)
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
