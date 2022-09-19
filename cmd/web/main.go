package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/codeninja/revision/pkg/config"
	"github.com/codeninja/revision/pkg/handlers"
	"github.com/codeninja/revision/pkg/render"
)

const portNumber = ":8080"

// func Home(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "this is the home page")
// 	RenderTemplate(w, "home.page.tmpl")
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is % f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil
// }

// func About(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "about.page.tmpl")
// 	// sum := addValues(2, 2)
// 	// _, _ = fmt.Fprintf(w, fmt.Sprintf("about us page and 2+2 is %d", sum))
// }

// func addValues(x, y int) int {
// 	// var sum int
// 	//sum = x+y
// 	//return sum, nil

//		return x + y
//	}
var app config.AppConfig
var session *scs.SessionManager

func main() {
	//change this true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	// http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
	//    n, err :=  fmt.Fprintf(w, "Hello ,World")
	//    if err != nil{
	//       fmt.Println(err)
	//    }
	//    fmt.Println(fmt.Sprintf("Number of bytes written: %d",n))
	// })
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Printf("starting application on port %s", portNumber))
	//_ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
