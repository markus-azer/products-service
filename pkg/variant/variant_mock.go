// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package variant is a generated GoMock package.
package variant

import (
	gomock "github.com/golang/mock/gomock"
	entity "github.com/markus-azer/products-service/pkg/entity"
	reflect "reflect"
)

// MockmessagesReader is a mock of messagesReader interface
type MockmessagesReader struct {
	ctrl     *gomock.Controller
	recorder *MockmessagesReaderMockRecorder
}

// MockmessagesReaderMockRecorder is the mock recorder for MockmessagesReader
type MockmessagesReaderMockRecorder struct {
	mock *MockmessagesReader
}

// NewMockmessagesReader creates a new mock instance
func NewMockmessagesReader(ctrl *gomock.Controller) *MockmessagesReader {
	mock := &MockmessagesReader{ctrl: ctrl}
	mock.recorder = &MockmessagesReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmessagesReader) EXPECT() *MockmessagesReaderMockRecorder {
	return m.recorder
}

// MockmessagesWriter is a mock of messagesWriter interface
type MockmessagesWriter struct {
	ctrl     *gomock.Controller
	recorder *MockmessagesWriterMockRecorder
}

// MockmessagesWriterMockRecorder is the mock recorder for MockmessagesWriter
type MockmessagesWriterMockRecorder struct {
	mock *MockmessagesWriter
}

// NewMockmessagesWriter creates a new mock instance
func NewMockmessagesWriter(ctrl *gomock.Controller) *MockmessagesWriter {
	mock := &MockmessagesWriter{ctrl: ctrl}
	mock.recorder = &MockmessagesWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockmessagesWriter) EXPECT() *MockmessagesWriterMockRecorder {
	return m.recorder
}

// SendMessage mocks base method
func (m_2 *MockmessagesWriter) SendMessage(m *entity.Message) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "SendMessage", m)
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockmessagesWriterMockRecorder) SendMessage(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockmessagesWriter)(nil).SendMessage), m)
}

// SendMessages mocks base method
func (m *MockmessagesWriter) SendMessages(messages []*entity.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessages", messages)
}

// SendMessages indicates an expected call of SendMessages
func (mr *MockmessagesWriterMockRecorder) SendMessages(messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockmessagesWriter)(nil).SendMessages), messages)
}

// MockMessagesRepository is a mock of MessagesRepository interface
type MockMessagesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessagesRepositoryMockRecorder
}

// MockMessagesRepositoryMockRecorder is the mock recorder for MockMessagesRepository
type MockMessagesRepositoryMockRecorder struct {
	mock *MockMessagesRepository
}

// NewMockMessagesRepository creates a new mock instance
func NewMockMessagesRepository(ctrl *gomock.Controller) *MockMessagesRepository {
	mock := &MockMessagesRepository{ctrl: ctrl}
	mock.recorder = &MockMessagesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessagesRepository) EXPECT() *MockMessagesRepositoryMockRecorder {
	return m.recorder
}

// SendMessage mocks base method
func (m_2 *MockMessagesRepository) SendMessage(m *entity.Message) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "SendMessage", m)
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockMessagesRepositoryMockRecorder) SendMessage(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockMessagesRepository)(nil).SendMessage), m)
}

// SendMessages mocks base method
func (m *MockMessagesRepository) SendMessages(messages []*entity.Message) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessages", messages)
}

// SendMessages indicates an expected call of SendMessages
func (mr *MockMessagesRepositoryMockRecorder) SendMessages(messages interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessages", reflect.TypeOf((*MockMessagesRepository)(nil).SendMessages), messages)
}

// MockstoreReader is a mock of storeReader interface
type MockstoreReader struct {
	ctrl     *gomock.Controller
	recorder *MockstoreReaderMockRecorder
}

// MockstoreReaderMockRecorder is the mock recorder for MockstoreReader
type MockstoreReaderMockRecorder struct {
	mock *MockstoreReader
}

// NewMockstoreReader creates a new mock instance
func NewMockstoreReader(ctrl *gomock.Controller) *MockstoreReader {
	mock := &MockstoreReader{ctrl: ctrl}
	mock.recorder = &MockstoreReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstoreReader) EXPECT() *MockstoreReaderMockRecorder {
	return m.recorder
}

// FindOneByID mocks base method
func (m *MockstoreReader) FindOneByID(id entity.ID) (*entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", id)
	ret0, _ := ret[0].(*entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID
func (mr *MockstoreReaderMockRecorder) FindOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockstoreReader)(nil).FindOneByID), id)
}

