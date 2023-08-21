package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateParentRel is the resolver for the CreateParentRel field.
func (r *mutationResolver) CreateParentRel(ctx context.Context, parentID string, childID string) (*models.Response, error) {
	atoiP, errP := strconv.Atoi(parentID)
	atoiC, errC := strconv.Atoi(childID)
	if errP != nil {
		r.loggers.Err.Printf("%s", errP.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if errC != nil {
		r.loggers.Err.Printf("%s", errC.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	_, err := r.parentRelService.CreateRel(uint(atoiP), uint(atoiC))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.Response{Ok: true}, nil
}

// DeleteParentRel is the resolver for the DeleteParentRel field.
func (r *mutationResolver) DeleteParentRel(ctx context.Context, parentID string, childID string) (*models.Response, error) {
	atoiP, errP := strconv.Atoi(parentID)
	atoiC, errC := strconv.Atoi(childID)
	if errP != nil {
		r.loggers.Err.Printf("%s", errP.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if errC != nil {
		r.loggers.Err.Printf("%s", errC.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": utils.ResponseError{
					Code:    http.StatusBadRequest,
					Message: consts.ErrAtoi,
				},
			},
		}
	}
	if err := r.parentRelService.DeleteRel(uint(atoiP), uint(atoiC)); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	} else {
		return &models.Response{Ok: true}, nil
	}
}

// GetChildrenByParent is the resolver for the GetChildrenByParent field.
func (r *queryResolver) GetChildrenByParent(ctx context.Context, parentID string) (*models.UsersList, error) {
	atoi, err := strconv.Atoi(parentID)
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
	children, err := r.parentRelService.GetChildrenByParentId(uint(atoi))
	if err != nil {
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.UsersList{
		Users:     models.FromUsersCore(children),
		CountRows: len(children),
	}, nil
}

// GetParentsByChild is the resolver for the GetParentsByChild field.
func (r *queryResolver) GetParentsByChild(ctx context.Context, childID string) (*models.UsersList, error) {
	panic(fmt.Errorf("not implemented: GetParentsByChild - GetParentsByChild"))
}
