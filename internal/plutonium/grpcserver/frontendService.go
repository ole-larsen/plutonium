package grpcserver

import (
	"context"

	"connectrpc.com/connect"
	frontendv1 "github.com/ole-larsen/plutonium/gen/frontend/v1"
	"github.com/ole-larsen/plutonium/gen/frontend/v1/frontendv1connect"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FrontendServiceServer implements the FrontendsService API.
type FrontendServiceServer struct {
	frontendv1connect.UnimplementedFrontendServiceHandler
	logger  *log.Logger
	storage storage.DBStorageInterface
}

func (s *FrontendServiceServer) SetLogger(logger *log.Logger) *FrontendServiceServer {
	s.logger = logger
	return s
}

func (s *FrontendServiceServer) SetStorage(store storage.DBStorageInterface) *FrontendServiceServer {
	s.storage = store
	return s
}

func (s *FrontendServiceServer) Menu(
	ctx context.Context,
	request *connect.Request[frontendv1.MenuRequest],
) (*connect.Response[frontendv1.MenuResponse], error) {
	item, err := s.storage.
		GetMenusRepository().
		GetMenuByProvider(ctx, request.Msg.Provider)

	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.MenuResponse{
		Response: &frontendv1.MenuResponse_Data{
			Data: &frontendv1.PublicMenu{
				Id: item.ID,
				Attributes: &frontendv1.PublicMenuAttributes{
					Name:    item.Attributes.Name,
					Link:    item.Attributes.Link,
					OrderBy: item.Attributes.OrderBy,
					Items:   NestPublicMenu(item.Attributes.Items),
				},
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) Page(
	ctx context.Context,
	request *connect.Request[frontendv1.PageRequest],
) (*connect.Response[frontendv1.PageResponse], error) {
	slug := request.Msg.Provider

	page, err := s.storage.
		GetPagesRepository().
		GetPageBySlug(ctx, slug)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.PageResponse{
		Response: &frontendv1.PageResponse_Data{
			Data: NestPublicPage(page),
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) Contact(
	ctx context.Context,
	request *connect.Request[frontendv1.ContactRequest],
) (*connect.Response[frontendv1.ContactResponse], error) {
	pageID := request.Msg.PageId

	contact, err := s.storage.
		GetContactsRepository().
		GetContactByPageID(ctx, pageID)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.ContactResponse{
		Response: &frontendv1.ContactResponse_Data{
			Data: NestPublicContact(contact),
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) Faq(
	ctx context.Context,
	_ *connect.Request[frontendv1.FaqRequest],
) (*connect.Response[frontendv1.FaqResponse], error) {
	items, err := s.storage.
		GetFaqsRepository().
		GetPublicFaqs(ctx)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.FaqResponse{
		Response: &frontendv1.FaqResponse_Data{
			Data: &frontendv1.SuccessFaq{
				Faq: NestPublicFaqItems(items),
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) HelpCenter(
	ctx context.Context,
	_ *connect.Request[frontendv1.HelpCenterRequest],
) (*connect.Response[frontendv1.HelpCenterResponse], error) {
	items, err := s.storage.
		GetHelpCenterRepository().
		GetPublicHelpCenter(ctx)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.HelpCenterResponse{
		Response: &frontendv1.HelpCenterResponse_Data{
			Data: &frontendv1.SuccessHelpCenter{
				HelpCenter: NestPublicHelpCenterItems(items),
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) PostContact(
	ctx context.Context,
	request *connect.Request[frontendv1.PostContactRequest],
) (*connect.Response[frontendv1.PostContactResponse], error) {
	body := request.Msg.Body
	contactForm := make(map[string]interface{})
	contactForm["page_id"] = body.PageId
	contactForm["name"] = body.Name
	contactForm["email"] = body.Email
	contactForm["subject"] = body.Subject
	contactForm["message"] = body.Message
	contactForm["provider"] = body.Provider
	contactForm["csrf"] = body.Csrf
	s.logger.Infoln("send data from contact form", contactForm)

	err := s.storage.
		GetContactFormsRepository().
		Create(ctx, contactForm)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.PostContactResponse{
		Response: &frontendv1.PostContactResponse_Data{
			Data: "message successfully sent",
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) Slider(
	ctx context.Context,
	request *connect.Request[frontendv1.SliderRequest],
) (*connect.Response[frontendv1.SliderResponse], error) {
	slider, err := s.storage.
		GetSlidersRepository().
		GetSliderByProvider(ctx, request.Msg.Provider)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.SliderResponse{
		Response: &frontendv1.SliderResponse_Data{
			Data: &frontendv1.PublicSlider{
				Id: slider.ID,
				Attributes: &frontendv1.PublicSliderAttributes{
					SliderItems: NestPublicSlides(slider.Attributes.SliderItems),
				},
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *FrontendServiceServer) PostSubscribe(
	ctx context.Context,
	request *connect.Request[frontendv1.PostSubscribeRequest],
) (*connect.Response[frontendv1.PostSubscribeResponse], error) {
	body := request.Msg.Body
	contactForm := make(map[string]interface{})
	contactForm["email"] = body.Email
	contactForm["csrf"] = body.Csrf
	contactForm["provider"] = "subscribe"
	s.logger.Infoln("send data from subscribe form", contactForm)

	err := s.storage.
		GetContactFormsRepository().
		CreateSubscribe(ctx, contactForm)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &frontendv1.PostSubscribeResponse{
		Response: &frontendv1.PostSubscribeResponse_Data{
			Data: "message successfully sent",
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}
