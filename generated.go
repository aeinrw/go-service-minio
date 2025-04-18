// Code generated by go generate via cmd/definitions; DO NOT EDIT.
package minio

import (
	"context"
	"io"
	"time"

	"github.com/beyondstorage/go-storage/v4/pkg/httpclient"
	"github.com/beyondstorage/go-storage/v4/services"
	. "github.com/beyondstorage/go-storage/v4/types"
)

var _ Storager
var _ services.ServiceError
var _ httpclient.Options
var _ time.Duration

// Type is the type for minio
const Type = "minio"

// ObjectSystemMetadata stores system metadata for object.
type ObjectSystemMetadata struct {
	// StorageClass
	StorageClass string
}

// GetObjectSystemMetadata will get ObjectSystemMetadata from Object.
//
// - This function should not be called by service implementer.
// - The returning ObjectServiceMetadata is read only and should not be modified.
func GetObjectSystemMetadata(o *Object) ObjectSystemMetadata {
	sm, ok := o.GetSystemMetadata()
	if ok {
		return sm.(ObjectSystemMetadata)
	}
	return ObjectSystemMetadata{}
}

// setObjectSystemMetadata will set ObjectSystemMetadata into Object.
//
// - This function should only be called once, please make sure all data has been written before set.
func setObjectSystemMetadata(o *Object, sm ObjectSystemMetadata) {
	o.SetSystemMetadata(sm)
}

// StorageSystemMetadata stores system metadata for storage meta.
type StorageSystemMetadata struct {
}

// GetStorageSystemMetadata will get SystemMetadata from StorageMeta.
//
// - The returning StorageSystemMetadata is read only and should not be modified.
func GetStorageSystemMetadata(s *StorageMeta) StorageSystemMetadata {
	sm, ok := s.GetSystemMetadata()
	if ok {
		return sm.(StorageSystemMetadata)
	}
	return StorageSystemMetadata{}
}

// setStorageSystemMetadata will set SystemMetadata into StorageMeta.
//
// - This function should only be called once, please make sure all data has been written before set.
func setStorageSystemMetadata(s *StorageMeta, sm StorageSystemMetadata) {
	s.SetSystemMetadata(sm)
}

// WithDefaultServicePairs will apply default_service_pairs value to Options.
//
// DefaultServicePairs set default pairs for service actions
func WithDefaultServicePairs(v DefaultServicePairs) Pair {
	return Pair{
		Key:   "default_service_pairs",
		Value: v,
	}
}

// WithDefaultStoragePairs will apply default_storage_pairs value to Options.
//
// DefaultStoragePairs set default pairs for storager actions
func WithDefaultStoragePairs(v DefaultStoragePairs) Pair {
	return Pair{
		Key:   "default_storage_pairs",
		Value: v,
	}
}

// WithServiceFeatures will apply service_features value to Options.
//
// ServiceFeatures set service features
func WithServiceFeatures(v ServiceFeatures) Pair {
	return Pair{
		Key:   "service_features",
		Value: v,
	}
}

// WithStorageClass will apply storage_class value to Options.
//
// StorageClass
func WithStorageClass(v string) Pair {
	return Pair{
		Key:   "storage_class",
		Value: v,
	}
}

// WithStorageFeatures will apply storage_features value to Options.
//
// StorageFeatures set storage features
func WithStorageFeatures(v StorageFeatures) Pair {
	return Pair{
		Key:   "storage_features",
		Value: v,
	}
}

var pairMap = map[string]string{
	"content_md5":           "string",
	"content_type":          "string",
	"context":               "context.Context",
	"continuation_token":    "string",
	"credential":            "string",
	"default_service_pairs": "DefaultServicePairs",
	"default_storage_pairs": "DefaultStoragePairs",
	"endpoint":              "string",
	"expire":                "time.Duration",
	"http_client_options":   "*httpclient.Options",
	"interceptor":           "Interceptor",
	"io_callback":           "func([]byte)",
	"list_mode":             "ListMode",
	"location":              "string",
	"multipart_id":          "string",
	"name":                  "string",
	"object_mode":           "ObjectMode",
	"offset":                "int64",
	"service_features":      "ServiceFeatures",
	"size":                  "int64",
	"storage_class":         "string",
	"storage_features":      "StorageFeatures",
	"work_dir":              "string",
}
var (
	_ Servicer = &Service{}
)

