// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	requests "github.com/Akshayvij07/ecommerce/pkg/helper/request"
	response "github.com/Akshayvij07/ecommerce/pkg/helper/respondse"
	domain "github.com/Akshayvij07/ecommerce/pkg/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AddAdress mocks base method.
func (m *MockUserRepository) AddAdress(ctx context.Context, UserID int, address requests.AddressReq) (domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAdress", ctx, UserID, address)
	ret0, _ := ret[0].(domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAdress indicates an expected call of AddAdress.
func (mr *MockUserRepositoryMockRecorder) AddAdress(ctx, UserID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdress", reflect.TypeOf((*MockUserRepository)(nil).AddAdress), ctx, UserID, address)
}

//SaveUserAdress mock base method.
func (m *MockUserRepository) SaveUserAddress(ctx context.Context, userAddress domain.UserAddress) error{
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m,"SaveUserAddress",ctx,userAddress)
	ret0,_ := ret[0].(error)
	return ret0
}
//SaveAddress indicated an expected call of SaveUserAddress
func (mr *MockUserRepositoryMockRecorder) SaveUserAddress(ctx, Useraddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdress", reflect.TypeOf((*MockUserRepository)(nil).SaveUserAddress), ctx, Useraddress)
}

// FindAllWishListItemsByUserID mocks base method.
func (m *MockUserRepository) FindAllWishListItemsByUserID(ctx context.Context, userID uint) ([]response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWishListItemsByUserID", ctx, userID)
	ret0, _ := ret[0].([]response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWishListItemsByUserID indicates an expected call of FindAllWishListItemsByUserID.
func (mr *MockUserRepositoryMockRecorder) FindAllWishListItemsByUserID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWishListItemsByUserID", reflect.TypeOf((*MockUserRepository)(nil).FindAllWishListItemsByUserID), ctx, userID)
}

// FindProduct mocks base method.
func (m *MockUserRepository) FindProduct(ctx context.Context, id uint) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProduct", ctx, id)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProduct indicates an expected call of FindProduct.
func (mr *MockUserRepositoryMockRecorder) FindProduct(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProduct", reflect.TypeOf((*MockUserRepository)(nil).FindProduct), ctx, id)
}

// FindWishListItem mocks base method.
func (m *MockUserRepository) FindWishListItem(ctx context.Context, productID, userID uint) (domain.WishList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWishListItem", ctx, productID, userID)
	ret0, _ := ret[0].(domain.WishList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWishListItem indicates an expected call of FindWishListItem.
func (mr *MockUserRepositoryMockRecorder) FindWishListItem(ctx, productID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWishListItem", reflect.TypeOf((*MockUserRepository)(nil).FindWishListItem), ctx, productID, userID)
}

// OtpLogin mocks base method.
func (m *MockUserRepository) OtpLogin(mbnum string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OtpLogin", mbnum)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OtpLogin indicates an expected call of OtpLogin.
func (mr *MockUserRepositoryMockRecorder) OtpLogin(mbnum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OtpLogin", reflect.TypeOf((*MockUserRepository)(nil).OtpLogin), mbnum)
}

// Find User mocks base method.
func (m *MockUserRepository) FindUser(ctx context.Context, UsersId int) (domain.Users, error)  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUser", ctx, UsersId)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUser indicates an expected call of User.
func (mr *MockUserRepositoryMockRecorder) FindUser(ctx,userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUser", reflect.TypeOf((*MockUserRepository)(nil).FindUser), ctx,userId)
}


// RemoveWishListItem mocks base method.
func (m *MockUserRepository) RemoveWishListItem(ctx context.Context, wishList domain.WishList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWishListItem", ctx, wishList)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveWishListItem indicates an expected call of RemoveWishListItem.
func (mr *MockUserRepositoryMockRecorder) RemoveWishListItem(ctx, wishList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWishListItem", reflect.TypeOf((*MockUserRepository)(nil).RemoveWishListItem), ctx, wishList)
}

// SaveWishListItem mocks base method.
func (m *MockUserRepository) SaveWishListItem(ctx context.Context, wishList domain.WishList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveWishListItem", ctx, wishList)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveWishListItem indicates an expected call of SaveWishListItem.
func (mr *MockUserRepositoryMockRecorder) SaveWishListItem(ctx, wishList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveWishListItem", reflect.TypeOf((*MockUserRepository)(nil).SaveWishListItem), ctx, wishList)
}

// UpdateAdress mocks base method.
func (m *MockUserRepository) UpdateAddress(ctx context.Context, address domain.Address) error  {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdress", ctx,  address)
	ret0, _ := ret[1].(error)
	return ret0
}

// UpdateAdress indicates an expected call of UpdateAdress.
func (mr *MockUserRepositoryMockRecorder) UpdateAdress(ctx, UserID, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), ctx, UserID, address)
}


//UpdateUserAddressj mock base method.
func (m *MockUserRepository) UpdateUserAdress(ctx context.Context, userAddress domain.UserAddress) error{
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m,"UpdateUserAddress",ctx,userAddress)
	ret0,_ := ret[0].(error)
	return ret0
}

//UpdateUserAddress indicates an expected call of UpdateUserAddress
func (mr *MockUserRepositoryMockRecorder) UpdateUserAdress(ctx, Useraddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserAdress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), ctx, Useraddress)
}

// UpdatePassword mock base method 
func (m *MockUserRepository) UpdatePassword(ctx context.Context, UserID int, Password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", ctx, UserID, Password)
	ret0, _ := ret[1].(error)
	return ret0
}
// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserRepositoryMockRecorder) UpdatePassword(ctx, UserID, Password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUserRepository)(nil).UpdatePassword), ctx, UserID, Password)
}

// UserLogin mocks base method.
func (m *MockUserRepository) UserLogin(ctx context.Context, Email string) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", ctx, Email)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserRepositoryMockRecorder) UserLogin(ctx, Email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserRepository)(nil).UserLogin), ctx, Email)
}

