// Code generated by capnpc-go. DO NOT EDIT.

package host

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	server "capnproto.org/go/capnp/v3/server"
	context "context"
)

type Negotiator struct{ Client *capnp.Client }

// Negotiator_TypeID is the unique identifier for the type Negotiator.
const Negotiator_TypeID = 0xe35a52b4e5c60a15

func (c Negotiator) ConfigService(ctx context.Context, params func(Negotiator_configService_Params) error) (Negotiator_configService_Results_Future, capnp.ReleaseFunc) {
	s := capnp.Send{
		Method: capnp.Method{
			InterfaceID:   0xe35a52b4e5c60a15,
			MethodID:      0,
			InterfaceName: "protocols/v1/host/negotiator.capnp:Negotiator",
			MethodName:    "configService",
		},
	}
	if params != nil {
		s.ArgsSize = capnp.ObjectSize{DataSize: 0, PointerCount: 0}
		s.PlaceArgs = func(s capnp.Struct) error { return params(Negotiator_configService_Params{Struct: s}) }
	}
	ans, release := c.Client.SendCall(ctx, s)
	return Negotiator_configService_Results_Future{Future: ans.Future()}, release
}

func (c Negotiator) AddRef() Negotiator {
	return Negotiator{
		Client: c.Client.AddRef(),
	}
}

func (c Negotiator) Release() {
	c.Client.Release()
}

// A Negotiator_Server is a Negotiator with a local implementation.
type Negotiator_Server interface {
	ConfigService(context.Context, Negotiator_configService) error
}

// Negotiator_NewServer creates a new Server from an implementation of Negotiator_Server.
func Negotiator_NewServer(s Negotiator_Server, policy *server.Policy) *server.Server {
	c, _ := s.(server.Shutdowner)
	return server.New(Negotiator_Methods(nil, s), s, c, policy)
}

// Negotiator_ServerToClient creates a new Client from an implementation of Negotiator_Server.
// The caller is responsible for calling Release on the returned Client.
func Negotiator_ServerToClient(s Negotiator_Server, policy *server.Policy) Negotiator {
	return Negotiator{Client: capnp.NewClient(Negotiator_NewServer(s, policy))}
}

// Negotiator_Methods appends Methods to a slice that invoke the methods on s.
// This can be used to create a more complicated Server.
func Negotiator_Methods(methods []server.Method, s Negotiator_Server) []server.Method {
	if cap(methods) == 0 {
		methods = make([]server.Method, 0, 1)
	}

	methods = append(methods, server.Method{
		Method: capnp.Method{
			InterfaceID:   0xe35a52b4e5c60a15,
			MethodID:      0,
			InterfaceName: "protocols/v1/host/negotiator.capnp:Negotiator",
			MethodName:    "configService",
		},
		Impl: func(ctx context.Context, call *server.Call) error {
			return s.ConfigService(ctx, Negotiator_configService{call})
		},
	})

	return methods
}

// Negotiator_configService holds the state for a server call to Negotiator.configService.
// See server.Call for documentation.
type Negotiator_configService struct {
	*server.Call
}

// Args returns the call's arguments.
func (c Negotiator_configService) Args() Negotiator_configService_Params {
	return Negotiator_configService_Params{Struct: c.Call.Args()}
}

// AllocResults allocates the results struct.
func (c Negotiator_configService) AllocResults() (Negotiator_configService_Results, error) {
	r, err := c.Call.AllocResults(capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Negotiator_configService_Results{Struct: r}, err
}

type Negotiator_configService_Params struct{ capnp.Struct }

// Negotiator_configService_Params_TypeID is the unique identifier for the type Negotiator_configService_Params.
const Negotiator_configService_Params_TypeID = 0xf9b86cbf3f6febfb

func NewNegotiator_configService_Params(s *capnp.Segment) (Negotiator_configService_Params, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Negotiator_configService_Params{st}, err
}

func NewRootNegotiator_configService_Params(s *capnp.Segment) (Negotiator_configService_Params, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0})
	return Negotiator_configService_Params{st}, err
}

func ReadRootNegotiator_configService_Params(msg *capnp.Message) (Negotiator_configService_Params, error) {
	root, err := msg.Root()
	return Negotiator_configService_Params{root.Struct()}, err
}

func (s Negotiator_configService_Params) String() string {
	str, _ := text.Marshal(0xf9b86cbf3f6febfb, s.Struct)
	return str
}

// Negotiator_configService_Params_List is a list of Negotiator_configService_Params.
type Negotiator_configService_Params_List = capnp.StructList[Negotiator_configService_Params]

// NewNegotiator_configService_Params creates a new list of Negotiator_configService_Params.
func NewNegotiator_configService_Params_List(s *capnp.Segment, sz int32) (Negotiator_configService_Params_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 0}, sz)
	return capnp.StructList[Negotiator_configService_Params]{List: l}, err
}

// Negotiator_configService_Params_Future is a wrapper for a Negotiator_configService_Params promised by a client call.
type Negotiator_configService_Params_Future struct{ *capnp.Future }

