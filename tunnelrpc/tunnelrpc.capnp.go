// Code generated by capnpc-go. DO NOT EDIT.

package tunnelrpc

import (
	context "golang.org/x/net/context"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
	server "zombiezen.com/go/capnproto2/server"
)

type Authentication struct{ capnp.Struct }

// Authentication_TypeID is the unique identifier for the type Authentication.
const Authentication_TypeID = 0xc082ef6e0d42ed1d

func NewAuthentication(s *capnp.Segment) (Authentication, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Authentication{st}, err
}

func NewRootAuthentication(s *capnp.Segment) (Authentication, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return Authentication{st}, err
}

func ReadRootAuthentication(msg *capnp.Message) (Authentication, error) {
	root, err := msg.RootPtr()
	return Authentication{root.Struct()}, err
}

func (s Authentication) String() string {
	str, _ := text.Marshal(0xc082ef6e0d42ed1d, s.Struct)
	return str
}

func (s Authentication) Key() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Authentication) HasKey() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Authentication) KeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Authentication) SetKey(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Authentication) Email() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Authentication) HasEmail() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Authentication) EmailBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Authentication) SetEmail(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Authentication) OriginCAKey() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Authentication) HasOriginCAKey() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Authentication) OriginCAKeyBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Authentication) SetOriginCAKey(v string) error {
	return s.Struct.SetText(2, v)
}

// Authentication_List is a list of Authentication.
type Authentication_List struct{ capnp.List }

// NewAuthentication creates a new list of Authentication.
func NewAuthentication_List(s *capnp.Segment, sz int32) (Authentication_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return Authentication_List{l}, err
}

func (s Authentication_List) At(i int) Authentication { return Authentication{s.List.Struct(i)} }

func (s Authentication_List) Set(i int, v Authentication) error { return s.List.SetStruct(i, v.Struct) }

// Authentication_Promise is a wrapper for a Authentication promised by a client call.
type Authentication_Promise struct{ *capnp.Pipeline }

func (p Authentication_Promise) Struct() (Authentication, error) {
	s, err := p.Pipeline.Struct()
	return Authentication{s}, err
}

type TunnelRegistration struct{ capnp.Struct }

// TunnelRegistration_TypeID is the unique identifier for the type TunnelRegistration.
const TunnelRegistration_TypeID = 0xf41a0f001ad49e46

func NewTunnelRegistration(s *capnp.Segment) (TunnelRegistration, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return TunnelRegistration{st}, err
}

func NewRootTunnelRegistration(s *capnp.Segment) (TunnelRegistration, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return TunnelRegistration{st}, err
}

func ReadRootTunnelRegistration(msg *capnp.Message) (TunnelRegistration, error) {
	root, err := msg.RootPtr()
	return TunnelRegistration{root.Struct()}, err
}

func (s TunnelRegistration) String() string {
	str, _ := text.Marshal(0xf41a0f001ad49e46, s.Struct)
	return str
}

func (s TunnelRegistration) Err() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s TunnelRegistration) HasErr() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) ErrBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s TunnelRegistration) SetErr(v string) error {
	return s.Struct.SetText(0, v)
}

func (s TunnelRegistration) Urls() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.TextList{List: p.List()}, err
}

func (s TunnelRegistration) HasUrls() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) SetUrls(v capnp.TextList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewUrls sets the urls field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s TunnelRegistration) NewUrls(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s TunnelRegistration) LogLines() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s TunnelRegistration) HasLogLines() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TunnelRegistration) SetLogLines(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewLogLines sets the logLines field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s TunnelRegistration) NewLogLines(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s TunnelRegistration) PermanentFailure() bool {
	return s.Struct.Bit(0)
}

func (s TunnelRegistration) SetPermanentFailure(v bool) {
	s.Struct.SetBit(0, v)
}

// TunnelRegistration_List is a list of TunnelRegistration.
type TunnelRegistration_List struct{ capnp.List }

// NewTunnelRegistration creates a new list of TunnelRegistration.
func NewTunnelRegistration_List(s *capnp.Segment, sz int32) (TunnelRegistration_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return TunnelRegistration_List{l}, err
}

func (s TunnelRegistration_List) At(i int) TunnelRegistration {
	return TunnelRegistration{s.List.Struct(i)}
}

func (s TunnelRegistration_List) Set(i int, v TunnelRegistration) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelRegistration_Promise is a wrapper for a TunnelRegistration promised by a client call.
type TunnelRegistration_Promise struct{ *capnp.Pipeline }

func (p TunnelRegistration_Promise) Struct() (TunnelRegistration, error) {
	s, err := p.Pipeline.Struct()
	return TunnelRegistration{s}, err
}

type RegistrationOptions struct{ capnp.Struct }

