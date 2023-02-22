package certmagic

import (
	"context"
	"log"
)

func magicEvent(ctx context.Context, evt string, data map[string]any) error {

	log.Printf("Certmagic Event: %s with data: %v\n", evt, data)
	return nil

}
