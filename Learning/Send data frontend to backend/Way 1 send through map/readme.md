Send a map data from frontend to backend

Steps to add another map data backend to frontend
1) Go to templates-->>about.page.tmpl Add line    <p>This came from the template: {{index .StringMap "state"}}</p>
2) Go to pkg-->> handlers-->> handlerd.go-->>about function  Add one map elemen as stringMap["state"] = "Gujarat" 



How to run it
go run ./cmd/web .


How to test it  
Go to http://localhost:8080/about This will render page and send data like below.

This is the about page
This is a paragraph of text

This is a paragraph of text

This came from the template: Hello, again

State is: Gujarat
