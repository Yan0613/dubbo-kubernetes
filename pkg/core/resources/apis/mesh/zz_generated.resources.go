// Generated by tools/resource-gen.
// Run "make generate" to update this file.

// nolint:whitespace
package mesh

import (
	"fmt"
)

import (
	mesh_proto "github.com/apache/dubbo-kubernetes/api/mesh/v1alpha1"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/model"
	"github.com/apache/dubbo-kubernetes/pkg/core/resources/registry"
)

const (
	AffinityRouteType model.ResourceType = "AffinityRoute"
)

var _ model.Resource = &AffinityRouteResource{}

type AffinityRouteResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.AffinityRoute
}

func NewAffinityRouteResource() *AffinityRouteResource {
	return &AffinityRouteResource{
		Spec: &mesh_proto.AffinityRoute{},
	}
}

func (t *AffinityRouteResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *AffinityRouteResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *AffinityRouteResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *AffinityRouteResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.AffinityRoute)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.AffinityRoute{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *AffinityRouteResource) Descriptor() model.ResourceTypeDescriptor {
	return AffinityRouteResourceTypeDescriptor
}

var _ model.ResourceList = &AffinityRouteResourceList{}

type AffinityRouteResourceList struct {
	Items      []*AffinityRouteResource
	Pagination model.Pagination
}

func (l *AffinityRouteResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *AffinityRouteResourceList) GetItemType() model.ResourceType {
	return AffinityRouteType
}

func (l *AffinityRouteResourceList) NewItem() model.Resource {
	return NewAffinityRouteResource()
}

func (l *AffinityRouteResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*AffinityRouteResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*AffinityRouteResource)(nil), r)
	}
}

func (l *AffinityRouteResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *AffinityRouteResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var AffinityRouteResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                AffinityRouteType,
	Resource:            NewAffinityRouteResource(),
	ResourceList:        &AffinityRouteResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.GlobalToAllZonesFlag,
	WsPath:              "affinityroutes",
	DubboctlArg:         "affinityroute",
	DubboctlListArg:     "affinityroutes",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Affinity Route",
	PluralDisplayName:   "Affinity Routes",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(AffinityRouteResourceTypeDescriptor)
}

const (
	ConditionRouteType model.ResourceType = "ConditionRoute"
)

var _ model.Resource = &ConditionRouteResource{}

type ConditionRouteResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.ConditionRoute
}

func NewConditionRouteResource() *ConditionRouteResource {
	return &ConditionRouteResource{
		Spec: &mesh_proto.ConditionRoute{},
	}
}

func (t *ConditionRouteResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *ConditionRouteResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *ConditionRouteResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *ConditionRouteResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.ConditionRoute)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.ConditionRoute{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *ConditionRouteResource) Descriptor() model.ResourceTypeDescriptor {
	return ConditionRouteResourceTypeDescriptor
}

var _ model.ResourceList = &ConditionRouteResourceList{}

type ConditionRouteResourceList struct {
	Items      []*ConditionRouteResource
	Pagination model.Pagination
}

func (l *ConditionRouteResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *ConditionRouteResourceList) GetItemType() model.ResourceType {
	return ConditionRouteType
}

func (l *ConditionRouteResourceList) NewItem() model.Resource {
	return NewConditionRouteResource()
}

func (l *ConditionRouteResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*ConditionRouteResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*ConditionRouteResource)(nil), r)
	}
}

func (l *ConditionRouteResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *ConditionRouteResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var ConditionRouteResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                ConditionRouteType,
	Resource:            NewConditionRouteResource(),
	ResourceList:        &ConditionRouteResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.GlobalToAllZonesFlag,
	WsPath:              "conditionroutes",
	DubboctlArg:         "conditionroute",
	DubboctlListArg:     "conditionroutes",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Condition Route",
	PluralDisplayName:   "Condition Routes",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(ConditionRouteResourceTypeDescriptor)
}

const (
	DataplaneType model.ResourceType = "Dataplane"
)

var _ model.Resource = &DataplaneResource{}

type DataplaneResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.Dataplane
}

func NewDataplaneResource() *DataplaneResource {
	return &DataplaneResource{
		Spec: &mesh_proto.Dataplane{},
	}
}

func (t *DataplaneResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *DataplaneResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *DataplaneResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *DataplaneResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.Dataplane)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.Dataplane{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *DataplaneResource) Descriptor() model.ResourceTypeDescriptor {
	return DataplaneResourceTypeDescriptor
}

