package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"practice/whetherinfo/workers/handlers"
	"practice/whetherinfo/workers/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestHandlersImpl_GetCity(t *testing.T) {

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	service := mocks.NewMockService(ctrl)
	handlersImpl := handlers.NewHandlerImpl(service)

	type args struct {
		Method string
		URL    string
		Body   []byte
	}
	tests := []struct {
		name       string
		handlers   *handlers.HandlersImpl
		args       args
		wantStatus int
		fn         func()
	}{
		{
			"Test 1 :",
			handlersImpl,
			args{
				"GET",
				"/city/cityName",
				nil,
			},
			200,
			func() {
				service.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()

			req, err := http.NewRequest(tt.args.Method, tt.args.URL, bytes.NewBuffer(tt.args.Body))
			if err != nil {
				t.Fatal(err)
			}
			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.handlers.GetCity)
			// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
			// directly and pass in our Request and ResponseRecorder.
			handler.ServeHTTP(rr, req)
			// Check the status code is what we expect.
			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatus)
			}

		})
	}
}
