package frontendapi

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/frontend"
)

type API struct {
	service *plutonium.Server
}

func NewFrontendAPI(s *plutonium.Server) *API {
	return &API{service: s}
}

func (a *API) GetMenuHandler(params frontend.GetFrontendMenuParams, _ *models.Principal) middleware.Responder {
	if params.Provider == nil {
		return frontend.NewGetFrontendMenuBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "provider is required",
		})
	}

	menu, err := a.service.GetStorage().GetMenusRepository().GetMenuByProvider(*params.Provider)
	if err != nil {
		if strings.Contains(err.Error(), MenuNotFound) {
			return frontend.NewGetFrontendMenuNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "[" + *params.Provider + "] " + MenuNotFound,
			})
		}

		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendMenuInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}

	payload := &models.PublicMenuResponse{
		Menu: menu,
	}

	return frontend.NewGetFrontendMenuOK().WithPayload(payload)
}

func (a *API) XTokenAuth(token string) (*models.Principal, error) {
	if token == a.service.GetSettings().XToken {
		principal := models.Principal(token)
		return &principal, nil
	}

	return nil, errors.New("incorrect api key authApi")
}
