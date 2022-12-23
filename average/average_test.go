package average

import "testing"

func TestCalculateAverage(t *testing.T)  {
	//Arrange
	expectedResultAverage := 2.5

	//Act
	result := CalculateAverage(1,2,3,4)

	//Assert
	if result != expectedResultAverage {
		t.Errorf("Prueba fallo")
	}
}