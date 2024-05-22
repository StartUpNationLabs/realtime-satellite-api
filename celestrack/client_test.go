package celestrack

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_scrap(t *testing.T) {
	m, err := scrap()
	if err != nil {
		return
	}

	log.Info(m)
}
