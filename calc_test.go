package unitTestingGolang

import "testing"
/**
func TestAdd(test *testing.T){
	result   := Add(2,2);
	expected :=  4;

	if result != expected{
		test.Errorf("Result: %d, Expected: %d", result, expected);
	}
}
func TestSub(test *testing.T){
	result   := Sub(4,2);
	expected :=  2;

	if result != expected{
		test.Errorf("Result: %d, Expected: %d", result, expected);
	}
}

func TestMult(test *testing.T){
	result   := Mult(4,5);
	expected :=  20;

	if result != expected{
		test.Errorf("Result: %d, Expected: %d", result, expected);
	}
}

func TestDiv(test *testing.T){
	result   := Div(4,2);
	expected :=  2;

	if result != expected{
		test.Errorf("Result: %d, Expected: %d", result, expected);
	}
}
**/
func TestCalc(test *testing.T){
checkExpectedResult := func(test *testing.T,result,expected int){
	test.Helper();
	if result != expected{
		test.Errorf("\nResult: %d\nExpected: %d",result,expected);
	}
  }

  test.Run("Add 2 numbers",func(test *testing.T){
	result   := Add(50,100);
	expected := 150;
	checkExpectedResult(test,result,expected);
  })

  test.Run("Sub 2 numbers",func(test *testing.T){
	result   := Sub(43,1);
	expected := 42;
	checkExpectedResult(test,result,expected);
  })
  test.Run("Mult 2 numbers",func(test *testing.T){
	result   := Mult(5,5);
	expected := 25;
	checkExpectedResult(test,result,expected);
  })
  test.Run("Div 2 numbers", func(test *testing.T){
	result   :=  Div(40,2);
	expected :=  20;
	checkExpectedResult(test, result,expected);
  })
}
