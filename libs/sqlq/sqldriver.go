package sqlq

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gearintellix/u2"
	"github.com/lib/pq"

	log "github.com/sirupsen/logrus"

	"bitbucket.org/finesys/finesys-utility/utils/utinterface"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
	"bitbucket.org/finesys/finesys-utility/utils/uttime"
)

// QColumn type
type QColumn []string

// QRaw type
type QRaw string

// QArray type
type QArray []interface{}

// QCast type
type QCast []interface{}

// SQLDriver type
type SQLDriver int

const (
	// DriverMySQL driver
	DriverMySQL SQLDriver = 1 + iota

	// DriverPostgreSQL driver
	DriverPostgreSQL

	// DriverBigQuery driver
	DriverBigQuery
)

type ConditionObj struct {
	Condition1 string
	Operator   string
	Condition2 string
}

func (QArray) FromStrings(v []string) QArray {
	var res QArray
	if len(v) > 0 {
		res = QArray{}
		for _, v := range v {
			res = append(res, v)
		}
	}
	return res
}

func (QArray) FromInt64s(v []int64) QArray {
	var res QArray
	if len(v) > 0 {
		res = QArray{}
		for _, v := range v {
			res = append(res, v)
		}
	}
	return res
}

func (QCast) Casting(val interface{}, cst string) QCast {
	return QCast{val, cst}
}

func (ox QCast) GetValue() interface{} {
	if len(ox) >= 0 {
		return ox[0]
	}

	return nil
}

func (ox QCast) GetType() string {
	if len(ox) >= 1 {
		return fmt.Sprintf("::%s", utinterface.ToString(ox[1]))
	}

	return ""
}

func (ox ConditionObj) Syntax() string {
	return fmt.Sprintf("%s %s %s", ox.Condition1, ox.Operator, ox.Condition2)
}

