package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyID_WhenCreateANew_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.isValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANew_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.isValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANew_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.isValid(), "invalid tax")
}

func TestGivenAnValidParams_WhenICallNewOrder_ThenIShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.isValid())
}
