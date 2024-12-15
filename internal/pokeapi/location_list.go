package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client)ListLocations(pageURL *string)(RespShallowLocations, error){
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok{
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil{
			return RespShallowLocations{}, err
		}
		fmt.Println("Cache Hit!")
		return locationResp, nil
	}
	fmt.Println("Cache Miss!")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer resp.Body.Close()
	
	dat,err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	locationResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}
	
	c.cache.Add(url, dat)
	
	return locationResp, nil
}

func (c *Client)GetLocationArea(locationAreaName string) (LocationArea, error){
	url := baseURL + "/location-area/" +  locationAreaName
	

	if val, ok := c.cache.Get(url); ok{
		locationAreaResp := LocationArea{}
		err := json.Unmarshal(val, &locationAreaResp)
		if err != nil{
			return LocationArea{}, err
		}
		fmt.Println("Cache Hit!")
		return locationAreaResp, nil
	}
	fmt.Println("Cache Miss!")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()
	
	dat,err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	locationAreaResp := LocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return LocationArea{}, err
	}
	
	c.cache.Add(url, dat)
	
	return locationAreaResp, nil
}

func (c *Client)GetPokemon(pokemon string) (Pokemon, error){
	url := baseURL + "/pokemon/" +  pokemon
	

	if val, ok := c.cache.Get(url); ok{
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil{
			return Pokemon{}, err
		}
		fmt.Println("Cache Hit!")
		return pokemonResp, nil
	}
	fmt.Println("Cache Miss!")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()
	
	dat,err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}
	
	c.cache.Add(url, dat)
	
	return pokemonResp, nil
}