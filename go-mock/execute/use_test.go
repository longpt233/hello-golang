package execute

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/longpt233/go-mock/mocks"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockHello := mocks.NewMockHelloRepo(mockCtrl)
	testUser := &usecase.HelloUc{Doer: mockHello}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call.
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
