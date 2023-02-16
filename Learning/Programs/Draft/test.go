// fetch multiple
// Append those user id to slice
// print the user details
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	userid := []string{"123", "456", "789"}

	userIDString := stringstrings.Join(userIDs, ",")
	url := fmt.Sprintf("https://endpoint.com/user/ids=%s", userIDString)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	data := []string(string(body))
	fmt.Println(data)

}
