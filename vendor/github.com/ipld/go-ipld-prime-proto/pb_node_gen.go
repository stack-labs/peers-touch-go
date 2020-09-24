package dagpb

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/impl/typed"
	"github.com/ipld/go-ipld-prime/schema"
)

// Code generated go-ipld-prime DO NOT EDIT.

type String struct{ x string }

func (x String) String() string {
	return x.x
}
type String__Content struct {
	Value string
}

func (b String__Content) Build() (String, error) {
	x := String{
		b.Value,
	}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b String__Content) MustBuild() String {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybeString struct {
	Maybe typed.Maybe
	Value String
}

func (m MaybeString) Must() String {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = String{}
var _ typed.Node = String{}

func (String) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (String) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_String
}
func (String) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (String) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (String) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}
}
func (String) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_String}
}
func (String) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "String", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}}
}
func (String) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "String", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}}
}
func (String) Length() int {
	return -1
}
func (String) IsUndefined() bool {
	return false
}
func (String) IsNull() bool {
	return false
}
func (String) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_String}
}
func (String) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_String}
}
func (String) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_String}
}
func (x String) AsString() (string, error) {
	return x.x, nil
}
func (String) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_String}
}
func (String) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_String}
}
func (String) NodeBuilder() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
type _String__NodeBuilder struct{}
func String__NodeBuilder() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
func (_String__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_String}
}
func (nb _String__NodeBuilder) CreateString(v string) (ipld.Node, error) {
	return String{v}, nil
}
func (_String__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_String}
}
func (_String__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "String.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_String}
}
func (String) Representation() ipld.Node {
	panic("TODO representation")
}
type Int struct{ x int }

func (x Int) Int() int {
	return x.x
}
// TODO generateKindInt.EmitNativeBuilder
type MaybeInt struct {
	Maybe typed.Maybe
	Value Int
}

func (m MaybeInt) Must() Int {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = Int{}
var _ typed.Node = Int{}

func (Int) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Int) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Int
}
func (Int) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (Int) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (Int) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}
}
func (Int) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Int}
}
func (Int) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Int", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}}
}
func (Int) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Int", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}}
}
func (Int) Length() int {
	return -1
}
func (Int) IsUndefined() bool {
	return false
}
func (Int) IsNull() bool {
	return false
}
func (Int) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Int}
}
func (x Int) AsInt() (int, error) {
	return x.x, nil
}
func (Int) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Int}
}
func (Int) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Int}
}
func (Int) NodeBuilder() ipld.NodeBuilder {
	return _Int__NodeBuilder{}
}
type _Int__NodeBuilder struct{}
func Int__NodeBuilder() ipld.NodeBuilder {
	return _Int__NodeBuilder{}
}
func (_Int__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Int}
}
func (nb _Int__NodeBuilder) CreateInt(v int) (ipld.Node, error) {
	return Int{v}, nil
}
func (_Int__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Int}
}
func (_Int__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Int.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Int}
}
func (Int) Representation() ipld.Node {
	panic("TODO representation")
}
type Bytes struct{ x []byte }

// TODO generateKindBytes.EmitNativeAccessors
// TODO generateKindBytes.EmitNativeBuilder
type MaybeBytes struct {
	Maybe typed.Maybe
	Value Bytes
}

func (m MaybeBytes) Must() Bytes {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = Bytes{}
var _ typed.Node = Bytes{}

func (Bytes) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Bytes) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Bytes
}
func (Bytes) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}}
}
func (Bytes) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}}
}
func (Bytes) Length() int {
	return -1
}
func (Bytes) IsUndefined() bool {
	return false
}
func (Bytes) IsNull() bool {
	return false
}
func (Bytes) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Bytes}
}
func (x Bytes) AsBytes() ([]byte, error) {
	return x.x, nil
}
func (Bytes) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) NodeBuilder() ipld.NodeBuilder {
	return _Bytes__NodeBuilder{}
}
type _Bytes__NodeBuilder struct{}
func Bytes__NodeBuilder() ipld.NodeBuilder {
	return _Bytes__NodeBuilder{}
}
func (_Bytes__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Bytes}
}
func (_Bytes__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Bytes}
}
func (nb _Bytes__NodeBuilder) CreateBytes(v []byte) (ipld.Node, error) {
	return Bytes{v}, nil
}
func (_Bytes__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Bytes.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Bytes}
}
func (Bytes) Representation() ipld.Node {
	panic("TODO representation")
}
type Link struct{ x ipld.Link }

