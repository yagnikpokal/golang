package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
)

func main() {
	client := graphql.NewClient("https://api.octopart.com/v3/graphql", &http.Client{})

	query := `
		query {
			items(queries: [{mpn: "LM555CN"}]) {
				hits {
					mpn
					brand {
						name
					}
					datasheets {
						url
					}
				}
			}
		}
	`
	variables := map[string]interface{}{}

	req := graphql.NewRequest(query, nil)
	req.SetVariables(variables)

	ctx := context.Background()
	var respData struct {
		Items []struct {
			Hits []struct {
				MPN        string
				Brand      struct{ Name string }
				Datasheets []struct{ URL string }
			}
		}
	}
	if err := client.Run(ctx, req, &respData); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("MPN:", respData.Items[0].Hits[0].MPN)
	fmt.Println("Brand:", respData.Items[0].Hits[0].Brand.Name)
	fmt.Println("Datasheets:")
	for _, ds := range respData.Items[0].Hits[0].Datasheets {
		fmt.Println("  ", ds.URL)
	}
}
