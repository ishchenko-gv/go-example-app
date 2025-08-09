package userid

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (u ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(u).String())
}

func New() ID {
	return ID(uuid.New())
}