// TODO generateKindLink.EmitNativeAccessors
// TODO generateKindLink.EmitNativeBuilder
type MaybeLink struct {
	Maybe typed.Maybe
	Value Link
}

func (m MaybeLink) Must() Link {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = Link{}
var _ typed.Node = Link{}

func (Link) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Link) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (Link) Lookup(ipld.Node) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "Lookup", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (Link) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}
}
func (Link) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "LookupSegment", AppropriateKind: ipld.ReprKindSet_Recursive, ActualKind: ipld.ReprKind_Link}
}
func (Link) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "Link", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}}
}
func (Link) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "Link", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}}
}
func (Link) Length() int {
	return -1
}
func (Link) IsUndefined() bool {
	return false
}
func (Link) IsNull() bool {
	return false
}
func (Link) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Link}
}
func (Link) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Link}
}
func (x Link) AsLink() (ipld.Link, error) {
	return x.x, nil
}
func (Link) NodeBuilder() ipld.NodeBuilder {
	return _Link__NodeBuilder{}
}
type _Link__NodeBuilder struct{}

func Link__NodeBuilder() ipld.NodeBuilder {
	return _Link__NodeBuilder{}
}
func (_Link__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Link}
}
func (_Link__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "Link.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Link}
}
func (nb _Link__NodeBuilder) CreateLink(v ipld.Link) (ipld.Node, error) {
	return Link{v}, nil
}
func (Link) Representation() ipld.Node {
	panic("TODO representation")
}
type PBLink struct{
	Hash *Link
	Name *String
	Tsize *Int
	
}

func (x Link) FieldHash() Link {
	// TODO going to tear through here with changes to Maybe system in a moment anyway
	return Link{}
}
func (x String) FieldName() String {
	// TODO going to tear through here with changes to Maybe system in a moment anyway
	return String{}
}
func (x Int) FieldTsize() Int {
	// TODO going to tear through here with changes to Maybe system in a moment anyway
	return Int{}
}


type PBLink__Content struct {
	// TODO
	// TODO
	// TODO
}

func (b PBLink__Content) Build() (PBLink, error) {
	x := PBLink{
		// TODO
	}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b PBLink__Content) MustBuild() PBLink {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybePBLink struct {
	Maybe typed.Maybe
	Value PBLink
}

func (m MaybePBLink) Must() PBLink {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = PBLink{}
var _ typed.Node = PBLink{}

func (PBLink) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (PBLink) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (x PBLink) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Hash":
		if x.Hash == nil {
			return ipld.Undef, nil
		}
		return *x.Hash, nil
	case "Name":
		if x.Name == nil {
			return ipld.Undef, nil
		}
		return *x.Name, nil
	case "Tsize":
		if x.Tsize == nil {
			return ipld.Undef, nil
		}
		return *x.Tsize, nil
	default:
		return nil, typed.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (x PBLink) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, ipld.ErrInvalidKey{"got " + key.ReprKind().String() + ", need string"}
	}
	return x.LookupString(ks)
}
func (PBLink) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n PBLink) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (x PBLink) MapIterator() ipld.MapIterator {
	return &_PBLink__Itr{&x, 0}
}

type _PBLink__Itr struct {
	node *PBLink
	idx  int
}

