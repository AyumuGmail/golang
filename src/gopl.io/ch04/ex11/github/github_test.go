package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestCreateIssue(t *testing.T) {
	issue := &Issue{
		Title:     "テストCreateIsuer",
		CreatedAt: time.Now(),
		Body:      "中身は日本語でテスト",
	}
	jsonBytes, _ := json.Marshal(issue)

	fmt.Printf("%s\n", string(jsonBytes))

	resp, err := http.Post(CreateIssueURL, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalf("err!!\n")
	}
	fmt.Printf("%v\n", resp)
}
