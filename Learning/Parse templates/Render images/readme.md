#Render the images on the go wenapplications


Steps to render the images
1) Add below items into routes.go
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

2) Add below line into the home.page.tmpl
   <img src="/static/images/yagnik.jpg" alt="My photo">
3) Create the static folder-->> images folder -->> Add the yagnik.jpg image on it


To run the app 
go run ./cmd/web .

Test the app 
http://localhost:8080 
This will load the yagnik.jpg image on the webpage