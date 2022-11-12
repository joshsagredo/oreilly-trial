package mail

const (
	emailRequestQuery = "{\"query\":\"mutation introduceSession($input: IntroduceSessionInput) { " +
		"introduceSession(input: $input) { " +
		"id " +
		"addresses { " +
		"address " +
		"} " +
		"expiresAt " +
		"} " +
		"}\",\"variables\":{\"input\":{\"withAddress\":true,\"domainId\":\"%s\"}}}"
	domainRequestQuery = "{\"query\":\"query { domains { id name introducedAt availableVia }}\",\"variables\":{}}"
	hostHeader         = "dropmail.p.rapidapi.com"
	contentType        = "application/json"
)

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
