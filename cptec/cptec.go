package cptec

type CptecV1 struct {
	bases map[string]string	
}


func V1() *CptecV1 {
	return &CptecV1{
		bases: getBasesForCptecRequests(),
	}
}


func getBasesForCptecRequests() map[string]string {

	bases := make(map[string]string)

	bases["locales"] = "https://brasilapi.com.br/api/cptec/v1/cidade"
	bases["current_weather"] = "https://brasilapi.com.br/api/cptec/v1/clima"
	bases["ocean"]     	     = "https://brasilapi.com.br/api/cptec/v1/ondas"

	return bases
}

func (c *CptecV1) Bases() map[string]string {
	return c.bases
}