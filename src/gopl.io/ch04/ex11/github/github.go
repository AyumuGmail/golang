package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
const CreateIssueURL = "https://api.github.com/repos/AyumuGmail/golang/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int    `json:"number,omitempty"`
	HTMLURL   string `json:"html_url,omitempty"`
	Title     string
	State     string    `json:"state,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//GitHubのイシュートラッカーへの問い合わせ
func SearchIsues(terms []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	//すべての実行パスでresp.Bodyを閉じる必要がある。この処理を簡単にする
	//'defer'を５章で説明
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query faild: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