// UserSignup mocks base method.
func (m *MockUserRepository) UserSignup(ctx context.Context, user requests.UserSign) (response.UserValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignup", ctx, user)
	ret0, _ := ret[0].(response.UserValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignup indicates an expected call of UserSignup.
func (mr *MockUserRepositoryMockRecorder) UserSignup(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignup", reflect.TypeOf((*MockUserRepository)(nil).UserSignup), ctx, user)
}

// VeiwAdress mocks base method.
func (m *MockUserRepository) VeiwAdress(ctx context.Context, UserID int) ([]response.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VeiwAdress", ctx, UserID)
	ret0, _ := ret[0].([]response.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VeiwAdress indicates an expected call of VeiwAdress.
func (mr *MockUserRepositoryMockRecorder) VeiwAdress(ctx, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VeiwAdress", reflect.TypeOf((*MockUserRepository)(nil).VeiwAdress), ctx, UserID)
}


//FindAdressByUserDetails mock base method.
func (m *MockUserRepository) FindAddressByUserDetails(ctx context.Context, AddressReq requests.AddressReq, UserID int)(domain.Address, error){
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddressByuserDetails", ctx, UserID)
	ret0, _ := ret[0].(domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

//FindAdressByUserDetails an expected call of FindAddressByUserDetails
func (mr *MockUserRepositoryMockRecorder) FindAddressByUserDetails(ctx, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddressByUserDetails", reflect.TypeOf((*MockUserRepository)(nil).FindAddressByUserDetails), ctx, UserID)
}

//FindUserAddressById mock base method.
func (m *MockUserRepository) FindUserAddressById(ctx context.Context, userId int) (domain.Address, error){
	m.ctrl.T.Helper()
	ret:= m.ctrl.Call(m,"FindUserAddressById",ctx,userId)
	ret0,_ := ret[0].(domain.Address)
	ret1,_ := ret[1].(error)
	return ret0,ret1
}
//FindUserAddressById an expected call of FindUserAddressById
func (mr *MockUserRepositoryMockRecorder) FindUserAddressById(ctx, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,"FindUserAddressById",reflect.TypeOf((*MockUserRepository)(nil).FindUserAddressById),ctx,UserID)
}

//GetInvoice mock base method.
func (m *MockUserRepository) GetInvoice(ctx context.Context, UserId int) (response.Invoice, error){
	m.ctrl.T.Helper()
	ret:= m.ctrl.Call(m,"FindUserAddressById",ctx,UserId)
	ret0,_ := ret[0].(response.Invoice)
	ret1,_ := ret[1].(error)
	return ret0,ret1
}

//GetInvoice an expected call of GetInvoice
func (mr *MockUserRepositoryMockRecorder) GetInvoice(ctx, UserID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock,"GetInvoice",reflect.TypeOf((*MockUserRepository)(nil).GetInvoice),ctx,UserID)
}