// RegistrationOptions_TypeID is the unique identifier for the type RegistrationOptions.
const RegistrationOptions_TypeID = 0xc793e50592935b4a

func NewRegistrationOptions(s *capnp.Segment) (RegistrationOptions, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return RegistrationOptions{st}, err
}

func NewRootRegistrationOptions(s *capnp.Segment) (RegistrationOptions, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5})
	return RegistrationOptions{st}, err
}

func ReadRootRegistrationOptions(msg *capnp.Message) (RegistrationOptions, error) {
	root, err := msg.RootPtr()
	return RegistrationOptions{root.Struct()}, err
}

func (s RegistrationOptions) String() string {
	str, _ := text.Marshal(0xc793e50592935b4a, s.Struct)
	return str
}

func (s RegistrationOptions) ClientId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s RegistrationOptions) HasClientId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) ClientIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetClientId(v string) error {
	return s.Struct.SetText(0, v)
}

func (s RegistrationOptions) Version() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s RegistrationOptions) HasVersion() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) VersionBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetVersion(v string) error {
	return s.Struct.SetText(1, v)
}

func (s RegistrationOptions) Os() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s RegistrationOptions) HasOs() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) OsBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetOs(v string) error {
	return s.Struct.SetText(2, v)
}

func (s RegistrationOptions) ExistingTunnelPolicy() ExistingTunnelPolicy {
	return ExistingTunnelPolicy(s.Struct.Uint16(0))
}

func (s RegistrationOptions) SetExistingTunnelPolicy(v ExistingTunnelPolicy) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s RegistrationOptions) PoolId() (string, error) {
	p, err := s.Struct.Ptr(3)
	return p.Text(), err
}

func (s RegistrationOptions) HasPoolId() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) PoolIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return p.TextBytes(), err
}

func (s RegistrationOptions) SetPoolId(v string) error {
	return s.Struct.SetText(3, v)
}

func (s RegistrationOptions) ExposeInternalHostname() bool {
	return s.Struct.Bit(16)
}

func (s RegistrationOptions) SetExposeInternalHostname(v bool) {
	s.Struct.SetBit(16, v)
}

func (s RegistrationOptions) Tags() (Tag_List, error) {
	p, err := s.Struct.Ptr(4)
	return Tag_List{List: p.List()}, err
}

func (s RegistrationOptions) HasTags() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s RegistrationOptions) SetTags(v Tag_List) error {
	return s.Struct.SetPtr(4, v.List.ToPtr())
}

// NewTags sets the tags field to a newly
// allocated Tag_List, preferring placement in s's segment.
func (s RegistrationOptions) NewTags(n int32) (Tag_List, error) {
	l, err := NewTag_List(s.Struct.Segment(), n)
	if err != nil {
		return Tag_List{}, err
	}
	err = s.Struct.SetPtr(4, l.List.ToPtr())
	return l, err
}

// RegistrationOptions_List is a list of RegistrationOptions.
type RegistrationOptions_List struct{ capnp.List }

// NewRegistrationOptions creates a new list of RegistrationOptions.
func NewRegistrationOptions_List(s *capnp.Segment, sz int32) (RegistrationOptions_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 5}, sz)
	return RegistrationOptions_List{l}, err
}

func (s RegistrationOptions_List) At(i int) RegistrationOptions {
	return RegistrationOptions{s.List.Struct(i)}
}

func (s RegistrationOptions_List) Set(i int, v RegistrationOptions) error {
	return s.List.SetStruct(i, v.Struct)
}

// RegistrationOptions_Promise is a wrapper for a RegistrationOptions promised by a client call.
type RegistrationOptions_Promise struct{ *capnp.Pipeline }

func (p RegistrationOptions_Promise) Struct() (RegistrationOptions, error) {
	s, err := p.Pipeline.Struct()
	return RegistrationOptions{s}, err
}

type Tag struct{ capnp.Struct }

// Tag_TypeID is the unique identifier for the type Tag.
const Tag_TypeID = 0xcbd96442ae3bb01a

func NewTag(s *capnp.Segment) (Tag, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tag{st}, err
}

func NewRootTag(s *capnp.Segment) (Tag, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return Tag{st}, err
}

func ReadRootTag(msg *capnp.Message) (Tag, error) {
	root, err := msg.RootPtr()
	return Tag{root.Struct()}, err
}

func (s Tag) String() string {
	str, _ := text.Marshal(0xcbd96442ae3bb01a, s.Struct)
	return str
}

func (s Tag) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Tag) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Tag) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Tag) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Tag) Value() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Tag) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Tag) ValueBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Tag) SetValue(v string) error {
	return s.Struct.SetText(1, v)
}

// Tag_List is a list of Tag.
type Tag_List struct{ capnp.List }

// NewTag creates a new list of Tag.
func NewTag_List(s *capnp.Segment, sz int32) (Tag_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return Tag_List{l}, err
}

