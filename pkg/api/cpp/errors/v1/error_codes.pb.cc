// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: errors/v1/error_codes.proto

#include "errors/v1/error_codes.pb.h"

#include <algorithm>
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/extension_set.h"
#include "google/protobuf/wire_format_lite.h"
#include "google/protobuf/descriptor.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/reflection_ops.h"
#include "google/protobuf/wire_format.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"
PROTOBUF_PRAGMA_INIT_SEG
namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = ::PROTOBUF_NAMESPACE_ID::internal;
namespace errors {
namespace v1 {
}  // namespace v1
}  // namespace errors
static const ::_pb::EnumDescriptor* file_level_enum_descriptors_errors_2fv1_2ferror_5fcodes_2eproto[1];
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_errors_2fv1_2ferror_5fcodes_2eproto = nullptr;
const ::uint32_t TableStruct_errors_2fv1_2ferror_5fcodes_2eproto::offsets[1] = {};
static constexpr ::_pbi::MigrationSchema* schemas = nullptr;
static constexpr ::_pb::Message* const* file_default_instances = nullptr;
const char descriptor_table_protodef_errors_2fv1_2ferror_5fcodes_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\033errors/v1/error_codes.proto\022\terrors.v1"
    "*\242\003\n\004Code\022\024\n\020CODE_UNSPECIFIED\020\000\022\013\n\007CODE_"
    "OK\020\001\022\022\n\016CODE_CANCELLED\020\002\022\020\n\014CODE_UNKNOWN"
    "\020\003\022\031\n\025CODE_INVALID_ARGUMENT\020\004\022\032\n\026CODE_DE"
    "ADLINE_EXCEEDED\020\005\022\022\n\016CODE_NOT_FOUND\020\006\022\027\n"
    "\023CODE_ALREADY_EXISTS\020\007\022\032\n\026CODE_PERMISSIO"
    "N_DENIED\020\010\022\030\n\024CODE_UNAUTHENTICATED\020\t\022\033\n\027"
    "CODE_RESOURCE_EXHAUSTED\020\n\022\034\n\030CODE_FAILED"
    "_PRECONDITION\020\013\022\020\n\014CODE_ABORTED\020\014\022\025\n\021COD"
    "E_OUT_OF_RANGE\020\r\022\026\n\022CODE_UNIMPLEMENTED\020\016"
    "\022\021\n\rCODE_INTERNAL\020\017\022\024\n\020CODE_UNAVAILABLE\020"
    "\020\022\022\n\016CODE_DATA_LOSS\020\021B\271\001\n)com.github.mxp"
    "lusb.pleiades.api.errors.v1B\017ErrorCodesP"
    "rotoP\001Z6github.com/mxplusb/pleiades/pkg/"
    "api/errors/v1;errorsv1\242\002\003EXX\252\002\tErrors.V1"
    "\312\002\tErrors\\V1\342\002\025Errors\\V1\\GPBMetadata\352\002\nE"
    "rrors::V1b\006proto3"
};
static ::absl::once_flag descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto = {
    false,
    false,
    657,
    descriptor_table_protodef_errors_2fv1_2ferror_5fcodes_2eproto,
    "errors/v1/error_codes.proto",
    &descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto_once,
    nullptr,
    0,
    0,
    schemas,
    file_default_instances,
    TableStruct_errors_2fv1_2ferror_5fcodes_2eproto::offsets,
    nullptr,
    file_level_enum_descriptors_errors_2fv1_2ferror_5fcodes_2eproto,
    file_level_service_descriptors_errors_2fv1_2ferror_5fcodes_2eproto,
};

// This function exists to be marked as weak.
// It can significantly speed up compilation by breaking up LLVM's SCC
// in the .pb.cc translation units. Large translation units see a
// reduction of more than 35% of walltime for optimized builds. Without
// the weak attribute all the messages in the file, including all the
// vtables and everything they use become part of the same SCC through
// a cycle like:
// GetMetadata -> descriptor table -> default instances ->
//   vtables -> GetMetadata
// By adding a weak function here we break the connection from the
// individual vtables back into the descriptor table.
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto_getter() {
  return &descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_errors_2fv1_2ferror_5fcodes_2eproto(&descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto);
namespace errors {
namespace v1 {
const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Code_descriptor() {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto);
  return file_level_enum_descriptors_errors_2fv1_2ferror_5fcodes_2eproto[0];
}
bool Code_IsValid(int value) {
  switch (value) {
    case 0:
    case 1:
    case 2:
    case 3:
    case 4:
    case 5:
    case 6:
    case 7:
    case 8:
    case 9:
    case 10:
    case 11:
    case 12:
    case 13:
    case 14:
    case 15:
    case 16:
    case 17:
      return true;
    default:
      return false;
  }
}
// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace errors
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
