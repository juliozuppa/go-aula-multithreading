package dto

type BrasilApiInput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type ViaCepInput struct {
	Cep        string `json:"cep"`
	Uf         string `json:"uf"`
	Localidade string `json:"localidade"`
	Bairro     string `json:"bairro"`
	Logradouro string `json:"logradouro"`
}

type SearchCepOutput struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func NewSearchCepOutput(cep, state, city, neighborhood, street string) SearchCepOutput {
	return SearchCepOutput{
		Cep:          cep,
		State:        state,
		City:         city,
		Neighborhood: neighborhood,
		Street:       street,
	}
}
