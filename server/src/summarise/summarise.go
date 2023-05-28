package summarise

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"summar/server/utils"
)

type SummariseBookmarkRequest struct {
	Url          string `json:"url"`
	NumSentences int    `json:"num_sentences"`
	IsDetailed   bool   `json:"is_detailed"`
}

type SummariseBookmarkResponse struct {
	Summary         []string `json:"summary"`
	ArticleText     string   `json:"article_text"`
	ArticleTitle    string   `json:"article_title"`
	ArticleAuthors  []string `json:"article_authors"`
	ArticleImage    string   `json:"article_image"`
	ArticlePubDate  string   `json:"article_pub_date"`
	ArticleUrl      string   `json:"article_url"`
	ArticleHtml     string   `json:"article_html"`
	ArticleAbstract string   `json:"article_abstract"`
}

func SummariseBookmark(url string) (*SummariseBookmarkResponse, error) {
	apiKey := os.Getenv("RAPID_API_KEY")

	endpoint := "https://tldrthis.p.rapidapi.com/v1/model/extractive/summarize-url/"

	payload := utils.JSONMarshal(&SummariseBookmarkRequest{
		Url:          url,
		NumSentences: 10,
		IsDetailed:   false,
	})

	req, err := http.NewRequest("POST", endpoint, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-rapidapi-key", apiKey)
	req.Header.Set("x-rapidapi-host", "tldrthis.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println(string(body))

	if res.StatusCode == 200 {
		return utils.JSONUnmarshal[SummariseBookmarkResponse](body)
	} else {
		return nil, fmt.Errorf("Error while summarising bookmark: %+v", string(body))
	}
}
