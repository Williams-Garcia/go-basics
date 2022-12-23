package animals

import "testing"

func TestDog(t *testing.T) {
	//Arrange
	dogs := 10.0
	expectedFood:= 100.0

	//Act
	result := Dog(dogs)

	//Assert
	if result != expectedFood {
		t.Errorf("Prueba fallo")
	}
}

func TestCat(t *testing.T) {
	//Arrange
	cats := 10.0
	expectedFood:= 50.0

	//Act
	result := Cat(cats)

	//Assert
	if result != expectedFood {
		t.Errorf("Prueba fallo")
	}
}

func TestMouse(t *testing.T) {
	//Arrange
	mouses := 5.0
	expectedFood:= 1.25

	//Act
	result := Mouse(mouses)

	//Assert
	if result != expectedFood {
		t.Errorf("Prueba fallo")
	}
}

func TestSpider(t *testing.T) {
	//Arrange
	spider := 8.0
	expectedFood:= 1.2

	//Act
	result := Spider(spider)

	//Assert
	if result != expectedFood {
		t.Errorf("Prueba fallo")
	}
}