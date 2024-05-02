package utfunc

import "github.com/jinzhu/copier"

func CopyWhenUpdate(toValue any, fromValue any) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
	})
}