// ToSQLValueQuery function
func (ox SQLDriver) ToSQLValueQuery(value interface{}, ccol *Column) (res string, ok bool) {
	opts := map[string]string{
		"null":  "NULL",
		"true":  "true",
		"false": "false",
	}

	switch ox {
	case DriverPostgreSQL:
		opts["true"] = "TRUE"
		opts["false"] = "FALSE"
	}

	if value == nil {
		res, ok = opts["null"], true
		return res, ok
	}

	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Ptr {
		valx := val.Elem()
		if valx.Kind() == reflect.Invalid {
			res, ok = opts["null"], true
			return res, ok
		}

		value = valx.Interface()
		if value == nil {
			res, ok = opts["null"], true
			return res, ok
		}

		val = reflect.ValueOf(value)
	}

	if value == nil {
		res, ok = opts["null"], true
		return res, ok
	}

	res, ok = "", false

	switch valx := value.(type) {
	case QRaw:
		res, ok = val.String(), true

	case QArray:
		resx, err := pq.Array(value).Value()
		if err != nil {
			log.Error(err, "Failed to parsing array")
			return res, ok
		}

		res, ok = ox.SQLValueEscape(fmt.Sprintf("%s", resx)), true

	case QCast:
		if res, ok = ox.ToSQLValueQuery(valx.GetValue(), ccol); !ok {
			return res, ok
		}
		res += valx.GetType()

	case QColumn:
		val1 := []string{}
		for i := 0; i < val.Len(); i++ {
			val1 = append(val1, ox.SQLColumnEscape(val.Index(i).String()))
		}
		res, ok = strings.Join(val1, "."), true

	default:
		switch val.Kind() {
		case reflect.Invalid:
			res, ok = opts["null"], true

		case reflect.String:
			res = utstring.Trim(fmt.Sprintf("%s", value))
			if ccol != nil && ccol.Size != 0 && len(res) > ccol.Size {
				res = utstring.Trim(res[:ccol.Size])
			}
			res, ok = ox.SQLValueEscape(res), true

		case reflect.Bool:
			res, ok = opts["false"], true
			if value.(bool) {
				res = opts["true"]
			}

		case reflect.Map:
			arrs := []string{}
			for _, key := range val.MapKeys() {
				val2, ok2 := ox.ToSQLValueQuery(val.MapIndex(key).Interface(), ccol)
				if ok2 {
					arrs = append(arrs, val2)
				}
			}
			res, ok = strings.Join(arrs, ", "), true

		case reflect.Slice:
			arrs := []string{}
			for i := 0; i < val.Len(); i++ {
				val2, ok2 := ox.ToSQLValueQuery(val.Index(i).Interface(), ccol)
				if ok2 {
					arrs = append(arrs, val2)
				}
			}
			res, ok = strings.Join(arrs, ", "), true

		default:
			typ := reflect.TypeOf(value).String()

			switch typ {
			case "int", "int8", "int16", "int32", "int64":
				res, ok = utstring.Int64ToString(val.Int()), true

			case "uint", "uint8", "uint16", "uint32", "uint64":
				res, ok = utstring.Uint64ToString(val.Uint()), true

			case "float32", "float64":
				res, ok = utstring.FloatToString(val.Float()), true

			case "time.Time":
				res, ok = ox.SQLValueEscape(uttime.Format(uttime.DefaultDateTimeFormat, value.(time.Time))), true

			case "sql.NullInt16":
				if value.(sql.NullInt16).Valid == true {
					res, ok = utstring.IntToString(int(value.(sql.NullInt16).Int16)), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullInt32":
				if value.(sql.NullInt32).Valid == true {
					res, ok = utstring.IntToString(int(value.(sql.NullInt32).Int32)), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullInt64":
				if value.(sql.NullInt64).Valid == true {
					res, ok = utstring.Int64ToString(value.(sql.NullInt64).Int64), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullFloat64":
				if value.(sql.NullFloat64).Valid == true {
					res, ok = utstring.FloatToString(value.(sql.NullFloat64).Float64), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullTime":
				if value.(sql.NullTime).Valid == true {
					res, ok = ox.SQLValueEscape(uttime.Format(uttime.DefaultDateTimeFormat, value.(sql.NullTime).Time)), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullString":
				if value.(sql.NullString).Valid == true {
					res, ok = ox.SQLValueEscape(value.(sql.NullString).String), true
				} else {
					res, ok = opts["null"], true
				}
			case "sql.NullBool":
				if value.(sql.NullBool).Valid == true {
					res, ok = utstring.BoolToString(value.(sql.NullBool).Bool), true
				} else {
					res, ok = opts["null"], true
				}
			}
		}
	}

	return res, ok
}

// SQLValueEscape function
func (ox SQLDriver) SQLValueEscape(value string) (res string) {
	quote := "'"

	switch ox {
	default:
		res = quote + strings.ReplaceAll(value, quote, strings.Repeat(quote, 2)) + quote
	}

	return res
}

// SQLColumnEscape function
func (ox SQLDriver) SQLColumnEscape(name string) (res string) {
	quote := "\""

	switch ox {
	default:
		res = quote + strings.ReplaceAll(name, quote, strings.Repeat(quote, 2)) + quote
	}

	return res
}

// ToSQLOperatorQuery function
func (ox SQLDriver) ToSQLOperatorQuery(opr Operator) (res string, ok bool) {
	val := string(opr)

	switch ox {
	default:
		ok = true
		switch opr {
		case OperatorAll:
			fallthrough
		case OperatorEmpty:
			return res, false
		}
	}

	res = val
	return res, ok
}

func (ox SQLDriver) ToSQLConditionObject(cond1 interface{}, opr Operator, cond2 interface{}) (res ConditionObj, ok bool) {
	v, ok := "", true
	if v, ok = ox.ToSQLValueQuery(cond1, nil); !ok {
		return res, false
	}
	res.Condition1 = v

	if v, ok = ox.ToSQLOperatorQuery(opr); !ok {
		return res, false
	}
	res.Operator = v

	ok = true
	switch opr {
	case OperatorBetween:
		ok = false
		ovals := reflect.ValueOf(cond2)
		if ovals.Kind().String() == "slice" {
			if ovals.Len() >= 2 {
				val1, val2 := "", ""

				var ok2 bool
				if val1, ok2 = ox.ToSQLValueQuery(ovals.Index(0).Interface(), nil); !ok2 {
					return res, false
				}

				if val2, ok2 = ox.ToSQLValueQuery(ovals.Index(1).Interface(), nil); !ok2 {
					return res, false
				}

				v = fmt.Sprintf("%s AND %s", val1, val2)
				ok = true
			}
		}
		if !ok {
			return ConditionObj{
				Condition1: "TRUE",
				Operator:   "",
				Condition2: "",
			}, true
		}

	case OperatorIn:
		fallthrough
	case OperatorNotIn:
		val1 := ""
		if val1, ok = ox.ToSQLValueQuery(cond2, nil); !ok {
			return res, false
		}

		if ok {
			v = fmt.Sprintf("(%s)", val1)
		}

	case OperatorIsNotNull:
		fallthrough
	case OperatorIsNull:
		v = ""

	default:
		if v, ok = ox.ToSQLValueQuery(cond2, nil); !ok {
			return res, false
		}
	}
	res.Condition2 = v

	return res, true
}

// ToSQLConditionQuery function
func (ox SQLDriver) ToSQLConditionQuery(cond1 interface{}, opr Operator, cond2 interface{}) (res string, ok bool) {
	obj, ok := ox.ToSQLConditionObject(cond1, opr, cond2)
	if ok {
		res = obj.Syntax()
	}

	return res, ok
}

// U2SQLBinding function
func (ox SQLDriver) U2SQLBinding(query string, values map[string]interface{}) (res string) {
	opts := map[string]string{
		"true":  "true",
		"false": "false",
	}

	switch ox {
	case DriverPostgreSQL:
		opts["true"] = "'t'"
		opts["false"] = "'f'"
	}

	vals := make(map[string]string)
	for k, dd := range values {
		dv, ok := ox.ToSQLValueQuery(dd, nil)
		if ok {
			vals[k] = dv
		}
	}

	res = u2.Binding(query, vals)
	return
}
