package models

type Healthz struct {
	Name    	string `json:"name"`
	Version 	string `json:"version"`
	Environment string `json:"environment"`
}
