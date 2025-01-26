package frontendapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
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

type FrontendAPI interface {
	GetMenuHandler(params frontend.GetFrontendMenuParams, _ *models.Principal) middleware.Responder
	GetSliderHandler(params frontend.GetFrontendSliderParams, _ *models.Principal) middleware.Responder
	GetFileHandler(params frontend.GetFrontendFilesFileParams) middleware.Responder
	GetCategoriesHandler(_ frontend.GetFrontendCategoriesParams, _ *models.Principal) middleware.Responder
	GetContractsHandler(_ frontend.GetFrontendContractsParams, _ *models.Principal) middleware.Responder
	GetUsersHandler(params frontend.GetFrontendUsersParams, _ *models.Principal) middleware.Responder
	GetPagesHandler(params frontend.GetFrontendPageSlugParams, principal *models.Principal) middleware.Responder
	GetContactsHandler(params frontend.GetFrontendContactParams, principal *models.Principal) middleware.Responder
	PostContactsHandler(params frontend.PostFrontendContactFormParams) middleware.Responder
	GetFaqHandler(params frontend.GetFrontendFaqParams, principal *models.Principal) middleware.Responder
	GetHelpCenterHandler(params frontend.GetFrontendHelpCenterParams, principal *models.Principal) middleware.Responder
	GetBlogsHandler(params frontend.GetFrontendBlogParams, principal *models.Principal) middleware.Responder
	GetBlogsSlugHandler(params frontend.GetFrontendBlogSlugParams, principal *models.Principal) middleware.Responder
	GetWalletConnectHandler(params frontend.GetFrontendWalletConnectParams, principal *models.Principal) middleware.Responder
	GetCreateAndSellHandler(params frontend.GetFrontendCreateAndSellParams, principal *models.Principal) middleware.Responder
	XTokenAuth(token string) (*models.Principal, error)
}

func NewFrontendAPI(s *plutonium.Server) FrontendAPI {
	return &API{service: s}
}

func (a *API) GetMenuHandler(params frontend.GetFrontendMenuParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Provider == nil {
		return frontend.NewGetFrontendMenuBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "menu provider is required",
		})
	}

	items, err := a.service.
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

	return frontend.NewGetFrontendMenuOK().WithPayload(items)
}

func (a *API) GetSliderHandler(params frontend.GetFrontendSliderParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Provider == nil {
		return frontend.NewGetFrontendSliderBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "slider provider is required",
		})
	}

	items, err := a.service.
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

		return frontend.NewGetFrontendSliderInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}

	return frontend.NewGetFrontendSliderOK().WithPayload(items)
}

func (a *API) GetFileHandler(params frontend.GetFrontendFilesFileParams) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
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

		if _, err := w.Write(buf); err != nil {
			a.service.GetLogger().Errorln(err)
		}
	})
}

func (a *API) GetCategoriesHandler(_ frontend.GetFrontendCategoriesParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.
		GetStorage().
		GetCategoriesRepository().
		GetPublicCollectibleCategories(ctx, a.service.GetStorage().GetUsersRepository())
	if err != nil {
		return frontend.NewGetFrontendCategoriesInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}

	fmt.Println(items)

	return frontend.NewGetFrontendCategoriesOK().WithPayload(items)
}
func (a *API) GetContractsHandler(_ frontend.GetFrontendContractsParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	web3Dialer := a.service.GetWeb3Dialer()

	err := web3Dialer.Load(ctx)
	if err != nil {
		return frontend.NewGetFrontendContractsInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: SomethingsWentWrong,
		})
	}

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

func (a *API) GetUsersHandler(params frontend.GetFrontendUsersParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Address == nil {
		return frontend.NewGetFrontendUsersInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "users provider is required",
		})
	}

	user, err := a.service.
		GetStorage().
		GetUsersRepository().
		GetUserByAddress(ctx, *params.Address)

	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendUsersInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	payload := &models.PublicUser{
		//Address:  user.Address,
		Email: user.Email,
		ID:    user.ID,
		//Nonce:    user.Nonce,
		//Username: user.Nonce,
		//UUID: user.UUID,
	}

	return frontend.NewGetFrontendUsersOK().WithPayload([]*models.PublicUser{payload})
}

