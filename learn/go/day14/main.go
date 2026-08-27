package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Session struct {
	Subject  string   `json:"subject"`
	Minutes  int      `json:"minutes"`
	Tags     []string `json:"tags,omitempty"`
	Note     *string  `json:"note,omitempty"`
	Complete bool     `json:"complete"`
}

func main() {
	note := "good progress"
	s := Session{
		Subject:  "Go",
		Minutes:  60,
		Tags:     []string{"fundamentals"},
		Note:     &note,
		Complete: true,
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println(string(data))

	input := []byte(`{"subject":"Math","minutes":45,"complete":false}`)
	var decoded Session
	if err := json.Unmarshal(input, &decoded); err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Printf("decoded: %+v\n", decoded)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetIndent("", "  ")
	_ = enc.Encode(decoded)
	fmt.Println(buf.String())
}
