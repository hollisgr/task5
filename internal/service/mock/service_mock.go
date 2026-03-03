package mock

import (
	"context"
	"task5/internal/model"

	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) Load(ctx context.Context, id int) (model.Movie, error) {
	args := m.Called(ctx, id)
	book, _ := args.Get(0).(model.Movie)
	return book, args.Error(1)
}

func (m *ServiceMock) Create(ctx context.Context, data model.Movie) (int, error) {
	args := m.Called(ctx, data)
	return args.Int(0), args.Error(1)
}
