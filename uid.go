package gormdatatypes

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UID uuid.UUID

//String -> String Representation of Binary16
func (u UID) String() string {
	return uuid.UUID(u).String()
}

//GormDataType -> sets type to binary(16)
func (u UID) GormDataType() string {
	return "binary(16)"
}

func (my UID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(my)
	str := "\"" + s.String() + "\""
	return []byte(str), nil
}

func (my *UID) UnmarshalJSON(by []byte) error {
	s, err := uuid.ParseBytes(by)
	*my = UID(s)
	return err
}

// Scan --> tells GORM how to receive from the database
func (my *UID) Scan(value interface{}) error {

	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*my = UID(parseByte)
	return err
}

// Value -> tells GORM how to save into the database
func (my UID) Value() (driver.Value, error) {
	return uuid.UUID(my).MarshalBinary()
}
