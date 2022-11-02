package main

import (
	web "Mado/asciiWeb/Web"
	"log"
	"net/http"
)

func main() {
	// var ShowText string

	mux := http.NewServeMux()
	mux.HandleFunc("/", web.Home)
	// mux.HandleFunc("/ascii-art-web", web.Ascii)

	// mux.HandleFunc("/snippet", web.ShowSnippet)
	// mux.HandleFunc("/snippet/create", web.CreateSnippet)

	log.Println("Запуск сервера на http://127.0.0.1:4050")
	err := http.ListenAndServe(":4050", mux)
	log.Fatal(err)
}

/*img1 := [][]int{{1, 1, 0}, {0, 1, 0}, {0, 1, 0}}
	img2 := [][]int{{0, 0, 0}, {0, 1, 1}, {0, 0, 1}}
	fmt.Println(largestOverlap(img1, img2))

func largestOverlap(img1 [][]int, img2 [][]int) int {
	largestOverlap := 0

	for row := -1*len(img1) + 1; row < len(img1); row++ {
		for col := -1*len(img1) + 1; col < len(img1); col++ {
			largestOverlap = max(largestOverlap, overlapOnes(img1, img2, row, col))
		}
	}
	return largestOverlap
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func overlapOnes(img1 [][]int, img2 [][]int, rowOff, colOff int) int {
	count := 0

	for row := 0; row < len(img1); row++ {
		for col := 0; col < len(img1[0]); col++ {
			if (row+rowOff < 0 || row+rowOff >= len(img1)) || (col+colOff < 0 || col+colOff >= len(img1)) {
				continue
			}
			if img2[row][col]+img1[row+rowOff][col+colOff] == 2 {
				count++
			}
		}
	}
	return count
}
*/
