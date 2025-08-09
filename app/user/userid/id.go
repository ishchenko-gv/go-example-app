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

func (id *ID) UnmarshalJSON(data []byte) error {
	parsed, err := FromString(string(data))
	if err != nil {
		return err
	}

	*id = parsed
	return nil
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