func (s Tag_List) At(i int) Tag { return Tag{s.List.Struct(i)} }

func (s Tag_List) Set(i int, v Tag) error { return s.List.SetStruct(i, v.Struct) }

// Tag_Promise is a wrapper for a Tag promised by a client call.
type Tag_Promise struct{ *capnp.Pipeline }

func (p Tag_Promise) Struct() (Tag, error) {
	s, err := p.Pipeline.Struct()
	return Tag{s}, err
}

type ExistingTunnelPolicy uint16

// ExistingTunnelPolicy_TypeID is the unique identifier for the type ExistingTunnelPolicy.
const ExistingTunnelPolicy_TypeID = 0x84cb9536a2cf6d3c

// Values of ExistingTunnelPolicy.
const (
	ExistingTunnelPolicy_ignore     ExistingTunnelPolicy = 0
	ExistingTunnelPolicy_disconnect ExistingTunnelPolicy = 1
	ExistingTunnelPolicy_balance    ExistingTunnelPolicy = 2
)

// String returns the enum's constant name.
func (c ExistingTunnelPolicy) String() string {
	switch c {
	case ExistingTunnelPolicy_ignore:
		return "ignore"
	case ExistingTunnelPolicy_disconnect:
		return "disconnect"
	case ExistingTunnelPolicy_balance:
		return "balance"

	default:
		return ""
	}
}

// ExistingTunnelPolicyFromString returns the enum value with a name,
// or the zero value if there's no such value.
func ExistingTunnelPolicyFromString(c string) ExistingTunnelPolicy {
	switch c {
	case "ignore":
		return ExistingTunnelPolicy_ignore
	case "disconnect":
		return ExistingTunnelPolicy_disconnect
	case "balance":
		return ExistingTunnelPolicy_balance

	default:
		return 0
	}
}

type ExistingTunnelPolicy_List struct{ capnp.List }

func NewExistingTunnelPolicy_List(s *capnp.Segment, sz int32) (ExistingTunnelPolicy_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return ExistingTunnelPolicy_List{l.List}, err
}

func (l ExistingTunnelPolicy_List) At(i int) ExistingTunnelPolicy {
	ul := capnp.UInt16List{List: l.List}
	return ExistingTunnelPolicy(ul.At(i))
}

func (l ExistingTunnelPolicy_List) Set(i int, v ExistingTunnelPolicy) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type ServerInfo struct{ capnp.Struct }

// ServerInfo_TypeID is the unique identifier for the type ServerInfo.
const ServerInfo_TypeID = 0xf2c68e2547ec3866

func NewServerInfo(s *capnp.Segment) (ServerInfo, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ServerInfo{st}, err
}

func NewRootServerInfo(s *capnp.Segment) (ServerInfo, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return ServerInfo{st}, err
}

func ReadRootServerInfo(msg *capnp.Message) (ServerInfo, error) {
	root, err := msg.RootPtr()
	return ServerInfo{root.Struct()}, err
}

func (s ServerInfo) String() string {
	str, _ := text.Marshal(0xf2c68e2547ec3866, s.Struct)
	return str
}

func (s ServerInfo) LocationName() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ServerInfo) HasLocationName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ServerInfo) LocationNameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ServerInfo) SetLocationName(v string) error {
	return s.Struct.SetText(0, v)
}

// ServerInfo_List is a list of ServerInfo.
type ServerInfo_List struct{ capnp.List }

// NewServerInfo creates a new list of ServerInfo.
func NewServerInfo_List(s *capnp.Segment, sz int32) (ServerInfo_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return ServerInfo_List{l}, err
}

func (s ServerInfo_List) At(i int) ServerInfo { return ServerInfo{s.List.Struct(i)} }

func (s ServerInfo_List) Set(i int, v ServerInfo) error { return s.List.SetStruct(i, v.Struct) }

// ServerInfo_Promise is a wrapper for a ServerInfo promised by a client call.
type ServerInfo_Promise struct{ *capnp.Pipeline }

func (p ServerInfo_Promise) Struct() (ServerInfo, error) {
	s, err := p.Pipeline.Struct()
	return ServerInfo{s}, err
}

type TunnelServer struct{ Client capnp.Client }

// TunnelServer_TypeID is the unique identifier for the type TunnelServer.
const TunnelServer_TypeID = 0xea58385c65416035

func (c TunnelServer) RegisterTunnel(ctx context.Context, params func(TunnelServer_registerTunnel_Params) error, opts ...capnp.CallOption) TunnelServer_registerTunnel_Results_Promise {
	if c.Client == nil {
		return TunnelServer_registerTunnel_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      0,
			InterfaceName: "tunnelrpc.capnp:TunnelServer",
			MethodName:    "registerTunnel",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 3}
		call.ParamsFunc = func(s capnp.Struct) error { return params(TunnelServer_registerTunnel_Params{Struct: s}) }
	}
	return TunnelServer_registerTunnel_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}
