package controllers

import (
    "context"
    "testing"

    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
)

// MockServices is a mock implementation of the Services struct
var MockServices = new(MockedServices)

// MockedServices is a mock type for the Services type
type MockedServices struct {
    mock.Mock
}

// Image2Image mocks the Image2Image function of the Forge service
func (m *MockedServices) Image2Image(ctx context.Context, image dto.ProcessImageRequest) (string, error) {
    args := m.Called(ctx, image)
    return args.String(0), args.Error(1)
}

// TestImageController tests the ImageController
func TestImageController(t *testing.T) {
    t.Run("ProcessImage successfully processes an image", func(t *testing.T) {
        ctx := context.Background()
        imageRequest := dto.ProcessImageRequest{Name: "test-image"}

        // Setup expectations for the mock services
        MockServices.On("Image2Image", ctx, imageRequest).Return("processed-image-data", nil)

        controller := NewImageController(MockServices, nil)
        result := controller.ProcessImage(ctx, imageRequest)

        require.Equal(t, "processed-image-data", result)
    })

    t.Run("ListImages retrieves all images", func(t *testing.T) {
        ctx := context.Background()

        // Mock repository
        mockRepo := new(MockedRepository)
        mockRepo.On("List", ctx).Return([]models.Image{{Name: "image1"}, {Name: "image2"}}, nil)

        controller := NewImageController(MockServices, mockRepo)
        images, err := controller.ListImages(ctx)

        require.NoError(t, err)
        require.Len(t, images, 2)
    })
}

// MockedRepository is a mock type for the Repositories type
type MockedRepository struct {
    mock.Mock
}

// List mocks the List function of the ImageRepository
func (m *MockedRepository) List(ctx context.Context) ([]models.Image, error) {
    args := m.Called(ctx)
    return args.Get(0).([]models.Image), args.Error(1)
}
