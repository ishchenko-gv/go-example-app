package userid

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

func New() ID {
	return ID(uuid.New())
}
