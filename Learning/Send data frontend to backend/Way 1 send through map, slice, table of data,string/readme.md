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




--------------------*************************------------------------------
Send a slice data from back end to front end on service page

1) Go to template-->> service.page.html   Add <p>{{index .StringSlice 0}}</p>
2) Go to pkg-->>models-->>templatedata.go Add StringSlice []string 
3) Go to pkg-->>handlers-->>handlers.go-->>Service handler   Add   stringslice := []string{"This came from slice"}
    Add inside the &models.TemplateData StringSlice: stringslice.


Run the application
go run ./cmd/web .

Test the application
http://localhost:8080/service
Service page
This came from backend This came from slice


---------------------********************---------------------------------
Send table of data from backend to frontend

1) Go to template-->> service.page.html   Add <p>{{index .StringSlice 0}}</p>


                        <p>Table of data</p>

                                    <table border="2">
                                        <thead>
                                            <tr>
                                                <th>Name</th>
                                                <th>Occupation</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                      
                                            <tr>
                                                <td>{{index .StringSlice 1}}</td>
                                                <td>{{index .StringSlice 2}}</td>
                                         
                                            </tr>
                                                                                 
                                            <tr>
                                                <td>{{index .StringSlice 3}}</td>
                                                <td>{{index .StringSlice 4}}</td>
                                            </tr>
                                        </tbody>
                                    </table>

2) Go to pkg-->>models-->>templatedata.go Add StringSlice []string 
3) Go to pkg-->>handlers-->>handlers.go-->>Service handler   Add   stringslice := []string{"This came from slice", "Yagnik", "Golang developer", "David", "JS developer"}
    Add inside the &models.TemplateData StringSlice: stringslice.



Run the application
go run ./cmd/web .


Test the application
http://localhost:8080/service
Service page
This came from backend This came from slice

Table of data

Name	Occupation
Yagnik	Golang developer
David	JS developer
