package mail

type DomainResponse struct {
	DomainData `json:"data"`
}

type DomainData struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	Name         string   `json:"name"`
	ID           string   `json:"id"`
	AvailableVia []string `json:"availableVia"`
}

type EmailResponse struct {
	EmailData `json:"data"`
}

type EmailData struct {
	IntroduceSession `json:"introduceSession"`
}

type IntroduceSession struct {
	ID        string    `json:"id"`
	ExpiresAt string    `json:"expiresAt"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Address string `json:"address"`
}
