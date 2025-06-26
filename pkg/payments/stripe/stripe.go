package stripe

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v82"
)

type StripeClient struct {
}

func NewClient() {
	stripeSK := os.Getenv("STRIPE_SECRET_KEY")
	if stripeSK =="" {
		log.Fatal("STRIPE_SECRET_KEY environment variable is not set")
	}
	stripe.Key = stripeSK
}