package spatial

import (
	"fmt"
	"os"

	"github.com/noypi/kv"
	"github.com/noypi/kv/gtreap"
)

type Options interface{}

type OptKVDir struct{ Dir string }
type OptKVName struct{ AnyKey interface{} }

// private configurations
type optKVOptions struct{ KVOptions []kv.Option }
type optDbname struct{ Name string }
type optEnableExtInfo struct{ Enable bool }
type optExtInfoFilePath struct{ Path string }

const (
	cOptKVDir = iota
	cOptKVName

	// private configurations
	cOptKVOptions
	cOptDbname
	cOptEnableExtInfo
	cOptExtInfoFilePath
)

func ParseOpts(opts []Options) (m map[int]interface{}, err error) {
	m = map[int]interface{}{}
	for _, opt := range opts {
		switch v := opt.(type) {
		case OptKVDir:
			m[cOptKVDir] = v.Dir
		case OptKVName:
			fmt.Println("kvname=", v)
			if "true" == fmt.Sprintf("%v", v) {
				panic("where coming from")
			}
			m[cOptKVName] = v.AnyKey
		case optKVOptions:
			m[cOptKVOptions] = v.KVOptions
		case optDbname:
			m[cOptDbname] = v.Name
		case optEnableExtInfo:
			m[cOptEnableExtInfo] = v.Enable
		case optExtInfoFilePath:
			m[cOptExtInfoFilePath] = v.Path
		}
	}
	err = validateOpts(m)
	return
}

func validateOpts(m map[int]interface{}) error {
	if _, has := m[cOptKVName]; !has {
		m[cOptKVName] = gtreap.Name
	}

	if v, has := m[cOptKVDir]; has {
		if info, err := os.Stat(v.(string)); nil != err {
			if os.IsNotExist(err) {
				os.MkdirAll(v.(string), os.ModePerm)
			} else {
				return err
			}
		} else if !info.IsDir() {
			return fmt.Errorf("OptKVDir is not a valid directory")
		}
	}

	if _, has := m[cOptDbname]; !has {
		m[cOptDbname] = "_1d.db"
	}

	if _, has := m[cOptEnableExtInfo]; !has {
		m[cOptEnableExtInfo] = true
	}

	if kvdir, has := m[cOptKVDir].(string); has {
		kvextinfo := fmt.Sprintf("%s/%s", kvdir, "extinfo.db")
		kvfilepath := fmt.Sprintf("%s/%s", kvdir, m[cOptDbname].(string))

		m[cOptKVOptions] = []kv.Option{kv.OptFilePath{kvfilepath}}
		m[cOptExtInfoFilePath] = kvextinfo
	}

	return nil
}
