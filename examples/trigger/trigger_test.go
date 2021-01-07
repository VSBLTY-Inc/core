package sample

import (
	"testing"

	"github.com/VSBLTY-Inc/core/support"
	"github.com/VSBLTY-Inc/core/trigger"
	"github.com/stretchr/testify/assert"
)

func TestTrigger_Register(t *testing.T) {

	ref := support.GetRef(&Trigger{})
	f := trigger.GetFactory(ref)
	assert.NotNil(t, f)
}
