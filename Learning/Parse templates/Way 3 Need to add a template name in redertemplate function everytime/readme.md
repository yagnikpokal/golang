Way 3 of render the templates
This will read template data everytime from the disk which is not a desied way.
Due to that it is not a desired way beacuse it take more stack/resources.
inside the func RenderTemplate function each and everytime we have to add a new render template

TO add a new template follow below steps.
1) Add a Service function in hendlers.go
2) Add a line in main.go http.HandleFunc("/service", handlers.Service)
3) Add service.page.tmpl in templates folder


Run the app
go run ./cmd/web .


Test the app
http://localhost:8080/service

This will give
This is the service page
This is some service
