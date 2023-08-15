package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"net/http"

	"github.com/skinnykaen/rpa_clone/internal/models"
	"github.com/skinnykaen/rpa_clone/pkg/utils"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// SetActivationByLink is the resolver for the SetActivationByLink field.
func (r *mutationResolver) SetActivationByLink(ctx context.Context, activationByLink bool) (*models.Response, error) {
	if err := r.settingsService.SetActivationByLink(activationByLink); err != nil {
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
	return &models.Response{
		Ok: true,
	}, nil
}

// GetSettings is the resolver for the GetSettings field.
func (r *queryResolver) GetSettings(ctx context.Context) (*models.Settings, error) {
	activationLink, err := r.settingsService.GetActivationByLink()
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
	return &models.Settings{
		ActivationByLink: activationLink,
	}, nil
}
