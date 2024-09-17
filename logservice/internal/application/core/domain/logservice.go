package domain

type Logservice struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	App  string `json:"app"`
	Data string `json:"data"`
}

func NewLogservice(app string, name string, data string) Logservice {
	return Logservice{
		Name: name,
		App:  app,
		Data: data,
	}
}
