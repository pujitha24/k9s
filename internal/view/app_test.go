// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of K9s

package view_test

import (
	"testing"

	"github.com/derailed/k9s/internal/config/mock"
	"github.com/derailed/k9s/internal/view"
	"github.com/derailed/tcell/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppNew(t *testing.T) {
	a := view.NewApp(mock.NewMockConfig(t))
	_ = a.Init("blee", 10)

	assert.Equal(t, 14, a.GetActions().Len())
}

func TestAppKeyboardLFAsEnter(t *testing.T) {
	a := view.NewApp(mock.NewMockConfig(t))
	_ = a.Init("blee", 10)

	uu := map[string]struct {
		in, e tcell.Key
	}{
		"lf":    {tcell.KeyLF, tcell.KeyEnter},
		"enter": {tcell.KeyEnter, tcell.KeyEnter},
		"other": {tcell.KeyDown, tcell.KeyDown},
	}
	for k, u := range uu {
		t.Run(k, func(t *testing.T) {
			evt := a.GetInputCapture()(tcell.NewEventKey(u.in, 0, tcell.ModNone))
			require.NotNil(t, evt)
			assert.Equal(t, u.e, evt.Key())
		})
	}
}
