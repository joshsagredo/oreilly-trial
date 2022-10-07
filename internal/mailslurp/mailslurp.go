package mailslurp

import (
	"context"

	"github.com/antihax/optional"
	mailslurp "github.com/mailslurp/mailslurp-client-go"
)

var ApiKey = "0ab31749c0af6c1f5519a4fddb0c0758dfdfdaee39453857fb0add61d2c0f3a3"

func newAPIClient() *mailslurp.APIClient {
	return mailslurp.NewAPIClient(mailslurp.NewConfiguration())
}

func GenerateTempMail() (string, error) {
	client := newAPIClient()
	ctx := context.WithValue(
		context.Background(),
		mailslurp.ContextAPIKey,
		mailslurp.APIKey{Key: ApiKey},
	)

	inbox, _, err := client.InboxControllerApi.CreateInbox(ctx, &mailslurp.CreateInboxOpts{ExpiresIn: optional.NewInt64(120000)})

	return inbox.EmailAddress, err
}
