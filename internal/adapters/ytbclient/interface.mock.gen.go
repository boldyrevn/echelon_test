// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -source interface.go -destination interface.mock.gen.go -package ytbclient -mock_names IYoutubeClient=YoutubeClientMock
//

// Package ytbclient is a generated GoMock package.
package ytbclient

import (
	context "context"
	model "echelon_task/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// YoutubeClientMock is a mock of IYoutubeClient interface.
type YoutubeClientMock struct {
	ctrl     *gomock.Controller
	recorder *YoutubeClientMockMockRecorder
}

// YoutubeClientMockMockRecorder is the mock recorder for YoutubeClientMock.
type YoutubeClientMockMockRecorder struct {
	mock *YoutubeClientMock
}

// NewYoutubeClientMock creates a new mock instance.
func NewYoutubeClientMock(ctrl *gomock.Controller) *YoutubeClientMock {
	mock := &YoutubeClientMock{ctrl: ctrl}
	mock.recorder = &YoutubeClientMockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *YoutubeClientMock) EXPECT() *YoutubeClientMockMockRecorder {
	return m.recorder
}

// GetThumbnail mocks base method.
func (m *YoutubeClientMock) GetThumbnail(ctx context.Context, videoID string, quality model.ThumbnailQuality) (*model.Thumbnail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThumbnail", ctx, videoID, quality)
	ret0, _ := ret[0].(*model.Thumbnail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThumbnail indicates an expected call of GetThumbnail.
func (mr *YoutubeClientMockMockRecorder) GetThumbnail(ctx, videoID, quality any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThumbnail", reflect.TypeOf((*YoutubeClientMock)(nil).GetThumbnail), ctx, videoID, quality)
}