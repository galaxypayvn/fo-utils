package packets

import (
	"encoding/json"
	"reflect"
	"strings"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/utils/utfunc"
	"bitbucket.org/finesys/finesys-utility/utils/utinterface"
)

func ToObj(value interface{}) *PObject {
	out := &PObject{}
	err := out.Pack(value)
	if err != nil {
		err.Panic()
	}

	return out
}

func ToObje(value interface{}) *PObject {
	out := &PObject{}
	_ = out.Pack(value)
	return out
}

func (ox *PObject) Pack(value interface{}) (errx serror.SError) {
	var err error

	ox.Type = "nil"
	ox.Body = []byte("null")

	if !utinterface.IsNil(value) {
		ox.Type = reflect.TypeOf(value).String()

		switch value.(type) {
		case int, int8, int16, int32, int64:
			ox.Type = "int64"

		case []int, []int8, []int16, []int32, []int64:
			ox.Type = "[]int64"

		case float32, float64:
			ox.Type = "float64"

		case []float32, []float64:
			ox.Type = "[]float64"
		}

		ox.Body, err = json.Marshal(value)
		if err != nil {
			return serror.NewFromErrorc(err, "while packaging object")
		}
	}

	return nil
}

func (ox *PObject) Unpack(out interface{}) (errx serror.SError) {
	err := json.Unmarshal(ox.Body, out)
	if err != nil {
		return serror.NewFromErrorc(err, "while unpacking object")
	}

	return nil
}

func (ox *PObject) UnpackV2(out interface{}, parentOut interface{}, fieldName string) (errx serror.SError) {
	errx = utfunc.Try(func() serror.SError {
		var tmp interface{}

		valo := reflect.ValueOf(out)
		if valo.Kind() != reflect.Ptr {
			return serror.New("Out must be pointer")
		}

		err := json.Unmarshal(ox.Body, &tmp)
		if err != nil {
			return serror.NewFromErrorc(err, "while unpacking object")
		}

		switch strings.ToLower(ox.Type) {
		case "int", "int8", "int16", "int32", "int64":
			tmp = utinterface.ToInt(tmp, 0)

		case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
			tmpx := []int64{}
			if utinterface.IsNil(tmp) {
				tmp = tmpx
			}

			tmpv := reflect.ValueOf(tmp)
			if tmpv.Kind() == reflect.Slice {
				for i := 0; i < tmpv.Len(); i++ {
					cur := tmpv.Index(i)
					tmpx = append(tmpx, utinterface.ToInt(cur.Interface(), 0))
				}
			}

			tmp = tmpx

		case "float32", "float64":
			tmp = utinterface.ToFloat(tmp, 0)

		case "[]float32", "[]float64":
			tmpx := []float64{}
			if utinterface.IsNil(tmp) {
				tmp = tmpx
			}

			tmpv := reflect.ValueOf(tmp)
			if tmpv.Kind() == reflect.Slice {
				for i := 0; i < tmpv.Len(); i++ {
					cur := tmpv.Index(i)
					tmpx = append(tmpx, utinterface.ToFloat(cur.Interface(), 0))
				}
			}

			tmp = tmpx
		}

		if tmp != nil {
			// opt0: cast by struct caster
			if parentOut != nil {
				if scst, ok := parentOut.(interface {
					Cast(string, interface{}) (interface{}, serror.SError)
				}); ok {
					tmp, errx = scst.Cast(fieldName, tmp)
					if errx != nil {
						errx.AddCommentf("while Cast (%s)", fieldName)
						return errx
					}
				}
			}

			// opt1: cast by field caster
			outx := valo.Elem().Interface()
			if cts, ok := outx.(ICopyToStruct); ok {
				var err error
				tmp, err = cts.Cast(tmp)
				if err != nil {
					errx = serror.NewFromErrorc(err, "Failed to casting value")
					return errx
				}
			}

			valo.Elem().Set(reflect.ValueOf(tmp))
		}

		return nil
	})
	if errx != nil {
		errx.AddComments("while try unpack pobject")
		return errx
	}

	return errx
}

func (ox *PObject) Extract() interface{} {
	var out interface{}
	_ = ox.Unpack(&out)
	return out
}

func (ox *PObject) ExtractV2() interface{} {
	var out interface{}
	_ = ox.UnpackV2(&out, nil, "@")
	return out
}

func (ox *PObject) MarshalJSON() ([]byte, error) {
	return ox.Body, nil
}

func (ox *PObject) UnmarshalJSON(data []byte) error {
	var out interface{}
	if err := json.Unmarshal(data, &out); err != nil {
		return err
	}

	ox.Type = reflect.TypeOf(out).String()
	ox.Body = data
	return nil
}