func (itr *_PBLink__Itr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = String{"Hash"}
		if itr.node.Hash == nil {
			v = ipld.Undef
			break
		}
		v = itr.node.Hash
	case 1:
		k = String{"Name"}
		if itr.node.Name == nil {
			v = ipld.Undef
			break
		}
		v = itr.node.Name
	case 2:
		k = String{"Tsize"}
		if itr.node.Tsize == nil {
			v = ipld.Undef
			break
		}
		v = itr.node.Tsize
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PBLink__Itr) Done() bool {
	return itr.idx >= 3
}

func (PBLink) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (PBLink) Length() int {
	return 3
}
func (PBLink) IsUndefined() bool {
	return false
}
func (PBLink) IsNull() bool {
	return false
}
func (PBLink) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (PBLink) NodeBuilder() ipld.NodeBuilder {
	return _PBLink__NodeBuilder{}
}
type _PBLink__NodeBuilder struct{}

func PBLink__NodeBuilder() ipld.NodeBuilder {
	return _PBLink__NodeBuilder{}
}
func (nb _PBLink__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return &_PBLink__MapBuilder{v:&PBLink{}}, nil
}

type _PBLink__MapBuilder struct{
	v *PBLink
}

func (mb *_PBLink__MapBuilder) Insert(k, v ipld.Node) error {
	ks, err := k.AsString()
	if err != nil {
		return ipld.ErrInvalidKey{"not a string: " + err.Error()}
	}
	switch ks {
	case "Hash":
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Link)
		if !ok {
			panic("field 'Hash' in type PBLink is type Link; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Hash = &x
	case "Name":
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(String)
		if !ok {
			panic("field 'Name' in type PBLink is type String; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Name = &x
	case "Tsize":
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Int)
		if !ok {
			panic("field 'Tsize' in type PBLink is type Int; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Tsize = &x
	default:
		return typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks}
	}
	return nil
}
func (mb *_PBLink__MapBuilder) Delete(k ipld.Node) error {
	panic("TODO later")
}
func (mb *_PBLink__MapBuilder) Build() (ipld.Node, error) {
	v := *mb.v
	mb = nil
	return v, nil
}
func (mb *_PBLink__MapBuilder) BuilderForKeys() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
func (mb *_PBLink__MapBuilder) BuilderForValue(ks string) ipld.NodeBuilder {
	switch ks {
	case "Hash":
		return Link__NodeBuilder()
	case "Name":
		return String__NodeBuilder()
	case "Tsize":
		return Int__NodeBuilder()
	default:
		panic(typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks})
	}
	return nil
}

func (nb _PBLink__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	panic("TODO later")
}
func (_PBLink__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (n PBLink) Representation() ipld.Node {
	return _PBLink__Repr{&n}
}
var _ ipld.Node = _PBLink__Repr{}

type _PBLink__Repr struct{
	n *PBLink
}

func (_PBLink__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (rn _PBLink__Repr) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Hash":
		if rn.n.Hash == nil {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return *rn.n.Hash, nil
	case "Name":
		if rn.n.Name == nil {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return *rn.n.Name, nil
	case "Tsize":
		if rn.n.Tsize == nil {
			return ipld.Undef, ipld.ErrNotExists{ipld.PathSegmentOfString(key)}
		}
		return *rn.n.Tsize, nil
	default:
		return nil, typed.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (rn _PBLink__Repr) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, ipld.ErrInvalidKey{"got " + key.ReprKind().String() + ", need string"}
	}
	return rn.LookupString(ks)
}
func (_PBLink__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n _PBLink__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (rn _PBLink__Repr) MapIterator() ipld.MapIterator {
	return &_PBLink__ReprItr{rn.n, 0}
}

type _PBLink__ReprItr struct {
	node *PBLink
	idx  int
}

func (itr *_PBLink__ReprItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	for {
		switch itr.idx {
		case 0:
			k = String{"Hash"}
			if itr.node.Hash == nil {
				itr.idx++
				continue
			}
			v = itr.node.Hash
		case 1:
			k = String{"Name"}
			if itr.node.Name == nil {
				itr.idx++
				continue
			}
			v = itr.node.Name
		case 2:
			k = String{"Tsize"}
			if itr.node.Tsize == nil {
				itr.idx++
				continue
			}
			v = itr.node.Tsize
		default:
			panic("unreachable")
		}
	}
	itr.idx++
	return
}
func (itr *_PBLink__ReprItr) Done() bool {
	return itr.idx >= 3
}

func (_PBLink__Repr) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (rn _PBLink__Repr) Length() int {
	l := 3
	if rn.n.Hash == nil {
		l--
	}
	if rn.n.Name == nil {
		l--
	}
	if rn.n.Tsize == nil {
		l--
	}
	return l
}
func (_PBLink__Repr) IsUndefined() bool {
	return false
}
func (_PBLink__Repr) IsNull() bool {
	return false
}
func (_PBLink__Repr) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__Repr) NodeBuilder() ipld.NodeBuilder {
	return _PBLink__ReprBuilder{}
}
type _PBLink__ReprBuilder struct{}

func PBLink__ReprBuilder() ipld.NodeBuilder {
	return _PBLink__ReprBuilder{}
}
func (nb _PBLink__ReprBuilder) CreateMap() (ipld.MapBuilder, error) {
	return &_PBLink__ReprMapBuilder{v:&PBLink{}}, nil
}

type _PBLink__ReprMapBuilder struct{
	v *PBLink
	Hash__isset bool
	Name__isset bool
	Tsize__isset bool
}

func (mb *_PBLink__ReprMapBuilder) Insert(k, v ipld.Node) error {
	ks, err := k.AsString()
	if err != nil {
		return ipld.ErrInvalidKey{"not a string: " + err.Error()}
	}
	switch ks {
	case "Hash":
		if mb.Hash__isset {
			panic("repeated assignment to field") // FIXME need an error type for this
		}
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Link)
		if !ok {
			panic("field 'Hash' (key: 'Hash') in type PBLink is type Link; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Hash = &x
		mb.Hash__isset = true
	case "Name":
		if mb.Name__isset {
			panic("repeated assignment to field") // FIXME need an error type for this
		}
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(String)
		if !ok {
			panic("field 'Name' (key: 'Name') in type PBLink is type String; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Name = &x
		mb.Name__isset = true
	case "Tsize":
		if mb.Tsize__isset {
			panic("repeated assignment to field") // FIXME need an error type for this
		}
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Int)
		if !ok {
			panic("field 'Tsize' (key: 'Tsize') in type PBLink is type Int; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Tsize = &x
		mb.Tsize__isset = true
	default:
		return typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks}
	}
	return nil
}
func (mb *_PBLink__ReprMapBuilder) Delete(k ipld.Node) error {
	panic("TODO later")
}
func (mb *_PBLink__ReprMapBuilder) Build() (ipld.Node, error) {
	v := mb.v
	mb = nil
	return v, nil
}
func (mb *_PBLink__ReprMapBuilder) BuilderForKeys() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
func (mb *_PBLink__ReprMapBuilder) BuilderForValue(ks string) ipld.NodeBuilder {
	switch ks {
	case "Hash":
		return Link__NodeBuilder()
	case "Name":
		return String__NodeBuilder()
	case "Tsize":
		return Int__NodeBuilder()
	default:
		panic(typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks})
	}
	return nil
}

func (nb _PBLink__ReprBuilder) AmendMap() (ipld.MapBuilder, error) {
	panic("TODO later")
}
func (_PBLink__ReprBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBLink__ReprBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLink.Representation.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
type PBLinks struct{
	x []PBLink
}
// TODO generateKindList.EmitNativeAccessors
// TODO generateKindList.EmitNativeBuilder
type MaybePBLinks struct {
	Maybe typed.Maybe
	Value PBLinks
}

func (m MaybePBLinks) Must() PBLinks {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = PBLinks{}
var _ typed.Node = PBLinks{}

func (PBLinks) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (PBLinks) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (PBLinks) LookupString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "LookupString", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}
}
func (x PBLinks) Lookup(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, ipld.ErrInvalidKey{"got " + key.ReprKind().String() + ", need Int"}
	}
	return x.LookupIndex(ki)
}
func (x PBLinks) LookupIndex(index int) (ipld.Node, error) {
	if index >= len(x.x) {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(index)}
	}
	return x.x[index], nil
}
func (n PBLinks) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	idx, err := seg.Index()
	if err != nil {
		return nil, err
	}
	return n.LookupIndex(idx)
}
func (PBLinks) MapIterator() ipld.MapIterator {
	return mapIteratorReject{ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "MapIterator", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}}
}
func (x PBLinks) ListIterator() ipld.ListIterator {
	return &_PBLinks__Itr{&x, 0}
}

type _PBLinks__Itr struct {
	node *PBLinks
	idx  int
}

func (itr *_PBLinks__Itr) Next() (idx int, value ipld.Node, _ error)	{
	if itr.idx >= len(itr.node.x) {
		return 0, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	value = itr.node.x[idx]
	itr.idx++
	return
}

func (itr *_PBLinks__Itr) Done() bool {
	return itr.idx >= len(itr.node.x)
}

func (x PBLinks) Length() int {
	return len(x.x)
}
func (PBLinks) IsUndefined() bool {
	return false
}
func (PBLinks) IsNull() bool {
	return false
}
func (PBLinks) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_List}
}
func (PBLinks) NodeBuilder() ipld.NodeBuilder {
	return _PBLinks__NodeBuilder{}
}
type _PBLinks__NodeBuilder struct{}

func PBLinks__NodeBuilder() ipld.NodeBuilder {
	return _PBLinks__NodeBuilder{}
}
func (_PBLinks__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "AmendMap", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: ipld.ReprKind_List}
}
func (nb _PBLinks__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return &_PBLinks__ListBuilder{v:&PBLinks{}}, nil
}

