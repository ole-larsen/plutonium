package frontendapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/frontend"
)

const UploadDir = "./uploads"

type API struct {
	service *plutonium.Server
}

func NewFrontendAPI(s *plutonium.Server) *API {
	return &API{service: s}
}

func (a *API) GetMenuHandler(params frontend.GetFrontendMenuParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if params.Provider == nil {
		return frontend.NewGetFrontendMenuBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "provider is required",
		})
	}

	menu, err := a.service.
		GetStorage().
		GetMenusRepository().
		GetMenuByProvider(ctx, *params.Provider)

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

	return frontend.NewGetFrontendMenuOK().WithPayload(menu)
}

func (a *API) GetSliderHandler(params frontend.GetFrontendSliderParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if params.Provider == nil {
		return frontend.NewGetFrontendSliderBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "provider is required",
		})
	}
	slider, err := a.service.
		GetStorage().
		GetSlidersRepository().
		GetSliderByProvider(ctx, *params.Provider)

	if err != nil {
		if strings.Contains(err.Error(), SliderNotFound) {
			return frontend.NewGetFrontendSliderNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "[" + *params.Provider + "] " + SliderNotFound,
			})
		}

		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendSliderInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}

	return frontend.NewGetFrontendSliderOK().WithPayload(slider)
}

func (a *API) GetFileHandler(params frontend.GetFrontendFilesFileParams) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, p runtime.Producer) {

		path := strings.Split(params.HTTPRequest.URL.RequestURI(), "/")

		encodedFilename := path[len(path)-1]

		filename, err := url.QueryUnescape(encodedFilename)
		if err != nil {
			a.service.GetLogger().Errorln(err)
		}
		buf, err := os.ReadFile(fmt.Sprintf("%s/%s", UploadDir, filename))

		if err != nil {
			a.service.GetLogger().Errorln(err)
		}
		ext := strings.Replace(filepath.Ext(filename), ".", "", 1)
		if ext == "svg" {
			w.Header().Set("Content-Type", fmt.Sprintf("image/%s+xml", ext))
		} else {
			w.Header().Set("Content-Type", fmt.Sprintf("image/%s", ext))
		}
		w.Write(buf)
	})
}

func (a *API) GetCategoriesHandler(params frontend.GetFrontendCategoriesParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	categories, err := a.service.GetStorage().GetCategoriesRepository().GetPublicCollectibleCategories(ctx, a.service.GetStorage().GetUsersRepository())
	if err != nil {
		return frontend.NewGetFrontendCategoriesInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}
	fmt.Println(categories)

	return frontend.NewGetFrontendCategoriesOK().WithPayload(categories)
}
func (a *API) GetContractsHandler(params frontend.GetFrontendContractsParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	web3Dialer := a.service.GetWeb3Dialer()
	web3Dialer.Load(ctx)
	collections := make(map[string]models.PublicContract)
	for collectionID, collection := range web3Dialer.Market.Collections {
		collections[collectionID] = models.PublicContract{
			Name:    collection.Name,
			Abi:     collection.ABI,
			Address: collection.Address.String(),
		}
	}
	auctions := make([]*models.PublicContract, len(web3Dialer.Market.Auctions))
	for i, auction := range web3Dialer.Market.Auctions {
		auctions[i] = &models.PublicContract{
			Name:    auction.Name,
			Abi:     auction.ABI,
			Address: auction.Address.String(),
		}
	}
	marketPlace := &models.PublicMarketplaceContract{
		Name:    web3Dialer.Market.Marketplace.Name,
		Abi:     web3Dialer.Market.Marketplace.ABI,
		Address: web3Dialer.Market.Marketplace.Address.String(),
		Fee:     *web3Dialer.Market.Marketplace.Fee,
		Owner:   web3Dialer.Market.Marketplace.Owner,
	}

	payload := &models.PublicContracts{
		Contracts: &models.PublicContractsContracts{
			Auctions:    auctions,
			Collections: collections,
			Marketplace: marketPlace,
		},
	}
	return frontend.NewGetFrontendContractsOK().WithPayload(payload)
}
func (a *API) XTokenAuth(token string) (*models.Principal, error) {
	if token == a.service.GetSettings().XToken {
		principal := models.Principal(token)
		return &principal, nil
	}

	return nil, errors.New("incorrect api key authApi")
}
