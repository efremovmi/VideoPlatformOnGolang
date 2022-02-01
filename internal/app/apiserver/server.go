package apiserver

import (
	"context"
	"database/sql"
	"errors"
	pb "github.com/genridarkbkru/registration/internal/app/grpc_server/github.com/genridarkbkru/registration"
	"github.com/genridarkbkru/registration/internal/app/model"
	"github.com/genridarkbkru/registration/internal/app/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"strconv"
)

const (
	sessionName = "VideoService"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
)

// Структура сервиса

type server struct {
	store store.Store
	pb.UnimplementedServiceRAServer
}

// Инициализация сервиса.
func NewServer(store store.Store) pb.ServiceRAServer {
	return &server{
		store: store,
	}
}

func (s *server) HandleUsersCreate(ctx context.Context, in *pb.RequestCreate) (*pb.ResponseCreate, error) {

	u := &model.User{
		Email:    in.Email,
		Password: in.Password,
	}

	if err := s.store.User().Create(u); err != nil {
		return &pb.ResponseCreate{Status: strconv.Itoa(http.StatusUnprocessableEntity), Id: -1, Email: "", Err: err.Error()}, status.New(codes.OK, "").Err()
	}
	u.Sanitize()
	return &pb.ResponseCreate{Status: strconv.Itoa(http.StatusCreated), Id: u.ID, Email: u.Email, Err: "No error"}, status.New(codes.OK, "").Err()
}

// Аутентификация ...

func (s *server) HandleSessionsCreate(ctx context.Context, in *pb.RequestCreateSessions) (*pb.ResponseCreateSessions, error) {

	u, err := s.store.User().FindByEmail(in.Email)
	if err != nil || !u.ComparePassword(in.Password) {
		return &pb.ResponseCreateSessions{Status: strconv.Itoa(http.StatusUnauthorized), Cookie: false, Err: err.Error()}, status.New(codes.OK, "").Err()
	}

	return &pb.ResponseCreateSessions{Status: strconv.Itoa(http.StatusOK), Cookie: true, Err: "No error"}, status.New(codes.OK, "").Err()
}

// Подключение к БД.
func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
