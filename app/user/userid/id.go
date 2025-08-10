package userid

import (
	"github.com/google/uuid"
	"github.com/ishchenko-gv/go-example-app/app/common/id"
)

type ID struct {
	id.ID
}

func New() ID {
	return ID{
		ID: id.ID(uuid.New()),
	}
}

func Zero() ID {
	return ID{
		ID: id.Zero(),
	}
}

func FromString(value string) (ID, error) {
	v, err := id.FromString(value)
	if err != nil {
		return Zero(), nil
	}

	return ID{
		ID: v,
	}, nil
}
