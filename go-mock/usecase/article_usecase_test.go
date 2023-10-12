package usecase_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/longpt233/go-mock/mocks"
	"github.com/longpt233/go-mock/usecase"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockArticleRepo := mocks.NewMockArticleRepo(mockCtrl)
	// Expect Do to be called once with 123 and "123" as parameters, and return hello from the mocked call.
	mockArticleRepo.EXPECT().Get(123).Return("hello").Times(1)

	// test vao business
	testUser := &usecase.ArticleUc{Repo: mockArticleRepo}
	out := testUser.Do()

	if out != "helodo" {
		t.Fail()
	}
}