var _ model.ResourceList = &DataplaneResourceList{}

type DataplaneResourceList struct {
	Items      []*DataplaneResource
	Pagination model.Pagination
}

func (l *DataplaneResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *DataplaneResourceList) GetItemType() model.ResourceType {
	return DataplaneType
}

func (l *DataplaneResourceList) NewItem() model.Resource {
	return NewDataplaneResource()
}

func (l *DataplaneResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*DataplaneResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*DataplaneResource)(nil), r)
	}
}

func (l *DataplaneResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *DataplaneResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var DataplaneResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                DataplaneType,
	Resource:            NewDataplaneResource(),
	ResourceList:        &DataplaneResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.ZoneToGlobalFlag,
	WsPath:              "dataplanes",
	DubboctlArg:         "dataplane",
	DubboctlListArg:     "dataplanes",
	AllowToInspect:      true,
	IsPolicy:            false,
	SingularDisplayName: "Dataplane",
	PluralDisplayName:   "Dataplanes",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(DataplaneResourceTypeDescriptor)
}

const (
	DataplaneInsightType model.ResourceType = "DataplaneInsight"
)

var _ model.Resource = &DataplaneInsightResource{}

type DataplaneInsightResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.DataplaneInsight
}

func NewDataplaneInsightResource() *DataplaneInsightResource {
	return &DataplaneInsightResource{
		Spec: &mesh_proto.DataplaneInsight{},
	}
}

func (t *DataplaneInsightResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *DataplaneInsightResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *DataplaneInsightResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *DataplaneInsightResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.DataplaneInsight)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.DataplaneInsight{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *DataplaneInsightResource) Descriptor() model.ResourceTypeDescriptor {
	return DataplaneInsightResourceTypeDescriptor
}

var _ model.ResourceList = &DataplaneInsightResourceList{}

type DataplaneInsightResourceList struct {
	Items      []*DataplaneInsightResource
	Pagination model.Pagination
}

func (l *DataplaneInsightResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *DataplaneInsightResourceList) GetItemType() model.ResourceType {
	return DataplaneInsightType
}

func (l *DataplaneInsightResourceList) NewItem() model.Resource {
	return NewDataplaneInsightResource()
}

func (l *DataplaneInsightResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*DataplaneInsightResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*DataplaneInsightResource)(nil), r)
	}
}

func (l *DataplaneInsightResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *DataplaneInsightResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var DataplaneInsightResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                DataplaneInsightType,
	Resource:            NewDataplaneInsightResource(),
	ResourceList:        &DataplaneInsightResourceList{},
	ReadOnly:            true,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.ZoneToGlobalFlag,
	WsPath:              "dataplane-insights",
	DubboctlArg:         "",
	DubboctlListArg:     "",
	AllowToInspect:      false,
	IsPolicy:            false,
	SingularDisplayName: "Dataplane Insight",
	PluralDisplayName:   "Dataplane Insights",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(DataplaneInsightResourceTypeDescriptor)
}

const (
	DynamicConfigType model.ResourceType = "DynamicConfig"
)

var _ model.Resource = &DynamicConfigResource{}

type DynamicConfigResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.DynamicConfig
}

func NewDynamicConfigResource() *DynamicConfigResource {
	return &DynamicConfigResource{
		Spec: &mesh_proto.DynamicConfig{},
	}
}

func (t *DynamicConfigResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *DynamicConfigResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *DynamicConfigResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *DynamicConfigResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.DynamicConfig)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.DynamicConfig{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *DynamicConfigResource) Descriptor() model.ResourceTypeDescriptor {
	return DynamicConfigResourceTypeDescriptor
}

var _ model.ResourceList = &DynamicConfigResourceList{}

type DynamicConfigResourceList struct {
	Items      []*DynamicConfigResource
	Pagination model.Pagination
}

func (l *DynamicConfigResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *DynamicConfigResourceList) GetItemType() model.ResourceType {
	return DynamicConfigType
}

func (l *DynamicConfigResourceList) NewItem() model.Resource {
	return NewDynamicConfigResource()
}

func (l *DynamicConfigResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*DynamicConfigResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*DynamicConfigResource)(nil), r)
	}
}

