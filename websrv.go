package main

import (
    "fmt"
    "net/http"
    "os"
    "io/ioutil"
)


type Page struct{
    Title string
    Body  []byte
    Loaded bool
}

func (p *Page) save() error{
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) load(title string) error{
//    filename := title + ".txt"
//    body, error := ioutil.ReadFile(filename)
    p.Title = title
    p.Body = []byte("abcdef")
    return nil
}

func (p *Page) loadBody() {
    filename := p.Title + ".txt"
    body, error := ioutil.ReadFile(filename)
    p.Loaded = error == nil
    if p.Loaded{
        p.Body = []byte(body)
    }
    return
}



func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL)
    
    for _, s := range os.Args{
        fmt.Fprintln(w, s)
    }
}

const
    lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request){
    title := r.URL.Path[lenPath:]
    p := &Page{Title: title}
    p.loadBody()
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>%v", p.Title, p.Body, p.Loaded)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