type _PBLinks__ListBuilder struct{
	v *PBLinks
}

func (lb *_PBLinks__ListBuilder) growList(k int) {
	oldLen := len(lb.v.x)
	minLen := k + 1
	if minLen > oldLen {
		// Grow.
		oldCap := cap(lb.v.x)
		if minLen > oldCap {
			// Out of cap; do whole new backing array allocation.
			//  Growth maths are per stdlib's reflect.grow.
			// First figure out how much growth to do.
			newCap := oldCap
			if newCap == 0 {
				newCap = minLen
			} else {
				for minLen > newCap {
					if minLen < 1024 {
						newCap += newCap
					} else {
						newCap += newCap / 4
					}
				}
			}
			// Now alloc and copy over old.
			newArr := make([]PBLink, minLen, newCap)
			copy(newArr, lb.v.x)
			lb.v.x = newArr
		} else {
			// Still have cap, just extend the slice.
			lb.v.x = lb.v.x[0:minLen]
		}
	}
}

func (lb *_PBLinks__ListBuilder) validate(v ipld.Node) error {
	if v.IsNull() {
		panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
	}
	tv, ok := v.(typed.Node)
	if !ok {
		panic("need typed.Node for insertion into struct") // FIXME need an error type for this
	}
	_, ok = v.(PBLink)
	if !ok {
		panic("value for type PBLinks is type PBLink; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
	}
	return nil
}

func (lb *_PBLinks__ListBuilder) unsafeSet(idx int, v ipld.Node) {
	x := v.(PBLink)
	lb.v.x[idx] = x
}

func (lb *_PBLinks__ListBuilder) AppendAll(vs []ipld.Node) error {
	for _, v := range vs {
		err := lb.validate(v)
		if err != nil {
			return err
		}
	}
	off := len(lb.v.x)
	new := off + len(vs)
	lb.growList(new-1)
	for _, v := range vs {
		lb.unsafeSet(off, v)
		off++
	}
	return nil
}

func (lb *_PBLinks__ListBuilder) Append(v ipld.Node) error {
	err := lb.validate(v)
	if err != nil {
		return err
	}
	off := len(lb.v.x)
	lb.growList(off)
	lb.unsafeSet(off, v)
	return nil
}
func (lb *_PBLinks__ListBuilder) Set(idx int, v ipld.Node) error {
	err := lb.validate(v)
	if err != nil {
		return err
	}
	lb.growList(idx)
	lb.unsafeSet(idx, v)
	return nil
}

func (lb *_PBLinks__ListBuilder) Build() (ipld.Node, error) {
	v := *lb.v
	lb = nil
	return v, nil
}

func (lb *_PBLinks__ListBuilder) BuilderForValue(_ int) ipld.NodeBuilder {
	return PBLink__NodeBuilder()
}

func (nb _PBLinks__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	panic("TODO later")
}
func (_PBLinks__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_List}
}
func (_PBLinks__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBLinks.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_List}
}
func (n PBLinks) Representation() ipld.Node {
	panic("TODO representation")
}
type PBNode struct{
	Links PBLinks
	Data Bytes
	
}

