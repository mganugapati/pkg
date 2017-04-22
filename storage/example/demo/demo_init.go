// NOTE: AUTO-GENERATED by midc, DON'T edit!!

package demo

import (
	"fmt"

	"github.com/mkideal/pkg/storage"
	"github.com/mkideal/pkg/typeconv"
	"gopkg.in/redis.v5"
)

var (
	_ = fmt.Printf
	_ = storage.Unused
	_ = typeconv.Unused
	_ = redis.Nil
)

func Init(eng storage.Engine) {
	eng.AddIndex(UserAgeIndexVar)
	eng.AddIndex(UserAddrIndexVar)
	eng.AddIndex(ProductPriceIndexVar)
}