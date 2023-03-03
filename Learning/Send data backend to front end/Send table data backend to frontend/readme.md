This app will add table on frontend
Send map of data from backend to front end

Steps to add a table
1) Add below into templatedata.go
	StringMap map[string]string

2) Add below into home function in handers.go
	stringMap := make(map[string]string)

	stringMap["Company_1"] = "Alfreds Futterkiste"
	stringMap["Contact_1"] = "Maria Anders"
	stringMap["Country_1"] = "Germany"

	stringMap["Company_2"] = "Centro comercial Moctezuma"
	stringMap["Contact_2"] = "Francisco Chang"
	stringMap["Country_2"] = "Mexico"

	stringMap["Company_3"] = "Ernst Handel"
	stringMap["Contact_3"] = "Maria Anders"
	stringMap["Country_3"] = "Austria"
		render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

3)Add below to homr.page.tmpl

								<h2>HTML Table</h2>
								<table>
								<tr>
									<th>Company</th>
									<th>Contact</th>
									<th>Country</th>
								</tr>
								{{if ne (index .StringMap "Company_1") ""}}
								<tr>
									<td>{{index .StringMap "Company_1"}}</td>
									<td>{{index .StringMap "Contact_1"}}</td>
									<td>{{index .StringMap "Country_1"}}</td>
								</tr>
								
								<tr>
									<td>{{index .StringMap "Company_2"}}</td>
									<td>{{index .StringMap "Contact_2"}}</td>
									<td>{{index .StringMap "Country_2"}}</td>
								</tr>
								<tr>
									<td>{{index .StringMap "Company_3"}}</td>
									<td>{{index .StringMap "Contact_3"}}</td>
									<td>{{index .StringMap "Country_3"}}</td>
									<hr>
								</tr>
								{{end}}

								<!--   If 4th table is not present then it will be hide by addinng {{if ne (index .StringMap "Company_4") ""}} {{end}}
								{{if ne (index .StringMap "Company_4") ""}}

									<tr>
									<td>{{index .StringMap "Company_4"}}</td>
									<td>{{index .StringMap "Contact_4"}}</td>
									<td>{{index .StringMap "Country_4"}}</td>
								</tr>
								{{end}}
								-->
									<hr>
									
								</table>








To run the app
go run ./cmd/web .

Test the app
http://localhost:8080

This is the home page
This is some text

HTML Table
Company	Contact	Country
Alfreds Futterkiste	Maria Anders	Germany
Centro comercial Moctezuma	Francisco Chang	Mexico
Ernst Handel	Maria Anders	Austria