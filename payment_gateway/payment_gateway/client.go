package payment_gateway

const (
	ErrClientDoesNotExist = "client does not exist"
)

type Client struct {
	ClientId       int
	ClientName     string
	PaymentGateway []*PaymentGateway
}
