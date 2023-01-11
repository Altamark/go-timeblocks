package timeblocks

import (
	"fmt"
	"time"
)

type Block struct {
	Start time.Time
	End   time.Time
}

func New(start, end time.Time, location *time.Location) *Block {
	return &Block{
		Start: start,
		End:   end,
	}
}

func (b *Block) String() string {
	return fmt.Sprintf("{ Start: %s, End: %s }", b.Start.Format(time.RFC3339), b.End.Format(time.RFC3339))
}

func (b *Block) MustNormalize(location *time.Location) *Block {
	if location == nil {
		panic(&ErrorNoLocation{})
	}

	b.Start = b.Start.In(location)
	b.End = b.End.In(location)

	// switch if end before start
	if b.End.Before(b.Start) {
		dur := b.Start.Sub(b.End)
		b.Start = b.End
		b.End = b.Start.Add(dur)
	}

	return b
}

func (b *Block) Duration() time.Duration {
	return b.End.Sub(b.Start)
}

func (b *Block) IsOverlap(other *Block) bool {
	thisBlock, otherBlock := *b, *other
	copyThis, copyOther := thisBlock, otherBlock
	copyThis.MustNormalize(copyThis.Start.Location())
	copyOther.MustNormalize(copyOther.Start.Location())

	if !copyThis.End.After(copyOther.Start) || !copyThis.Start.Before(copyOther.End) {
		return false
	}

	return true
}

func (b *Block) AddDuration(d time.Duration) *Block {
	return &Block{
		Start: b.Start.Add(d),
		End:   b.End.Add(d),
	}
}

func (b *Block) AddDate(years, months, days int) *Block {
	return &Block{
		Start: b.Start.AddDate(years, months, days),
		End:   b.End.AddDate(years, months, days),
	}
}