// FindOneByAttribute mocks base method
func (m *MockstoreReader) FindOneByAttribute(product entity.ID, attributes map[string]string) (*entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByAttribute", product, attributes)
	ret0, _ := ret[0].(*entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByAttribute indicates an expected call of FindOneByAttribute
func (mr *MockstoreReaderMockRecorder) FindOneByAttribute(product, attributes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByAttribute", reflect.TypeOf((*MockstoreReader)(nil).FindOneByAttribute), product, attributes)
}

// MockstoreWriter is a mock of storeWriter interface
type MockstoreWriter struct {
	ctrl     *gomock.Controller
	recorder *MockstoreWriterMockRecorder
}

// MockstoreWriterMockRecorder is the mock recorder for MockstoreWriter
type MockstoreWriterMockRecorder struct {
	mock *MockstoreWriter
}

// NewMockstoreWriter creates a new mock instance
func NewMockstoreWriter(ctrl *gomock.Controller) *MockstoreWriter {
	mock := &MockstoreWriter{ctrl: ctrl}
	mock.recorder = &MockstoreWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockstoreWriter) EXPECT() *MockstoreWriterMockRecorder {
	return m.recorder
}

// StoreCommand mocks base method
func (m *MockstoreWriter) StoreCommand(c *entity.Command) (*entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCommand", c)
	ret0, _ := ret[0].(*entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreCommand indicates an expected call of StoreCommand
func (mr *MockstoreWriterMockRecorder) StoreCommand(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCommand", reflect.TypeOf((*MockstoreWriter)(nil).StoreCommand), c)
}

// Create mocks base method
func (m *MockstoreWriter) Create(variant *entity.Variant) (*entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", variant)
	ret0, _ := ret[0].(*entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockstoreWriterMockRecorder) Create(variant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockstoreWriter)(nil).Create), variant)
}

// UpdateOne mocks base method
func (m *MockstoreWriter) UpdateOne(id entity.ID, variant *entity.UpdateVariant, version entity.Version) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", id, variant, version)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne
func (mr *MockstoreWriterMockRecorder) UpdateOne(id, variant, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockstoreWriter)(nil).UpdateOne), id, variant, version)
}

// DeleteOne mocks base method
func (m *MockstoreWriter) DeleteOne(id entity.ID, version entity.Version) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", id, version)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne
func (mr *MockstoreWriterMockRecorder) DeleteOne(id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockstoreWriter)(nil).DeleteOne), id, version)
}

// MockStoreRepository is a mock of StoreRepository interface
type MockStoreRepository struct {
	ctrl     *gomock.Controller
	recorder *MockStoreRepositoryMockRecorder
}

// MockStoreRepositoryMockRecorder is the mock recorder for MockStoreRepository
type MockStoreRepositoryMockRecorder struct {
	mock *MockStoreRepository
}

// NewMockStoreRepository creates a new mock instance
func NewMockStoreRepository(ctrl *gomock.Controller) *MockStoreRepository {
	mock := &MockStoreRepository{ctrl: ctrl}
	mock.recorder = &MockStoreRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStoreRepository) EXPECT() *MockStoreRepositoryMockRecorder {
	return m.recorder
}

// FindOneByID mocks base method
func (m *MockStoreRepository) FindOneByID(id entity.ID) (*entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", id)
	ret0, _ := ret[0].(*entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID
func (mr *MockStoreRepositoryMockRecorder) FindOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockStoreRepository)(nil).FindOneByID), id)
}

// FindOneByAttribute mocks base method
func (m *MockStoreRepository) FindOneByAttribute(product entity.ID, attributes map[string]string) (*entity.Variant, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByAttribute", product, attributes)
	ret0, _ := ret[0].(*entity.Variant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByAttribute indicates an expected call of FindOneByAttribute
func (mr *MockStoreRepositoryMockRecorder) FindOneByAttribute(product, attributes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByAttribute", reflect.TypeOf((*MockStoreRepository)(nil).FindOneByAttribute), product, attributes)
}

// StoreCommand mocks base method
func (m *MockStoreRepository) StoreCommand(c *entity.Command) (*entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreCommand", c)
	ret0, _ := ret[0].(*entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreCommand indicates an expected call of StoreCommand
func (mr *MockStoreRepositoryMockRecorder) StoreCommand(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreCommand", reflect.TypeOf((*MockStoreRepository)(nil).StoreCommand), c)
}

// Create mocks base method
func (m *MockStoreRepository) Create(variant *entity.Variant) (*entity.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", variant)
	ret0, _ := ret[0].(*entity.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockStoreRepositoryMockRecorder) Create(variant interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStoreRepository)(nil).Create), variant)
}

// UpdateOne mocks base method
func (m *MockStoreRepository) UpdateOne(id entity.ID, variant *entity.UpdateVariant, version entity.Version) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", id, variant, version)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne
func (mr *MockStoreRepositoryMockRecorder) UpdateOne(id, variant, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockStoreRepository)(nil).UpdateOne), id, variant, version)
}

