package url_test

import (
	"testing"

	"github.com/alphakta/URLShortenerProject/domain/url"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (*MockRepository) GetTotalLinks() (int, error) {
	panic("unimplemented")
}

func (*MockRepository) IncrementClicks(shortURL string) error {
	panic("unimplemented")
}

func (m *MockRepository) GetClicksByID(id string) (int, error) {
	panic("unimplemented")
}

func (m *MockRepository) AddURL(longURL string) (string, error) {
	args := m.Called(longURL)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) FindLongURL(shortURL string) (string, error) {
	args := m.Called(shortURL)
	return args.String(0), args.Error(1)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	service := url.NewService(mockRepo)

	mockRepo.On("AddURL", "http://example.com").Return("abc123", nil)
	shortURL, err := service.Create("http://example.com")

	assert.NoError(t, err)
	assert.Equal(t, "abc123", shortURL)
	mockRepo.AssertExpectations(t)
}
