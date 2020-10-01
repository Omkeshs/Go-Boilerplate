package services_test

import (
	"context"
	"practice/whetherinfo/workers/mocks"
	service "practice/whetherinfo/workers/service"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestServiceImpl_GetCity(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositories := mocks.NewMockRepository(ctrl)
	svc := service.NewServiceImpl(repositories)

	type args struct {
		ctx      context.Context
		cityName string
	}
	tests := []struct {
		name    string
		service service.Service
		args    args
		// wantResp *models.Cities
		wantErr bool
		fn      func()
	}{
		// TODO: Add test cases.
		{
			"test 1 :",
			svc,
			args{
				context.Background(),
				"nashik",
			},
			false,
			func() {
				repositories.EXPECT().GetCity(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn()
			_, err := tt.service.GetCity(tt.args.ctx, tt.args.cityName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServiceImpl.GetCity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotResp, tt.wantResp) {
			// 	t.Errorf("ServiceImpl.GetCity() = %v, want %v", gotResp, tt.wantResp)
			// }
		})
	}
}
