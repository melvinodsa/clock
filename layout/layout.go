package layout

import (
	"context"

	"github.com/melvinodsa/clock/timer"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

type Layout interface {
	DrawLayout(t *tcell.Terminal, sd *segmentdisplay.SegmentDisplay) (*container.Container, error)
	UpdateTime(ctx context.Context, sd *segmentdisplay.SegmentDisplay, timer *timer.Timer)
}
