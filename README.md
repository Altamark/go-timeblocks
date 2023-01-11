# go-timeblocks
Working with blocks of time and calendars

# Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/prodsaga/go-timeblocks/timeblocks"
)

var locationNewYork, _ = time.LoadLocation("America/New_York")

func main() {
	block1 := timeblocks.New(
		time.Date(2023, 1, 1, 1, 0, 0, 0, time.Local),
		time.Date(2023, 1, 3, 1, 0, 0, 0, time.Local),
		locationNewYork,
	)
	block2 := timeblocks.New(
		time.Date(2023, 1, 2, 1, 0, 0, 0, time.Local),
		time.Date(2023, 1, 4, 1, 0, 0, 0, time.Local),
		locationNewYork,
	)
	fmt.Printf("Block 1: %s\nBlock 2: %s\n", block1, block2)
	fmt.Printf("Block 1 overlaps Block 2: %v\n", block1.IsOverlap(block2))

	block3 := timeblocks.New(
		time.Date(2023, 1, 2, 1, 0, 0, 0, time.Local),
		time.Date(2023, 1, 3, 1, 0, 0, 0, time.Local),
		locationNewYork,
	)

	block4 := timeblocks.New(
		time.Date(2023, 1, 3, 1, 0, 0, 0, time.Local),
		time.Date(2023, 1, 4, 1, 0, 0, 0, time.Local),
		locationNewYork,
	)
	fmt.Printf("Block 3: %s\nBlock 4: %s\n", block3, block4)
	fmt.Printf("Block 3 overlaps Block 4: %v\n", block3.IsOverlap(block4))

	fmt.Printf("Block 1 duration: %s\n", block1.Duration())
}
```

Program output:
```
Block 1: { Start: 2023-01-01T01:00:00-05:00, End: 2023-01-03T01:00:00-05:00 }
Block 2: { Start: 2023-01-02T01:00:00-05:00, End: 2023-01-04T01:00:00-05:00 }
Block 1 overlaps Block 2: true
Block 3: { Start: 2023-01-02T01:00:00-05:00, End: 2023-01-03T01:00:00-05:00 }
Block 4: { Start: 2023-01-03T01:00:00-05:00, End: 2023-01-04T01:00:00-05:00 }
Block 3 overlaps Block 4: false
Block 1 duration: 48h0m0s
```
