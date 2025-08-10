package id

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/google/uuid"
)

type ID uuid.UUID

func (id ID) String() string {
	return uuid.UUID(id).String()
}

func (id *ID) Scan(value any) error {
	if value == nil || value == "" {
		*id = Zero()
		return nil
	}

	v, ok := value.([]uint8)
	if !ok {
		*id = Zero()
		return nil
	}

	var err error
	*id, err = FromString(string(v))
	if err != nil {
		*id = Zero()
		return nil
	}

	return nil
}

func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
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

func FromString(id string) (ID, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return Zero(), err
	}
	return ID(parsed), nil
}

func Zero() ID {
	return ID(uuid.UUID{})
}
