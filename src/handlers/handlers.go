package handlers

import (
    "fmt"
    "net/http"
		"regexp"
    "encoding/json"
)

// validates path
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")


// func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
//     p, err := loadPage(title)
//     if err != nil {
//         http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//         return
//     }
//     renderTemplate(w, "view", p)
// }
//
// func editHandler(w http.ResponseWriter, r *http.Request, title string) {
//     p, err := loadPage(title)
//     if err != nil {
//         p = &Page{Title: title}
//     }
//     renderTemplate(w, "edit", p)
// }
//
// func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
//     body := r.FormValue("body")
//     p := &Page{Title: title, Body: []byte(body)}
//     err := p.save()
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
//     http.Redirect(w, r, "/view/"+title, http.StatusFound)
// }

// Function literal
// validation and error checking on every request
// returns a closure
// fn will be one of save/edit/view handlers
// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//     return func(w http.ResponseWriter, r *http.Request) {
//         m := validPath.FindStringSubmatch(r.URL.Path)
//         if m == nil {
//             http.NotFound(w, r)
//             return
//         }
//         fn(w, r, m[2])
//     }
// }

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func fetchJson(w http.ResponseWriter, r *http.Request) {

    p1 := &Page{}
    var vr = p1.load()

    fmt.Fprintf(w, "Hi there, I love %s!", vr)
}

func fetchJsonArray(w http.ResponseWriter, r *http.Request) {

    p1 := &Page{}
    var vr = p1.loadComplexPage()

    fmt.Fprintf(w, "Hi there, I love %s!", vr)
}

func returnJson(w http.ResponseWriter, r *http.Request) {

  p1 := &Page{Title: "WHAAAT", Body: []byte("This is a sample Page.")}

  js, err := json.Marshal(p1)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func returnJsonArray(w http.ResponseWriter, r *http.Request) {

  p1 := Page{Title: "So what", Body: []byte("This is a sample Page.")}
  p2 := Page{Title: "When", Body: []byte("This is a sample Page.")}

  dataArray := [2]Page{p1, p2} // build array of Page's

  js, err := json.Marshal(dataArray)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func SetupHandlers() {
		// Creates an initial page
		p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()

		// setup route handler
    http.HandleFunc("/", handler)
    http.HandleFunc("/fetch-json/", fetchJson)
    http.HandleFunc("/fetch-jsonarray/", fetchJsonArray)
    http.HandleFunc("/return-json/", returnJson)
    http.HandleFunc("/return-jsonarray/", returnJsonArray)

    // http.HandleFunc("/view/", makeHandler(viewHandler))
    // http.HandleFunc("/edit/", makeHandler(editHandler))
    // http.HandleFunc("/save/", makeHandler(saveHandler))
    
    http.ListenAndServe(":8080", nil)
}
