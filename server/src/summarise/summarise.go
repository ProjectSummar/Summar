package summarise

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

func SummariseBookmark(bookmark *types.Bookmark) (*SmApiResponse, error) {
	apiKey := os.Getenv("API_KEY")

	apiUrl := fmt.Sprintf(
		"https://api.smmry.com/&SM_API_KEY=%s&SM_URL=%s",
		apiKey,
		bookmark.Url,
	)

	req, err := http.NewRequest("POST", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("expect", "100-continue")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	log.Println("summarise api response body:\n", string(body))

	if res.StatusCode == 200 {
		return utils.JSONUnmarshal[SmApiResponse](body)
	} else {
		return nil, fmt.Errorf("Error while summarising bookmark: %+v", string(body))
	}
}
