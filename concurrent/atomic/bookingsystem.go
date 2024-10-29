package atomic

import (
	"sort"
)

type Flight struct {
	Origin, Dest string
	SeatLeft     int
	Locker       *SpinLock // spinlock
}

func NewFlight(origin, dest string) *Flight {
	return &Flight{
		Origin:   origin,
		Dest:     dest,
		SeatLeft: 200,
		Locker:   NewSpinLock(),
	}
}

func Book(flights []*Flight, seatsToBook int) bool {
	bookable := true
	sort.Slice(flights, func(a, b int) bool {
		flightA := flights[a].Origin + flights[a].Dest
		flightB := flights[b].Origin + flights[b].Dest
		return flightA < flightB
	})

	for _, f := range flights {
		f.Locker.Lock()
	}

	for i := 0; i < len(flights) && bookable; i++ {
		if flights[i].SeatLeft < seatsToBook {
			bookable = false
		}
	}

	for i := 0; i < len(flights) && bookable; i++ {
		flights[i].SeatLeft -= seatsToBook
	}

	for _, f := range flights {
		f.Locker.Unlock()
	}
	return bookable
}
