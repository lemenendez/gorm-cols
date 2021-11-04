package gormdatatypes_test

import (
	"github.com/google/uuid"
	gtypes "github.com/lemenendez/gormdatatypes"
	"testing"
)

func TestUID(t *testing.T) {
	uuid := uuid.MustParse("b5fa1ee1-033f-4fab-9473-5fcaab553c0b")

	uid := gtypes.UID(uuid)
	if uid.String() != "b5fa1ee1-033f-4fab-9473-5fcaab553c0b" {
		t.Fail()
	}
	t.Log("tested")
}