type ServiceFeatures struct {
}

// pairServiceNew is the parsed struct
type pairServiceNew struct {
	pairs []Pair

	// Required pairs
	HasCredential bool
	Credential    string
	HasEndpoint   bool
	Endpoint      string
	// Optional pairs
	HasDefaultServicePairs bool
	DefaultServicePairs    DefaultServicePairs
	HasServiceFeatures     bool
	ServiceFeatures        ServiceFeatures
}

// parsePairServiceNew will parse Pair slice into *pairServiceNew
func parsePairServiceNew(opts []Pair) (pairServiceNew, error) {
	result := pairServiceNew{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		case "credential":
			if result.HasCredential {
				continue
			}
			result.HasCredential = true
			result.Credential = v.Value.(string)
		case "endpoint":
			if result.HasEndpoint {
				continue
			}
			result.HasEndpoint = true
			result.Endpoint = v.Value.(string)
		// Optional pairs
		case "default_service_pairs":
			if result.HasDefaultServicePairs {
				continue
			}
			result.HasDefaultServicePairs = true
			result.DefaultServicePairs = v.Value.(DefaultServicePairs)
		case "service_features":
			if result.HasServiceFeatures {
				continue
			}
			result.HasServiceFeatures = true
			result.ServiceFeatures = v.Value.(ServiceFeatures)
		}
	}
	if !result.HasCredential {
		return pairServiceNew{}, services.PairRequiredError{Keys: []string{"credential"}}
	}
	if !result.HasEndpoint {
		return pairServiceNew{}, services.PairRequiredError{Keys: []string{"endpoint"}}
	}

	return result, nil
}

// DefaultServicePairs is default pairs for specific action
type DefaultServicePairs struct {
	Create []Pair
	Delete []Pair
	Get    []Pair
	List   []Pair
}

// pairServiceCreate is the parsed struct
type pairServiceCreate struct {
	pairs []Pair
}

