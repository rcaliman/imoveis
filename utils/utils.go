package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func MascaraDocumento(d string) string {
	var comp string
	var doc string
	d, comp, doc = mascaraCPF(d, comp, doc)
	doc = mascaraCNPJ(d, comp, doc)
	return doc
}

func mascaraCNPJ(d string, comp string, doc string) string {
	if len(d) > 11 {
		comp = strings.Repeat("0", 14-len(d))
		d = comp + d
	}
	if len(d) == 14 {
		dSlice := strings.SplitN(d, "", -1)
		doc = fmt.Sprintf("%s.%s.%s/%s-%s",
			strings.Join(dSlice[0:2], ""),
			strings.Join(dSlice[2:5], ""),
			strings.Join(dSlice[5:8], ""),
			strings.Join(dSlice[8:12], ""),
			strings.Join(dSlice[12:14], ""),
		)
	}
	return doc
}

func mascaraCPF(d string, comp string, doc string) (string, string, string) {
	if len(d) < 11 {
		comp = strings.Repeat("0", 11-len(d))
		d = comp + d
	}
	if len(d) == 11 {

		dSlice := strings.SplitN(d, "", -1)
		doc = fmt.Sprintf("%s.%s.%s-%s",
			strings.Join(dSlice[0:3], ""),
			strings.Join(dSlice[3:6], ""),
			strings.Join(dSlice[6:9], ""),
			strings.Join(dSlice[9:11], ""),
		)
	}
	return d, comp, doc
}

func FormataDinheiro(v float64) string {
	s := fmt.Sprintf("%.2f", v)
	return strings.Replace(s, ".", ",", 1)
}

func ParseDate(d string) time.Time {
	if len(d) == 0 {
		d = "1900-01-01"
	}
	data := strings.Split(d, "-")
	dia, _ := strconv.Atoi(data[2])
	mes := func() time.Month {
		d, _ := strconv.Atoi(data[1])
		return time.Month(d)
	}()
	ano, _ := strconv.Atoi(data[0])
	return time.Date(ano, mes, dia, 0, 0, 0, 0, time.Local)
}

func StringToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func FloatToString(f float64) string {
	return fmt.Sprintf("%g", f)
}

func StringToFloat(s string) float64 {
	n, _ := strconv.ParseFloat(s, 64)
	return n
}

func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func ExtraiData(data time.Time) string {
	return strings.Split(data.String(), " ")[0]
}

func FormataData(data time.Time) string {
	d := strings.Split(ExtraiData(data), "-")
	data_formatada := fmt.Sprintf("%s/%s/%s", d[2], d[1], d[0])
	return data_formatada
}

func InputSelectMeses() string {
	mesAtual := int(time.Now().Month())
	meses := []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"}
	selectMeses := "<select name='recibo_mes' class='inputgerarrecibo' id='recibo_mes'>"
	for i, m := range meses {
		if (i + 1) == mesAtual {
			selectMeses += fmt.Sprintf("<option selected>%s</option>", m)
		} else {
			selectMeses += fmt.Sprintf("<option>%s</option>", m)
		}
	}
	selectMeses += "</select>"
	return selectMeses
}

func InputSelectAnos() string {
	anoAtual := time.Now().Year()
	anos := []int{anoAtual - 1, anoAtual, anoAtual + 1}
	selectAnos := "<select name='recibo_ano' class='inputgerarrecibo' id='recibo_ano'>"
	for _, a := range anos {
		if a != anoAtual {
			selectAnos += fmt.Sprintf("<option>%d</option>", a)
		} else {
			selectAnos += fmt.Sprintf("<option selected>%d</option>", a)
		}
	}
	selectAnos += "</select>"
	return selectAnos
}

func DizMesAnterior(mes, ano string) (string, string) {
	switch mes {
	case "janeiro":
		return "fevereiro", ano
	case "fevereiro":
		return "março", ano
	case "março":
		return "abril", ano
	case "abril":
		return "maio", ano
	case "maio":
		return "junho", ano
	case "junho":
		return "julho", ano
	case "julho":
		return "agosto", ano
	case "agosto":
		return "setembro", ano
	case "setembro":
		return "outubro", ano
	case "outubro":
		return "novembro", ano
	case "novembro":
		return "dezembro", ano
	case "dezembro":
		return "janeiro", func() string {
			a, _ := strconv.Atoi(ano)
			a = a + 1
			s := strconv.Itoa(a)
			return s
		}()
	}
	return "mes", "ano"
}

func Arredonda2(numero float64) float64 {
	return math.Round(numero*100) / 100
}
