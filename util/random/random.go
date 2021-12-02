package random

import "github.com/gofrs/uuid"

func UUID() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}
