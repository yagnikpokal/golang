This app will automatically and efficiently load the templates with cache data. Look at the rendertemplate and createtemplatecache function inside the render.go

To add a new template
1) Add Service on the handlers.go
2) Add service.page.tmpl in template folder
3) Add http.HandleFunc("/service", handlers.Service) in main.go

To run the app
go run ./cmd/web .

Test the app
http://localhost:8080/service
This is the home page
This is some text