package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"

	"graphql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	return r.Store.Service.CreateUser(ctx, input)
}

func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	return r.Store.Service.CreateCompany(input)
}
func (r *mutationResolver) CreateJobForCompany(ctx context.Context, input string, newJob []*model.NewJob) ([]*model.Job, error) {
	return r.Store.Service.CreateByCompanyId(newJob, input)
}

// FetchAllUser is the resolver for the fetchAllUser field.
func (r *queryResolver) FetchAllUser(ctx context.Context) ([]*model.User, error) {
	return r.Store.Service.ViewUser(ctx)
	// panic(fmt.Errorf("not implemented: FetchAllUser - fetchAllUser"))
}

// FetchCompanies is the resolver for the fetchCompanies field.
func (r *queryResolver) FetchCompanies(ctx context.Context) ([]*model.Company, error) {
	return r.Store.Service.ViewCompanies(ctx)
	// panic(fmt.Errorf("not implemented: FetchCompanies - fetchCompanies"))
}

// FetchCompaniesByID is the resolver for the fetchCompaniesById field.
func (r *queryResolver) FetchCompaniesByID(ctx context.Context, input string) (*model.Company, error) {
	return r.Store.Service.FetchCompanyByID(ctx, input)
}

// FetchJobByCompanyID is the resolver for the fetchJobByCompanyId field.
func (r *queryResolver) FetchJobByCompanyID(ctx context.Context, input string) ([]*model.Job, error) {
	return r.Store.Service.FetchJobByCompanyId(ctx, input)
}

func (r *queryResolver) FetchByJobID(ctx context.Context, input string) (*model.Job, error) {
	return r.Store.Service.GetJobById(ctx, input)
}

// FetchAllJobs is the resolver for the fetchAllJobs field.
func (r *queryResolver) FetchAllJobs(ctx context.Context) ([]*model.Job, error) {
	return r.Store.Service.GetAllJobs(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }