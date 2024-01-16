package types

import (
	"strconv"
)

type Context struct {
	Location string
	Level    uint32
}

var (
	PRIME_CTX  = Context{"prime", 0}
	REGION_CTX = Context{"region", 1}
	ZONE_CTX   = Context{"zone", 2}
)

//*Definitions and implementations for a SliceID type*//

type SliceID struct {
	Context Context
	Region  uint32
	Zone    uint32
}

func (sliceID SliceID) String() string {
	return strconv.Itoa(int(sliceID.Region)) + "." + strconv.Itoa(int(sliceID.Zone))
}