// DeleteOne mocks base method
func (m *MockStoreRepository) DeleteOne(id entity.ID, version entity.Version) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", id, version)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne
func (mr *MockStoreRepositoryMockRecorder) DeleteOne(id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockStoreRepository)(nil).DeleteOne), id, version)
}

// Mockreader is a mock of reader interface
type Mockreader struct {
	ctrl     *gomock.Controller
	recorder *MockreaderMockRecorder
}

// MockreaderMockRecorder is the mock recorder for Mockreader
type MockreaderMockRecorder struct {
	mock *Mockreader
}

// NewMockreader creates a new mock instance
func NewMockreader(ctrl *gomock.Controller) *Mockreader {
	mock := &Mockreader{ctrl: ctrl}
	mock.recorder = &MockreaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockreader) EXPECT() *MockreaderMockRecorder {
	return m.recorder
}

// Mockwriter is a mock of writer interface
type Mockwriter struct {
	ctrl     *gomock.Controller
	recorder *MockwriterMockRecorder
}

// MockwriterMockRecorder is the mock recorder for Mockwriter
type MockwriterMockRecorder struct {
	mock *Mockwriter
}

// NewMockwriter creates a new mock instance
func NewMockwriter(ctrl *gomock.Controller) *Mockwriter {
	mock := &Mockwriter{ctrl: ctrl}
	mock.recorder = &MockwriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Mockwriter) EXPECT() *MockwriterMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *Mockwriter) Create(createVariantDTO CreateVariantDTO) (entity.ID, int32, []entity.ClientError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createVariantDTO)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].([]entity.ClientError)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockwriterMockRecorder) Create(createVariantDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*Mockwriter)(nil).Create), createVariantDTO)
}

// UpdateOne mocks base method
func (m *Mockwriter) UpdateOne(id entity.ID, version int32, updateVariantDTO UpdateVariantDTO) (int32, []entity.ClientError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", id, version, updateVariantDTO)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].([]entity.ClientError)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne
func (mr *MockwriterMockRecorder) UpdateOne(id, version, updateVariantDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*Mockwriter)(nil).UpdateOne), id, version, updateVariantDTO)
}

// Delete mocks base method
func (m *Mockwriter) Delete(id entity.ID, version int32) *entity.ClientError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, version)
	ret0, _ := ret[0].(*entity.ClientError)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockwriterMockRecorder) Delete(id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Mockwriter)(nil).Delete), id, version)
}

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUseCase) Create(createVariantDTO CreateVariantDTO) (entity.ID, int32, []entity.ClientError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createVariantDTO)
	ret0, _ := ret[0].(entity.ID)
	ret1, _ := ret[1].(int32)
	ret2, _ := ret[2].([]entity.ClientError)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockUseCaseMockRecorder) Create(createVariantDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), createVariantDTO)
}

// UpdateOne mocks base method
func (m *MockUseCase) UpdateOne(id entity.ID, version int32, updateVariantDTO UpdateVariantDTO) (int32, []entity.ClientError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOne", id, version, updateVariantDTO)
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].([]entity.ClientError)
	return ret0, ret1
}

// UpdateOne indicates an expected call of UpdateOne
func (mr *MockUseCaseMockRecorder) UpdateOne(id, version, updateVariantDTO interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOne", reflect.TypeOf((*MockUseCase)(nil).UpdateOne), id, version, updateVariantDTO)
}

// Delete mocks base method
func (m *MockUseCase) Delete(id entity.ID, version int32) *entity.ClientError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, version)
	ret0, _ := ret[0].(*entity.ClientError)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUseCaseMockRecorder) Delete(id, version interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), id, version)
}