func (l *DynamicConfigResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *DynamicConfigResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var DynamicConfigResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                DynamicConfigType,
	Resource:            NewDynamicConfigResource(),
	ResourceList:        &DynamicConfigResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.GlobalToAllZonesFlag,
	WsPath:              "dynamicconfigs",
	DubboctlArg:         "dynamicconfig",
	DubboctlListArg:     "dynamicconfigs",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Dynamic Config",
	PluralDisplayName:   "Dynamic Configs",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(DynamicConfigResourceTypeDescriptor)
}

const (
	MappingType model.ResourceType = "Mapping"
)

var _ model.Resource = &MappingResource{}

type MappingResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.Mapping
}

func NewMappingResource() *MappingResource {
	return &MappingResource{
		Spec: &mesh_proto.Mapping{},
	}
}

func (t *MappingResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MappingResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MappingResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MappingResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.Mapping)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.Mapping{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MappingResource) Descriptor() model.ResourceTypeDescriptor {
	return MappingResourceTypeDescriptor
}

var _ model.ResourceList = &MappingResourceList{}

type MappingResourceList struct {
	Items      []*MappingResource
	Pagination model.Pagination
}

func (l *MappingResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MappingResourceList) GetItemType() model.ResourceType {
	return MappingType
}

func (l *MappingResourceList) NewItem() model.Resource {
	return NewMappingResource()
}

func (l *MappingResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MappingResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MappingResource)(nil), r)
	}
}

func (l *MappingResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MappingResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MappingResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                MappingType,
	Resource:            NewMappingResource(),
	ResourceList:        &MappingResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.ZoneToGlobalFlag | model.GlobalToAllButOriginalZoneFlag,
	WsPath:              "mappings",
	DubboctlArg:         "mapping",
	DubboctlListArg:     "mappings",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Mapping",
	PluralDisplayName:   "Mappings",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(MappingResourceTypeDescriptor)
}

const (
	MeshType model.ResourceType = "Mesh"
)

var _ model.Resource = &MeshResource{}

type MeshResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.Mesh
}

func NewMeshResource() *MeshResource {
	return &MeshResource{
		Spec: &mesh_proto.Mesh{},
	}
}

func (t *MeshResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.Mesh)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.Mesh{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshResourceTypeDescriptor
}

var _ model.ResourceList = &MeshResourceList{}

type MeshResourceList struct {
	Items      []*MeshResource
	Pagination model.Pagination
}

func (l *MeshResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshResourceList) GetItemType() model.ResourceType {
	return MeshType
}

func (l *MeshResourceList) NewItem() model.Resource {
	return NewMeshResource()
}

func (l *MeshResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshResource)(nil), r)
	}
}

func (l *MeshResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MeshResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MeshResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                MeshType,
	Resource:            NewMeshResource(),
	ResourceList:        &MeshResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	DDSFlags:            model.GlobalToAllZonesFlag,
	WsPath:              "meshes",
	DubboctlArg:         "mesh",
	DubboctlListArg:     "meshes",
	AllowToInspect:      true,
	IsPolicy:            false,
	SingularDisplayName: "Mesh",
	PluralDisplayName:   "Meshes",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(MeshResourceTypeDescriptor)
}

const (
	MeshInsightType model.ResourceType = "MeshInsight"
)

var _ model.Resource = &MeshInsightResource{}

type MeshInsightResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.MeshInsight
}

func NewMeshInsightResource() *MeshInsightResource {
	return &MeshInsightResource{
		Spec: &mesh_proto.MeshInsight{},
	}
}

func (t *MeshInsightResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MeshInsightResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MeshInsightResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MeshInsightResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.MeshInsight)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.MeshInsight{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MeshInsightResource) Descriptor() model.ResourceTypeDescriptor {
	return MeshInsightResourceTypeDescriptor
}

var _ model.ResourceList = &MeshInsightResourceList{}

type MeshInsightResourceList struct {
	Items      []*MeshInsightResource
	Pagination model.Pagination
}

func (l *MeshInsightResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MeshInsightResourceList) GetItemType() model.ResourceType {
	return MeshInsightType
}

func (l *MeshInsightResourceList) NewItem() model.Resource {
	return NewMeshInsightResource()
}

func (l *MeshInsightResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MeshInsightResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MeshInsightResource)(nil), r)
	}
}

