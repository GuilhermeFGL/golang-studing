package forms

import (
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rec := Rectangle{Height: 5, Width: 10}
		expectedArea := float64(50)
		receivedArea := rec.Area()
		if receivedArea != expectedArea {
			t.Errorf("Rectangle area is incorrect: expected: %.2f, got: %.2f", expectedArea, receivedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{Radius: 5}
		expectedArea := float64(78.53981633974483)
		receivedArea := circle.Area()
		if receivedArea != expectedArea {
			t.Errorf("Cicle area is incorrect: expected: %.2f, got: %.2f", expectedArea, receivedArea)
		}
	})
}
