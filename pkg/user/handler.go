package user

import (
	"context"
	"database/sql"
	"log"

	pb "github.com/danielgyu/go-ecommerce/internal/proto"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	pb.UnimplementedUserServiceServer
	repo    *userRepository
	session map[int]string
}

func NewUserHandler(db *sql.DB) *userHandler {
	r := NewUserRepository(db)
	s := make(map[int]string)
	return &userHandler{repo: r, session: s}
}

func (h *userHandler) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest) (out *pb.HealthCheckResponse, err error) {
	return &pb.HealthCheckResponse{StatusCode: 200}, nil
}

func (h *userHandler) SignUp(ctx context.Context, in *pb.SignUpRequest) (out *pb.SignUpResponse, err error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(in.Password), 8)
	if err != nil {
		return &pb.SignUpResponse{}, err
	}

	err = h.repo.RegisterUser(in.Username, string(hashedPw))
	if err != nil {
		log.Printf("signup error %v\n", err)
		return &pb.SignUpResponse{}, err
	}
	return &pb.SignUpResponse{Success: true}, nil
}

func (h *userHandler) LogIn(ctx context.Context, in *pb.LogInRequest) (out *pb.LogInResponse, err error) {
	userId, err := h.repo.LogInUser(in.Username, in.Password)
	if err != nil {
		return &pb.LogInResponse{}, err
	}

	sessionToken := uuid.New().String()
	h.session[userId] = sessionToken

	return &pb.LogInResponse{Token: sessionToken}, nil
}

func (h *userHandler) AddCredit(ctx context.Context, in *pb.AddCreditRequest) (out *pb.AddCreditResponse, err error) {
	return
}
