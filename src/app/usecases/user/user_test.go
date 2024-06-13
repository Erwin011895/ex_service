package user

import (
	"errors"
	mockDTO "ex_service/mocks/app/dto/user"
	"testing"

	mockRepo "ex_service/mocks/infra/persistence/postgres/user"
	dto "ex_service/src/app/dto/user"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserUseCase struct {
	mock.Mock
}

type UserUseCaseList struct {
	suite.Suite
	mockDTO        *mockDTO.MockUsersDTO
	mockRepo       *mockRepo.MockUsersRepo
	useCase        UserUCInterface
	dtoRegisterReq *dto.RegisterReqDTO
	dtoLoginReq    *dto.LoginReqDTO
	dtoResp        *dto.RegisterRespDTO
}

func (suite *UserUseCaseList) SetupTest() {
	suite.mockDTO = new(mockDTO.MockUsersDTO)
	suite.mockRepo = new(mockRepo.MockUsersRepo)
	suite.useCase = NewUserUseCase(suite.mockRepo)

	suite.dtoRegisterReq = &dto.RegisterReqDTO{
		UserName: "jody",
		Password: "12345678",
	}

	suite.dtoLoginReq = &dto.LoginReqDTO{
		UserName: "jody",
		Password: "12345678",
	}

	suite.dtoResp = &dto.RegisterRespDTO{
		ID:       1,
		UserName: "jody",
		Password: "$2a$10$yTNi4oboulw7rTKj9XnjTug5n5ou0J.krCfogk/beqZzeNF8YodB2", // 12345678
		WalletID: 3,
		Token:    "asdfghjkl",
	}
}

func (u *UserUseCaseList) TestRegisterSuccess() {
	u.mockRepo.Mock.On("Register", u.dtoRegisterReq).Return(u.dtoResp, nil)
	_, err := u.useCase.Register(u.dtoRegisterReq)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestRegisterFailed() {
	u.mockRepo.Mock.On("Register", u.dtoRegisterReq).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Register(u.dtoRegisterReq)
	u.Equal(errors.New(mock.Anything), err)
}

func (u *UserUseCaseList) TestLoginSuccess() {
	u.mockRepo.Mock.On("Login", u.dtoLoginReq).Return(u.dtoResp, nil)
	_, err := u.useCase.Login(u.dtoLoginReq)
	u.Equal(nil, err)
}

func (u *UserUseCaseList) TestLoginFailed() {
	u.mockRepo.Mock.On("Login", u.dtoLoginReq).Return(nil, errors.New(mock.Anything))
	_, err := u.useCase.Login(u.dtoLoginReq)
	u.Equal(errors.New(mock.Anything), err)
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(UserUseCaseList))
}
