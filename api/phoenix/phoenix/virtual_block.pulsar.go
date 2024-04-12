// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package phoenix

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_VirtualBlock            protoreflect.MessageDescriptor
	fd_VirtualBlock_id         protoreflect.FieldDescriptor
	fd_VirtualBlock_time       protoreflect.FieldDescriptor
	fd_VirtualBlock_header     protoreflect.FieldDescriptor
	fd_VirtualBlock_previous   protoreflect.FieldDescriptor
	fd_VirtualBlock_parent     protoreflect.FieldDescriptor
	fd_VirtualBlock_leftChild  protoreflect.FieldDescriptor
	fd_VirtualBlock_rightChild protoreflect.FieldDescriptor
	fd_VirtualBlock_vdfProof   protoreflect.FieldDescriptor
)

func init() {
	file_phoenix_phoenix_virtual_block_proto_init()
	md_VirtualBlock = File_phoenix_phoenix_virtual_block_proto.Messages().ByName("VirtualBlock")
	fd_VirtualBlock_id = md_VirtualBlock.Fields().ByName("id")
	fd_VirtualBlock_time = md_VirtualBlock.Fields().ByName("time")
	fd_VirtualBlock_header = md_VirtualBlock.Fields().ByName("header")
	fd_VirtualBlock_previous = md_VirtualBlock.Fields().ByName("previous")
	fd_VirtualBlock_parent = md_VirtualBlock.Fields().ByName("parent")
	fd_VirtualBlock_leftChild = md_VirtualBlock.Fields().ByName("leftChild")
	fd_VirtualBlock_rightChild = md_VirtualBlock.Fields().ByName("rightChild")
	fd_VirtualBlock_vdfProof = md_VirtualBlock.Fields().ByName("vdfProof")
}

var _ protoreflect.Message = (*fastReflection_VirtualBlock)(nil)

type fastReflection_VirtualBlock VirtualBlock

func (x *VirtualBlock) ProtoReflect() protoreflect.Message {
	return (*fastReflection_VirtualBlock)(x)
}

