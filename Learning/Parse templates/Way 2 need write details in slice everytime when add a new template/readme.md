Way 2 add a mew tempate/page in go

To create a new template follow below steps let us say if i want add a service page then
1) go to handlers.go --> add Service function
2) go to main.go -->> add http.HandleFunc("/service", handlers.Service)
3) go templates folder -->> Add service.page.tmpl


To run the application go root level of the folder
run 
go run ./cmd/web .

Go to web browser
http://localhost:8080/service

It will display 
This is the service page
This is some text of service page

Function RenderTemplate will create a cache and for the first time it will load the data from disk then after everytime it will use the data from cache.
Can check these data from below

go run ./cmd/web .
Staring application on port :8080
2023/02/16 12:45:15 creating template and adding to cache  #refresh the page
2023/02/16 12:45:15 using cached template