func (x PBLinks) FieldLinks() PBLinks {
	// TODO going to tear through here with changes to Maybe system in a moment anyway
	return PBLinks{}
}
func (x Bytes) FieldData() Bytes {
	// TODO going to tear through here with changes to Maybe system in a moment anyway
	return Bytes{}
}


type PBNode__Content struct {
	// TODO
	// TODO
}

func (b PBNode__Content) Build() (PBNode, error) {
	x := PBNode{
		// TODO
	}
	// FUTURE : want to support customizable validation.
	//   but 'if v, ok := x.(schema.Validatable); ok {' doesn't fly: need a way to work on concrete types.
	return x, nil
}
func (b PBNode__Content) MustBuild() PBNode {
	if x, err := b.Build(); err != nil {
		panic(err)
	} else {
		return x
	}
}

type MaybePBNode struct {
	Maybe typed.Maybe
	Value PBNode
}

func (m MaybePBNode) Must() PBNode {
	if m.Maybe != typed.Maybe_Value {
		panic("unbox of a maybe rejected")
	}
	return m.Value
}

var _ ipld.Node = PBNode{}
var _ typed.Node = PBNode{}

func (PBNode) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (PBNode) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (x PBNode) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Links":
		return x.Links, nil
	case "Data":
		return x.Data, nil
	default:
		return nil, typed.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (x PBNode) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, ipld.ErrInvalidKey{"got " + key.ReprKind().String() + ", need string"}
	}
	return x.LookupString(ks)
}
func (PBNode) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n PBNode) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (x PBNode) MapIterator() ipld.MapIterator {
	return &_PBNode__Itr{&x, 0}
}

