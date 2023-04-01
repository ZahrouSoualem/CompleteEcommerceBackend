// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/zahrou/ecommerce/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/zahrou/ecommerce/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateBrand mocks base method.
func (m *MockStore) CreateBrand(arg0 context.Context, arg1 db.CreateBrandParams) (db.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBrand", arg0, arg1)
	ret0, _ := ret[0].(db.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBrand indicates an expected call of CreateBrand.
func (mr *MockStoreMockRecorder) CreateBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBrand", reflect.TypeOf((*MockStore)(nil).CreateBrand), arg0, arg1)
}

// CreateCategory mocks base method.
func (m *MockStore) CreateCategory(arg0 context.Context, arg1 string) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockStoreMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockStore)(nil).CreateCategory), arg0, arg1)
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(arg0 context.Context, arg1 db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), arg0, arg1)
}

// CreateMarket mocks base method.
func (m *MockStore) CreateMarket(arg0 context.Context, arg1 db.CreateMarketParams) (db.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMarket", arg0, arg1)
	ret0, _ := ret[0].(db.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMarket indicates an expected call of CreateMarket.
func (mr *MockStoreMockRecorder) CreateMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMarket", reflect.TypeOf((*MockStore)(nil).CreateMarket), arg0, arg1)
}

// CreateOrder mocks base method.
func (m *MockStore) CreateOrder(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockStoreMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockStore)(nil).CreateOrder), arg0, arg1)
}

// CreateProduct mocks base method.
func (m *MockStore) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStore)(nil).CreateProduct), arg0, arg1)
}

// CreateReview mocks base method.
func (m *MockStore) CreateReview(arg0 context.Context, arg1 db.CreateReviewParams) (db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", arg0, arg1)
	ret0, _ := ret[0].(db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockStoreMockRecorder) CreateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockStore)(nil).CreateReview), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// Createordersproduct mocks base method.
func (m *MockStore) Createordersproduct(arg0 context.Context, arg1 db.CreateordersproductParams) (db.Ordersproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Createordersproduct", arg0, arg1)
	ret0, _ := ret[0].(db.Ordersproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Createordersproduct indicates an expected call of Createordersproduct.
func (mr *MockStoreMockRecorder) Createordersproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Createordersproduct", reflect.TypeOf((*MockStore)(nil).Createordersproduct), arg0, arg1)
}

// DeleteBrand mocks base method.
func (m *MockStore) DeleteBrand(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBrand", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBrand indicates an expected call of DeleteBrand.
func (mr *MockStoreMockRecorder) DeleteBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBrand", reflect.TypeOf((*MockStore)(nil).DeleteBrand), arg0, arg1)
}

// DeleteCategories mocks base method.
func (m *MockStore) DeleteCategories(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategories", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategories indicates an expected call of DeleteCategories.
func (mr *MockStoreMockRecorder) DeleteCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategories", reflect.TypeOf((*MockStore)(nil).DeleteCategories), arg0, arg1)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), arg0, arg1)
}

// DeleteMarket mocks base method.
func (m *MockStore) DeleteMarket(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMarket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMarket indicates an expected call of DeleteMarket.
func (mr *MockStoreMockRecorder) DeleteMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMarket", reflect.TypeOf((*MockStore)(nil).DeleteMarket), arg0, arg1)
}

// DeleteOrder mocks base method.
func (m *MockStore) DeleteOrder(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockStoreMockRecorder) DeleteOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockStore)(nil).DeleteOrder), arg0, arg1)
}

// DeleteProducts mocks base method.
func (m *MockStore) DeleteProducts(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProducts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProducts indicates an expected call of DeleteProducts.
func (mr *MockStoreMockRecorder) DeleteProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducts", reflect.TypeOf((*MockStore)(nil).DeleteProducts), arg0, arg1)
}

// DeleteUsers mocks base method.
func (m *MockStore) DeleteUsers(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUsers", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUsers indicates an expected call of DeleteUsers.
func (mr *MockStoreMockRecorder) DeleteUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUsers", reflect.TypeOf((*MockStore)(nil).DeleteUsers), arg0, arg1)
}

// Deleteordersproduct mocks base method.
func (m *MockStore) Deleteordersproduct(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deleteordersproduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deleteordersproduct indicates an expected call of Deleteordersproduct.
func (mr *MockStoreMockRecorder) Deleteordersproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deleteordersproduct", reflect.TypeOf((*MockStore)(nil).Deleteordersproduct), arg0, arg1)
}

// GetBrand mocks base method.
func (m *MockStore) GetBrand(arg0 context.Context, arg1 int64) (db.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBrand", arg0, arg1)
	ret0, _ := ret[0].(db.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBrand indicates an expected call of GetBrand.
func (mr *MockStoreMockRecorder) GetBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBrand", reflect.TypeOf((*MockStore)(nil).GetBrand), arg0, arg1)
}

// GetCategory mocks base method.
func (m *MockStore) GetCategory(arg0 context.Context, arg1 int64) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockStoreMockRecorder) GetCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockStore)(nil).GetCategory), arg0, arg1)
}

