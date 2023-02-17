package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func main() {
	jsonStr := `{"id":"cmpl-6kZ1z4bnN8BHc6cJvTTRmF6CX3xYE","object":"text_completion","created":1676555767,"model":"davinci-codex","choices":[{"text":"\n#what is python\n#what is python\n#what is python\n#what is python\n#what is python\n\n#what is python\n#what is python\n#what is python\n#what is python\n#what is","index":0,"logprobs":null,"finish_reason":"length"}],"usage":{"prompt_tokens":3,"completion_tokens":50,"total_tokens":53}}`

	var resp Response
	if err := json.Unmarshal([]byte(jsonStr), &resp); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(resp.Choices[0].Text)
}
