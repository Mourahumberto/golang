package main

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json: "idade"`
}

func main() {
	cachorroJson := `{"nome":"Rex","raca":"Dalmata, "idade":3}`
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}
