package pokeapi

import(
	"encoding/json"
	"net/http"
	"io"
	"fmt"
)


func (c *Client) GetLocation(locationName string) (Location, error){
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok{
		fmt.Println("Hello1")
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil{
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return Location{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return Location{}, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil{
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(res, &locationResp)
	if err != nil{
		return Location{}, err
	}

	c.cache.Add(url, res)

	return locationResp, nil

}