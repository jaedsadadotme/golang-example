# Golang-Template-Html

### ตัวอย่างการเขียน
---
#### 1. Add Packages
```go
...
import (
	"log"
	"html/template"
	"net/http"
	"path"
)
...
```

#### 2. เขียน Function Render 

```go
func HelloTemplate(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
```

#### 3. เขียน Main

```go
func main() {
	http.HandleFunc("/", HelloTemplate) // เรียกใช้ function Render

	log.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

```

#### 4. Run

```sh
$ go run main.go
```

#### 5. Open Browser

```sh
http://localhost:9000
```