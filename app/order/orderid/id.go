package orderid

import (
	"encoding/json"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (o ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(uuid.UUID(o).String())
}

func New() ID {
	return ID(uuid.New())
}

func Zero() ID {
	return ID(uuid.UUID{})
}

func FromString(id string) (ID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return Zero(), err
	}
	return ID(parsed), nil
}
