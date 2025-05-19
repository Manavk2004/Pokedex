package pokeapi

import(
	"net/http"
	"io"
	"encoding/json"
)

func (c *Client) GetPokemon(poke string) (Pokemon, error){
	url := baseURL + "/pokemon/" + poke

	if val, ok := c.cache.Get(url); ok{
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil{
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil{
		return Pokemon{}, err
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil{
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(res, &pokemonResp)
	if err != nil{
		return Pokemon{}, err
	}

	c.cache.Add(url, res)
	return pokemonResp, nil
}