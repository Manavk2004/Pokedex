package pokeapi

import(
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil{
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok{
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil{
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return RespShallowLocations{}, err
	}


	resp, err := c.httpClient.Do(req)
	if err != nil{
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return RespShallowLocations{}, err
	}

	var locations RespShallowLocations
	err = json.Unmarshal(body, &locations)
	if err != nil{
		return RespShallowLocations{}, err
	}


	c.cache.Add(url, body)
	return locations, nil
}