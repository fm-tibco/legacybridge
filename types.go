package legacybridge

import (
	"encoding/json"
	"errors"

	legacyData "github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/project-flogo/core/data"
)

// ToTypeEnum get the data type that corresponds to the specified name
func ToNewTypeFromLegacy(legacyType legacyData.Type) (data.Type, error) {

	switch legacyType {
	case legacyData.TypeAny:
		return data.TypeAny, nil
	case legacyData.TypeString:
		return data.TypeString, nil
	case legacyData.TypeInteger:
		return data.TypeInt, nil
	case legacyData.TypeLong:
		return data.TypeInt64, nil
	case legacyData.TypeDouble:
		return data.TypeFloat64, nil
	case legacyData.TypeBoolean:
		return data.TypeBool, nil
	case legacyData.TypeObject:
		return data.TypeObject, nil
	case legacyData.TypeParams:
		return data.TypeParams, nil
	case legacyData.TypeArray:
		return data.TypeArray, nil
	case legacyData.TypeComplexObject:
		return data.TypeObject, nil
	default:
		return 0, errors.New("unknown type: " + legacyType.String())
	}
}

// ToTypeEnum get the data type that corresponds to the specified name
func ToLegacyFromNewType(dataType data.Type) (legacyData.Type, error) {

	switch dataType {
	case data.TypeAny:
		return legacyData.TypeAny, nil
	case data.TypeString:
		return legacyData.TypeString, nil
	case data.TypeInt:
		return legacyData.TypeInteger, nil
	case data.TypeInt64:
		return legacyData.TypeLong, nil
	case data.TypeFloat64:
		return legacyData.TypeDouble, nil
	case data.TypeBool:
		return legacyData.TypeBoolean, nil
	case data.TypeObject:
		return legacyData.TypeObject, nil
	case data.TypeParams:
		return legacyData.TypeParams, nil
	case data.TypeArray:
		return legacyData.TypeArray, nil
	default:
		return 0, errors.New("unknown type: " + dataType.String())
	}
}

func GetComplexObjectInfo(val interface{}) (interface{}, string, bool) {

	switch t := val.(type) {
	case string:
		if val == "" {
			return nil, "", false
		} else {
			var complexMap map[string]interface{}
			err := json.Unmarshal([]byte(t), &complexMap)
			if err != nil {
				return nil, "", false
			}

			v, hasVal := complexMap["value"]
			mdI, hasMd := complexMap["metadata"]
			md := ""
			if hasMd {
				md, hasMd = mdI.(string)
			}
			if hasVal && hasMd {
				return v, md, true
			}
		}
	case map[string]interface{}:
		v, hasVal := t["value"]
		mdI, hasMd := t["metadata"]
		md := ""
		if hasMd {
			md, hasMd = mdI.(string)
		}

		if hasVal || hasMd {
			return v, md, true
		}
	case *legacyData.ComplexObject:
		return t.Value, t.Metadata, true
	default:
		return nil, "", false
	}

	return nil, "", false
}
