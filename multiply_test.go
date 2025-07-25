package main
import "testing"
func TestProduct(t *testing.T){
	result:=Product(4,5)
	expected:=20
	if result!=expected{
		t.Errorf("Expected %d but got %d",expected,result)
	}
}