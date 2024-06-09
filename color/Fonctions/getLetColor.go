package Fonctions

import (
	"strings"
)

// Déterminer les couleurs de chaque élément du string entré

func GetLetColor(color, ltc []string, str string) []string {
	resulat := make([]string, len(str))
	for j := len(ltc) - 1; j >= 0; j-- {
		index := strings.Index(str, ltc[j])
		ix := index
		for ix != -1 {
			for i := index; i < len(ltc[j])+index; i++ {
				resulat[i] = color[j]
			}
			ix = strings.Index(str[index+len(ltc[j]):], ltc[j])

			index += ix + len(ltc[j])
		}
	}
	return resulat
}