// parsePairServiceCreate will parse Pair slice into *pairServiceCreate
func (s *Service) parsePairServiceCreate(opts []Pair) (pairServiceCreate, error) {
	result := pairServiceCreate{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairServiceCreate{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairServiceDelete is the parsed struct
type pairServiceDelete struct {
	pairs []Pair
}

// parsePairServiceDelete will parse Pair slice into *pairServiceDelete
func (s *Service) parsePairServiceDelete(opts []Pair) (pairServiceDelete, error) {
	result := pairServiceDelete{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairServiceDelete{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairServiceGet is the parsed struct
type pairServiceGet struct {
	pairs []Pair
}

// parsePairServiceGet will parse Pair slice into *pairServiceGet
func (s *Service) parsePairServiceGet(opts []Pair) (pairServiceGet, error) {
	result := pairServiceGet{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairServiceGet{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairServiceList is the parsed struct
type pairServiceList struct {
	pairs []Pair
}

// parsePairServiceList will parse Pair slice into *pairServiceList
func (s *Service) parsePairServiceList(opts []Pair) (pairServiceList, error) {
	result := pairServiceList{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairServiceList{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// Create will create a new storager instance.
//
// This function will create a context by default.
func (s *Service) Create(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.CreateWithContext(ctx, name, pairs...)
}

// CreateWithContext will create a new storager instance.
func (s *Service) CreateWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	defer func() {
		err = s.formatError("create", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairServiceCreate

	opt, err = s.parsePairServiceCreate(pairs)
	if err != nil {
		return
	}

	return s.create(ctx, name, opt)
}

// Delete will delete a storager instance.
//
// This function will create a context by default.
func (s *Service) Delete(name string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, name, pairs...)
}

// DeleteWithContext will delete a storager instance.
func (s *Service) DeleteWithContext(ctx context.Context, name string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("delete", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairServiceDelete

	opt, err = s.parsePairServiceDelete(pairs)
	if err != nil {
		return
	}

	return s.delete(ctx, name, opt)
}

// Get will get a valid storager instance for service.
//
// This function will create a context by default.
func (s *Service) Get(name string, pairs ...Pair) (store Storager, err error) {
	ctx := context.Background()
	return s.GetWithContext(ctx, name, pairs...)
}

// GetWithContext will get a valid storager instance for service.
func (s *Service) GetWithContext(ctx context.Context, name string, pairs ...Pair) (store Storager, err error) {
	defer func() {
		err = s.formatError("get", err, name)
	}()

	pairs = append(pairs, s.defaultPairs.Get...)
	var opt pairServiceGet

	opt, err = s.parsePairServiceGet(pairs)
	if err != nil {
		return
	}

	return s.get(ctx, name, opt)
}

// List will list all storager instances under this service.
//
// This function will create a context by default.
func (s *Service) List(pairs ...Pair) (sti *StoragerIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, pairs...)
}

// ListWithContext will list all storager instances under this service.
func (s *Service) ListWithContext(ctx context.Context, pairs ...Pair) (sti *StoragerIterator, err error) {
	defer func() {

		err = s.formatError("list", err, "")
	}()

	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairServiceList

	opt, err = s.parsePairServiceList(pairs)
	if err != nil {
		return
	}

	return s.list(ctx, opt)
}

var (
	_ Copier   = &Storage{}
	_ Reacher  = &Storage{}
	_ Storager = &Storage{}
)

type StorageFeatures struct {
	// VirtualDir virtual_dir feature is designed for a service that doesn't have native dir support but wants to provide simulated operations.
	//
	// - If this feature is disabled (the default behavior), the service will behave like it doesn't have any dir support.
	// - If this feature is enabled, the service will support simulated dir behavior in create_dir, create, list, delete, and so on.
	//
	// This feature was introduced in GSP-109.
	VirtualDir bool
}

// pairStorageNew is the parsed struct
type pairStorageNew struct {
	pairs []Pair

	// Required pairs
	HasName bool
	Name    string
	// Optional pairs
	HasDefaultStoragePairs bool
	DefaultStoragePairs    DefaultStoragePairs
	HasStorageFeatures     bool
	StorageFeatures        StorageFeatures
	HasWorkDir             bool
	WorkDir                string
}

// parsePairStorageNew will parse Pair slice into *pairStorageNew
func parsePairStorageNew(opts []Pair) (pairStorageNew, error) {
	result := pairStorageNew{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		// Required pairs
		case "name":
			if result.HasName {
				continue
			}
			result.HasName = true
			result.Name = v.Value.(string)
		// Optional pairs
		case "default_storage_pairs":
			if result.HasDefaultStoragePairs {
				continue
			}
			result.HasDefaultStoragePairs = true
			result.DefaultStoragePairs = v.Value.(DefaultStoragePairs)
		case "storage_features":
			if result.HasStorageFeatures {
				continue
			}
			result.HasStorageFeatures = true
			result.StorageFeatures = v.Value.(StorageFeatures)
		case "work_dir":
			if result.HasWorkDir {
				continue
			}
			result.HasWorkDir = true
			result.WorkDir = v.Value.(string)
		}
	}
	if !result.HasName {
		return pairStorageNew{}, services.PairRequiredError{Keys: []string{"name"}}
	}

	return result, nil
}

// DefaultStoragePairs is default pairs for specific action
type DefaultStoragePairs struct {
	Copy     []Pair
	Create   []Pair
	Delete   []Pair
	List     []Pair
	Metadata []Pair
	Reach    []Pair
	Read     []Pair
	Stat     []Pair
	Write    []Pair
}

// pairStorageCopy is the parsed struct
type pairStorageCopy struct {
	pairs []Pair
}

// parsePairStorageCopy will parse Pair slice into *pairStorageCopy
func (s *Storage) parsePairStorageCopy(opts []Pair) (pairStorageCopy, error) {
	result := pairStorageCopy{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageCopy{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageCreate is the parsed struct
type pairStorageCreate struct {
	pairs         []Pair
	HasObjectMode bool
	ObjectMode    ObjectMode
}

// parsePairStorageCreate will parse Pair slice into *pairStorageCreate
func (s *Storage) parsePairStorageCreate(opts []Pair) (pairStorageCreate, error) {
	result := pairStorageCreate{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
			continue
		default:
			return pairStorageCreate{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageDelete is the parsed struct
type pairStorageDelete struct {
	pairs         []Pair
	HasObjectMode bool
	ObjectMode    ObjectMode
}

// parsePairStorageDelete will parse Pair slice into *pairStorageDelete
func (s *Storage) parsePairStorageDelete(opts []Pair) (pairStorageDelete, error) {
	result := pairStorageDelete{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
			continue
		default:
			return pairStorageDelete{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageList is the parsed struct
type pairStorageList struct {
	pairs       []Pair
	HasListMode bool
	ListMode    ListMode
}

// parsePairStorageList will parse Pair slice into *pairStorageList
func (s *Storage) parsePairStorageList(opts []Pair) (pairStorageList, error) {
	result := pairStorageList{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "list_mode":
			if result.HasListMode {
				continue
			}
			result.HasListMode = true
			result.ListMode = v.Value.(ListMode)
			continue
		default:
			return pairStorageList{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageMetadata is the parsed struct
type pairStorageMetadata struct {
	pairs []Pair
}

// parsePairStorageMetadata will parse Pair slice into *pairStorageMetadata
func (s *Storage) parsePairStorageMetadata(opts []Pair) (pairStorageMetadata, error) {
	result := pairStorageMetadata{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		default:
			return pairStorageMetadata{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageReach is the parsed struct
type pairStorageReach struct {
	pairs     []Pair
	HasExpire bool
	Expire    time.Duration
}

// parsePairStorageReach will parse Pair slice into *pairStorageReach
func (s *Storage) parsePairStorageReach(opts []Pair) (pairStorageReach, error) {
	result := pairStorageReach{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "expire":
			if result.HasExpire {
				continue
			}
			result.HasExpire = true
			result.Expire = v.Value.(time.Duration)
			continue
		default:
			return pairStorageReach{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageRead is the parsed struct
type pairStorageRead struct {
	pairs         []Pair
	HasIoCallback bool
	IoCallback    func([]byte)
	HasOffset     bool
	Offset        int64
	HasSize       bool
	Size          int64
}

// parsePairStorageRead will parse Pair slice into *pairStorageRead
func (s *Storage) parsePairStorageRead(opts []Pair) (pairStorageRead, error) {
	result := pairStorageRead{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
			continue
		case "offset":
			if result.HasOffset {
				continue
			}
			result.HasOffset = true
			result.Offset = v.Value.(int64)
			continue
		case "size":
			if result.HasSize {
				continue
			}
			result.HasSize = true
			result.Size = v.Value.(int64)
			continue
		default:
			return pairStorageRead{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageStat is the parsed struct
type pairStorageStat struct {
	pairs         []Pair
	HasObjectMode bool
	ObjectMode    ObjectMode
}

// parsePairStorageStat will parse Pair slice into *pairStorageStat
func (s *Storage) parsePairStorageStat(opts []Pair) (pairStorageStat, error) {
	result := pairStorageStat{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "object_mode":
			if result.HasObjectMode {
				continue
			}
			result.HasObjectMode = true
			result.ObjectMode = v.Value.(ObjectMode)
			continue
		default:
			return pairStorageStat{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// pairStorageWrite is the parsed struct
type pairStorageWrite struct {
	pairs           []Pair
	HasContentMd5   bool
	ContentMd5      string
	HasContentType  bool
	ContentType     string
	HasIoCallback   bool
	IoCallback      func([]byte)
	HasStorageClass bool
	StorageClass    string
}

// parsePairStorageWrite will parse Pair slice into *pairStorageWrite
func (s *Storage) parsePairStorageWrite(opts []Pair) (pairStorageWrite, error) {
	result := pairStorageWrite{
		pairs: opts,
	}

	for _, v := range opts {
		switch v.Key {
		case "content_md5":
			if result.HasContentMd5 {
				continue
			}
			result.HasContentMd5 = true
			result.ContentMd5 = v.Value.(string)
			continue
		case "content_type":
			if result.HasContentType {
				continue
			}
			result.HasContentType = true
			result.ContentType = v.Value.(string)
			continue
		case "io_callback":
			if result.HasIoCallback {
				continue
			}
			result.HasIoCallback = true
			result.IoCallback = v.Value.(func([]byte))
			continue
		case "storage_class":
			if result.HasStorageClass {
				continue
			}
			result.HasStorageClass = true
			result.StorageClass = v.Value.(string)
			continue
		default:
			return pairStorageWrite{}, services.PairUnsupportedError{Pair: v}
		}
	}

	// Check required pairs.

	return result, nil
}

// Copy will copy an Object or multiple object in the service.
//
// ## Behavior
//
// - Copy only copy one and only one object.
//   - Service DON'T NEED to support copy a non-empty directory or copy files recursively.
//   - User NEED to implement copy a non-empty directory and copy recursively by themself.
//   - Copy a file to a directory SHOULD return `ErrObjectModeInvalid`.
// - Copy SHOULD NOT return an error as dst object exists.
//   - Service that has native support for `overwrite` doesn't NEED to check the dst object exists or not.
//   - Service that doesn't have native support for `overwrite` SHOULD check and delete the dst object if exists.
// - A successful copy opration should be complete, which means the dst object's content and metadata should be the same as src object.
//
// This function will create a context by default.
func (s *Storage) Copy(src string, dst string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.CopyWithContext(ctx, src, dst, pairs...)
}

// CopyWithContext will copy an Object or multiple object in the service.
//
// ## Behavior
//
// - Copy only copy one and only one object.
//   - Service DON'T NEED to support copy a non-empty directory or copy files recursively.
//   - User NEED to implement copy a non-empty directory and copy recursively by themself.
//   - Copy a file to a directory SHOULD return `ErrObjectModeInvalid`.
// - Copy SHOULD NOT return an error as dst object exists.
//   - Service that has native support for `overwrite` doesn't NEED to check the dst object exists or not.
//   - Service that doesn't have native support for `overwrite` SHOULD check and delete the dst object if exists.
// - A successful copy opration should be complete, which means the dst object's content and metadata should be the same as src object.
func (s *Storage) CopyWithContext(ctx context.Context, src string, dst string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("copy", err, src, dst)
	}()

	pairs = append(pairs, s.defaultPairs.Copy...)
	var opt pairStorageCopy

	opt, err = s.parsePairStorageCopy(pairs)
	if err != nil {
		return
	}

	return s.copy(ctx, src, dst, opt)
}

// Create will create a new object without any api call.
//
// ## Behavior
//
// - Create SHOULD NOT send any API call.
// - Create SHOULD accept ObjectMode pair as object mode.
//
// This function will create a context by default.
func (s *Storage) Create(path string, pairs ...Pair) (o *Object) {
	pairs = append(pairs, s.defaultPairs.Create...)
	var opt pairStorageCreate

	// Ignore error while handling local funtions.
	opt, _ = s.parsePairStorageCreate(pairs)

	return s.create(path, opt)
}

// Delete will delete an object from service.
//
// ## Behavior
//
// - Delete only delete one and only one object.
//   - Service DON'T NEED to support remove all.
//   - User NEED to implement remove_all by themself.
// - Delete is idempotent.
//   - Successful delete always return nil error.
//   - Delete SHOULD never return `ObjectNotExist`
//   - Delete DON'T NEED to check the object exist or not.
//
// This function will create a context by default.
func (s *Storage) Delete(path string, pairs ...Pair) (err error) {
	ctx := context.Background()
	return s.DeleteWithContext(ctx, path, pairs...)
}

// DeleteWithContext will delete an object from service.
//
// ## Behavior
//
// - Delete only delete one and only one object.
//   - Service DON'T NEED to support remove all.
//   - User NEED to implement remove_all by themself.
// - Delete is idempotent.
//   - Successful delete always return nil error.
//   - Delete SHOULD never return `ObjectNotExist`
//   - Delete DON'T NEED to check the object exist or not.
func (s *Storage) DeleteWithContext(ctx context.Context, path string, pairs ...Pair) (err error) {
	defer func() {
		err = s.formatError("delete", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Delete...)
	var opt pairStorageDelete

	opt, err = s.parsePairStorageDelete(pairs)
	if err != nil {
		return
	}

	return s.delete(ctx, path, opt)
}

// List will return list a specific path.
//
// ## Behavior
//
// - Service SHOULD support default `ListMode`.
// - Service SHOULD implement `ListModeDir` without the check for `VirtualDir`.
// - Service DON'T NEED to `Stat` while in `List`.
//
// This function will create a context by default.
func (s *Storage) List(path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	ctx := context.Background()
	return s.ListWithContext(ctx, path, pairs...)
}

// ListWithContext will return list a specific path.
//
// ## Behavior
//
// - Service SHOULD support default `ListMode`.
// - Service SHOULD implement `ListModeDir` without the check for `VirtualDir`.
// - Service DON'T NEED to `Stat` while in `List`.
func (s *Storage) ListWithContext(ctx context.Context, path string, pairs ...Pair) (oi *ObjectIterator, err error) {
	defer func() {
		err = s.formatError("list", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.List...)
	var opt pairStorageList

	opt, err = s.parsePairStorageList(pairs)
	if err != nil {
		return
	}

	return s.list(ctx, path, opt)
}

// Metadata will return current storager metadata.
//
// This function will create a context by default.
func (s *Storage) Metadata(pairs ...Pair) (meta *StorageMeta) {
	pairs = append(pairs, s.defaultPairs.Metadata...)
	var opt pairStorageMetadata

	// Ignore error while handling local funtions.
	opt, _ = s.parsePairStorageMetadata(pairs)

	return s.metadata(opt)
}

// Reach will provide a way, which can reach the object.
//
// This function will create a context by default.
func (s *Storage) Reach(path string, pairs ...Pair) (url string, err error) {
	ctx := context.Background()
	return s.ReachWithContext(ctx, path, pairs...)
}

// ReachWithContext will provide a way, which can reach the object.
func (s *Storage) ReachWithContext(ctx context.Context, path string, pairs ...Pair) (url string, err error) {
	defer func() {
		err = s.formatError("reach", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Reach...)
	var opt pairStorageReach

	opt, err = s.parsePairStorageReach(pairs)
	if err != nil {
		return
	}

	return s.reach(ctx, path, opt)
}

// Read will read the file's data.
//
// This function will create a context by default.
func (s *Storage) Read(path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.ReadWithContext(ctx, path, w, pairs...)
}

// ReadWithContext will read the file's data.
func (s *Storage) ReadWithContext(ctx context.Context, path string, w io.Writer, pairs ...Pair) (n int64, err error) {
	defer func() {
		err = s.formatError("read", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Read...)
	var opt pairStorageRead

	opt, err = s.parsePairStorageRead(pairs)
	if err != nil {
		return
	}

	return s.read(ctx, path, w, opt)
}

// Stat will stat a path to get info of an object.
//
// ## Behavior
//
// - Stat SHOULD accept ObjectMode pair as hints.
//   - Service COULD have different implementations for different object mode.
//   - Service SHOULD check if returning ObjectMode is match
//
// This function will create a context by default.
func (s *Storage) Stat(path string, pairs ...Pair) (o *Object, err error) {
	ctx := context.Background()
	return s.StatWithContext(ctx, path, pairs...)
}

// StatWithContext will stat a path to get info of an object.
//
// ## Behavior
//
// - Stat SHOULD accept ObjectMode pair as hints.
//   - Service COULD have different implementations for different object mode.
//   - Service SHOULD check if returning ObjectMode is match
func (s *Storage) StatWithContext(ctx context.Context, path string, pairs ...Pair) (o *Object, err error) {
	defer func() {
		err = s.formatError("stat", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Stat...)
	var opt pairStorageStat

	opt, err = s.parsePairStorageStat(pairs)
	if err != nil {
		return
	}

	return s.stat(ctx, path, opt)
}

// Write will write data into a file.
//
// ## Behavior
//
// - Write SHOULD NOT return an error as the object exist.
//   - Service that has native support for `overwrite` doesn't NEED to check the object exists or not.
//   - Service that doesn't have native support for `overwrite` SHOULD check and delete the object if exists.
// - A successful write operation SHOULD be complete, which means the object's content and metadata should be the same as specified in write request.
//
// This function will create a context by default.
func (s *Storage) Write(path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	ctx := context.Background()
	return s.WriteWithContext(ctx, path, r, size, pairs...)
}

// WriteWithContext will write data into a file.
//
// ## Behavior
//
// - Write SHOULD NOT return an error as the object exist.
//   - Service that has native support for `overwrite` doesn't NEED to check the object exists or not.
//   - Service that doesn't have native support for `overwrite` SHOULD check and delete the object if exists.
// - A successful write operation SHOULD be complete, which means the object's content and metadata should be the same as specified in write request.
func (s *Storage) WriteWithContext(ctx context.Context, path string, r io.Reader, size int64, pairs ...Pair) (n int64, err error) {
	defer func() {
		err = s.formatError("write", err, path)
	}()

	pairs = append(pairs, s.defaultPairs.Write...)
	var opt pairStorageWrite

	opt, err = s.parsePairStorageWrite(pairs)
	if err != nil {
		return
	}

	return s.write(ctx, path, r, size, opt)
}

func init() {
	services.RegisterServicer(Type, NewServicer)
	services.RegisterStorager(Type, NewStorager)
	services.RegisterSchema(Type, pairMap)
}
