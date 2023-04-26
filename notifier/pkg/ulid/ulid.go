package ulid

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/goccy/go-json"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type ID ulid.ULID

func NewID() ID {
	return ID(ulid.Make())
}

func ParseID(v string) (ID, error) {
	if len(v) == 0 || strings.Compare(v, "0") == 0 {
		return ID(ulid.ULID{}), nil
	}

	id, err := ulid.Parse(v)
	return ID(id), err
}

func (id ID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, id.String())), nil
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}

	v, err := ParseID(str)
	if err != nil {
		return err
	}
	*id = v
	return nil
}

func (id ID) String() string {
	return ulid.ULID(id).String()
}

// Scan Implement gorm.Scanner interface
func (id *ID) Scan(value interface{}) error {
	if value == nil {
		*id = ID{}
		return nil
	}

	data, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan ID")
	}

	v, err := ParseID(string(data[:]))
	if err != nil {
		return err
	}
	*id = v
	return nil
}

// Value Implement gorm.Valuer interface
func (id ID) Value() (driver.Value, error) {
	return id.String(), nil
}
