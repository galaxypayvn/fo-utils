package packets

import (
	fmt "fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gearintellix/structs"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"bitbucket.org/finesys/finesys-utility/libs/serror"
	"bitbucket.org/finesys/finesys-utility/utils/utarray"
	"bitbucket.org/finesys/finesys-utility/utils/utstring"
	"bitbucket.org/finesys/finesys-utility/utils/uttime"

	log "github.com/sirupsen/logrus"
)

type ViewArguments struct {
	Offset     int64                  `json:"offset,omitempty"`
	Limit      int                    `json:"limit,omitempty"`
	Sorting    []string               `json:"sorting,omitempty"`
	Conditions map[string]interface{} `json:"conditions,omitempty"`
	Fields     []string               `json:"fields,omitempty"`
	Groups     []string               `json:"groups,omitempty"`
}

type Error struct {
	Code     int32  `json:"code,omitempty"`
	Key      string `json:"key,omitempty"`
	Message  string `json:"message"`
	Comment  string `json:"comment,omitempty"`
	MoreInfo string `json:"moreInfo,omitempty"`
}

type ViewAttribute struct {
	ID          int64       `json:"id,omitempty"`
	Key         string      `json:"name"`
	Group       string      `json:"group,omitempty"`
	Description string      `json:"description,omitempty"`
	Value       interface{} `json:"value"`
}

type State struct {
	Success bool    `json:"success"`
	Errors  []Error `json:"errors,omitempty"`
}

type ViewData struct {
	Fields     []string        `json:"fields"`
	Attributes []ViewAttribute `json:"attributes"`
}

type InputData struct {
	ID          int64  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value,omitempty"`
}

type InputAnyData struct {
	ID          int64       `json:"id,omitempty"`
	Description string      `json:"description,omitempty"`
	Value       interface{} `json:"value,omitempty"`
}

type ViewResponse struct {
	State State      `json:"state"`
	Datas []ViewData `json:"datas"`
}

