package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"net/http"
	"strconv"

	"github.com/skinnykaen/rpa_clone/graph"
	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/utils"
)

// CreateUser is the resolver for the CreateUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (*models.UserHTTP, error) {
	// middlename не обязательное поле и может быть nil
	var middlename string
	if input.Middlename != nil {
		middlename = *input.Middlename
	}
	user := models.UserCore{
		Email:      input.Email,
		Password:   input.Password,
		Firstname:  input.Firstname,
		Lastname:   input.Lastname,
		Middlename: middlename,
		Nickname:   input.Nickname,
		IsActive:   true,
		Role:       input.Role,
	}
	newUser, err := r.userService.CreateUser(user, ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	userHttp := models.UserHTTP{}
	userHttp.FromCore(newUser)
	return &userHttp, nil
}

// UpdateUser is the resolver for the UpdateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUser) (*models.UserHTTP, error) {
	atoi, err := strconv.Atoi(input.ID)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	// TODO not required field
	user := models.UserCore{
		ID:         uint(atoi),
		Email:      input.Email,
		Firstname:  input.Firstname,
		Lastname:   input.Lastname,
		Middlename: input.Middlename,
		Nickname:   input.Nickname,
	}
	updatedUser, err := r.userService.UpdateUser(user, ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	userHttp := models.UserHTTP{}
	userHttp.FromCore(updatedUser)
	return &userHttp, nil
}

// DeleteUser is the resolver for the DeleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.Response, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	err = r.userService.DeleteUser(uint(atoi))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// SetUserIsActive is the resolver for the SetUserIsActive field.
func (r *mutationResolver) SetUserIsActive(ctx context.Context, id string, isActive bool) (*models.Response, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if err := r.userService.SetIsActive(uint(atoi), isActive); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// GetUserByAccessToken is the resolver for the GetUserByAccessToken field.
func (r *queryResolver) GetUserByAccessToken(ctx context.Context) (*models.UserHTTP, error) {
	user, err := r.userService.GetUserById(ctx.Value(consts.KeyId).(uint), ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	userHttp := models.UserHTTP{}
	userHttp.FromCore(user)
	return &userHttp, nil
}

// GetUserByID is the resolver for the GetUserById field.
func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*models.UserHTTP, error) {
	atoi, err := strconv.Atoi(id)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	user, err := r.userService.GetUserById(uint(atoi), ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	userHttp := models.UserHTTP{}
	userHttp.FromCore(user)
	return &userHttp, nil
}

// GetAllUsers is the resolver for the GetAllUsers field.
func (r *queryResolver) GetAllUsers(ctx context.Context, page *int, pageSize *int, active bool, roles []models.Role) (*models.UsersList, error) {
	users, countRows, err := r.userService.GetAllUsers(page, pageSize, active, roles, ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return &models.UsersList{}, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				},
			},
		}
	}
	return &models.UsersList{
		Users:     models.FromUsersCore(users),
		CountRows: int(countRows),
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
