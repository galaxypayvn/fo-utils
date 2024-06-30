package utfunc

import (
	"time"

	"code.finan.cc/finan-one-be/fo-utils/utils/customtype"
	"github.com/jinzhu/copier"
)

var customTimeToTimeConverter = copier.TypeConverter{
	SrcType: customtype.Time{},
	DstType: time.Time{},
	Fn: func(src interface{}) (interface{}, error) {
		customTm, ok := src.(customtype.Time)
		if !ok {
			return time.Time{}, nil
		}

		return customTm.Time.UTC(), nil
	},
}

func CopyWhenUpdate(toValue any, fromValue any) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		Converters:  []copier.TypeConverter{customTimeToTimeConverter},
	})
}

func CopyWhenInsert(toValue any, fromValue any) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		Converters: []copier.TypeConverter{customTimeToTimeConverter},
	})
}
