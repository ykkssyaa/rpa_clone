package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"net/http"
	"strconv"

	"github.com/skinnykaen/rpa_clone/internal/consts"
	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// CreateRobboGroup is the resolver for the CreateRobboGroup field.
func (r *mutationResolver) CreateRobboGroup(ctx context.Context, input models.NewRobboGroup) (*models.RobboGroupHTTP, error) {
	atoi, err := strconv.Atoi(input.RobboUnitID)
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
	robboGroup, err := r.robboGroupService.CreateRobboGroup(models.RobboGroupCore{Name: input.Name, RobboUnitID: uint(atoi)})
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	robboGroupHttp := models.RobboGroupHTTP{}
	robboGroupHttp.FromCore(robboGroup)
	return &robboGroupHttp, nil
}

// UpdateRobboGroup is the resolver for the UpdateRobboGroup field.
func (r *mutationResolver) UpdateRobboGroup(ctx context.Context, input models.UpdateRobboGroup) (*models.RobboGroupHTTP, error) {
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
	robboGroup := models.RobboGroupCore{
		ID:   uint(atoi),
		Name: input.Name,
	}
	updatedRobboGroup, err := r.robboGroupService.UpdateRobboGroup(robboGroup)
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	robboGroupHttp := models.RobboGroupHTTP{}
	robboGroupHttp.FromCore(updatedRobboGroup)
	return &robboGroupHttp, nil
}

// DeleteRobboGroup is the resolver for the DeleteRobboGroup field.
func (r *mutationResolver) DeleteRobboGroup(ctx context.Context, id string) (*models.Response, error) {
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
	if err := r.robboGroupService.DeleteRobboGroup(uint(atoi)); err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.Response{Ok: true}, err
}

// GetRobboGroupByID is the resolver for the GetRobboGroupById field.
func (r *queryResolver) GetRobboGroupByID(ctx context.Context, id string) (*models.RobboGroupHTTP, error) {
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
	robboGroup, err := r.robboGroupService.GetRobboGroupById(uint(atoi))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	robboGroupHttp := models.RobboGroupHTTP{}
	robboGroupHttp.FromCore(robboGroup)
	return &robboGroupHttp, nil
}

// GetAllRobboGroupByAccessToken is the resolver for the GetAllRobboGroupByAccessToken field.
func (r *queryResolver) GetAllRobboGroupByAccessToken(ctx context.Context, page *int, pageSize *int) (*models.RobboGroupHTTPList, error) {
	robboGroups, countRows, err := r.robboGroupService.GetAllRobboGroups(page, pageSize, ctx.Value(consts.KeyRole).(models.Role))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: models.FromRobboGroupsCore(robboGroups),
		CountRows:   int(countRows),
	}, nil
}

// GetRobboGroupsByRobboUnitID is the resolver for the GetRobboGroupsByRobboUnitId field.
func (r *queryResolver) GetRobboGroupsByRobboUnitID(ctx context.Context, page *int, pageSize *int, robboUnitID string) (*models.RobboGroupHTTPList, error) {
	atoi, err := strconv.Atoi(robboUnitID)
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
	robboGroups, countRows, err := r.robboGroupService.GetRobboGroupsByRobboUnitById(page, pageSize, uint(atoi))
	if err != nil {
		r.loggers.Err.Printf("%s", err.Error())
		return nil, &gqlerror.Error{
			Extensions: map[string]interface{}{
				"err": err,
			},
		}
	}
	return &models.RobboGroupHTTPList{
		RobboGroups: models.FromRobboGroupsCore(robboGroups),
		CountRows:   int(countRows),
	}, nil
}
