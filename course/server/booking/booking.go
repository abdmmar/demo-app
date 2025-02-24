package booking

import (
	"context"

	"github.com/imrenagicom/demo-app/course/booking"
	v1 "github.com/imrenagicom/demo-app/pkg/apiclient/course/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func New(svc Service) *Server {
	return &Server{
		service: svc,
	}
}

type Service interface {
	CreateBooking(ctx context.Context, req *v1.CreateBookingRequest) (*booking.Booking, error)
	ReserveBooking(ctx context.Context, req *v1.ReserveBookingRequest) (*booking.Booking, error)
	GetBooking(ctx context.Context, req *v1.GetBookingRequest) (*booking.Booking, error)
	ExpireBooking(ctx context.Context, req *v1.ExpireBookingRequest) error
	ListBookings(ctx context.Context, req *v1.ListBookingsRequest) ([]booking.Booking, string, error)
}

type Server struct {
	v1.UnimplementedBookingServiceServer

	service Service
}

func (s Server) CreateBooking(ctx context.Context, req *v1.CreateBookingRequest) (*v1.Booking, error) {
	b, err := s.service.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return b.ApiV1(), nil
}

func (s Server) ReserveBooking(ctx context.Context, req *v1.ReserveBookingRequest) (*v1.ReserveBookingResponse, error) {
	_, err := s.service.ReserveBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.ReserveBookingResponse{}, nil
}

func (s Server) GetBooking(ctx context.Context, req *v1.GetBookingRequest) (*v1.Booking, error) {
	b, err := s.service.GetBooking(ctx, req)
	if err != nil {
		return nil, err
	}
	return b.ApiV1(), nil
}

func (s Server) ExpireBooking(ctx context.Context, req *v1.ExpireBookingRequest) (*v1.ExpireBookingResponse, error) {
	logger := zerolog.Ctx(ctx)
	errLog := logger.Error().Str("domain", "booking_handler")
	requestID, ok := ctx.Value("request_id").(string)
	if !ok || requestID == "" {
		errLog.Str("status", "400").Msgf("request_id is empty")
		return nil, booking.ErrBookingAlreadyCompleted
	}

	err := s.service.ExpireBooking(ctx, req)
	if err != nil {
		log.Error().Str("request_id", requestID).Msg(err.Error())
		return nil, err
	}

	return &v1.ExpireBookingResponse{}, nil
}

func (s Server) ListBookings(ctx context.Context, req *v1.ListBookingsRequest) (*v1.ListBookingsResponse, error) {
	logger := zerolog.Ctx(ctx)

	bookings, _, err := s.service.ListBookings(ctx, req)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, err
	}
	var bks []*v1.Booking
	for _, b := range bookings {
		bks = append(bks, b.ApiV1())
	}
	return &v1.ListBookingsResponse{
		Bookings: bks,
	}, nil
}
