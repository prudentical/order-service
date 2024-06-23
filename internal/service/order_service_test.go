package service_test

import (
	"errors"
	"order-service/internal/dto"
	"order-service/internal/model"
	"order-service/internal/persistence"
	mock_persistence "order-service/internal/persistence/mock"
	"order-service/internal/service"
	mock_service "order-service/internal/service/mock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/shopspring/decimal"
	"go.uber.org/mock/gomock"
)

var _ = Describe("Order service", Label("order"), func() {

	var orders service.OrderService
	var dao *mock_persistence.MockOrderDAO
	var positions *mock_service.MockPositionService
	var ctrl *gomock.Controller

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		dao = mock_persistence.NewMockOrderDAO(ctrl)
		positions = mock_service.NewMockPositionService(ctrl)
		orders = service.NewOrderService(dao, positions)
	})
	AfterEach(func() {
		ctrl.Finish()
	})
	Describe("Create order", func() {
		Context("with invalid order type", func() {
			It("should return an invalid-order-type error", func() {
				order := dto.OrderDTO{Type: "hey"}
				_, err := orders.Create(order)
				Expect(err).To(MatchError(service.InvalidOrderTypeError{"hey"}))
			})
		})
		Context("buy", func() {
			Context("with failing to create the position", func() {
				It("should return an error", func() {
					positions.EXPECT().Create(gomock.Any()).Return(model.Position{}, errors.New(""))
					order := dto.OrderDTO{Type: "Buy"}
					_, err := orders.Create(order)
					Expect(err).NotTo(BeNil())
				})
			})
		})
		Context("Sell", func() {
			Context("with no position id", func() {
				It("should return a field-required error", func() {
					order := dto.OrderDTO{Type: "Sell"}
					_, err := orders.Create(order)
					Expect(err).To(MatchError(service.FieldRequiredError{Field: "PositionId", Type: "sell orders"}))
				})
			})
			Context("with no position", func() {
				It("should return a not-found error", func() {
					positions.EXPECT().GetById(gomock.Any()).Return(model.Position{}, persistence.RecordNotFoundError{})
					order := dto.OrderDTO{Type: "Sell", PositionId: 1}
					_, err := orders.Create(order)
					Expect(err).To(MatchError(service.NotFoundError{Type: model.Position{}, Id: order.PositionId}.Error()))
				})
			})
			Context("with closed position", func() {
				It("should return a modifying-closed-position error", func() {
					position := model.Position{Status: model.Closed}
					position.ID = 1
					positions.EXPECT().GetById(gomock.Any()).Return(position, nil)
					order := dto.OrderDTO{Type: "Sell", PositionId: position.ID}
					_, err := orders.Create(order)
					Expect(err).To(MatchError(service.ModifyingClosedPositionError{order.PositionId}))
				})
			})
			Context("with not enough amount to close the position", func() {
				It("should return the created order", func() {
					position := model.Position{
						Status: model.Open,
						Orders: []model.Order{
							{
								Type:         model.Buy,
								FilledAmount: decimal.NewFromInt(100),
							},
						},
					}
					positions.EXPECT().GetById(gomock.Any()).Return(position, nil)
					dao.EXPECT().Create(gomock.Any()).Return(model.Order{}, nil)
					order := dto.OrderDTO{
						Type:         "Sell",
						PositionId:   1,
						FilledAmount: decimal.NewFromInt(99),
					}
					_, err := orders.Create(order)
					Expect(err).To(BeNil())
				})
			})
			Context("with enough amount to close the position but fails to close", func() {
				It("should return an error", func() {
					position := model.Position{
						Status: model.Open,
						Orders: []model.Order{{Type: model.Buy, FilledAmount: decimal.NewFromInt(100)}},
					}
					positions.EXPECT().GetById(gomock.Any()).Return(position, nil)
					positions.EXPECT().Close(gomock.Any()).Return(errors.New(""))
					order := dto.OrderDTO{
						Type:         "Sell",
						PositionId:   1,
						FilledAmount: decimal.NewFromInt(100),
					}
					_, err := orders.Create(order)
					Expect(err).ToNot(BeNil())
				})
			})
			Context("with enough amount to close the position", func() {
				It("should close the position and return the created order", func() {
					position := model.Position{
						Status: model.Open,
						Orders: []model.Order{{Type: model.Buy, FilledAmount: decimal.NewFromInt(100)}},
					}
					positions.EXPECT().GetById(gomock.Any()).Return(position, nil)
					dao.EXPECT().Create(gomock.Any()).Return(model.Order{}, nil)
					positions.EXPECT().Close(gomock.Any()).Return(nil)
					order := dto.OrderDTO{
						Type:         "Sell",
						PositionId:   1,
						FilledAmount: decimal.NewFromInt(100),
					}
					_, err := orders.Create(order)
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