func (l *MeshInsightResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MeshInsightResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MeshInsightResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                MeshInsightType,
	Resource:            NewMeshInsightResource(),
	ResourceList:        &MeshInsightResourceList{},
	ReadOnly:            true,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	WsPath:              "mesh-insights",
	DubboctlArg:         "",
	DubboctlListArg:     "",
	AllowToInspect:      false,
	IsPolicy:            false,
	SingularDisplayName: "Mesh Insight",
	PluralDisplayName:   "Mesh Insights",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(MeshInsightResourceTypeDescriptor)
}

const (
	MetaDataType model.ResourceType = "MetaData"
)

var _ model.Resource = &MetaDataResource{}

type MetaDataResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.MetaData
}

func NewMetaDataResource() *MetaDataResource {
	return &MetaDataResource{
		Spec: &mesh_proto.MetaData{},
	}
}

func (t *MetaDataResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *MetaDataResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *MetaDataResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *MetaDataResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.MetaData)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.MetaData{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *MetaDataResource) Descriptor() model.ResourceTypeDescriptor {
	return MetaDataResourceTypeDescriptor
}

var _ model.ResourceList = &MetaDataResourceList{}

type MetaDataResourceList struct {
	Items      []*MetaDataResource
	Pagination model.Pagination
}

func (l *MetaDataResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *MetaDataResourceList) GetItemType() model.ResourceType {
	return MetaDataType
}

func (l *MetaDataResourceList) NewItem() model.Resource {
	return NewMetaDataResource()
}

func (l *MetaDataResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*MetaDataResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*MetaDataResource)(nil), r)
	}
}

func (l *MetaDataResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *MetaDataResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var MetaDataResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                MetaDataType,
	Resource:            NewMetaDataResource(),
	ResourceList:        &MetaDataResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.ZoneToGlobalFlag,
	WsPath:              "metadatas",
	DubboctlArg:         "metadata",
	DubboctlListArg:     "metadatas",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Meta Data",
	PluralDisplayName:   "Meta Datas",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(MetaDataResourceTypeDescriptor)
}

const (
	TagRouteType model.ResourceType = "TagRoute"
)

var _ model.Resource = &TagRouteResource{}

type TagRouteResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.TagRoute
}

func NewTagRouteResource() *TagRouteResource {
	return &TagRouteResource{
		Spec: &mesh_proto.TagRoute{},
	}
}

func (t *TagRouteResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *TagRouteResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *TagRouteResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *TagRouteResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.TagRoute)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.TagRoute{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *TagRouteResource) Descriptor() model.ResourceTypeDescriptor {
	return TagRouteResourceTypeDescriptor
}

var _ model.ResourceList = &TagRouteResourceList{}

type TagRouteResourceList struct {
	Items      []*TagRouteResource
	Pagination model.Pagination
}

func (l *TagRouteResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *TagRouteResourceList) GetItemType() model.ResourceType {
	return TagRouteType
}

func (l *TagRouteResourceList) NewItem() model.Resource {
	return NewTagRouteResource()
}

func (l *TagRouteResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*TagRouteResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*TagRouteResource)(nil), r)
	}
}

func (l *TagRouteResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *TagRouteResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var TagRouteResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                TagRouteType,
	Resource:            NewTagRouteResource(),
	ResourceList:        &TagRouteResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeMesh,
	DDSFlags:            model.GlobalToAllZonesFlag,
	WsPath:              "tagroutes",
	DubboctlArg:         "tagroute",
	DubboctlListArg:     "tagroutes",
	AllowToInspect:      true,
	IsPolicy:            true,
	SingularDisplayName: "Tag Route",
	PluralDisplayName:   "Tag Routes",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(TagRouteResourceTypeDescriptor)
}

const (
	ZoneEgressType model.ResourceType = "ZoneEgress"
)

var _ model.Resource = &ZoneEgressResource{}

type ZoneEgressResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.ZoneEgress
}

func NewZoneEgressResource() *ZoneEgressResource {
	return &ZoneEgressResource{
		Spec: &mesh_proto.ZoneEgress{},
	}
}

func (t *ZoneEgressResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *ZoneEgressResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *ZoneEgressResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *ZoneEgressResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.ZoneEgress)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.ZoneEgress{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *ZoneEgressResource) Descriptor() model.ResourceTypeDescriptor {
	return ZoneEgressResourceTypeDescriptor
}

var _ model.ResourceList = &ZoneEgressResourceList{}

