/*
#######
##                  _    __
##       __ ____ __(_)__/ /
##      / // / // / / _  /
##      \_,_/\_,_/_/\_,_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package uuid

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"sync"
	"time"
)

const (
	_uuidSize = 16
)

var (
	_once sync.Once
)

func generate(reader io.Reader) string {
	buf := make([]byte, _uuidSize)

	if _, err := io.ReadFull(reader, buf); err != nil {
		_once.Do(func() {
			mathrand.Seed(time.Now().UnixNano())
		})

		mathrand.Read(buf)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

// New generates random strings in UUID format.
func New() string {
	return generate(rand.Reader)
}

/*
######################################################################################################## @(°_°)@ #######
*/
