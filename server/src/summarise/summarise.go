package summarise

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"summar/server/types"
	"summar/server/utils"
)

type SmApiResponse struct {
	Message        string `json:"sm_api_message"`
	CharacterCount string `json:"sm_api_character_count"`
	Title          string `json:"sm_api_title"`
	Content        string `json:"sm_api_content"`
	KeywordArray   string `json:"sm_api_keyword_array"`
	Error          string `json:"sm_api_error"`
}

func SummariseBookmark(bookmark types.Bookmark) (SmApiResponse, error) {
	apiBaseURL := os.Getenv("API_BASE_URL")
	apiKey := os.Getenv("API_KEY")

	apiUrl, _ := url.Parse(apiBaseURL)

	apiQuery := apiUrl.Query()
	apiQuery.Set("SM_API_KEY", apiKey)
	apiQuery.Set("SM_URL", bookmark.Url)

	apiUrl.RawQuery = apiQuery.Encode()

	req, err := http.NewRequest("POST", apiUrl.String(), nil)
	if err != nil {
		return SmApiResponse{}, err
	}

	req.Header.Set("expect", "100-continue")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return SmApiResponse{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return SmApiResponse{}, err
	}

	log.Println("summarise api response body:\n", string(body))

	if res.StatusCode == 200 {
		return utils.JSONUnmarshal[SmApiResponse](body)
	} else {
		return SmApiResponse{}, fmt.Errorf("Error while summarising bookmark: %+v", string(body))
	}
}
