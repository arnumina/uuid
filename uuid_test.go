/*
#######
##                  _    __
##       __ ____ __(_)__/ /
##      / // / // / / _  /
##      \_,_/\_,_/_/\_,_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package uuid_test

import (
	"regexp"
	"testing"

	"github.com/arnumina/uuid"
)

func check(t *testing.T, m string, f func() string) {
	ms := make(map[string]bool)
	re := regexp.MustCompile(`^[a-f\d]{8}(-[a-f\d]{4}){3}-[a-f\d]{12}$`)

	for i := 0; i < 1000; i++ {
		uuid := f()

		if !re.MatchString(uuid) {
			t.Errorf("%s() => this value is not a UUID: %s", m, uuid)
			return
		}

		if ms[uuid] {
			t.Errorf("%s() => this uuid is not unique: %s", m, uuid)
			return
		}

		ms[uuid] = true
	}
}

func TestNew(t *testing.T) {
	check(t, "New", uuid.New)
	check(t, "GenerateWithFallback", uuid.GenerateWithFallback)
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = uuid.New()
	}
}

func BenchmarkGenerateWithFallback(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = uuid.GenerateWithFallback()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
