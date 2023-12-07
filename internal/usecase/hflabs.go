package usecase

import (
	"GoHFLabsParcer/internal/entity"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParceTable(url string) ([]string, []string, error) {
	data, err := requestToEndPoint(url)
	if err != nil {
		return nil, nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data.Body.Storage.Value))
	if err != nil {
		return nil, nil, err
	}

	var status []string
	var desc []string
	total := 0
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				total++
				if total%2 != 0 {
					status = append(status, tablecell.Text())
				} else {
					desc = append(desc, tablecell.Text())
				}
			})
		})
	})

	return status, desc, nil
}

func requestToEndPoint(url string) (entity.Response, error) {
	var data entity.Response

	resp, err := http.Get(url)
	if err != nil {
		return data, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return data, err
	}

	return data, nil
}