func (c TunnelServer) GetServerInfo(ctx context.Context, params func(TunnelServer_getServerInfo_Params) error, opts ...capnp.CallOption) TunnelServer_getServerInfo_Results_Promise {
	if c.Client == nil {
		return TunnelServer_getServerInfo_Results_Promise{Pipeline: capnp.NewPipeline(capnp.ErrorAnswer(capnp.ErrNullClient))}
	}
	call := &capnp.Call{
		Ctx: ctx,
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      1,
			InterfaceName: "tunnelrpc.capnp:TunnelServer",
			MethodName:    "getServerInfo",
		},
		Options: capnp.NewCallOptions(opts),
	}
	if params != nil {
		call.ParamsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		call.ParamsFunc = func(s capnp.Struct) error { return params(TunnelServer_getServerInfo_Params{Struct: s}) }
	}
	return TunnelServer_getServerInfo_Results_Promise{Pipeline: capnp.NewPipeline(c.Client.Call(call))}
}

type TunnelServer_Server interface {
	RegisterTunnel(TunnelServer_registerTunnel) error

	GetServerInfo(TunnelServer_getServerInfo) error
}

func TunnelServer_ServerToClient(s TunnelServer_Server) TunnelServer {
	c, _ := s.(server.Closer)
	return TunnelServer{Client: server.New(TunnelServer_Methods(nil, s), c)}
}

func TunnelServer_Methods(methods []server.Method, s TunnelServer_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 2)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      0,
			InterfaceName: "tunnelrpc.capnp:TunnelServer",
			MethodName:    "registerTunnel",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := TunnelServer_registerTunnel{c, opts, TunnelServer_registerTunnel_Params{Struct: p}, TunnelServer_registerTunnel_Results{Struct: r}}
			return s.RegisterTunnel(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xea58385c65416035,
			MethodID:      1,
			InterfaceName: "tunnelrpc.capnp:TunnelServer",
			MethodName:    "getServerInfo",
		},
		Impl: func(c context.Context, opts capnp.CallOptions, p, r capnp.Struct) error {
			call := TunnelServer_getServerInfo{c, opts, TunnelServer_getServerInfo_Params{Struct: p}, TunnelServer_getServerInfo_Results{Struct: r}}
			return s.GetServerInfo(call)
		},
		ResultsSize: capnp.ObjectSize{DataSize: 0, PointerCount: 1},
	})

	return methods
}

// TunnelServer_registerTunnel holds the arguments for a server call to TunnelServer.registerTunnel.
type TunnelServer_registerTunnel struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  TunnelServer_registerTunnel_Params
	Results TunnelServer_registerTunnel_Results
}

// TunnelServer_getServerInfo holds the arguments for a server call to TunnelServer.getServerInfo.
type TunnelServer_getServerInfo struct {
	Ctx     context.Context
	Options capnp.CallOptions
	Params  TunnelServer_getServerInfo_Params
	Results TunnelServer_getServerInfo_Results
}

type TunnelServer_registerTunnel_Params struct{ capnp.Struct }

// TunnelServer_registerTunnel_Params_TypeID is the unique identifier for the type TunnelServer_registerTunnel_Params.
const TunnelServer_registerTunnel_Params_TypeID = 0xb70431c0dc014915

func NewTunnelServer_registerTunnel_Params(s *capnp.Segment) (TunnelServer_registerTunnel_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return TunnelServer_registerTunnel_Params{st}, err
}

func NewRootTunnelServer_registerTunnel_Params(s *capnp.Segment) (TunnelServer_registerTunnel_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3})
	return TunnelServer_registerTunnel_Params{st}, err
}

func ReadRootTunnelServer_registerTunnel_Params(msg *capnp.Message) (TunnelServer_registerTunnel_Params, error) {
	root, err := msg.RootPtr()
	return TunnelServer_registerTunnel_Params{root.Struct()}, err
}

func (s TunnelServer_registerTunnel_Params) String() string {
	str, _ := text.Marshal(0xb70431c0dc014915, s.Struct)
	return str
}

func (s TunnelServer_registerTunnel_Params) Auth() (Authentication, error) {
	p, err := s.Struct.Ptr(0)
	return Authentication{Struct: p.Struct()}, err
}

func (s TunnelServer_registerTunnel_Params) HasAuth() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) SetAuth(v Authentication) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewAuth sets the auth field to a newly
// allocated Authentication struct, preferring placement in s's segment.
func (s TunnelServer_registerTunnel_Params) NewAuth() (Authentication, error) {
	ss, err := NewAuthentication(s.Struct.Segment())
	if err != nil {
		return Authentication{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s TunnelServer_registerTunnel_Params) Hostname() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s TunnelServer_registerTunnel_Params) HasHostname() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) HostnameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s TunnelServer_registerTunnel_Params) SetHostname(v string) error {
	return s.Struct.SetText(1, v)
}

