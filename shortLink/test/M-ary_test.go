package test

import (
	"testing"
	"log"
	"github.com/mr_litt/go-case/shortLink/app"
)

func TestAppInit(t *testing.T) {
	b64 := app.DecToB64(10000)
	log.Println(b64)
	dec := app.B64ToDec(b64)
	log.Println(dec)
}