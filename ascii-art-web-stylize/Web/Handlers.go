package web

import (
	"Mado/asciiWeb/Web/FileReadWork"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// fmt.Println(1)
		ts, err := template.ParseFiles("./Ui/Html/Home.page.tmpl")
		if err != nil {
			fmt.Println("dsfsd")
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		//	fmt.Println(2)
		err = r.ParseForm()
		fmt.Println(err)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		// fmt.Println(3)
		text, ok := r.Form["textForAscii"]
		fmt.Println(4)
		if !ok {
			//	fmt.Println("Here")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// fmt.Println(5)

		var masText []rune
		for _, v := range text[0] {
			masText = append(masText, rune(v))
		}
		// fmt.Println(masText)
		// fmt.Println(text)
		fileName, ok := r.Form["fileName"]
		if !ok {
			// fmt.Println("Here2")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// fmt.Println(fileName)

		checkFileName := fileName[0]
		fileName[0] = "./Web/FileReadWork" + "/" + fileName[0] + ".txt"
		// fmt.Println(fileName)

		massivString, errF := FileReadWork.FileReadWork(fileName[0], checkFileName) // reading file
		if errF != nil {
			fmt.Println(errF)
			os.Exit(0)
		}

		TextForShow := PrintArt(massivString, text[0])
		// fmt.Println(text)
		// sfmt.Fprintf(w, "%s ", text)

		errWhenExecute := ts.Execute(w, TextForShow)
		if errWhenExecute != nil {
			fmt.Println(errWhenExecute)
			// fmt.Println("Error when Execute Some text")
			os.Exit(0)
		}
	case "GET":
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if r.Method != http.MethodGet {
			errorHandler(w, r, http.StatusNotFound)
			http.Error(w, "Method not Allowed, 405", http.StatusMethodNotAllowed)
			return
		}
		ts, err := template.ParseFiles("./Ui/Html/Home.page.tmpl")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}
	default:
		fmt.Println(1)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

// Используем функцию template.ParseFiles() для чтения файла шаблона.
// Если возникла ошибка, мы запишем детальное сообщение ошибки и
// используя функцию http.Error() мы отправим пользователю
// ответ: 500 Internal Server Error (Внутренняя ошибка на сервере)

// Затем мы используем метод Execute() для записи содержимого
// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
// возможность отправки динамических данных в шаблон.
// func Ascii(w http.ResponseWriter, r *http.Request) {

func PrintArt(massivString []string, args string) string {
	text := ""

	// args = strings.ReplaceAll(args, "\n", "\\n")

	splitArgs := strings.Split(args, "\n") // sozder mojno vmesto args

	for _, arg := range splitArgs {

		if arg == "" {
			fmt.Println()
			text += "\n"
			continue
		}

		for i := 0; i < 8; i++ {

			for index, v := range arg {
				if v == '\n' || v == '\r' {
					continue
				}
				if !(v >= ' ' && v <= '~') {
					fmt.Printf("DONT CORRECT ARGS IN %d", index)
					break
				}

				text += strings.Split(massivString[v-' '], "\n")[i]
			}

			text += "\n"
		}
	}
	return text
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
