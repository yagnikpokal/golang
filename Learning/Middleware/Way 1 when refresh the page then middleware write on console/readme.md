This app will wite middleware
When i refresh the app then it will write text on the console.


To add a new template
1) Create the file middleware.go and add a function Writetoconsole
2) Add 	r.Use(Writetoconsole) in the routes.go file

To run the app
go run ./cmd/web .

Test the app
http://localhost:8080/
Refresh the page

Staring application on port :8080
Hit the page


Again if refresh the page then it shows
Staring application on port :8080
Hit the page
Hit the page