package utils

import "testing"

func TestDizMesPosterior(t *testing.T) {
	mes, ano := DizMesPosterior("fevereiro", "2020")
	if mes != "março" && ano != "2020" {
		t.Errorf("o mes calculado deveria ser março e foi %s e o ano deveria ser 2020 e foi %s", mes, ano)
	}
	mes, ano = DizMesPosterior("janeiro", "2020")
	if mes != "dezembro" && ano != "2020" {
		t.Errorf("o mes calculado deveria ser janeiro e foi %s e o ano deveria ser 2021 e foi %s", mes, ano)
	}
}