// GetComment mocks base method.
func (m *MockStore) GetComment(arg0 context.Context, arg1 int64) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockStoreMockRecorder) GetComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockStore)(nil).GetComment), arg0, arg1)
}

// GetMarket mocks base method.
func (m *MockStore) GetMarket(arg0 context.Context, arg1 int64) (db.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarket", arg0, arg1)
	ret0, _ := ret[0].(db.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarket indicates an expected call of GetMarket.
func (mr *MockStoreMockRecorder) GetMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarket", reflect.TypeOf((*MockStore)(nil).GetMarket), arg0, arg1)
}

// GetOrders mocks base method.
func (m *MockStore) GetOrders(arg0 context.Context, arg1 int64) (db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0, arg1)
	ret0, _ := ret[0].(db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockStoreMockRecorder) GetOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockStore)(nil).GetOrders), arg0, arg1)
}

// GetProducts mocks base method.
func (m *MockStore) GetProducts(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockStoreMockRecorder) GetProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockStore)(nil).GetProducts), arg0, arg1)
}

// GetReview mocks base method.
func (m *MockStore) GetReview(arg0 context.Context, arg1 int64) (db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReview", arg0, arg1)
	ret0, _ := ret[0].(db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReview indicates an expected call of GetReview.
func (mr *MockStoreMockRecorder) GetReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReview", reflect.TypeOf((*MockStore)(nil).GetReview), arg0, arg1)
}

// GetUsers mocks base method.
func (m *MockStore) GetUsers(arg0 context.Context, arg1 int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockStoreMockRecorder) GetUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockStore)(nil).GetUsers), arg0, arg1)
}

// Getordersproduct mocks base method.
func (m *MockStore) Getordersproduct(arg0 context.Context, arg1 int64) (db.Ordersproduct, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Getordersproduct", arg0, arg1)
	ret0, _ := ret[0].(db.Ordersproduct)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Getordersproduct indicates an expected call of Getordersproduct.
func (mr *MockStoreMockRecorder) Getordersproduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Getordersproduct", reflect.TypeOf((*MockStore)(nil).Getordersproduct), arg0, arg1)
}

// ListBrands mocks base method.
func (m *MockStore) ListBrands(arg0 context.Context, arg1 db.ListBrandsParams) ([]db.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBrands", arg0, arg1)
	ret0, _ := ret[0].([]db.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBrands indicates an expected call of ListBrands.
func (mr *MockStoreMockRecorder) ListBrands(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBrands", reflect.TypeOf((*MockStore)(nil).ListBrands), arg0, arg1)
}

// ListCategories mocks base method.
func (m *MockStore) ListCategories(arg0 context.Context, arg1 db.ListCategoriesParams) ([]db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", arg0, arg1)
	ret0, _ := ret[0].([]db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockStoreMockRecorder) ListCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockStore)(nil).ListCategories), arg0, arg1)
}

// ListComments mocks base method.
func (m *MockStore) ListComments(arg0 context.Context, arg1 db.ListCommentsParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", arg0, arg1)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListComments indicates an expected call of ListComments.
func (mr *MockStoreMockRecorder) ListComments(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockStore)(nil).ListComments), arg0, arg1)
}

// ListJoinOrderProducts mocks base method.
func (m *MockStore) ListJoinOrderProducts(arg0 context.Context) ([]db.ListJoinOrderProductsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJoinOrderProducts", arg0)
	ret0, _ := ret[0].([]db.ListJoinOrderProductsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJoinOrderProducts indicates an expected call of ListJoinOrderProducts.
func (mr *MockStoreMockRecorder) ListJoinOrderProducts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJoinOrderProducts", reflect.TypeOf((*MockStore)(nil).ListJoinOrderProducts), arg0)
}

// ListJoinProducts mocks base method.
func (m *MockStore) ListJoinProducts(arg0 context.Context, arg1 db.ListJoinProductsParams) ([]db.ListJoinProductsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListJoinProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.ListJoinProductsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListJoinProducts indicates an expected call of ListJoinProducts.
func (mr *MockStoreMockRecorder) ListJoinProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListJoinProducts", reflect.TypeOf((*MockStore)(nil).ListJoinProducts), arg0, arg1)
}

// ListMarkets mocks base method.
func (m *MockStore) ListMarkets(arg0 context.Context, arg1 db.ListMarketsParams) ([]db.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMarkets", arg0, arg1)
	ret0, _ := ret[0].([]db.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMarkets indicates an expected call of ListMarkets.
func (mr *MockStoreMockRecorder) ListMarkets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMarkets", reflect.TypeOf((*MockStore)(nil).ListMarkets), arg0, arg1)
}

// ListOrders mocks base method.
func (m *MockStore) ListOrders(arg0 context.Context, arg1 db.ListOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockStoreMockRecorder) ListOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockStore)(nil).ListOrders), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStore) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStore)(nil).ListProducts), arg0, arg1)
}