type ZoneEgressResourceList struct {
	Items      []*ZoneEgressResource
	Pagination model.Pagination
}

func (l *ZoneEgressResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *ZoneEgressResourceList) GetItemType() model.ResourceType {
	return ZoneEgressType
}

func (l *ZoneEgressResourceList) NewItem() model.Resource {
	return NewZoneEgressResource()
}

func (l *ZoneEgressResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*ZoneEgressResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*ZoneEgressResource)(nil), r)
	}
}

func (l *ZoneEgressResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *ZoneEgressResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var ZoneEgressResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                ZoneEgressType,
	Resource:            NewZoneEgressResource(),
	ResourceList:        &ZoneEgressResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	DDSFlags:            model.ZoneToGlobalFlag | model.GlobalToAllButOriginalZoneFlag,
	WsPath:              "zoneegresses",
	DubboctlArg:         "zoneegress",
	DubboctlListArg:     "zoneegresses",
	AllowToInspect:      true,
	IsPolicy:            false,
	SingularDisplayName: "Zone Egress",
	PluralDisplayName:   "Zone Egresses",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(ZoneEgressResourceTypeDescriptor)
}

const (
	ZoneEgressInsightType model.ResourceType = "ZoneEgressInsight"
)

var _ model.Resource = &ZoneEgressInsightResource{}

type ZoneEgressInsightResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.ZoneEgressInsight
}

func NewZoneEgressInsightResource() *ZoneEgressInsightResource {
	return &ZoneEgressInsightResource{
		Spec: &mesh_proto.ZoneEgressInsight{},
	}
}

func (t *ZoneEgressInsightResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *ZoneEgressInsightResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *ZoneEgressInsightResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *ZoneEgressInsightResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.ZoneEgressInsight)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.ZoneEgressInsight{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *ZoneEgressInsightResource) Descriptor() model.ResourceTypeDescriptor {
	return ZoneEgressInsightResourceTypeDescriptor
}

var _ model.ResourceList = &ZoneEgressInsightResourceList{}

type ZoneEgressInsightResourceList struct {
	Items      []*ZoneEgressInsightResource
	Pagination model.Pagination
}

func (l *ZoneEgressInsightResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *ZoneEgressInsightResourceList) GetItemType() model.ResourceType {
	return ZoneEgressInsightType
}

func (l *ZoneEgressInsightResourceList) NewItem() model.Resource {
	return NewZoneEgressInsightResource()
}

func (l *ZoneEgressInsightResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*ZoneEgressInsightResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*ZoneEgressInsightResource)(nil), r)
	}
}

func (l *ZoneEgressInsightResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *ZoneEgressInsightResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var ZoneEgressInsightResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                ZoneEgressInsightType,
	Resource:            NewZoneEgressInsightResource(),
	ResourceList:        &ZoneEgressInsightResourceList{},
	ReadOnly:            true,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	DDSFlags:            model.ZoneToGlobalFlag,
	WsPath:              "zoneegressinsights",
	DubboctlArg:         "",
	DubboctlListArg:     "",
	AllowToInspect:      false,
	IsPolicy:            false,
	SingularDisplayName: "Zone Egress Insight",
	PluralDisplayName:   "Zone Egress Insights",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(ZoneEgressInsightResourceTypeDescriptor)
}

const (
	ZoneIngressType model.ResourceType = "ZoneIngress"
)

var _ model.Resource = &ZoneIngressResource{}

type ZoneIngressResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.ZoneIngress
}

func NewZoneIngressResource() *ZoneIngressResource {
	return &ZoneIngressResource{
		Spec: &mesh_proto.ZoneIngress{},
	}
}

func (t *ZoneIngressResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *ZoneIngressResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *ZoneIngressResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *ZoneIngressResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.ZoneIngress)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.ZoneIngress{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *ZoneIngressResource) Descriptor() model.ResourceTypeDescriptor {
	return ZoneIngressResourceTypeDescriptor
}

var _ model.ResourceList = &ZoneIngressResourceList{}

type ZoneIngressResourceList struct {
	Items      []*ZoneIngressResource
	Pagination model.Pagination
}

func (l *ZoneIngressResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *ZoneIngressResourceList) GetItemType() model.ResourceType {
	return ZoneIngressType
}

func (l *ZoneIngressResourceList) NewItem() model.Resource {
	return NewZoneIngressResource()
}