func (x *VirtualBlock) slowProtoReflect() protoreflect.Message {
	mi := &file_phoenix_phoenix_virtual_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_VirtualBlock_messageType fastReflection_VirtualBlock_messageType
var _ protoreflect.MessageType = fastReflection_VirtualBlock_messageType{}

type fastReflection_VirtualBlock_messageType struct{}

func (x fastReflection_VirtualBlock_messageType) Zero() protoreflect.Message {
	return (*fastReflection_VirtualBlock)(nil)
}
func (x fastReflection_VirtualBlock_messageType) New() protoreflect.Message {
	return new(fastReflection_VirtualBlock)
}
func (x fastReflection_VirtualBlock_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_VirtualBlock
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_VirtualBlock) Descriptor() protoreflect.MessageDescriptor {
	return md_VirtualBlock
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_VirtualBlock) Type() protoreflect.MessageType {
	return _fastReflection_VirtualBlock_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_VirtualBlock) New() protoreflect.Message {
	return new(fastReflection_VirtualBlock)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_VirtualBlock) Interface() protoreflect.ProtoMessage {
	return (*VirtualBlock)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_VirtualBlock) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_VirtualBlock_id, value) {
			return
		}
	}
	if x.Time != int64(0) {
		value := protoreflect.ValueOfInt64(x.Time)
		if !f(fd_VirtualBlock_time, value) {
			return
		}
	}
	if len(x.Header) != 0 {
		value := protoreflect.ValueOfBytes(x.Header)
		if !f(fd_VirtualBlock_header, value) {
			return
		}
	}
	if len(x.Previous) != 0 {
		value := protoreflect.ValueOfBytes(x.Previous)
		if !f(fd_VirtualBlock_previous, value) {
			return
		}
	}
	if len(x.Parent) != 0 {
		value := protoreflect.ValueOfBytes(x.Parent)
		if !f(fd_VirtualBlock_parent, value) {
			return
		}
	}
	if len(x.LeftChild) != 0 {
		value := protoreflect.ValueOfBytes(x.LeftChild)
		if !f(fd_VirtualBlock_leftChild, value) {
			return
		}
	}
	if len(x.RightChild) != 0 {
		value := protoreflect.ValueOfBytes(x.RightChild)
		if !f(fd_VirtualBlock_rightChild, value) {
			return
		}
	}
	if len(x.VdfProof) != 0 {
		value := protoreflect.ValueOfBytes(x.VdfProof)
		if !f(fd_VirtualBlock_vdfProof, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_VirtualBlock) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		return x.Id != uint64(0)
	case "phoenix.phoenix.VirtualBlock.time":
		return x.Time != int64(0)
	case "phoenix.phoenix.VirtualBlock.header":
		return len(x.Header) != 0
	case "phoenix.phoenix.VirtualBlock.previous":
		return len(x.Previous) != 0
	case "phoenix.phoenix.VirtualBlock.parent":
		return len(x.Parent) != 0
	case "phoenix.phoenix.VirtualBlock.leftChild":
		return len(x.LeftChild) != 0
	case "phoenix.phoenix.VirtualBlock.rightChild":
		return len(x.RightChild) != 0
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		return len(x.VdfProof) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VirtualBlock) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		x.Id = uint64(0)
	case "phoenix.phoenix.VirtualBlock.time":
		x.Time = int64(0)
	case "phoenix.phoenix.VirtualBlock.header":
		x.Header = nil
	case "phoenix.phoenix.VirtualBlock.previous":
		x.Previous = nil
	case "phoenix.phoenix.VirtualBlock.parent":
		x.Parent = nil
	case "phoenix.phoenix.VirtualBlock.leftChild":
		x.LeftChild = nil
	case "phoenix.phoenix.VirtualBlock.rightChild":
		x.RightChild = nil
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		x.VdfProof = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_VirtualBlock) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "phoenix.phoenix.VirtualBlock.time":
		value := x.Time
		return protoreflect.ValueOfInt64(value)
	case "phoenix.phoenix.VirtualBlock.header":
		value := x.Header
		return protoreflect.ValueOfBytes(value)
	case "phoenix.phoenix.VirtualBlock.previous":
		value := x.Previous
		return protoreflect.ValueOfBytes(value)
	case "phoenix.phoenix.VirtualBlock.parent":
		value := x.Parent
		return protoreflect.ValueOfBytes(value)
	case "phoenix.phoenix.VirtualBlock.leftChild":
		value := x.LeftChild
		return protoreflect.ValueOfBytes(value)
	case "phoenix.phoenix.VirtualBlock.rightChild":
		value := x.RightChild
		return protoreflect.ValueOfBytes(value)
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		value := x.VdfProof
		return protoreflect.ValueOfBytes(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VirtualBlock) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		x.Id = value.Uint()
	case "phoenix.phoenix.VirtualBlock.time":
		x.Time = value.Int()
	case "phoenix.phoenix.VirtualBlock.header":
		x.Header = value.Bytes()
	case "phoenix.phoenix.VirtualBlock.previous":
		x.Previous = value.Bytes()
	case "phoenix.phoenix.VirtualBlock.parent":
		x.Parent = value.Bytes()
	case "phoenix.phoenix.VirtualBlock.leftChild":
		x.LeftChild = value.Bytes()
	case "phoenix.phoenix.VirtualBlock.rightChild":
		x.RightChild = value.Bytes()
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		x.VdfProof = value.Bytes()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VirtualBlock) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		panic(fmt.Errorf("field id of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.time":
		panic(fmt.Errorf("field time of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.header":
		panic(fmt.Errorf("field header of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.previous":
		panic(fmt.Errorf("field previous of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.parent":
		panic(fmt.Errorf("field parent of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.leftChild":
		panic(fmt.Errorf("field leftChild of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.rightChild":
		panic(fmt.Errorf("field rightChild of message phoenix.phoenix.VirtualBlock is not mutable"))
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		panic(fmt.Errorf("field vdfProof of message phoenix.phoenix.VirtualBlock is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_VirtualBlock) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "phoenix.phoenix.VirtualBlock.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "phoenix.phoenix.VirtualBlock.time":
		return protoreflect.ValueOfInt64(int64(0))
	case "phoenix.phoenix.VirtualBlock.header":
		return protoreflect.ValueOfBytes(nil)
	case "phoenix.phoenix.VirtualBlock.previous":
		return protoreflect.ValueOfBytes(nil)
	case "phoenix.phoenix.VirtualBlock.parent":
		return protoreflect.ValueOfBytes(nil)
	case "phoenix.phoenix.VirtualBlock.leftChild":
		return protoreflect.ValueOfBytes(nil)
	case "phoenix.phoenix.VirtualBlock.rightChild":
		return protoreflect.ValueOfBytes(nil)
	case "phoenix.phoenix.VirtualBlock.vdfProof":
		return protoreflect.ValueOfBytes(nil)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: phoenix.phoenix.VirtualBlock"))
		}
		panic(fmt.Errorf("message phoenix.phoenix.VirtualBlock does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_VirtualBlock) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in phoenix.phoenix.VirtualBlock", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_VirtualBlock) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_VirtualBlock) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_VirtualBlock) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_VirtualBlock) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*VirtualBlock)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		if x.Time != 0 {
			n += 1 + runtime.Sov(uint64(x.Time))
		}
		l = len(x.Header)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Previous)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Parent)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.LeftChild)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.RightChild)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.VdfProof)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*VirtualBlock)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.VdfProof) > 0 {
			i -= len(x.VdfProof)
			copy(dAtA[i:], x.VdfProof)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.VdfProof)))
			i--
			dAtA[i] = 0x42
		}
		if len(x.RightChild) > 0 {
			i -= len(x.RightChild)
			copy(dAtA[i:], x.RightChild)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.RightChild)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.LeftChild) > 0 {
			i -= len(x.LeftChild)
			copy(dAtA[i:], x.LeftChild)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.LeftChild)))
			i--
			dAtA[i] = 0x32
		}
		if len(x.Parent) > 0 {
			i -= len(x.Parent)
			copy(dAtA[i:], x.Parent)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Parent)))
			i--
			dAtA[i] = 0x2a
		}
		if len(x.Previous) > 0 {
			i -= len(x.Previous)
			copy(dAtA[i:], x.Previous)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Previous)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Header) > 0 {
			i -= len(x.Header)
			copy(dAtA[i:], x.Header)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Header)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Time != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Time))
			i--
			dAtA[i] = 0x10
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*VirtualBlock)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VirtualBlock: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: VirtualBlock: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
				}
				x.Time = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Time |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Header = append(x.Header[:0], dAtA[iNdEx:postIndex]...)
				if x.Header == nil {
					x.Header = []byte{}
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Previous", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Previous = append(x.Previous[:0], dAtA[iNdEx:postIndex]...)
				if x.Previous == nil {
					x.Previous = []byte{}
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Parent", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Parent = append(x.Parent[:0], dAtA[iNdEx:postIndex]...)
				if x.Parent == nil {
					x.Parent = []byte{}
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field LeftChild", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.LeftChild = append(x.LeftChild[:0], dAtA[iNdEx:postIndex]...)
				if x.LeftChild == nil {
					x.LeftChild = []byte{}
				}
				iNdEx = postIndex
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field RightChild", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.RightChild = append(x.RightChild[:0], dAtA[iNdEx:postIndex]...)
				if x.RightChild == nil {
					x.RightChild = []byte{}
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field VdfProof", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.VdfProof = append(x.VdfProof[:0], dAtA[iNdEx:postIndex]...)
				if x.VdfProof == nil {
					x.VdfProof = []byte{}
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: phoenix/phoenix/virtual_block.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VirtualBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time       int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Header     []byte `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Previous   []byte `protobuf:"bytes,4,opt,name=previous,proto3" json:"previous,omitempty"`
	Parent     []byte `protobuf:"bytes,5,opt,name=parent,proto3" json:"parent,omitempty"`
	LeftChild  []byte `protobuf:"bytes,6,opt,name=leftChild,proto3" json:"leftChild,omitempty"`
	RightChild []byte `protobuf:"bytes,7,opt,name=rightChild,proto3" json:"rightChild,omitempty"`
	VdfProof   []byte `protobuf:"bytes,8,opt,name=vdfProof,proto3" json:"vdfProof,omitempty"`
}

func (x *VirtualBlock) Reset() {
	*x = VirtualBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_phoenix_phoenix_virtual_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirtualBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirtualBlock) ProtoMessage() {}

// Deprecated: Use VirtualBlock.ProtoReflect.Descriptor instead.
func (*VirtualBlock) Descriptor() ([]byte, []int) {
	return file_phoenix_phoenix_virtual_block_proto_rawDescGZIP(), []int{0}
}

func (x *VirtualBlock) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VirtualBlock) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *VirtualBlock) GetHeader() []byte {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *VirtualBlock) GetPrevious() []byte {
	if x != nil {
		return x.Previous
	}
	return nil
}

func (x *VirtualBlock) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *VirtualBlock) GetLeftChild() []byte {
	if x != nil {
		return x.LeftChild
	}
	return nil
}

func (x *VirtualBlock) GetRightChild() []byte {
	if x != nil {
		return x.RightChild
	}
	return nil
}

func (x *VirtualBlock) GetVdfProof() []byte {
	if x != nil {
		return x.VdfProof
	}
	return nil
}

var File_phoenix_phoenix_virtual_block_proto protoreflect.FileDescriptor

var file_phoenix_phoenix_virtual_block_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2f, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69,
	0x78, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2e, 0x70,
	0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x22, 0xd8, 0x01, 0x0a, 0x0c, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x66, 0x74, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x6c, 0x65, 0x66, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x69, 0x67, 0x68, 0x74, 0x43, 0x68,
	0x69, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x69, 0x67, 0x68, 0x74,
	0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x64, 0x66, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x76, 0x64, 0x66, 0x50, 0x72, 0x6f, 0x6f,
	0x66, 0x42, 0xa7, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69,
	0x78, 0x2e, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x42, 0x11, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x20,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x2f, 0x70, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78,
	0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa, 0x02, 0x0f, 0x50, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78,
	0x2e, 0x50, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0xca, 0x02, 0x0f, 0x50, 0x68, 0x6f, 0x65, 0x6e,
	0x69, 0x78, 0x5c, 0x50, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0xe2, 0x02, 0x1b, 0x50, 0x68, 0x6f,
	0x65, 0x6e, 0x69, 0x78, 0x5c, 0x50, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x50, 0x68, 0x6f, 0x65, 0x6e,
	0x69, 0x78, 0x3a, 0x3a, 0x50, 0x68, 0x6f, 0x65, 0x6e, 0x69, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_phoenix_phoenix_virtual_block_proto_rawDescOnce sync.Once
	file_phoenix_phoenix_virtual_block_proto_rawDescData = file_phoenix_phoenix_virtual_block_proto_rawDesc
)

func file_phoenix_phoenix_virtual_block_proto_rawDescGZIP() []byte {
	file_phoenix_phoenix_virtual_block_proto_rawDescOnce.Do(func() {
		file_phoenix_phoenix_virtual_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_phoenix_phoenix_virtual_block_proto_rawDescData)
	})
	return file_phoenix_phoenix_virtual_block_proto_rawDescData
}

var file_phoenix_phoenix_virtual_block_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_phoenix_phoenix_virtual_block_proto_goTypes = []interface{}{
	(*VirtualBlock)(nil), // 0: phoenix.phoenix.VirtualBlock
}
var file_phoenix_phoenix_virtual_block_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_phoenix_phoenix_virtual_block_proto_init() }
func file_phoenix_phoenix_virtual_block_proto_init() {
	if File_phoenix_phoenix_virtual_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_phoenix_phoenix_virtual_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirtualBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_phoenix_phoenix_virtual_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_phoenix_phoenix_virtual_block_proto_goTypes,
		DependencyIndexes: file_phoenix_phoenix_virtual_block_proto_depIdxs,
		MessageInfos:      file_phoenix_phoenix_virtual_block_proto_msgTypes,
	}.Build()
	File_phoenix_phoenix_virtual_block_proto = out.File
	file_phoenix_phoenix_virtual_block_proto_rawDesc = nil
	file_phoenix_phoenix_virtual_block_proto_goTypes = nil
	file_phoenix_phoenix_virtual_block_proto_depIdxs = nil
}
