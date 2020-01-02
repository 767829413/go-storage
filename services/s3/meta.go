// Code generated by go generate via internal/cmd/meta; DO NOT EDIT.
package s3

import (
	"github.com/Xuanwo/storage"
	"github.com/Xuanwo/storage/pkg/credential"
	"github.com/Xuanwo/storage/pkg/endpoint"
	"github.com/Xuanwo/storage/pkg/segment"
	"github.com/Xuanwo/storage/types"
	"github.com/Xuanwo/storage/types/pairs"
)

var _ credential.Provider
var _ endpoint.Provider
var _ segment.Segment
var _ storage.Storager

// Type is the type for s3
const Type = "s3"

var allowedStoragePairs = map[string]map[string]struct{}{
	"init": {
		"work_dir": struct{}{},
	},
	"init_segment": {
		"part_size": struct{}{},
	},
	"list": {
		"dir_func":  struct{}{},
		"file_func": struct{}{},
	},
	"list_segments": {
		"segment_func": struct{}{},
	},
	"write": {
		"checksum":      struct{}{},
		"size":          struct{}{},
		"storage_class": struct{}{},
	},
}

var allowedServicePairs = map[string]map[string]struct{}{
	"create": {
		"location": struct{}{},
	},
	"delete": {
		"location": struct{}{},
	},
	"get": {
		"location": struct{}{},
	},
	"init": {
		"credential": struct{}{},
		"endpoint":   struct{}{},
	},
	"list": {
		"storager_func": struct{}{},
	},
}

type pairStorageInit struct {
	HasWorkDir bool
	WorkDir    string
}

func parseStoragePairInit(opts ...*types.Pair) (*pairStorageInit, error) {
	result := &pairStorageInit{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["init"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["init"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.WorkDir]
	if ok {
		result.HasWorkDir = true
		result.WorkDir = v.(string)
	}
	return result, nil
}

type pairStorageInitSegment struct {
	HasPartSize bool
	PartSize    int64
}

func parseStoragePairInitSegment(opts ...*types.Pair) (*pairStorageInitSegment, error) {
	result := &pairStorageInitSegment{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["init_segment"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["init_segment"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.PartSize]
	if !ok {
		return nil, types.NewErrPairRequired(pairs.PartSize)
	}
	if ok {
		result.HasPartSize = true
		result.PartSize = v.(int64)
	}
	return result, nil
}

type pairStorageList struct {
	HasDirFunc  bool
	DirFunc     types.ObjectFunc
	HasFileFunc bool
	FileFunc    types.ObjectFunc
}

func parseStoragePairList(opts ...*types.Pair) (*pairStorageList, error) {
	result := &pairStorageList{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["list"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["list"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.DirFunc]
	if ok {
		result.HasDirFunc = true
		result.DirFunc = v.(types.ObjectFunc)
	}
	v, ok = values[pairs.FileFunc]
	if ok {
		result.HasFileFunc = true
		result.FileFunc = v.(types.ObjectFunc)
	}
	return result, nil
}

type pairStorageListSegments struct {
	HasSegmentFunc bool
	SegmentFunc    segment.Func
}

func parseStoragePairListSegments(opts ...*types.Pair) (*pairStorageListSegments, error) {
	result := &pairStorageListSegments{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["list_segments"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["list_segments"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.SegmentFunc]
	if ok {
		result.HasSegmentFunc = true
		result.SegmentFunc = v.(segment.Func)
	}
	return result, nil
}

type pairStorageWrite struct {
	HasChecksum     bool
	Checksum        string
	HasSize         bool
	Size            int64
	HasStorageClass bool
	StorageClass    string
}

func parseStoragePairWrite(opts ...*types.Pair) (*pairStorageWrite, error) {
	result := &pairStorageWrite{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedStoragePairs["write"]; !ok {
			continue
		}
		if _, ok := allowedStoragePairs["write"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Checksum]
	if ok {
		result.HasChecksum = true
		result.Checksum = v.(string)
	}
	v, ok = values[pairs.Size]
	if !ok {
		return nil, types.NewErrPairRequired(pairs.Size)
	}
	if ok {
		result.HasSize = true
		result.Size = v.(int64)
	}
	v, ok = values[pairs.StorageClass]
	if ok {
		result.HasStorageClass = true
		result.StorageClass = v.(string)
	}
	return result, nil
}

type pairServiceCreate struct {
	HasLocation bool
	Location    string
}

func parseServicePairCreate(opts ...*types.Pair) (*pairServiceCreate, error) {
	result := &pairServiceCreate{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedServicePairs["create"]; !ok {
			continue
		}
		if _, ok := allowedServicePairs["create"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Location]
	if !ok {
		return nil, types.NewErrPairRequired(pairs.Location)
	}
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceDelete struct {
	HasLocation bool
	Location    string
}

func parseServicePairDelete(opts ...*types.Pair) (*pairServiceDelete, error) {
	result := &pairServiceDelete{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedServicePairs["delete"]; !ok {
			continue
		}
		if _, ok := allowedServicePairs["delete"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceGet struct {
	HasLocation bool
	Location    string
}

func parseServicePairGet(opts ...*types.Pair) (*pairServiceGet, error) {
	result := &pairServiceGet{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedServicePairs["get"]; !ok {
			continue
		}
		if _, ok := allowedServicePairs["get"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Location]
	if ok {
		result.HasLocation = true
		result.Location = v.(string)
	}
	return result, nil
}

type pairServiceInit struct {
	HasCredential bool
	Credential    *credential.Provider
	HasEndpoint   bool
	Endpoint      endpoint.Provider
}

func parseServicePairInit(opts ...*types.Pair) (*pairServiceInit, error) {
	result := &pairServiceInit{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedServicePairs["init"]; !ok {
			continue
		}
		if _, ok := allowedServicePairs["init"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.Credential]
	if !ok {
		return nil, types.NewErrPairRequired(pairs.Credential)
	}
	if ok {
		result.HasCredential = true
		result.Credential = v.(*credential.Provider)
	}
	v, ok = values[pairs.Endpoint]
	if ok {
		result.HasEndpoint = true
		result.Endpoint = v.(endpoint.Provider)
	}
	return result, nil
}

type pairServiceList struct {
	HasStoragerFunc bool
	StoragerFunc    storage.StoragerFunc
}

func parseServicePairList(opts ...*types.Pair) (*pairServiceList, error) {
	result := &pairServiceList{}

	values := make(map[string]interface{})
	for _, v := range opts {
		if _, ok := allowedServicePairs["list"]; !ok {
			continue
		}
		if _, ok := allowedServicePairs["list"][v.Key]; !ok {
			continue
		}
		values[v.Key] = v.Value
	}
	var v interface{}
	var ok bool
	v, ok = values[pairs.StoragerFunc]
	if ok {
		result.HasStoragerFunc = true
		result.StoragerFunc = v.(storage.StoragerFunc)
	}
	return result, nil
}