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
	"fmt"
)

type (
	badReader struct{}
)

func (br *badReader) Read([]byte) (int, error) {
	return 0,
		fmt.Errorf("unexpected error")
}

func GenerateWithFallback() string {
	return generate(&badReader{})
}

/*
######################################################################################################## @(°_°)@ #######
*/
