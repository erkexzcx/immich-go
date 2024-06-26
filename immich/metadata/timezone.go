package metadata

import (
	"time"

	"github.com/simulot/immich-go/helpers/tzone"
)

var local *time.Location

func init() {
	var err error
	local, err = tzone.Local()
	if err != nil {
		panic(err)
	}
}