func (l *ZoneIngressResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*ZoneIngressResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*ZoneIngressResource)(nil), r)
	}
}

func (l *ZoneIngressResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *ZoneIngressResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var ZoneIngressResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                ZoneIngressType,
	Resource:            NewZoneIngressResource(),
	ResourceList:        &ZoneIngressResourceList{},
	ReadOnly:            false,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	DDSFlags:            model.ZoneToGlobalFlag | model.GlobalToAllButOriginalZoneFlag,
	WsPath:              "zoneingresses",
	DubboctlArg:         "zoneingress",
	DubboctlListArg:     "zoneingresses",
	AllowToInspect:      true,
	IsPolicy:            false,
	SingularDisplayName: "Zone Ingress",
	PluralDisplayName:   "Zone Ingresses",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(ZoneIngressResourceTypeDescriptor)
}

const (
	ZoneIngressInsightType model.ResourceType = "ZoneIngressInsight"
)

var _ model.Resource = &ZoneIngressInsightResource{}

type ZoneIngressInsightResource struct {
	Meta model.ResourceMeta
	Spec *mesh_proto.ZoneIngressInsight
}

func NewZoneIngressInsightResource() *ZoneIngressInsightResource {
	return &ZoneIngressInsightResource{
		Spec: &mesh_proto.ZoneIngressInsight{},
	}
}

func (t *ZoneIngressInsightResource) GetMeta() model.ResourceMeta {
	return t.Meta
}

func (t *ZoneIngressInsightResource) SetMeta(m model.ResourceMeta) {
	t.Meta = m
}

func (t *ZoneIngressInsightResource) GetSpec() model.ResourceSpec {
	return t.Spec
}

func (t *ZoneIngressInsightResource) SetSpec(spec model.ResourceSpec) error {
	protoType, ok := spec.(*mesh_proto.ZoneIngressInsight)
	if !ok {
		return fmt.Errorf("invalid type %T for Spec", spec)
	} else {
		if protoType == nil {
			t.Spec = &mesh_proto.ZoneIngressInsight{}
		} else {
			t.Spec = protoType
		}
		return nil
	}
}

func (t *ZoneIngressInsightResource) Descriptor() model.ResourceTypeDescriptor {
	return ZoneIngressInsightResourceTypeDescriptor
}

var _ model.ResourceList = &ZoneIngressInsightResourceList{}

type ZoneIngressInsightResourceList struct {
	Items      []*ZoneIngressInsightResource
	Pagination model.Pagination
}

func (l *ZoneIngressInsightResourceList) GetItems() []model.Resource {
	res := make([]model.Resource, len(l.Items))
	for i, elem := range l.Items {
		res[i] = elem
	}
	return res
}

func (l *ZoneIngressInsightResourceList) GetItemType() model.ResourceType {
	return ZoneIngressInsightType
}

func (l *ZoneIngressInsightResourceList) NewItem() model.Resource {
	return NewZoneIngressInsightResource()
}

func (l *ZoneIngressInsightResourceList) AddItem(r model.Resource) error {
	if trr, ok := r.(*ZoneIngressInsightResource); ok {
		l.Items = append(l.Items, trr)
		return nil
	} else {
		return model.ErrorInvalidItemType((*ZoneIngressInsightResource)(nil), r)
	}
}

func (l *ZoneIngressInsightResourceList) GetPagination() *model.Pagination {
	return &l.Pagination
}

func (l *ZoneIngressInsightResourceList) SetPagination(p model.Pagination) {
	l.Pagination = p
}

var ZoneIngressInsightResourceTypeDescriptor = model.ResourceTypeDescriptor{
	Name:                ZoneIngressInsightType,
	Resource:            NewZoneIngressInsightResource(),
	ResourceList:        &ZoneIngressInsightResourceList{},
	ReadOnly:            true,
	AdminOnly:           false,
	Scope:               model.ScopeGlobal,
	DDSFlags:            model.ZoneToGlobalFlag,
	WsPath:              "zone-ingress-insights",
	DubboctlArg:         "",
	DubboctlListArg:     "",
	AllowToInspect:      false,
	IsPolicy:            false,
	SingularDisplayName: "Zone Ingress Insight",
	PluralDisplayName:   "Zone Ingress Insights",
	IsExperimental:      false,
}

func init() {
	registry.RegisterType(ZoneIngressInsightResourceTypeDescriptor)
}
