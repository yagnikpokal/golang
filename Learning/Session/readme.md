This app will add support of session
The app will rembember the the sign in and sign out of the user

The package use is
go get github.com/alexedwards/scs/v2



Steps to add a new session
1) Add below in main.go
	app.Inproduction = false

	//Session data
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction
    app.Session = session

2) Add 	below line in routes.go
	mux.Use(SessionLoad)
3) Add below function in the middleware
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}



Below will be use to test wether the session is working or not. Here we will be adding the IP adress on variable from the home page. And then we go to another page and see that they have the IP or not?

1) Add below into About function in handlers.go
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP


2) Add below into home function in handers.go
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

3)Add below to about.page.tmpl

                <p>
                {{if ne (index .StringMap "remote_ip") ""}}
                    Your remote IP address is {{index .StringMap "remote_ip"}}
                {{else}}
                    I Dont know your IP. Visit <a href = "/">Home page</a> So i can send.
                {{end}}
                </p>






To run the app
go run ./cmd/web .

Test the app
http://localhost:8080/about
This is the about page
This is a paragraph of text

This is a paragraph of text

This came from the template: Hello, again

I Dont know your IP. Visit Home page So i can send.




Then go to http://localhost:8080/
This will send IP adress


http://localhost:8080/about
This is the about page
This is a paragraph of text

This is a paragraph of text

This came from the template: Hello, again

Your remote IP address is [::1]:58993