// ListReviews mocks base method.
func (m *MockStore) ListReviews(arg0 context.Context, arg1 db.ListReviewsParams) ([]db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReviews", arg0, arg1)
	ret0, _ := ret[0].([]db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListReviews indicates an expected call of ListReviews.
func (mr *MockStoreMockRecorder) ListReviews(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReviews", reflect.TypeOf((*MockStore)(nil).ListReviews), arg0, arg1)
}

// ListUserOrders mocks base method.
func (m *MockStore) ListUserOrders(arg0 context.Context, arg1 db.ListUserOrdersParams) ([]db.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUserOrders", arg0, arg1)
	ret0, _ := ret[0].([]db.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUserOrders indicates an expected call of ListUserOrders.
func (mr *MockStoreMockRecorder) ListUserOrders(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUserOrders", reflect.TypeOf((*MockStore)(nil).ListUserOrders), arg0, arg1)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(arg0 context.Context, arg1 db.ListUsersParams) ([]db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", arg0, arg1)
	ret0, _ := ret[0].([]db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), arg0, arg1)
}

// Listordersproducts mocks base method.
func (m *MockStore) Listordersproducts(arg0 context.Context, arg1 db.ListordersproductsParams) ([]db.ListordersproductsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Listordersproducts", arg0, arg1)
	ret0, _ := ret[0].([]db.ListordersproductsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Listordersproducts indicates an expected call of Listordersproducts.
func (mr *MockStoreMockRecorder) Listordersproducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Listordersproducts", reflect.TypeOf((*MockStore)(nil).Listordersproducts), arg0, arg1)
}

// TransactionTX mocks base method.
func (m *MockStore) TransactionTX(arg0 context.Context, arg1 db.TransactionTxPrams) (db.TxResultParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionTX", arg0, arg1)
	ret0, _ := ret[0].(db.TxResultParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionTX indicates an expected call of TransactionTX.
func (mr *MockStoreMockRecorder) TransactionTX(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionTX", reflect.TypeOf((*MockStore)(nil).TransactionTX), arg0, arg1)
}

// UpdateBrand mocks base method.
func (m *MockStore) UpdateBrand(arg0 context.Context, arg1 db.UpdateBrandParams) (db.Brand, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBrand", arg0, arg1)
	ret0, _ := ret[0].(db.Brand)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBrand indicates an expected call of UpdateBrand.
func (mr *MockStoreMockRecorder) UpdateBrand(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBrand", reflect.TypeOf((*MockStore)(nil).UpdateBrand), arg0, arg1)
}

// UpdateCategory mocks base method.
func (m *MockStore) UpdateCategory(arg0 context.Context, arg1 db.UpdateCategoryParams) (db.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(db.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockStoreMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockStore)(nil).UpdateCategory), arg0, arg1)
}

// UpdateComment mocks base method.
func (m *MockStore) UpdateComment(arg0 context.Context, arg1 db.UpdateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockStoreMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockStore)(nil).UpdateComment), arg0, arg1)
}

// UpdateMarket mocks base method.
func (m *MockStore) UpdateMarket(arg0 context.Context, arg1 db.UpdateMarketParams) (db.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMarket", arg0, arg1)
	ret0, _ := ret[0].(db.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateMarket indicates an expected call of UpdateMarket.
func (mr *MockStoreMockRecorder) UpdateMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMarket", reflect.TypeOf((*MockStore)(nil).UpdateMarket), arg0, arg1)
}

// UpdateProductName mocks base method.
func (m *MockStore) UpdateProductName(arg0 context.Context, arg1 db.UpdateProductNameParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductName", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProductName indicates an expected call of UpdateProductName.
func (mr *MockStoreMockRecorder) UpdateProductName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductName", reflect.TypeOf((*MockStore)(nil).UpdateProductName), arg0, arg1)
}

// UpdateReview mocks base method.
func (m *MockStore) UpdateReview(arg0 context.Context, arg1 db.UpdateReviewParams) (db.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReview", arg0, arg1)
	ret0, _ := ret[0].(db.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReview indicates an expected call of UpdateReview.
func (mr *MockStoreMockRecorder) UpdateReview(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReview", reflect.TypeOf((*MockStore)(nil).UpdateReview), arg0, arg1)
}

// UpdateSoldProduct mocks base method.
func (m *MockStore) UpdateSoldProduct(arg0 context.Context, arg1 db.UpdateSoldProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSoldProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSoldProduct indicates an expected call of UpdateSoldProduct.
func (mr *MockStoreMockRecorder) UpdateSoldProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSoldProduct", reflect.TypeOf((*MockStore)(nil).UpdateSoldProduct), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}