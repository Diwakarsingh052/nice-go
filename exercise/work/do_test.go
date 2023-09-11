package work

import (
	mock_data "excercise/interface/data/mock"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	g := gomock.NewController(t)
	m := mock_data.NewMockClient(g)
	r, _ := http.NewRequest(http.MethodGet, "http:localhost:8080", nil)
	m.EXPECT().DoReq(r).Return("hello")
}
