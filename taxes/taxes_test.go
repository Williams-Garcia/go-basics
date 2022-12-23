package taxes

import "testing"

func TestCalculateTaxesLess50(t *testing.T)  {
	//Arrange
	salary := 45000.0
	expectedResultTaxe := 0.0

	//Act
	result := calculateTaxes(salary)

	//Assert
	if result != expectedResultTaxe{
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateTaxesMore50(t *testing.T)  {
	//Arrange
	salary := 70000.0
	expectedResultTaxe := 11900.0

	//Act
	result := calculateTaxes(salary)

	//Assert
	if result != expectedResultTaxe {
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateTaxesMore150(t *testing.T)  {
	//Arrange
	salary := 270000.0
	expectedResultTaxe := 45900.0

	//Act
	result := calculateTaxes(salary)

	//Assert
	if result != expectedResultTaxe {
		t.Errorf("Prueba fallo")
	}
}