func (s TunnelServer_registerTunnel_Params) Options() (RegistrationOptions, error) {
	p, err := s.Struct.Ptr(2)
	return RegistrationOptions{Struct: p.Struct()}, err
}

func (s TunnelServer_registerTunnel_Params) HasOptions() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Params) SetOptions(v RegistrationOptions) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewOptions sets the options field to a newly
// allocated RegistrationOptions struct, preferring placement in s's segment.
func (s TunnelServer_registerTunnel_Params) NewOptions() (RegistrationOptions, error) {
	ss, err := NewRegistrationOptions(s.Struct.Segment())
	if err != nil {
		return RegistrationOptions{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_registerTunnel_Params_List is a list of TunnelServer_registerTunnel_Params.
type TunnelServer_registerTunnel_Params_List struct{ capnp.List }

// NewTunnelServer_registerTunnel_Params creates a new list of TunnelServer_registerTunnel_Params.
func NewTunnelServer_registerTunnel_Params_List(s *capnp.Segment, sz int32) (TunnelServer_registerTunnel_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 3}, sz)
	return TunnelServer_registerTunnel_Params_List{l}, err
}

func (s TunnelServer_registerTunnel_Params_List) At(i int) TunnelServer_registerTunnel_Params {
	return TunnelServer_registerTunnel_Params{s.List.Struct(i)}
}

func (s TunnelServer_registerTunnel_Params_List) Set(i int, v TunnelServer_registerTunnel_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_registerTunnel_Params_Promise is a wrapper for a TunnelServer_registerTunnel_Params promised by a client call.
type TunnelServer_registerTunnel_Params_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_registerTunnel_Params_Promise) Struct() (TunnelServer_registerTunnel_Params, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_registerTunnel_Params{s}, err
}

func (p TunnelServer_registerTunnel_Params_Promise) Auth() Authentication_Promise {
	return Authentication_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

func (p TunnelServer_registerTunnel_Params_Promise) Options() RegistrationOptions_Promise {
	return RegistrationOptions_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

type TunnelServer_registerTunnel_Results struct{ capnp.Struct }

// TunnelServer_registerTunnel_Results_TypeID is the unique identifier for the type TunnelServer_registerTunnel_Results.
const TunnelServer_registerTunnel_Results_TypeID = 0xf2c122394f447e8e

func NewTunnelServer_registerTunnel_Results(s *capnp.Segment) (TunnelServer_registerTunnel_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_registerTunnel_Results{st}, err
}

func NewRootTunnelServer_registerTunnel_Results(s *capnp.Segment) (TunnelServer_registerTunnel_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_registerTunnel_Results{st}, err
}

func ReadRootTunnelServer_registerTunnel_Results(msg *capnp.Message) (TunnelServer_registerTunnel_Results, error) {
	root, err := msg.RootPtr()
	return TunnelServer_registerTunnel_Results{root.Struct()}, err
}

func (s TunnelServer_registerTunnel_Results) String() string {
	str, _ := text.Marshal(0xf2c122394f447e8e, s.Struct)
	return str
}

func (s TunnelServer_registerTunnel_Results) Result() (TunnelRegistration, error) {
	p, err := s.Struct.Ptr(0)
	return TunnelRegistration{Struct: p.Struct()}, err
}

func (s TunnelServer_registerTunnel_Results) HasResult() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_registerTunnel_Results) SetResult(v TunnelRegistration) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResult sets the result field to a newly
// allocated TunnelRegistration struct, preferring placement in s's segment.
func (s TunnelServer_registerTunnel_Results) NewResult() (TunnelRegistration, error) {
	ss, err := NewTunnelRegistration(s.Struct.Segment())
	if err != nil {
		return TunnelRegistration{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_registerTunnel_Results_List is a list of TunnelServer_registerTunnel_Results.
type TunnelServer_registerTunnel_Results_List struct{ capnp.List }

// NewTunnelServer_registerTunnel_Results creates a new list of TunnelServer_registerTunnel_Results.
func NewTunnelServer_registerTunnel_Results_List(s *capnp.Segment, sz int32) (TunnelServer_registerTunnel_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TunnelServer_registerTunnel_Results_List{l}, err
}

func (s TunnelServer_registerTunnel_Results_List) At(i int) TunnelServer_registerTunnel_Results {
	return TunnelServer_registerTunnel_Results{s.List.Struct(i)}
}

func (s TunnelServer_registerTunnel_Results_List) Set(i int, v TunnelServer_registerTunnel_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_registerTunnel_Results_Promise is a wrapper for a TunnelServer_registerTunnel_Results promised by a client call.
type TunnelServer_registerTunnel_Results_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_registerTunnel_Results_Promise) Struct() (TunnelServer_registerTunnel_Results, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_registerTunnel_Results{s}, err
}

func (p TunnelServer_registerTunnel_Results_Promise) Result() TunnelRegistration_Promise {
	return TunnelRegistration_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type TunnelServer_getServerInfo_Params struct{ capnp.Struct }

// TunnelServer_getServerInfo_Params_TypeID is the unique identifier for the type TunnelServer_getServerInfo_Params.
const TunnelServer_getServerInfo_Params_TypeID = 0xdc3ed6801961e502

func NewTunnelServer_getServerInfo_Params(s *capnp.Segment) (TunnelServer_getServerInfo_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TunnelServer_getServerInfo_Params{st}, err
}

func NewRootTunnelServer_getServerInfo_Params(s *capnp.Segment) (TunnelServer_getServerInfo_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return TunnelServer_getServerInfo_Params{st}, err
}

func ReadRootTunnelServer_getServerInfo_Params(msg *capnp.Message) (TunnelServer_getServerInfo_Params, error) {
	root, err := msg.RootPtr()
	return TunnelServer_getServerInfo_Params{root.Struct()}, err
}

func (s TunnelServer_getServerInfo_Params) String() string {
	str, _ := text.Marshal(0xdc3ed6801961e502, s.Struct)
	return str
}

// TunnelServer_getServerInfo_Params_List is a list of TunnelServer_getServerInfo_Params.
type TunnelServer_getServerInfo_Params_List struct{ capnp.List }

// NewTunnelServer_getServerInfo_Params creates a new list of TunnelServer_getServerInfo_Params.
func NewTunnelServer_getServerInfo_Params_List(s *capnp.Segment, sz int32) (TunnelServer_getServerInfo_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return TunnelServer_getServerInfo_Params_List{l}, err
}

func (s TunnelServer_getServerInfo_Params_List) At(i int) TunnelServer_getServerInfo_Params {
	return TunnelServer_getServerInfo_Params{s.List.Struct(i)}
}

func (s TunnelServer_getServerInfo_Params_List) Set(i int, v TunnelServer_getServerInfo_Params) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_getServerInfo_Params_Promise is a wrapper for a TunnelServer_getServerInfo_Params promised by a client call.
type TunnelServer_getServerInfo_Params_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_getServerInfo_Params_Promise) Struct() (TunnelServer_getServerInfo_Params, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_getServerInfo_Params{s}, err
}

type TunnelServer_getServerInfo_Results struct{ capnp.Struct }

// TunnelServer_getServerInfo_Results_TypeID is the unique identifier for the type TunnelServer_getServerInfo_Results.
const TunnelServer_getServerInfo_Results_TypeID = 0xe3e37d096a5b564e

func NewTunnelServer_getServerInfo_Results(s *capnp.Segment) (TunnelServer_getServerInfo_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_getServerInfo_Results{st}, err
}

func NewRootTunnelServer_getServerInfo_Results(s *capnp.Segment) (TunnelServer_getServerInfo_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return TunnelServer_getServerInfo_Results{st}, err
}

func ReadRootTunnelServer_getServerInfo_Results(msg *capnp.Message) (TunnelServer_getServerInfo_Results, error) {
	root, err := msg.RootPtr()
	return TunnelServer_getServerInfo_Results{root.Struct()}, err
}

func (s TunnelServer_getServerInfo_Results) String() string {
	str, _ := text.Marshal(0xe3e37d096a5b564e, s.Struct)
	return str
}

func (s TunnelServer_getServerInfo_Results) Result() (ServerInfo, error) {
	p, err := s.Struct.Ptr(0)
	return ServerInfo{Struct: p.Struct()}, err
}

func (s TunnelServer_getServerInfo_Results) HasResult() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s TunnelServer_getServerInfo_Results) SetResult(v ServerInfo) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResult sets the result field to a newly
// allocated ServerInfo struct, preferring placement in s's segment.
func (s TunnelServer_getServerInfo_Results) NewResult() (ServerInfo, error) {
	ss, err := NewServerInfo(s.Struct.Segment())
	if err != nil {
		return ServerInfo{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// TunnelServer_getServerInfo_Results_List is a list of TunnelServer_getServerInfo_Results.
type TunnelServer_getServerInfo_Results_List struct{ capnp.List }

// NewTunnelServer_getServerInfo_Results creates a new list of TunnelServer_getServerInfo_Results.
func NewTunnelServer_getServerInfo_Results_List(s *capnp.Segment, sz int32) (TunnelServer_getServerInfo_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return TunnelServer_getServerInfo_Results_List{l}, err
}

func (s TunnelServer_getServerInfo_Results_List) At(i int) TunnelServer_getServerInfo_Results {
	return TunnelServer_getServerInfo_Results{s.List.Struct(i)}
}

func (s TunnelServer_getServerInfo_Results_List) Set(i int, v TunnelServer_getServerInfo_Results) error {
	return s.List.SetStruct(i, v.Struct)
}

// TunnelServer_getServerInfo_Results_Promise is a wrapper for a TunnelServer_getServerInfo_Results promised by a client call.
type TunnelServer_getServerInfo_Results_Promise struct{ *capnp.Pipeline }

func (p TunnelServer_getServerInfo_Results_Promise) Struct() (TunnelServer_getServerInfo_Results, error) {
	s, err := p.Pipeline.Struct()
	return TunnelServer_getServerInfo_Results{s}, err
}

func (p TunnelServer_getServerInfo_Results_Promise) Result() ServerInfo_Promise {
	return ServerInfo_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_db8274f9144abc7e = "x\xda\x94To\x88\x14e\x18\x7f~\xef\xbb3\xabp" +
	"\xb6;\xcc\x0aut\x08b\x90\x82\xe6e\x86\x99\xb4\xe7" +
	"\xa5\xe6^\xa7\xb7\xaf]a\xe9\x07\xc7\xbd\xf7\xf6\xc6f" +
	"g\xb6\x99\xd9\xcb\x8b\xd4\x92 \x0c2R\xfb\xd2\x87\xc8" +
	"\xfbf`%\x14E\x18\x9c\xd0\x1f\xc1\"\x02\x0b\xad\xeb" +
	"C\x98\x04RH\xd2\x87\x0cb\xe2\x9d\xbd\xd9\x99\xee\x94" +
	"\xf0\xdb\xfb\xe7y\x7f\xef\xef\xf9=\xcf\xf3[9\xc1\xfa" +
	"X\xafv\x8fF$\xd6iz\xb4\xae\xf1\xcd\xe4\xfdo" +
	"\x9c{\x89\x8c\"\x8b\xf6\x9f\x1e(]\x0f\x0f\xfeH\x84" +
	"U\xfb\xd82\x98\xaf\xb2<\x91y\x88\x0d\x11\xa2\x85\x15" +
	"LO\xf5\xe6>\"\xa3\x07D\x1a\xcf\x13\xad:\xce\xde" +
	"\x04\xc1<\xc5\xde#D=\xbf\xf7/p\xaf\x1e\x9c\"" +
	"\xa3\x88\x14*\x0e4\xb7\xf0\xbf\xcd'\xe3\xd5\xe3\\\xc5" +
	"\x0e\xec8zD\xbb|\xf4K\x12Ed\x835\x85\xfa" +
	"\x07_\x0c\x139\xb5\xfc\x87\xbf\x06B\xd4\xfd\xfe\x83\xef" +
	"\xf6\x8f\\<7\x0b:fwB\x9b4O\xa9w\xe6" +
	"I\xedYB\xc4.[w\xbc\xf0\xfdC\xd3m\x9e1" +
	"\xca|\xfd\x08(\x17m}b\xc7\x9e\xf9\xfb.]\x9a" +
	"\xc9\x00\xea\xea\xba\x16g0_/\x13\xa2\xd5\xbb\xd6\xcb" +
	"\x9dk\xb6_!\xa3\xc8\xb3b\x98K\xf5+\xe6j]" +
	"\xfd\xd1\xab\xbfl\x1eR\xab\xe8\xf0\xfe\x0dC\x0f,>" +
	"s-\x8b\xf6\x8c>\xa9\xd0^\x8c\xd1F\xd7\xfc\xf6\xc8" +
	"]\x87\xbf\xb86\x8b\xb4\x0a4\x8f\xeb?\x98'c\xc0" +
	"\x13*\xf6\xea\xa6\xb7\xcew\x17\xba\xff\x9c\xa5F\xac\xf1" +
	"\xd7z7\xcc\x9f\xe2\xd8\x8b\xfa\xaf\xb4<\x0a[\xae+" +
	"\x1d\xbf\xc9k+jV\xd3m\xae\xdd\xb8\xd7\x0eB\xdb" +
	"\xad\x0f\xc7\x17U\xaf\xe0\xd8\xb5\x89* \xba\xc0\x88\x8c" +
	"\x9e\xb5D\x80\xb1\xf0)\"0\xc3\xe8'*\xdbu\xd7" +
	"\xf3e4b\x075\xcfu%\xf1Zx`\xb7\xe5X" +
	"nMv\xe0\xb5\x04\xbe\x0d\xfb\x98\xf4\xc7\xa5\xbf\xc2\x97" +
	"u;\x08\xa5\xdf>\\R\xb5|\x8b7\x02\xd1\xc5s" +
	"D9\x10\x19\x1b\x97\x11\x89>\x0e1\xc8`\x00%\xa8" +
	"\xc3\xca\x00\x91\xd8\xcc!\x86\x19\x0c\xc6J1/\xd1O" +
	"$\x069\xc4v\x86\x82\xd5\x0a\xc7PL{\x88\x80\"" +
	"!\x1a\xf3\x82\xd0\xb5\x1a\x92\x88\xd0E\x0c]\x84\x03^" +
	"3\xb4=7@1\xed\xa2\x99\xe8\x84:K\xa8\xafo" +
	"\x85c\xd2\x0d\xedr\xcdRobMR\xa6\x8bo\xc4" +
	"\xf4^\"\xb1\x81CT3L\xb7\xecN\x99\xe6\x9f\x96" +
	"\x13\x09\x95E\xb2a\xd9N\xb2\x8b<\xdf\xae\xdb\xee\xc3" +
	"\xeb)\xffh\x1a3\xb7\\\xdbb\x09\xfd\x98\xd1P3" +
	"\xb4\xf3\x9e\x1b(fwv\x98}\xa8\xe4\xfa\x80CL" +
	"e\x98}\xaa\xe4\xfa\x98C|\x96av\xa6\x9bH\x9c" +
	"\xe6\x10g\x19\xc0K\xe0D\xc6\xe7\xef\x10\x89\xb3\x1c\xe2" +
	"<\x83\x91\xe3%\xe4\x88\x8co\xd7\x12\x89\xaf8\xc4\x05" +
	"\x06C+\x96\xa0\x11\x19\xdf}B$.p\x88_\x18" +
	"\x0c=W\x82Nd\xfc\xac\x0a8\xcd!\xfeb\x88j" +
	"\x8e-\xdd\xb02\x92\xd5\x7f\\\xfa\x81\xed\xb9\xc9\x9e{" +
	"A'W9\xd3\x89\xf8O+\xa2\x90\xda\x0c\x01\x05B" +
	"\xb9\xe9yNe$\xf3\xae\xe9\x05\xb2\xe2\"\x94\xbek" +
	"9\x9b\xbdr\xbb\xec\x001\x80P\x08\xadz\x80\xdb\x08" +
	"U\x0e\x14S; \xa8\xc3\x8e\xc4H$\xce\x0f[u" +
	"%\xe9\xbc\x8e\xa4KUVK8\xc4\xca\x8c\xa4\xcbU" +
	"\xb1\xef\xe6\x10\xf71\x14\xe2\xff\x92\xc2\x8e[NK\xce" +
	")\xe1\x8dG\xa2.\xc3\xf6\xaa\xe2\x8ez\xf1D4\x10" +
	"\xdc\xd2\x9bm2h9<\x0cD\xae\xc3w\x81\xaa\xd7" +
	"<\x0eQb(\xfb\xea>D1\xb5\x94\x9b5|\xf2" +
	"IAa\xb7\x15\xd0\x88:\xde\x8d\xc4\xb4\x8c\xde\xe7\x88" +
	"\x19K\xf3H\xfd\x12\x89=\x1a=>1ca>J" +
	"f\x9d\xcam\xd8>D\x09oZ\x143\xefC\x15\xb8" +
	"5\xc7P\xb9\xe6\x9d\xff\xcf5\xb1\xc4\x9be\x9a\xc8\xc7" +
	"G=\x95g\x06m\x0f\x91\xe8\xe2\x10\xb73D\x8e\xd7" +
	"\x9e|*l\xcd\x94w\xeeL\xb6\xc9\xa5\x93\xc9\xdbf" +
	"Q\xec\xa0Z\xca,vr\x88\xb1L\xffH\xd5T\xbb" +
	"8\xc4\xf3\x99\x91\x9cP\xc3\xbb\x97C\x1cKG\xf2\xf5" +
	"W\x88\xc41\x0e\xf16C^\xfa~B\xa4\xd0\xf2\x9d" +
	"N_\xab3\xd5\xcd\x8eW\x1f\xb4]\x19\xa8\x99\x9bu" +
	"\xd5\x94~\xc3r\xa5\x8bp\x93e;-_\xaaFh" +
	"\x8f\xc8\xbf\x01\x00\x00\xff\xff\x80\xf4\x060"

func init() {
	schemas.Register(schema_db8274f9144abc7e,
		0x84cb9536a2cf6d3c,
		0xb70431c0dc014915,
		0xc082ef6e0d42ed1d,
		0xc793e50592935b4a,
		0xcbd96442ae3bb01a,
		0xdc3ed6801961e502,
		0xe3e37d096a5b564e,
		0xea58385c65416035,
		0xf2c122394f447e8e,
		0xf2c68e2547ec3866,
		0xf41a0f001ad49e46)
}