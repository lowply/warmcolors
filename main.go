package main

import (
	"html/template"
	"log"
	"os"
)

func handler(page string) {
	template_path := "templates/" + page + ".tmpl"
	_, err := os.Stat(template_path)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseFiles(
		template_path,
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

func handler_station(num string, sub string) {
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
	template_path := "templates/profile_" + name + ".tmpl"
	_, err := os.Stat(template_path)
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.ParseFiles(
		template_path,
		"templates/profile_base.tmpl",
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

	err = t.ExecuteTemplate(file, "base", nil)
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
	handler_profile("kga")
	handler_profile("lowply")
	handler_profile("nami")
	handler_profile("acchi")
	handler_profile("sav")
	handler("profile_guest")
	handler_profile("sower")
	handler_profile("kiyomi")
	handler_profile("makoto")
	handler_profile("deeizm")
	handler_profile("velocity")
	handler_profile("inza")
	handler("radio")
	handler_station("01", "lowply - studio mix - dec.04")
	handler_station("02", "akira - studio mix - aug.04 for glo.jp")
	handler_station("03", "nami - PtMIX ginger - mar.05")
	handler_station("04", "velocity - commix promotion studio mix - mar.05")
	handler_station("05", "dj ray - klockworks - mar.05")
	handler_station("06", "nami - PtMIX caramel - aug.05 [new]")
	handler("chart")
	handler("chart_before")
	handler("link")
	handler("topics050325")
}
