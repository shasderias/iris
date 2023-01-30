package evt_test

import (
	"testing"

	"github.com/shasderias/iris/context"
	"github.com/shasderias/iris/evt"
)

func TestSanity(t *testing.T) {
	ctx := context.NewBase()

	be := evt.Basic(ctx, evt.OType(1), evt.OValue(3), evt.OFloatValue(1.0))

	if be.Type != 1 || be.Value != 3 || be.FloatValue != 1.0 {
		t.Fail()
	}

	evt.Basic(ctx, evt.LightOff)
}
