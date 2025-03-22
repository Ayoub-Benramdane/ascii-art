package Fonctions

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetAllBanners() []string {
	b := []string{"standard", "shadow", "thinkertoy"}
	banners := make([]string, 0)
	for _, s := range b {
		url := "https://raw.githubusercontent.com/01-edu/public/refs/heads/master/subjects/ascii-art/" + s + ".txt"
		response, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		banners = append(banners, string(body))
	}
	return banners
}

func SplittedBanner(a map[string]string) [][]string {

	banners := GetAllBanners()

	banners[2] = strings.Replace(banners[2], "\r\n", "\n", -1)
	splitt := make([][]string, 0)
	for _, banner := range banners {
		splitted := strings.Split(banner, "\n\n")
		splitted[0] = splitted[0][1:]
		splitt = append(splitt, splitted)
	}
	for j := 0; j < len(splitt); j++ {
		for i := 32; i <= 126; i++ {
			a[splitt[j][(i-32)]] = string(byte(i))

		}
	}
	return splitt
}

func Reverse(filename string) {
	res, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
	a := make(map[string]string)
	text, banners := string(res), SplittedBanner(a)
	text = strings.ReplaceAll(text, "\r\n", "\n")
	if strings.Count(text, "\n") == len(text) {
		if len(text) == 0 {
			return
		}
		text = strings.ReplaceAll(text, "\n", "\\n")
		fmt.Println(text)
		return
	}
	splitted := strings.Split(text, "\n")
	ff, l := make(map[int]bool), false
	s := ""
	for i := 0; i < len(splitted); i++ {
		if splitted[i] == "" && !ff[i] {
			s += "\\n"
			ff[i] = true
			l = true
			continue
		}
		for bb, banner := range banners {
			for k := 0; k < len(banner); k++ {
				splbanner := strings.Split(banner[k], "\n")
				maxlen := 0
				for _, v := range splbanner {
					if len(v) > maxlen {
						maxlen = len(v)
					}
				}
				count := 0
				ww := i
				for j := 0; j < len(splbanner); j++ {
					if ww >= len(splitted) {
						fmt.Println("Error: the file is errored")
						os.Exit(0)
					}
					if len(splbanner[j]) > len(splitted[ww]) {
						break
					}
					if splbanner[j] == splitted[ww][:len(splbanner[j])] {
						ww++
						count++
					} else {
						break
					}
				}
				if count == 8 {
					count = 0
					s += a[banner[k]]
					for j := i; j < i+len(splbanner); j++ {
						splitted[j] = splitted[j][maxlen:]
						ff[i] = true
					}
					if len(splitted[i]) == 0 {
						s += "\\n"
						k = -1
						i += 7
						break
					}
					k = -1
				}
				if bb == len(banners)-1 && k == len(banner)-1 && len(splitted[i]) > 1 {
					fmt.Println("Error: the file is errored\nthe latest result was:", s)
					os.Exit(0)
				}
			}
		}
	}
	if l {
		fmt.Println(s[:len(s)-4])
	} else {
		fmt.Println(s)
	}

}
