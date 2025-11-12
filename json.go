package main
import(
	"net/http"
	"fmt"
	"io"
	"encoding/json"
)

type locationAreaList struct{
	Count	int
	Next	*string
	Previous *string
	Results	[]struct{
		Name 	string
		URL		string
	}
}

func fetchLocationAreaPage(url string) (locationAreaList, error){
	res, err := http.Get(url)
	if err != nil{
		return locationAreaList{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return locationAreaList{}, fmt.Errorf("Bad status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err!= nil{
		return locationAreaList{}, err
	}

	var page locationAreaList

	if err := json.Unmarshal(body, &page); err !=nil{
		return locationAreaList{}, err
	}

	return page, nil
} 