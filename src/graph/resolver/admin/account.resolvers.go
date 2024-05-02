package resolver_admin

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	generated_admin "git.ctisoftware.vn/back-end/base/src/graph/generated/admin"
	graph_model "git.ctisoftware.vn/back-end/base/src/graph/generated/model"
	"git.ctisoftware.vn/back-end/base/src/network"
	service_account "git.ctisoftware.vn/back-end/base/src/service/account"
)

// AccountAdd is the resolver for the accountAdd field.
func (r *mutationResolver) AccountAdd(ctx context.Context, data graph_model.AccountAdd) (*graph_model.Account, error) {
	input := &service_account.AccountAddCommand{
		WorkspaceID: network.WorkspaceID(ctx),
		ActionBy:    network.Username(ctx),
		ActionByID:  network.AccountID(ctx),
		Name:        data.Name,
		Gender:      data.Gender,
		Phone:       data.Phone,
		Avatar:      data.Avatar,
		Username:    data.Username,
		Email:       data.Email,
		Password:    data.Password,
		Type:        data.Type,
	}

	result, err := service_account.AccountAdd(ctx, input)
	if err != nil {
		return &graph_model.Account{}, err
	}

	return result.ConvertToModelGraph(), nil
}

// AccountUpdate is the resolver for the accountUpdate field.
func (r *mutationResolver) AccountUpdate(ctx context.Context, data graph_model.AccountUpdate) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountUpdate - accountUpdate"))
}

// AccountChangePassword is the resolver for the accountChangePassword field.
func (r *mutationResolver) AccountChangePassword(ctx context.Context, data graph_model.AccountChangePassword) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountChangePassword - accountChangePassword"))
}

// AccountSetPassword is the resolver for the accountSetPassword field.
func (r *mutationResolver) AccountSetPassword(ctx context.Context, data graph_model.AccountSetPassword) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountSetPassword - accountSetPassword"))
}

// AccountDelete is the resolver for the accountDelete field.
func (r *mutationResolver) AccountDelete(ctx context.Context, data graph_model.AccountDelete) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountDelete - accountDelete"))
}

// AccountMe is the resolver for the accountMe field.
func (r *queryResolver) AccountMe(ctx context.Context) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountMe - accountMe"))
}

// AccountPagination is the resolver for the accountPagination field.
func (r *queryResolver) AccountPagination(ctx context.Context, page int, limit int, orderBy *string, search map[string]interface{}) (*graph_model.AccountPagination, error) {
	panic(fmt.Errorf("not implemented: AccountPagination - accountPagination"))
}

// AccountDetailBySearch is the resolver for the accountDetailBySearch field.
func (r *queryResolver) AccountDetailBySearch(ctx context.Context, key string) (*graph_model.Account, error) {
	panic(fmt.Errorf("not implemented: AccountDetailBySearch - accountDetailBySearch"))
}

// Mutation returns generated_admin.MutationResolver implementation.
func (r *Resolver) Mutation() generated_admin.MutationResolver { return &mutationResolver{r} }

// Query returns generated_admin.QueryResolver implementation.
func (r *Resolver) Query() generated_admin.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }