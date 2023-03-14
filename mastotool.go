package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	// c := &Config{
	// 	Hostname: "https://stranger.social",
	// 	AppWebsiteURI: "https://mysite.com",
	// }
	// err := c.CreateApplication()
	// if err != nil {
	// 	fmt.Print(err)
	// }

	a := app.New()
	w := a.NewWindow("TODO List")
	// w.SetContent(makeUI())
	w.ShowAndRun()

}