type InputRequest struct {
	Conditions map[string]interface{} `json:"conditions,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Metas      Metas                  `json:"metas,omitempty"`
}

type UpdateResponse struct {
	State         State           `json:"state"`
	AffectedCount int64           `json:"affectedCount,omitempty"`
	AffectedIDs   map[int64]int64 `json:"affectedIDs,omitempty"`
}

type InsertResponse struct {
	State State                  `json:"state"`
	IDs   map[int64]int64        `json:"ids,omitempty"`
	Metas map[int64]ViewResponse `json:"metas,omitempty"`
}

type ICopyToStruct interface {
	Cast(val interface{}) (res interface{}, err error)
}

func (ox *PInputRequest) Convert() *InputRequest {
	out := &InputRequest{
		Conditions: make(map[string]interface{}),
		Attributes: make(map[string]interface{}),
		Metas:      make(Metas),
	}

	for k, v := range ox.Conditions {
		out.Conditions[k] = v.ExtractV2()
	}

	for k, v := range ox.Attributes {
		out.Attributes[k] = v.ExtractV2()
	}

	for k, v := range ox.Metas {
		out.Metas[k] = InputData{
			ID:          v.ID,
			Description: v.Description,
			Value:       v.Value,
		}
	}

	return out
}

func (ox *PViewRequest) Convert() *ViewArguments {
	out := &ViewArguments{
		Offset:     ox.Offset,
		Limit:      int(ox.Limit),
		Sorting:    ox.Sorting,
		Conditions: make(map[string]interface{}),
		Fields:     ox.Fields,
		Groups:     ox.Groups,
	}

	for k, v := range ox.Conditions {
		out.Conditions[k] = v.ExtractV2()
	}
	return out
}

func (ox *PViewResponse) Convert() *ViewResponse {
	out := &ViewResponse{
		State: State{
			Success: ox.Status.Success,
			Errors:  []Error{},
		},
		Datas: []ViewData{},
	}

	for _, v := range ox.Status.Errors {
		out.State.Errors = append(out.State.Errors, Error{
			Code:    v.Code,
			Key:     v.Key,
			Message: v.Message,
			Comment: v.Comment,
		})
	}

	for _, v := range ox.Datas {
		d := ViewData{
			Fields:     v.Fields,
			Attributes: []ViewAttribute{},
		}

		for _, v2 := range v.Attributes {
			d.Attributes = append(d.Attributes, ViewAttribute{
				ID:          v2.ID,
				Key:         v2.Key,
				Group:       v2.Group,
				Description: v2.Description,
				Value:       v2.Value.Extract(),
			})
		}

		out.Datas = append(out.Datas, d)
	}

	return out
}

func (ox *PViewResponse) ConvertV2() *ViewResponse {
	out := &ViewResponse{
		State: State{
			Success: ox.Status.Success,
			Errors:  []Error{},
		},
		Datas: []ViewData{},
	}

	for _, v := range ox.Status.Errors {
		out.State.Errors = append(out.State.Errors, Error{
			Code:    v.Code,
			Key:     v.Key,
			Message: v.Message,
			Comment: v.Comment,
		})
	}

	for _, v := range ox.Datas {
		d := ViewData{
			Fields:     v.Fields,
			Attributes: []ViewAttribute{},
		}

		for _, v2 := range v.Attributes {
			d.Attributes = append(d.Attributes, ViewAttribute{
				ID:          v2.ID,
				Key:         v2.Key,
				Group:       v2.Group,
				Description: v2.Description,
				Value:       v2.Value.ExtractV2(),
			})
		}

		out.Datas = append(out.Datas, d)
	}

	return out
}

func (ox ViewResponse) Convert() *PViewResponse {
	out := &PViewResponse{
		Status: &PStatus{
			Success: ox.State.Success,
			Errors:  []*PError{},
		},
		Datas: []*PViewData{},
	}

	for _, v := range ox.State.Errors {
		out.Status.Errors = append(out.Status.Errors, &PError{
			Code:    v.Code,
			Key:     v.Key,
			Message: v.Message,
			Comment: v.Comment,
		})
	}

	for _, v := range ox.Datas {
		d := &PViewData{
			Fields:     v.Fields,
			Attributes: []*PViewAttribute{},
		}

		for _, v2 := range v.Attributes {
			d.Attributes = append(d.Attributes, &PViewAttribute{
				ID:          v2.ID,
				Key:         v2.Key,
				Group:       v2.Group,
				Description: v2.Description,
				Value:       ToObje(v2.Value),
			})
		}

		out.Datas = append(out.Datas, d)
	}

	return out
}

func (ox *PStatus) CaptureSError(errx serror.SError) {
	if ox.Errors == nil {
		ox.Errors = []*PError{}
	}

	if errx == nil {
		return
	}

	switch errx.Key() {
	case "!":
		log.Error(errx)

	default:
		log.Warn(errx)

		moreInfo := fmt.Sprintf("file://%s:%d", errx.File(), errx.Line())
		if !utarray.IsExist(errx.Key(), []string{"", "-"}) {
			moreInfo = fmt.Sprintf("const://%s", errx.Key())
		}

		ox.Errors = append(ox.Errors, &PError{
			Code:     int32(errx.Code()),
			Key:      errx.Key(),
			Message:  errx.Title(),
			Comment:  errx.SimpleString(),
			MoreInfo: moreInfo,
		})
	}
}

func (ox *State) CaptureSError(errx serror.SError) {
	if ox.Errors == nil {
		ox.Errors = []Error{}
	}

	if errx == nil {
		return
	}

	switch errx.Key() {
	case "!":
		log.Error(errx)

	default:
		log.Warn(errx)

		moreInfo := fmt.Sprintf("file://%s:%d", errx.File(), errx.Line())
		if !utarray.IsExist(errx.Key(), []string{"", "-"}) {
			moreInfo = fmt.Sprintf("const://%s", errx.Key())
		}

		ox.Errors = append(ox.Errors, Error{
			Code:     int32(errx.Code()),
			Key:      errx.Key(),
			Message:  errx.Title(),
			Comment:  errx.SimpleString(),
			MoreInfo: moreInfo,
		})
	}
}

func (ox InsertResponse) Convert() *PInsertResponse {
	resp := &PInsertResponse{
		Status: &PStatus{
			Success: ox.State.Success,
			Errors:  []*PError{},
		},
		IDs:   ox.IDs,
		Metas: make(map[int64]*PViewResponse),
	}

	for _, v := range ox.State.Errors {
		resp.Status.Errors = append(resp.Status.Errors, &PError{
			Code:    v.Code,
			Key:     v.Key,
			Message: v.Message,
			Comment: v.Comment,
		})
	}

	for k, v := range ox.Metas {
		v2 := &PViewResponse{
			Status: &PStatus{
				Success: v.State.Success,
				Errors:  []*PError{},
			},
			Datas: []*PViewData{},
		}

		for _, v3 := range v.State.Errors {
			v2.Status.Errors = append(v2.Status.Errors, &PError{
				Code:    v3.Code,
				Key:     v3.Key,
				Message: v3.Message,
				Comment: v3.Comment,
			})
		}

		for _, v3 := range v.Datas {
			v4 := &PViewData{
				Fields:     v3.Fields,
				Attributes: []*PViewAttribute{},
			}
			for _, v5 := range v3.Attributes {
				v4.Attributes = append(v4.Attributes, &PViewAttribute{
					ID:          v5.ID,
					Key:         v5.Key,
					Group:       v5.Group,
					Description: v5.Description,
					Value:       ToObje(v5.Value),
				})
			}
			v2.Datas = append(v2.Datas, v4)
		}

		resp.Metas[k] = v2
	}

	return resp
}

func (ox ViewResponse) CopyToStruct(obj interface{}) (errx serror.SError) {
	if obj == nil {
		return serror.Newc("Object cannot be null", "while check obj")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "while check obj")
	}

	if len(ox.Datas) > 0 {
		errx = ox.Datas[0].CopyToStruct(obj)
		if errx != nil {
			return errx
		}
	}

	fields := structs.Fields(obj)
	for _, v := range fields {
		var err error
		switch strings.ToLower(utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())) {
		case "@success":
			err = v.Set(ox.State.Success)
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, "while set metas")
			return errx
		}
	}

	return errx
}

func (ox *PViewResponse) CopyToStruct(obj interface{}) (errx serror.SError) {
	if obj == nil {
		return serror.Newc("Object cannot be null", "while check obj")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "while check obj")
	}

	if ox.Datas != nil && len(ox.Datas) > 0 && ox.Datas[0] != nil {
		errx = ox.Datas[0].CopyToStruct(obj)
		if errx != nil {
			errx.AddComments("while copy to struct")
			return errx
		}
	}

	fields := structs.Fields(obj)
	for _, v := range fields {
		var err error
		switch strings.ToLower(utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())) {
		case "@success":
			if ox.Status != nil {
				err = v.Set(ox.Status.Success)
			}
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, "while set metas")
			return errx
		}
	}

	return errx
}

func (ox *PViewResponse) CopyToStructV2(isStrict bool, obj interface{}) (errx serror.SError) {
	if obj == nil {
		errx = serror.New("Object cannot be null")
		return
	}

	if !structs.IsStruct(obj) {
		errx = serror.New("Object is not struct")
		return
	}

	if ox.Datas != nil && len(ox.Datas) > 0 && ox.Datas[0] != nil {
		errx = ox.Datas[0].CopyToStructV2(isStrict, obj)
		if errx != nil {
			errx.AddComments("while copy to struct")
			return
		}
	}

	for _, v := range structs.Fields(obj) {
		var (
			err error
			nm  string
		)
		switch nm = strings.ToLower(utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())); nm {
		case "@success":
			if ox.Status != nil {
				err = v.Set(ox.Status.Success)
			}
		}

		if err != nil {
			erry := serror.NewFromErrorc(err, fmt.Sprintf("Failed to set value %s", nm))
			if !isStrict {
				log.Warn(erry)
				continue
			}

			errx = erry
			return
		}
	}

	return
}

func (ox ViewData) CopyToStruct(obj interface{}) (errx serror.SError) {
	var err error

	if obj == nil {
		return serror.Newc("Object cannot be null", "while check obj")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "while check obj")
	}

	nms := make(map[string]int)
	fields := structs.Fields(obj)
	for k, v := range fields {
		nms[utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())] = k
	}

	setValue := func(idx int, val interface{}) {
		defer func() {
			if err := recover(); err != nil {
				errx = serror.Newc(err.(string), "while set value")
			}
		}()

		origin := fields[idx].Value()
		model := reflect.TypeOf((*ICopyToStruct)(nil)).Elem()
		if reflect.TypeOf(origin).Implements(model) {
			val, err = origin.(ICopyToStruct).Cast(val)
		}

		err = fields[idx].Set(val)
	}

	for _, v := range ox.Attributes {
		if idx, ok := nms[fmt.Sprintf("%s.%s", v.Group, v.Key)]; ok {
			setValue(idx, v.Value)

		} else if idx, ok := nms[v.Key]; ok {
			if v.Group != "@" && utarray.IsExist(fmt.Sprintf("@.%s", v.Key), ox.Fields) {
				continue
			}

			setValue(idx, v.Value)
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, "while set from view data")
			return errx
		}
	}

	return errx
}

func (ox *PViewData) CopyToStruct(obj interface{}) (errx serror.SError) {
	var err error

	if obj == nil {
		return serror.Newc("Object cannot be null", "while check obj")
	}

	if !structs.IsStruct(obj) {
		return serror.Newc("Object is not struct", "while check obj")
	}

	nms := make(map[string]int)
	fields := structs.Fields(obj)
	for k, v := range fields {
		nms[utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())] = k
	}

	for _, v := range ox.Attributes {
		if idx, ok := nms[fmt.Sprintf("%s.%s", v.Group, v.Key)]; ok {
			if v.Value != nil {
				_ = v.Value.Unpack(fields[idx].RawValue().Addr().Interface())
			}

		} else if idx, ok := nms[v.Key]; ok {
			if v.Group != "@" && utarray.IsExist(fmt.Sprintf("@.%s", v.Key), ox.Fields) {
				continue
			}

			if v.Value != nil {
				_ = v.Value.Unpack(fields[idx].RawValue().Addr().Interface())
			}
		}

		if err != nil {
			errx = serror.NewFromErrorc(err, "while set from view data")
			return errx
		}
	}

	return errx
}

func (ox *PViewData) CopyToStructV2(isStrict bool, obj interface{}) (errx serror.SError) {
	if obj == nil {
		errx = serror.New("Object cannot be null")
		return
	}

	if !structs.IsStruct(obj) {
		errx = serror.New("Object is not struct")
		return
	}

	var (
		nms    = make(map[string]int)
		fields = structs.Fields(obj)
	)
	for k, v := range fields {
		nms[utstring.Chains(v.Tag("key"), (strings.Split(v.Tag("json"), ",")[0]), v.Name())] = k
	}

	for _, v := range ox.Attributes {
		var (
			currentField *structs.Field
			erry         serror.SError
		)
		if idx, ok := nms[fmt.Sprintf("%s.%s", v.Group, v.Key)]; ok {
			currentField = fields[idx]
			if v.Value != nil {
				erry = v.Value.UnpackV2(currentField.RawValue().Addr().Interface(), obj, currentField.Name())
			}

		} else if idx, ok := nms[v.Key]; ok {
			currentField = fields[idx]
			if v.Group != "@" && utarray.IsExist(fmt.Sprintf("@.%s", v.Key), ox.Fields) {
				continue
			}

			if v.Value != nil {
				erry = v.Value.UnpackV2(currentField.RawValue().Addr().Interface(), obj, currentField.Name())
			}
		}

		if erry != nil {
			erry.AddCommentf("while set value of %s.%s", v.Group, v.Key)
			if !isStrict {
				log.Warn(erry)
				continue
			}

			errx = erry
			return
		}
	}

	return
}

func (ox *ViewData) AddAttributesWithCast(grp string, attrs []ViewAttribute) {
	for k, v := range attrs {
		if v.Group == "@" {
			attrs[k].Group = utstring.Chains(grp, v.Group)
		}
	}
	ox.AddAttributes(attrs)
}

func (ox *ViewData) AddAttributes(attrs []ViewAttribute) {
	for _, v := range attrs {
		key := v.Key
		if v.Group != "" {
			key = fmt.Sprintf("%s.%s", v.Group, key)
		}

		if exist, index := utarray.IsExists(key, (*ox).Fields); exist {
			(*ox).Attributes[index] = v
			continue
		}

		(*ox).Fields = append((*ox).Fields, key)
		(*ox).Attributes = append((*ox).Attributes, v)
	}
}

func (ox *PViewData) AddAttributesWithCast(grp string, attrs []*PViewAttribute) {
	npvie := []*PViewAttribute{}
	for _, v := range attrs {
		cgrp := v.Group
		if cgrp == "@" {
			cgrp = utstring.Chains(grp, cgrp)
		}
		npvie = append(npvie, &PViewAttribute{
			ID:    v.ID,
			Key:   v.Key,
			Group: cgrp,
			Value: v.Value,
		})
	}
	ox.AddAttributes(npvie)
}

func (ox *PViewData) AddAttributes(attrs []*PViewAttribute) {
	for _, v := range attrs {
		key := v.Key
		if v.Group != "" {
			key = fmt.Sprintf("%s.%s", v.Group, key)
		}

		if exist, index := utarray.IsExists(key, (*ox).Fields); exist {
			(*ox).Attributes[index] = v
			continue
		}

		(*ox).Fields = append((*ox).Fields, key)
		(*ox).Attributes = append((*ox).Attributes, v)
	}
}

func (ox *ViewResponse) AddDatas(datas [][]ViewAttribute) {
	for _, v := range datas {
		row := ViewData{}
		row.AddAttributes(v)

		ox.Datas = append(ox.Datas, row)
	}
}

func (ox *PViewResponse) AddDatas(datas [][]*PViewAttribute) {
	for _, v := range datas {
		row := &PViewData{}
		row.AddAttributes(v)

		ox.Datas = append(ox.Datas, row)
	}
}

func (ox *InsertResponse) SetResponseRow(index int64, newID int64, meta ViewResponse) {
	ox.IDs[index] = newID
	ox.Metas[index] = meta
}

func (ox *PInsertResponse) SetResponseRow(index int64, newID int64, meta *PViewResponse) {
	ox.IDs[index] = newID
	ox.Metas[index] = meta
}

func StrToTimex(val string) (tim *timestamp.Timestamp, errx serror.SError) {
	var rtim time.Time
	rtim, errx = uttime.ParseFromString(val)
	if errx != nil {
		errx.AddComments("while parse from string")
		return tim, errx
	}

	tim, errx = ToTimex(rtim)
	if errx != nil {
		errx.AddComments("while cast to time")
		return tim, errx
	}
	return tim, errx
}

func StrToTime(val string) *timestamp.Timestamp {
	tim, errx := StrToTimex(val)
	if errx != nil {
		errx.Panic()
	}

	return tim
}

func StrToTimee(val string) *timestamp.Timestamp {
	tim, errx := StrToTimex(val)
	if errx != nil {
		return nil
	}

	return tim
}

func ToTimex(val time.Time) (tim *timestamp.Timestamp, errx serror.SError) {
	var err error
	tim, err = ptypes.TimestampProto(val)
	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to convert time to proto timestamp")
		return tim, errx
	}

	return tim, errx
}

func ToTime(val time.Time) *timestamp.Timestamp {
	tim, errx := ToTimex(val)
	if errx != nil {
		errx.Panic()
	}

	return tim
}

func ToTimee(val time.Time) *timestamp.Timestamp {
	tim, errx := ToTimex(val)
	if errx != nil {
		return nil
	}

	return tim
}

func FromTimex(val *timestamp.Timestamp) (tim time.Time, errx serror.SError) {
	tim = time.Now()
	if val == nil {
		errx = serror.New("Timestamp is nil")
		return tim, errx
	}

	var err error
	tim, err = ptypes.Timestamp(val)
	if err != nil {
		errx = serror.NewFromErrorc(err, "Failed to convert proto timestamp to time")
		return tim, errx
	}

	tim = tim.In(time.Local)
	return tim, errx
}

func FromTime(val *timestamp.Timestamp) time.Time {
	tim, errx := FromTimex(val)
	if errx != nil {
		errx.Panic()
	}

	return tim
}

func FromTimee(val *timestamp.Timestamp) time.Time {
	tim, err := FromTimex(val)
	if err != nil {
		return time.Time{}
	}

	return tim
}
