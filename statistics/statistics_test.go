package statistics

import "testing"

func TestCalculateMinimum(t *testing.T) {
	//Arrange
	nums := []int{1, 403249, 42, 10, 82}
	expectedResultMinimum := 1

	//Act
	result := findMinimum(nums)

	//Assert
	if result != expectedResultMinimum {
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateMaximum(t *testing.T) {
	//Arrange
	nums := []int{1, 403249, 42, 10, 82}
	expectedResultMaximum := 403249

	//Act
	result := findMaximum(nums)

	//Assert
	if result != expectedResultMaximum {
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateAverage(t *testing.T) {
	//Arrange
	nums := []int{1, 403249, 42, 10, 82}
	expectedResultAvg := 80677

	//Act
	result := averageFunc(nums)

	//Assert
	if result != expectedResultAvg {
		t.Errorf("Prueba fallo")
	}
}
