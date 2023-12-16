package vozatexto_ia_test

import (
	"rdcarranza/vozatexto_ia"
	"testing"
)

func Test_Vozatexto_ia(t *testing.T) {
	got := vozatexto_ia.VTxt_Estado()
	if got == false {
		t.Logf("Se espera un verdadero, estado FALSO")
		t.Fail()
	}
}