type _PBNode__Itr struct {
	node *PBNode
	idx  int
}

func (itr *_PBNode__Itr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = String{"Links"}
		v = itr.node.Links
	case 1:
		k = String{"Data"}
		v = itr.node.Data
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PBNode__Itr) Done() bool {
	return itr.idx >= 2
}

func (PBNode) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (PBNode) Length() int {
	return 2
}
func (PBNode) IsUndefined() bool {
	return false
}
func (PBNode) IsNull() bool {
	return false
}
func (PBNode) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (PBNode) NodeBuilder() ipld.NodeBuilder {
	return _PBNode__NodeBuilder{}
}
type _PBNode__NodeBuilder struct{}

func PBNode__NodeBuilder() ipld.NodeBuilder {
	return _PBNode__NodeBuilder{}
}
func (nb _PBNode__NodeBuilder) CreateMap() (ipld.MapBuilder, error) {
	return &_PBNode__MapBuilder{v:&PBNode{}}, nil
}

type _PBNode__MapBuilder struct{
	v *PBNode
	Links__isset bool
	Data__isset bool
}

func (mb *_PBNode__MapBuilder) Insert(k, v ipld.Node) error {
	ks, err := k.AsString()
	if err != nil {
		return ipld.ErrInvalidKey{"not a string: " + err.Error()}
	}
	switch ks {
	case "Links":
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(PBLinks)
		if !ok {
			panic("field 'Links' in type PBNode is type PBLinks; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Links = x
		mb.Links__isset = true
	case "Data":
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Bytes)
		if !ok {
			panic("field 'Data' in type PBNode is type Bytes; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Data = x
		mb.Data__isset = true
	default:
		return typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks}
	}
	return nil
}
func (mb *_PBNode__MapBuilder) Delete(k ipld.Node) error {
	panic("TODO later")
}
func (mb *_PBNode__MapBuilder) Build() (ipld.Node, error) {
	if !mb.Links__isset {
		panic("missing required field 'Links' in building struct PBNode") // FIXME need an error type for this
	}
	if !mb.Data__isset {
		panic("missing required field 'Data' in building struct PBNode") // FIXME need an error type for this
	}
	v := *mb.v
	mb = nil
	return v, nil
}
func (mb *_PBNode__MapBuilder) BuilderForKeys() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
func (mb *_PBNode__MapBuilder) BuilderForValue(ks string) ipld.NodeBuilder {
	switch ks {
	case "Links":
		return PBLinks__NodeBuilder()
	case "Data":
		return Bytes__NodeBuilder()
	default:
		panic(typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks})
	}
	return nil
}

func (nb _PBNode__NodeBuilder) AmendMap() (ipld.MapBuilder, error) {
	panic("TODO later")
}
func (_PBNode__NodeBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__NodeBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (n PBNode) Representation() ipld.Node {
	return _PBNode__Repr{&n}
}
var _ ipld.Node = _PBNode__Repr{}

type _PBNode__Repr struct{
	n *PBNode
}

func (_PBNode__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (rn _PBNode__Repr) LookupString(key string) (ipld.Node, error) {
	switch key {
	case "Links":
		return rn.n.Links, nil
	case "Data":
		return rn.n.Data, nil
	default:
		return nil, typed.ErrNoSuchField{Type: nil /*TODO*/, FieldName: key}
	}
}
func (rn _PBNode__Repr) Lookup(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, ipld.ErrInvalidKey{"got " + key.ReprKind().String() + ", need string"}
	}
	return rn.LookupString(ks)
}
func (_PBNode__Repr) LookupIndex(idx int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "LookupIndex", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (n _PBNode__Repr) LookupSegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupString(seg.String())
}
func (rn _PBNode__Repr) MapIterator() ipld.MapIterator {
	return &_PBNode__ReprItr{rn.n, 0}
}

type _PBNode__ReprItr struct {
	node *PBNode
	idx  int
}

func (itr *_PBNode__ReprItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	for {
		switch itr.idx {
		case 0:
			k = String{"Links"}
			v = itr.node.Links
		case 1:
			k = String{"Data"}
			v = itr.node.Data
		default:
			panic("unreachable")
		}
	}
	itr.idx++
	return
}
func (itr *_PBNode__ReprItr) Done() bool {
	return itr.idx >= 2
}

func (_PBNode__Repr) ListIterator() ipld.ListIterator {
	return listIteratorReject{ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "ListIterator", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}}
}
func (rn _PBNode__Repr) Length() int {
	l := 2
	return l
}
func (_PBNode__Repr) IsUndefined() bool {
	return false
}
func (_PBNode__Repr) IsNull() bool {
	return false
}
func (_PBNode__Repr) AsBool() (bool, error) {
	return false, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsInt() (int, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsFloat() (float64, error) {
	return 0, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsString() (string, error) {
	return "", ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsBytes() ([]byte, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) AsLink() (ipld.Link, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation", MethodName: "AsLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__Repr) NodeBuilder() ipld.NodeBuilder {
	return _PBNode__ReprBuilder{}
}
type _PBNode__ReprBuilder struct{}

func PBNode__ReprBuilder() ipld.NodeBuilder {
	return _PBNode__ReprBuilder{}
}
func (nb _PBNode__ReprBuilder) CreateMap() (ipld.MapBuilder, error) {
	return &_PBNode__ReprMapBuilder{v:&PBNode{}}, nil
}

type _PBNode__ReprMapBuilder struct{
	v *PBNode
	Links__isset bool
	Data__isset bool
}

func (mb *_PBNode__ReprMapBuilder) Insert(k, v ipld.Node) error {
	ks, err := k.AsString()
	if err != nil {
		return ipld.ErrInvalidKey{"not a string: " + err.Error()}
	}
	switch ks {
	case "Links":
		if mb.Links__isset {
			panic("repeated assignment to field") // FIXME need an error type for this
		}
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(PBLinks)
		if !ok {
			panic("field 'Links' (key: 'Links') in type PBNode is type PBLinks; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Links = x
		mb.Links__isset = true
	case "Data":
		if mb.Data__isset {
			panic("repeated assignment to field") // FIXME need an error type for this
		}
		if v.IsNull() {
			panic("type mismatch on struct field assignment: cannot assign null to non-nullable field") // FIXME need an error type for this
		}
		tv, ok := v.(typed.Node)
		if !ok {
			panic("need typed.Node for insertion into struct") // FIXME need an error type for this
		}
		x, ok := v.(Bytes)
		if !ok {
			panic("field 'Data' (key: 'Data') in type PBNode is type Bytes; cannot assign "+tv.Type().Name()) // FIXME need an error type for this
		}
		mb.v.Data = x
		mb.Data__isset = true
	default:
		return typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks}
	}
	return nil
}
func (mb *_PBNode__ReprMapBuilder) Delete(k ipld.Node) error {
	panic("TODO later")
}
func (mb *_PBNode__ReprMapBuilder) Build() (ipld.Node, error) {
	if !mb.Links__isset {
		panic("missing required field 'Links' (key: 'Links') in building struct PBNode") // FIXME need an error type for this
	}
	if !mb.Data__isset {
		panic("missing required field 'Data' (key: 'Data') in building struct PBNode") // FIXME need an error type for this
	}
	v := mb.v
	mb = nil
	return v, nil
}
func (mb *_PBNode__ReprMapBuilder) BuilderForKeys() ipld.NodeBuilder {
	return _String__NodeBuilder{}
}
func (mb *_PBNode__ReprMapBuilder) BuilderForValue(ks string) ipld.NodeBuilder {
	switch ks {
	case "Links":
		return PBLinks__NodeBuilder()
	case "Data":
		return Bytes__NodeBuilder()
	default:
		panic(typed.ErrNoSuchField{Type: nil /*TODO:typelit*/, FieldName: ks})
	}
	return nil
}

func (nb _PBNode__ReprBuilder) AmendMap() (ipld.MapBuilder, error) {
	panic("TODO later")
}
func (_PBNode__ReprBuilder) CreateList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) AmendList() (ipld.ListBuilder, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "AmendList", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateNull() (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateNull", AppropriateKind: ipld.ReprKindSet_JustNull, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateBool(bool) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateBool", AppropriateKind: ipld.ReprKindSet_JustBool, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateInt(int) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateInt", AppropriateKind: ipld.ReprKindSet_JustInt, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateFloat(float64) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateFloat", AppropriateKind: ipld.ReprKindSet_JustFloat, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateString(string) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateString", AppropriateKind: ipld.ReprKindSet_JustString, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateBytes([]byte) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateBytes", AppropriateKind: ipld.ReprKindSet_JustBytes, ActualKind: ipld.ReprKind_Map}
}
func (_PBNode__ReprBuilder) CreateLink(ipld.Link) (ipld.Node, error) {
	return nil, ipld.ErrWrongKind{TypeName: "PBNode.Representation.Builder", MethodName: "CreateLink", AppropriateKind: ipld.ReprKindSet_JustLink, ActualKind: ipld.ReprKind_Map}
}
