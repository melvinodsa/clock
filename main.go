// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Binary segmentdisplaydemo shows the functionality of a segment display.
package main

import (
	"context"
	"time"

	"github.com/melvinodsa/clock/layout"
	"github.com/melvinodsa/clock/timer"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

func renderClock(ctx context.Context, l layout.Layout, t *tcell.Terminal, s *segmentdisplay.SegmentDisplay) (*container.Container, error) {
	c, err := l.DrawLayout(t, s)
	if err != nil {
		return nil, err
	}
	tm := timer.New()
	go l.UpdateTime(ctx, s, tm)
	tm.Start()
	return c, nil
}

func main() {
	t, err := tcell.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())
	clockSD, err := segmentdisplay.New()
	if err != nil {
		panic(err)
	}
	c, err := renderClock(ctx, layout.NewClassicLayout(), t, clockSD)
	if err != nil {
		panic(err)
	}

	quitter := func(k *terminalapi.Keyboard) {
		if k.Key == 'q' || k.Key == 'Q' {
			cancel()
		}
	}

	if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(1*time.Second)); err != nil {
		panic(err)
	}
}
