package internal_test

import (
	"fmt"
	"github.com/sanderaido/maroto/internal"
	"github.com/sanderaido/maroto/internal/mocks"
	"github.com/sanderaido/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewSignature(t *testing.T) {
	signature := internal.NewSignature(&mocks.Pdf{}, &mocks.Math{}, &mocks.Text{})

	assert.NotNil(t, signature)
	assert.Equal(t, fmt.Sprintf("%T", signature), "*internal.signature")
}

func TestSignature_AddSpaceFor_DefaultMargins(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
	pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(50.0)

	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	signature := internal.NewSignature(pdf, math, text)

	// Act
	signature.AddSpaceFor("label", internal.Cell{5, 5, 2, 0}, props.Text{Size: 10.0})

	// Assert
	pdf.AssertNumberOfCalls(t, "Line", 1)
	pdf.AssertCalled(t, "Line", 19.0, 15.0, 13.0, 15.0)
	text.AssertNumberOfCalls(t, "Add", 1)
	text.AssertCalled(t, "Add", "label", internal.Cell{5.0, 7.0, 2.0, 0.0}, props.Text{Size: 10.0})
}

func TestSignature_AddSpaceFor_NotDefaultMargins(t *testing.T) {
	// Arrange
	pdf := &mocks.Pdf{}
	pdf.On("GetMargins").Return(20.0, 10.0, 10.0, 10.0)
	pdf.On("Line", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	math := &mocks.Math{}
	math.On("GetWidthPerCol", mock.Anything).Return(50.0)

	text := &mocks.Text{}
	text.On("Add", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	signature := internal.NewSignature(pdf, math, text)

	// Act
	signature.AddSpaceFor("label", internal.Cell{2, 5, 5, 0}, props.Text{Size: 10.0})

	// Assert
	pdf.AssertNumberOfCalls(t, "Line", 1)
	pdf.AssertCalled(t, "Line", 26.0, 15.0, 23.0, 15.0)
	text.AssertNumberOfCalls(t, "Add", 1)
	text.AssertCalled(t, "Add", "label", internal.Cell{2.0, 7.0, 5.0, 0.0}, props.Text{Size: 10.0})
}
