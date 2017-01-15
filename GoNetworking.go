package main

import (
	"fmt"
	"encoding/json"
	"html/template"
	"html"
	"net/http"
	"net/url"
	"io/ioutil"
	"bytes"
)

type Page struct  {
	Title string
}

type Person struct {
	Name string `json:"nombre"`
	Age int	    `json:"edad"`
	Car string  `json:"coche"`
	City string `json:"ciudad"`
}

type People [] Person

type Toue struct {
	Myfield People
}

//type mensaje struct {
//	msg string
//}
//
//func(m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request){
//	fmt.Fprintf(w, m.msg);
//}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main()  {
	//mux := http.NewServeMux()

	//fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	//http.Handle("/static/", fileServer)

	//staticHandler := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/style/", staticHandler))

	http.Handle("/css/", http.FileServer(http.Dir("./xfklm")))
	http.Handle("/js/", http.FileServer(http.Dir("./xfklm")))
	http.Handle("/img/", http.FileServer(http.Dir("./xfklm")))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	//mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.SetCookie(res, &http.Cookie{
			Name:  "my-main-cookie",
			Value: "Cesar",
			HttpOnly: true,
		})
		templates.ExecuteTemplate(res, "main", &Page{Title: "Home"})
	})
	http.HandleFunc("/car", func(res http.ResponseWriter, req *http.Request) {
		//http://localhost:8080/car?q="dog"
		val := req.URL.Query().Get("q")
		fmt.Println(val)
	})
	http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			val := req.FormValue("username")
			fmt.Println(html.EscapeString(val))
			cookie, _ := req.Cookie("my-cookie")
			fmt.Println(cookie)
		}
		http.SetCookie(res, &http.Cookie{
			Name:  "my-cookie",
			Value: "some value",
			HttpOnly: true,
		})
		templates.ExecuteTemplate(res, "login", &Page{Title: "Login"})
	})
	http.HandleFunc("/sendjson", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST"{

			//xx := `{"Myfield":[{"nombre": "Nuria", "edad": 40, "coche": "Audi", "ciudad": "Santander"}, {"nombre": "Natalia", "edad": 29, "coche": "BMW", "ciudad": "Santander"}]}`
			//res.Write([]byte(xx))
			xx := Toue{
				Myfield: People{
					Person{Name: "Nuria", Age: 40, Car: "Audi", City: "Santander"},
					Person{Name: "Natalia", Age: 29, Car: "BMW", City: "Santander"},
				},
			}
			json.NewEncoder(res).Encode(xx)
		}else {
			//girl := Person{Name: "Isabel", Age: 30, Car: "Ferrari", City: "Santander"}
			manyGirls := People{
				Person{Name: "Nuria", Age: 40, Car: "Audi", City: "Santander"},
				Person{Name: "Natalia", Age: 29, Car: "BMW", City: "Santander"},
			}
			json.NewEncoder(res).Encode(manyGirls)
		}
	})
	http.HandleFunc("/toflask", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST"{
			json.NewEncoder(res).Encode(&Person{Name: "Natalia", Age: 29, Car: "BMW", City: "Santander"})
		}
	})
	//REDIRECT!!
	http.HandleFunc("/mychip", func(res http.ResponseWriter, req *http.Request) {

		http.Redirect(res, req, "http://127.0.0.1:5000", http.StatusFound)
	})
	//WORKS!!
	http.HandleFunc("/senda", func(res http.ResponseWriter, req *http.Request) {
		urldata := url.Values{}
		urldata.Add("car", `BMW`)
		hc := &http.Client{}
		resp, _ := hc.PostForm("http://127.0.0.1:5000/golang", urldata)
		defer resp.Body.Close()
		var p Page
		xx, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(xx, &p)
		fmt.Println(p.Title)
	})
	//WORKS!!
	http.HandleFunc("/sendb", func(res http.ResponseWriter, req *http.Request) {
		urlData := url.Values{}
		urlData.Set("car", "BMW")
		client := &http.Client{}
		resp, _ := client.PostForm("http://127.0.0.1:5000/golang", urlData)
		defer resp.Body.Close()
		var p Page
		err := json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {fmt.Println(err)}
		fmt.Println(p.Title)
	})
	//WORKS!!
	http.HandleFunc("/sendc", func(res http.ResponseWriter, req *http.Request) {
		resp, _ := http.PostForm("http://127.0.0.1:5000/golang", url.Values{"car": {"BMW"}, "city": {"Germany"}})
		defer resp.Body.Close()
		var p Page
		json.NewDecoder(resp.Body).Decode(&p)
		fmt.Println(p.Title)
	})
	//WORKS!!
	http.HandleFunc("/myget", func(res http.ResponseWriter, req *http.Request) {
		resp, _ := http.Get("http://127.0.0.1:5000")
		defer resp.Body.Close()
		xx, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(xx))
	})
	//WORKS!!
	http.HandleFunc("/mypost", func(res http.ResponseWriter, req *http.Request) {
		values := map[string]string{"name": "Cesar", "car": "BMW"}
		jsonValue, _ := json.Marshal(values)
		resp, _ := http.Post("http://127.0.0.1:5000/ue", "application/json", bytes.NewBuffer(jsonValue))
		defer resp.Body.Close()
		xx, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(xx))
	})
	//WORKS!!
	http.HandleFunc("/myrequest", func(res http.ResponseWriter, req *http.Request) {
		url := "http://127.0.0.1:5000/ue"
		var jsonStr = []byte(`{"name": "Cesar", "car": "BMW"}`)
		reque, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		//reque.Header.Set("X-Custom-Header", "myvalue")
		reque.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(reque)
		defer resp.Body.Close()
		if err != nil {
			panic(err)
		}
		xx, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(xx))

	})

	//myMsg := mensaje{
	//	msg: "Hola, soy Cesar",
	//}

	//http.Handle("/mensaje", myMsg);

	//myServer := &http.Server{
	//	Addr: ":8080",
	//	Handler: nil,
	//	ReadTimeout: 10 * time.Second,
	//	WriteTimeout: 10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//myServer.ListenAndServe()
	http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", http.RedirectHandler("http://127.0.0.1:5000", 301))
}
