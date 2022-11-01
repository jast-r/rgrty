package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/main.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/head.block.tmpl",
		"./ui/html/footer.block.tmpl",
		"./ui/html/menu.block.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"current_time": time.Now().UTC().Format("02.01.2006 15:04:05.99"),
	}
	if err = ts.Execute(w, data); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func lab1(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/lab1.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if err = ts.Execute(w, nil); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func multiTableHTML(rows int, cols int, color string) string {
	mulTableTemplate := "<table id='mult_table'>"

	for row := 0; row <= rows; row++ {
		mulTableTemplate += "<tr>"
		for col := 0; col <= cols; col++ {
			elem := fmt.Sprintf("<td>%d</td>", row*col)
			if row == 0 {
				elem = fmt.Sprintf("<td style='background-color: %s';>%d</td>", color, col)
			}
			if col == 0 {
				elem = fmt.Sprintf("<td style='background-color: %s';>%d</td>", color, row)
			}
			mulTableTemplate += elem
		}
		mulTableTemplate += "</tr>"
	}

	mulTableTemplate = mulTableTemplate + "</table>"
	return mulTableTemplate
}

func numberSystemConverter(source_num string, source_pow, res_pow int) string {
	num, err := strconv.ParseInt(source_num, source_pow, 64)
	if err != nil {
		return ""
	}
	return strconv.FormatInt(num, res_pow)
}

func symbolCounter(str string) int {
	counter := 0
	for _, s := range str {
		if s == 'а' || s == 'Н' || s == 'к' {
			counter++
		}
	}

	return counter
}

func count(array []interface{}) int {
	counter := 0
	for range array {
		counter++
	}
	return counter
}

func lab2(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/lab2.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	first_idz_str := "съешь ещё этих мягких фраНцузских булок да выпей чаю"
	second_idz_arr := []interface{}{
		1, 2, "asd", 0.12355, "丠", "🎃", "🦉", []int{1, 2, 3, 4},
	}

	source_num := "6743"
	source_pow := 10
	res_pow := 2
	result := numberSystemConverter(source_num, source_pow, res_pow)

	rows := 10
	cols := 10
	color := "yellowgreen"
	mt := template.HTML(multiTableHTML(rows, cols, color))

	if err = ts.Execute(w, map[string]interface{}{
		"rows":           rows,
		"cols":           cols,
		"color":          color,
		"mult_table":     mt,
		"source_num":     source_num,
		"source_pow":     source_pow,
		"res_pow":        res_pow,
		"result":         result,
		"first_idz_str":  first_idz_str,
		"first_idz_res":  symbolCounter(first_idz_str),
		"second_idz_arr": second_idz_arr,
		"second_idz_res": count(second_idz_arr),
	}); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/lab1", lab1)
	mux.HandleFunc("/lab2", lab2)

	fileServer := http.FileServer(neuteredFileSystem{fs: http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Сервер запущен по адресу: http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}
	s, _ := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}
