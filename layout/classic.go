package layout

import (
	"context"

	"github.com/melvinodsa/clock/timer"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
)

type classicLayout struct {
}

const (
	borderColor = linestyle.Light
)

var (
	textColor  = segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGreen))
	colonColor = segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorRed))
)

func NewClassicLayout() *classicLayout {
	return &classicLayout{}
}

func (c *classicLayout) DrawLayout(t *tcell.Terminal, sd *segmentdisplay.SegmentDisplay) (*container.Container, error) {
	return container.New(
		t,
		container.Border(borderColor),
		container.BorderTitle("PRESS Q TO QUIT"),
		container.SplitHorizontal(
			container.Top(),
			container.Bottom(
				container.PlaceWidget(sd),
			),
			container.SplitPercent(10),
		),
	)
}

func (c *classicLayout) UpdateTime(ctx context.Context, sd *segmentdisplay.SegmentDisplay, timer *timer.Timer) {
	for {
		select {
		case <-timer.Tick:
			tm := "PM"
			if timer.IsAM {
				tm = "AM"
			}
			chunks := []*segmentdisplay.TextChunk{
				segmentdisplay.NewChunk(timer.H, textColor),
				segmentdisplay.NewChunk(":", colonColor),
				segmentdisplay.NewChunk(timer.M, textColor),
				segmentdisplay.NewChunk(":", colonColor),
				segmentdisplay.NewChunk(timer.S, textColor),
				segmentdisplay.NewChunk(" "),
				segmentdisplay.NewChunk(tm, textColor),
			}
			if err := sd.Write(chunks); err != nil {
				panic(err)
			}

		case <-ctx.Done():
			timer.Stop()
			return
		}
	}
}