func (a *API) GetPagesHandler(params frontend.GetFrontendPageSlugParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	path := strings.Split(params.HTTPRequest.URL.RequestURI(), "/")
	slug := path[len(path)-1]
	fmt.Println(slug)

	page, err := a.service.
		GetStorage().
		GetPagesRepository().
		GetPageBySlug(ctx, slug)
	if err != nil {
		a.service.GetLogger().Error(err)

		if err.Error() == "[repository]: page not found" {
			return frontend.NewGetFrontendPageSlugNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "page not found",
			})
		}

		return frontend.NewGetFrontendPageSlugInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendPageSlugOK().WithPayload(page)
}

func (a *API) XTokenAuth(token string) (*models.Principal, error) {
	if token == a.service.GetSettings().XToken {
		principal := models.Principal(token)
		return &principal, nil
	}

	return nil, errors.New("incorrect api key authApi")
}

func (a *API) GetContactsHandler(params frontend.GetFrontendContactParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var pageIDParams = ""
	if params.ID != nil {
		pageIDParams = *params.ID
	}

	if pageIDParams == "" {
		return frontend.NewGetFrontendContactBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid pageId",
		})
	}

	pageID, err := strconv.ParseInt(pageIDParams, 10, 64)
	if err != nil {
		return frontend.NewGetFrontendContactInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	items, err := a.service.
		GetStorage().
		GetContactsRepository().
		GetContactByPageID(ctx, pageID)

	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendContactInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendContactOK().WithPayload(items)
}

func (a *API) PostContactsHandler(params frontend.PostFrontendContactFormParams) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Body == nil {
		return frontend.NewGetFrontendContactBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "invalid body",
		})
	}

	body := *params.Body
	attributes := make(map[string]interface{})
	attributes["page_id"] = body.PageID
	attributes["name"] = body.Name
	attributes["email"] = body.Email
	attributes["subject"] = body.Subject
	attributes["message"] = body.Message
	attributes["provider"] = body.Provider

	err := a.service.GetStorage().
		GetContactFormsRepository().
		Create(ctx, attributes)

	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewPostFrontendContactFormInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewPostFrontendContactFormOK().WithPayload(&models.FormSuccess{Message: "message sent"})
}

func (a *API) GetFaqHandler(_ frontend.GetFrontendFaqParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.GetStorage().
		GetFaqsRepository().
		GetPublicFaqs(ctx)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendFaqInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendFaqOK().WithPayload(items)
}

func (a *API) GetHelpCenterHandler(_ frontend.GetFrontendHelpCenterParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.GetStorage().
		GetHelpCenterRepository().
		GetPublicHelpCenter(ctx)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendHelpCenterInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendHelpCenterOK().WithPayload(items)
}

func (a *API) GetBlogsHandler(_ frontend.GetFrontendBlogParams, _ *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.GetStorage().
		GetBlogsRepository().
		GetPublicBlogs(ctx)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendBlogInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendBlogOK().WithPayload(items)
}

func (a *API) GetBlogsSlugHandler(params frontend.GetFrontendBlogSlugParams, _ *models.Principal) middleware.Responder {
	path := strings.Split(params.HTTPRequest.URL.RequestURI(), "/")
	slug := path[len(path)-1]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	item, err := a.service.GetStorage().
		GetBlogsRepository().
		GetPublicBlogItem(ctx, slug)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendBlogSlugInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendBlogSlugOK().WithPayload(item)
}

func (a *API) GetWalletConnectHandler(params frontend.GetFrontendWalletConnectParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.GetStorage().
		GetWalletsRepository().
		GetPublicWalletConnect(ctx)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendWalletConnectOK().WithPayload(items)
}
func (a *API) GetCreateAndSellHandler(params frontend.GetFrontendCreateAndSellParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	items, err := a.service.GetStorage().
		GetCreateAndSellRepository().
		GetPublicCreateAndSell(ctx)
	if err != nil {
		a.service.GetLogger().Error(err.Error())

		return frontend.NewGetFrontendCreateAndSellInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return frontend.NewGetFrontendCreateAndSellOK().WithPayload(items)
}
