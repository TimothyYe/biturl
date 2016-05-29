package utils

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestShortenURL(t *testing.T) {
	Convey("Get a long URL to test", t, func() {
		url := "http://www.xiaozhou.net/got-hhkb-pro-type-s-2016-02-23.html"

		Convey("Pass this long URL to shorten func", func() {
			shortens := ShortenURL(url)

			Convey("It should return 4 shorten URLs", func() {
				So(len(shortens), ShouldEqual, 4)

				Convey("Each URL should be unique", func() {
					fmt.Println("Shortens[0] is:" + shortens[0])
					for i := 1; i < 4; i++ {
						fmt.Printf("Shortens[%d] is:%s \r\n", i, shortens[i])
						So(shortens[0], ShouldNotEqual, shortens[i])
					}
				})
			})
		})
	})
}