func (p Negotiator_configService_Params_Future) Struct() (Negotiator_configService_Params, error) {
	s, err := p.Future.Struct()
	return Negotiator_configService_Params{s}, err
}

type Negotiator_configService_Results struct{ capnp.Struct }

// Negotiator_configService_Results_TypeID is the unique identifier for the type Negotiator_configService_Results.
const Negotiator_configService_Results_TypeID = 0xd663f8eebd0e079f

func NewNegotiator_configService_Results(s *capnp.Segment) (Negotiator_configService_Results, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Negotiator_configService_Results{st}, err
}

func NewRootNegotiator_configService_Results(s *capnp.Segment) (Negotiator_configService_Results, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Negotiator_configService_Results{st}, err
}

func ReadRootNegotiator_configService_Results(msg *capnp.Message) (Negotiator_configService_Results, error) {
	root, err := msg.Root()
	return Negotiator_configService_Results{root.Struct()}, err
}

func (s Negotiator_configService_Results) String() string {
	str, _ := text.Marshal(0xd663f8eebd0e079f, s.Struct)
	return str
}

func (s Negotiator_configService_Results) Svc() ConfigService {
	p, _ := s.Struct.Ptr(0)
	return ConfigService{Client: p.Interface().Client()}
}

func (s Negotiator_configService_Results) HasSvc() bool {
	return s.Struct.HasPtr(0)
}

func (s Negotiator_configService_Results) SetSvc(v ConfigService) error {
	if !v.Client.IsValid() {
		return s.Struct.SetPtr(0, capnp.Ptr{})
	}
	seg := s.Segment()
	in := capnp.NewInterface(seg, seg.Message().AddCap(v.Client))
	return s.Struct.SetPtr(0, in.ToPtr())
}

// Negotiator_configService_Results_List is a list of Negotiator_configService_Results.
type Negotiator_configService_Results_List = capnp.StructList[Negotiator_configService_Results]

// NewNegotiator_configService_Results creates a new list of Negotiator_configService_Results.
func NewNegotiator_configService_Results_List(s *capnp.Segment, sz int32) (Negotiator_configService_Results_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return capnp.StructList[Negotiator_configService_Results]{List: l}, err
}

// Negotiator_configService_Results_Future is a wrapper for a Negotiator_configService_Results promised by a client call.
type Negotiator_configService_Results_Future struct{ *capnp.Future }

func (p Negotiator_configService_Results_Future) Struct() (Negotiator_configService_Results, error) {
	s, err := p.Future.Struct()
	return Negotiator_configService_Results{s}, err
}

func (p Negotiator_configService_Results_Future) Svc() ConfigService {
	return ConfigService{Client: p.Future.Field(0, nil).Client()}
}

const schema_90858bf9ae63c319 = "x\xda\xa4\x90!HCA\x1c\xc6\xbf\xff{w;\xa7" +
	"L8\x9e\x06-\x83i\xde\x1ck\x96\xad\xd9\xc6\xde\xc4" +
	"b\x91\xc7\xf1\x9c\x83\xb9\x1bw\xcf\xd9E\x93\xc5b\x16" +
	"\x93EM\x06\x8b(\x08\xb6U\x9bA\x86I\x10l\x0e" +
	"\xc3\x89\x93\xc9\x9a\xc1\xfa\x85\xef\xfb\xfd\xbe\xa53\xaa\xb0" +
	"bfc\x0a^\xb8\xcfS\xeeDL\xdf\xbc}\xa8G" +
	"\xc8<\x01\x9c\x04P:\xe69\x0f\x14\\\xf02\xc8\xcd" +
	"N>\xbc\\\xd5\xd7\xfb\x90\x0b\xbe\x9b\xbbW\x97\x83\xc3" +
	"\x83#\x80J=n(x\xe6\x02\x08\x9e\xf8J\x90N" +
	"\x09\xc0}\xbe\xea\xf2m\xebz\xf0S\xc7\xbe\xdb\xde\xf9" +
	"\xbc\x07\x87\xaa\xeb\x18\x9dh\xa5[\xc2\x16\xba\xc5\xc2\x96" +
	"\xb6I\xa1\x1d7t\xd2\x8c\x12m\xf2*\xea\xb4;\xcb" +
	"\xd5\xb1@\xb77\x9b\x8d\xd5\xd8t\x9b*^\xac\xc7v" +
	"G\xb4\x12\x1b2\x9f\x01\x8c\x00\x99\xc9\x01\xe1\x84O\xe1" +
	"\x8cG\xc2v\x15I\xb7{\xbewz\xd7_\xeb\x01D" +
	"\x12\xf4;\xc9\xfe\x9a\xcc\x0e\x93\x1aQ\xc8|>\xe6A" +
	"\xa3\x7f\xa44\xf0dZ\xb8\x11\x16\xb2C\xb0\x0a\xd5\x88" +
	"\xfe\xa7V\x8bL\xe4o\xdb\xaf\x00\x00\x00\xff\xff\xd2\x90" +
	"\x87n"

func init() {
	schemas.Register(schema_90858bf9ae63c319,
		0xd663f8eebd0e079f,
		0xe35a52b4e5c60a15,
		0xf9b86cbf3f6febfb)
}
