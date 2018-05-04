package main

import (
	"html/template"
	"log"
	"os"
)

func handler(page string) {
	_, err := os.Stat("templates/" + page + ".tmpl")
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseFiles(
		"templates/"+page+".tmpl",
		"templates/base.tmpl",
		"templates/_header.tmpl",
		"templates/_footer.tmpl",
		"templates/_copy.tmpl",
	))

	file, err := os.Create("docs/" + page + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = t.ExecuteTemplate(file, "base", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler_station(sub string, num string) {
	type Station struct {
		Subject string
		Number  string
	}

	t := template.Must(template.ParseFiles(
		"templates/station.tmpl",
	))

	file, err := os.Create("docs/station" + num + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Station{
		Subject: sub,
		Number:  num,
	}

	err = t.ExecuteTemplate(file, "station", data)
	if err != nil {
		log.Fatal(err)
	}
}

func handler_profile(name string) {
	type Profile struct {
		Name string
	}

	t := template.Must(template.ParseFiles(
		"templates/profile_each.tmpl",
		"templates/base.tmpl",
		"templates/_header.tmpl",
		"templates/_footer.tmpl",
		"templates/_copy.tmpl",
	))

	file, err := os.Create("docs/profile_" + name + ".html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := Profile{
		Name: name,
	}

	err = t.ExecuteTemplate(file, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	handler("index")
	handler("news")
	handler("event")
	handler("event_before")
	handler("event_clubnostyle")
	handler("profile")
	handler_profile("akira")
	handler("radio")
	handler_station("lowply - studio mix - dec.04", "01")
	handler_station("akira - studio mix - aug.04 for glo.jp", "02")
	handler_station("nami - PtMIX ginger - mar.05", "03")
	handler_station("velocity - commix promotion studio mix - mar.05", "04")
	handler_station("dj ray - klockworks - mar.05", "05")
	handler_station("nami - PtMIX caramel - aug.05 [new]", "06")
	handler("chart")
	handler("chart_before")
	handler("link")
}
