package response

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PageResp struct {
	PageNo   int
	PageSize int
	Count    int64
	Data     interface{}
}

type PageReq struct {
	PageNo   int
	PageSize int
}

func CheckErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	logrus.Error(msg)
	return err
}

func ErrRecordNotFound(err error, msg string) error {
	if err == gorm.ErrRecordNotFound || err == nil {
		return nil
	}
	logrus.Error(msg)
	return err
}

func CopyStruct(dst, src interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	if dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("destination must be a pointer to a struct")
	}
	dstElem := dstVal.Elem()

	if dstElem.Kind() != reflect.Struct {
		return fmt.Errorf("destination must be a struct")
	}

	// Handle both struct and pointer to struct for src
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if srcVal.Kind() != reflect.Struct {
		return fmt.Errorf("source must be a struct or a pointer to a struct")
	}

	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldType := srcType.Field(i)
		if srcFieldType.PkgPath != "" { // Field is unexportable
			continue
		}

		dstField := dstElem.FieldByName(srcFieldType.Name)
		if !dstField.IsValid() {
			fmt.Printf("No such field: %s in destination\n", srcFieldType.Name)
			continue
		}
		if !dstField.CanSet() {
			fmt.Printf("Cannot set field: %s in destination\n", srcFieldType.Name)
			continue
		}
		if dstField.Type() != srcField.Type() {
			fmt.Printf("Type mismatch %s: %s != %s\n", srcFieldType.Name, dstField.Type(), srcField.Type())
			continue
		}

		dstField.Set(srcField)
	}
	return nil
}