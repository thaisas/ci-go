import "testing"

func TestSoma(t *testing.T){
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Error: Expected result as %d, it got %d", 30, total)
	}
}