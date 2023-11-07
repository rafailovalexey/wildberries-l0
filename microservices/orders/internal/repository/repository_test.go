package repository

import (
	mockRepository "github.com/emptyhopes/orders/internal/repository/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mockRepository.NewMockOrderRepositoryInterface(ctrl)

	id := "4ca5aa9b-ced2-4f9f-8ffb-526bf1ab9469"

	order, isExist := repository.GetOrderCache(id)

	require.Nil(t, order)
	require.False(t, isExist)
}
