package packets

import (
	fmt "fmt"
	"reflect"
	"strings"

	"github.com/gearintellix/structs"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
)

type Metas map[string]InputData

type AnyMetas map[string]InputAnyData

func (ox *Metas) FromPackets(data map[string]*PInputData) {
	if ox == nil {
		(*ox) = make(map[string]InputData)
	}

	for k, v := range data {
		if v == nil {
			continue
		}

		(*ox)[k] = InputData{
			ID:          v.ID,
			Description: v.Description,
			Value:       v.Value,
		}
	}
}

func (ox Metas) CopyToStruct(obj interface{}) (errx serror.SError) {
	var err error

	if obj == nil {
		return serror.Newc("Object cannot be null", "@")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "@")
	}

	nms := make(map[string]int)
	fields := structs.Fields(obj)
	for k, v := range fields {
		nms[utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())] = k
	}

	setValue := func(idx int, val interface{}) {
		defer func() {
			if err != nil {
				return
			}

			if err := recover(); err != nil {
				errx = serror.Newc(err.(string), "Failed to set struct")
			}
		}()

		origin := fields[idx].Value()
		model := reflect.TypeOf((*ICopyToStruct)(nil)).Elem()
		if reflect.TypeOf(origin).Implements(model) {
			val, err = origin.(ICopyToStruct).Cast(val)
		}

		err = fields[idx].Set(val)
	}

	for k, v := range ox {
		paths := utstring.CleanSpit(k, ".")
		group, name := "", k
		if len(paths) > 1 {
			group, name = paths[0], strings.Join(paths[1:], ".")
		}

		if idx, ok := nms[fmt.Sprintf("%s.%s", group, name)]; ok {
			setValue(idx, v.Value)

		} else if idx, ok := nms[name]; ok {
			if group != "@" {
				if _, ok := ox[fmt.Sprintf("@.%s", name)]; ok {
					continue
				}
			}

			setValue(idx, v.Value)
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, fmt.Sprintf("Failed to set field '%s'", k))
			return errx
		}
	}

	return errx
}

func (ox *AnyMetas) FromPackets(data map[string]*PInputAnyData) {
	if ox == nil {
		(*ox) = make(map[string]InputAnyData)
	}

	for k, v := range data {
		if v == nil {
			continue
		}

		(*ox)[k] = InputAnyData{
			ID:          v.ID,
			Description: v.Description,
			Value:       v.Value.Extract(),
		}
	}
}

func (ox AnyMetas) CopyToStruct(obj interface{}) (errx serror.SError) {
	var err error

	if obj == nil {
		return serror.Newc("Object cannot be null", "@")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "@")
	}

	nms := make(map[string]int)
	fields := structs.Fields(obj)
	for k, v := range fields {
		nms[utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())] = k
	}

	setValue := func(idx int, val interface{}) {
		defer func() {
			if err != nil {
				return
			}

			if err := recover(); err != nil {
				errx = serror.Newc(err.(string), "Something when wrong")
			}
		}()

		origin := fields[idx].Value()
		model := reflect.TypeOf((*ICopyToStruct)(nil)).Elem()
		if reflect.TypeOf(origin).Implements(model) {
			val, err = origin.(ICopyToStruct).Cast(val)
		}

		err = fields[idx].Set(val)
	}

	for k, v := range ox {
		paths := utstring.CleanSpit(k, ".")
		group, name := "", k
		if len(paths) > 1 {
			group, name = paths[0], strings.Join(paths[1:], ".")
		}

		if idx, ok := nms[fmt.Sprintf("%s.%s", group, name)]; ok {
			setValue(idx, v.Value)

		} else if idx, ok := nms[name]; ok {
			if group != "@" {
				if _, ok := ox[fmt.Sprintf("@.%s", name)]; ok {
					continue
				}
			}

			setValue(idx, v.Value)
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, fmt.Sprintf("Failed to set field '%s'", k))
			return errx
		}
	}

	return errx
}
