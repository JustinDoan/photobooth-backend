package controllers

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/justindoan/photobooth-backend/dto"
	"github.com/justindoan/photobooth-backend/models"
	mock_repository "github.com/justindoan/photobooth-backend/repository/mock"
	mock_services "github.com/justindoan/photobooth-backend/services/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepositories(ctrl)
	mockService := mock_services.NewMockServices(ctrl)

	controller := NewImageController(mockService, mockRepo)

	ctx := context.Background()
	imageRequest := dto.ProcessImageRequest{Name: "testImage"}

	mockRepo.EXPECT().ImageRepository.Create(ctx, &models.Image{Name: "testImage"}).Times(1)
	mockService.EXPECT().Forge.Image2Image(ctx, imageRequest).Return("processedImageData", nil)

	result := controller.ProcessImage(ctx, imageRequest)

	assert.Equal(t, "processedImageData", result)
}

func TestListImages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockRepositories(ctrl)
	controller := NewImageController(nil, mockRepo)

	ctx := context.Background()
	mockImages := []models.Image{{Name: "image1"}, {Name: "image2"}}

	mockRepo.EXPECT().ImageRepository.List(ctx).Return(mockImages, nil).Times(1)

	result, err := controller.ListImages(ctx)

	assert.NoError(t, err)
	assert.Equal(t, mockImages, result)
}
