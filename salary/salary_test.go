package salary

import "testing"

func TestCalculateSalaryTypeC(t *testing.T)  {
	//Arrange
	hoursInMinutes := 150
	expectedResultSalaryTypeC := 2500.0

	//Act
	result := IdentifySalary("c", hoursInMinutes)

	//Assert
	if result != expectedResultSalaryTypeC {
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateSalaryTypeB(t *testing.T)  {
	//Arrange
	hoursInMinutes := 150
	expectedResultSalaryTypeC := 51750.0

	//Act
	result := IdentifySalary("b", hoursInMinutes)

	//Assert
	if result != expectedResultSalaryTypeC {
		t.Errorf("Prueba fallo")
	}
}

func TestCalculateSalaryTypeA(t *testing.T)  {
	//Arrange
	hoursInMinutes := 150
	expectedResultSalaryTypeC := 247500.0

	//Act
	result := IdentifySalary("a", hoursInMinutes)

	//Assert
	if result != expectedResultSalaryTypeC {
		t.Errorf("Prueba fallo")
	}
}

