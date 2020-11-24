package parser

import (
	"errors"
	"fmt"
	"reflect"
)

/***
1、基本结构体: string, int, bool, float
2、不支持解析time.Time类型
3、struct
4、ptr
 */

type Parser struct {
	tagKey string 
}

func NewParser(key string) *Parser {
	return &Parser{tagKey:key}
}

// 解析入口
func (p *Parser) ParseEntry(m map[string]interface{}, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("解析目标不是*struct类型")
	}

	return p.ParsePtr(m, rv)
}

// 解析指针
func (p *Parser) ParsePtr(m map[string]interface{}, rv reflect.Value) error {
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("解析目标须为ptr类型")
	}

	rve := rv.Elem()
	switch rve.Kind() {
	case reflect.Ptr:
		return p.ParsePtr(m,rve)
	case reflect.Struct:
		return p.ParseStruct(m, rve)
	case reflect.Invalid:
		rv = reflect.New(rv.Type())

		return p.ParsePtr(m, rv)
	default:
		return fmt.Errorf("无效的类型")
	}
}

func (p *Parser) ParseStruct(m map[string]interface{}, rv reflect.Value) error {
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("解析目标须为struct类型")
	}

	rt := rv.Type()
	num := rt.NumField()
	for i := 0; i < num; i++ {
		fType := rt.Field(i)
		fValue := rv.Field(i)

		p.parseField(m, fType, fValue)
	}

	return nil
}

func (p *Parser) parseField(m map[string]interface{}, rf reflect.StructField, rv reflect.Value) error {
	key := rf.Name
	if len(rf.Tag) != 0 {
		key = rf.Tag.Get(p.tagKey)
	}

	if val, ok := m[key]; ok {
		vt := reflect.TypeOf(val)
		fmt.Println(vt, rv.Type())
		switch {
		case vt == rv.Type():
			// 通用类型转换
			if rv.CanSet() {
				rv.Set(reflect.ValueOf(val))
			}
			return nil
		case vt.Kind() == reflect.Float64:
			// 针对反序列化出的float64类型进行转换
			val, err := p.convertNum(val, rv)
			if err != nil {
				return err
			}
			m[key] = val
			p.parseField(m, rf, rv)
		case vt.Kind() == reflect.Map && rv.Kind() == reflect.Struct:
			// 处理struct类型的数据
			mval := val.(map[string]interface{})
			p.ParseStruct(mval, rv)
		case vt.Kind() == reflect.Map && rv.Kind() == reflect.Ptr:
			mval := val.(map[string]interface{})
			p.ParsePtr(mval, rv)
		}
	}

	return nil
}


func (p *Parser) convertNum(val interface{}, rv reflect.Value) (interface{}, error) {
	if reflect.TypeOf(val).Kind() != reflect.Float64 {
		return nil, errors.New("解析类型须为float64类型")
	}
	v := reflect.ValueOf(val).Float()
	var ival interface{}
	if rv.CanSet() {
		switch rv.Kind() {
		case reflect.Int:
			ival = int(v)
		case reflect.Int8:
			ival = int8(v)
		case reflect.Int16:
			ival = int16(v)
		case reflect.Int32:
			ival = int32(v)
		case reflect.Int64:
			ival = int64(v)
		case reflect.Uint:
			ival = uint(v)
		case reflect.Uint8:
			ival = uint8(v)
		case reflect.Uint16:
			ival = uint16(v)
		case reflect.Uint32:
			ival = uint32(v)
		case reflect.Uint64:
			ival = uint64(v)
		case reflect.Float32:
			ival = float32(v)
		default:
			return nil, fmt.Errorf("type incorrect")
		}
	}

	return ival, nil
}

