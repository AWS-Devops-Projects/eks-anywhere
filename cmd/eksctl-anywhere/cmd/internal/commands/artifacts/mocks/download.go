// Code generated by MockGen. DO NOT EDIT.
// Source: cmd/eksctl-anywhere/cmd/internal/commands/artifacts/download.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/aws/eks-anywhere/release/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
)

// MockReader is a mock of Reader interface.
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader.
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance.
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// ReadBundlesForVersion mocks base method.
func (m *MockReader) ReadBundlesForVersion(eksaVersion string) (*v1alpha1.Bundles, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBundlesForVersion", eksaVersion)
	ret0, _ := ret[0].(*v1alpha1.Bundles)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadBundlesForVersion indicates an expected call of ReadBundlesForVersion.
func (mr *MockReaderMockRecorder) ReadBundlesForVersion(eksaVersion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBundlesForVersion", reflect.TypeOf((*MockReader)(nil).ReadBundlesForVersion), eksaVersion)
}

// ReadChartsFromBundles mocks base method.
func (m *MockReader) ReadChartsFromBundles(bundles *v1alpha1.Bundles) []v1alpha1.Image {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadChartsFromBundles", bundles)
	ret0, _ := ret[0].([]v1alpha1.Image)
	return ret0
}

// ReadChartsFromBundles indicates an expected call of ReadChartsFromBundles.
func (mr *MockReaderMockRecorder) ReadChartsFromBundles(bundles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadChartsFromBundles", reflect.TypeOf((*MockReader)(nil).ReadChartsFromBundles), bundles)
}

// ReadImagesFromBundles mocks base method.
func (m *MockReader) ReadImagesFromBundles(bundles *v1alpha1.Bundles) ([]v1alpha1.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadImagesFromBundles", bundles)
	ret0, _ := ret[0].([]v1alpha1.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadImagesFromBundles indicates an expected call of ReadImagesFromBundles.
func (mr *MockReaderMockRecorder) ReadImagesFromBundles(bundles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadImagesFromBundles", reflect.TypeOf((*MockReader)(nil).ReadImagesFromBundles), bundles)
}

// MockImageMover is a mock of ImageMover interface.
type MockImageMover struct {
	ctrl     *gomock.Controller
	recorder *MockImageMoverMockRecorder
}

// MockImageMoverMockRecorder is the mock recorder for MockImageMover.
type MockImageMoverMockRecorder struct {
	mock *MockImageMover
}

// NewMockImageMover creates a new mock instance.
func NewMockImageMover(ctrl *gomock.Controller) *MockImageMover {
	mock := &MockImageMover{ctrl: ctrl}
	mock.recorder = &MockImageMoverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageMover) EXPECT() *MockImageMoverMockRecorder {
	return m.recorder
}

// Move mocks base method.
func (m *MockImageMover) Move(ctx context.Context, artifacts ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range artifacts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Move", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Move indicates an expected call of Move.
func (mr *MockImageMoverMockRecorder) Move(ctx interface{}, artifacts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, artifacts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Move", reflect.TypeOf((*MockImageMover)(nil).Move), varargs...)
}

// MockChartDownloader is a mock of ChartDownloader interface.
type MockChartDownloader struct {
	ctrl     *gomock.Controller
	recorder *MockChartDownloaderMockRecorder
}

// MockChartDownloaderMockRecorder is the mock recorder for MockChartDownloader.
type MockChartDownloaderMockRecorder struct {
	mock *MockChartDownloader
}

// NewMockChartDownloader creates a new mock instance.
func NewMockChartDownloader(ctrl *gomock.Controller) *MockChartDownloader {
	mock := &MockChartDownloader{ctrl: ctrl}
	mock.recorder = &MockChartDownloaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChartDownloader) EXPECT() *MockChartDownloaderMockRecorder {
	return m.recorder
}

// Download mocks base method.
func (m *MockChartDownloader) Download(ctx context.Context, artifacts ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range artifacts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Download", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Download indicates an expected call of Download.
func (mr *MockChartDownloaderMockRecorder) Download(ctx interface{}, artifacts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, artifacts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockChartDownloader)(nil).Download), varargs...)
}

// MockPackager is a mock of Packager interface.
type MockPackager struct {
	ctrl     *gomock.Controller
	recorder *MockPackagerMockRecorder
}

// MockPackagerMockRecorder is the mock recorder for MockPackager.
type MockPackagerMockRecorder struct {
	mock *MockPackager
}

// NewMockPackager creates a new mock instance.
func NewMockPackager(ctrl *gomock.Controller) *MockPackager {
	mock := &MockPackager{ctrl: ctrl}
	mock.recorder = &MockPackagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackager) EXPECT() *MockPackagerMockRecorder {
	return m.recorder
}

// Package mocks base method.
func (m *MockPackager) Package(folder, dstFile string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Package", folder, dstFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// Package indicates an expected call of Package.
func (mr *MockPackagerMockRecorder) Package(folder, dstFile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Package", reflect.TypeOf((*MockPackager)(nil).Package), folder, dstFile)
}
