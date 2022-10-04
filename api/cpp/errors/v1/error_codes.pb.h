// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: errors/v1/error_codes.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferror_5fcodes_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferror_5fcodes_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021006 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_errors_2fv1_2ferror_5fcodes_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_errors_2fv1_2ferror_5fcodes_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_errors_2fv1_2ferror_5fcodes_2eproto;
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
namespace errors {
namespace v1 {

enum Code : int {
  CODE_UNSPECIFIED = 0,
  CODE_OK = 1,
  CODE_CANCELLED = 2,
  CODE_UNKNOWN = 3,
  CODE_INVALID_ARGUMENT = 4,
  CODE_DEADLINE_EXCEEDED = 5,
  CODE_NOT_FOUND = 6,
  CODE_ALREADY_EXISTS = 7,
  CODE_PERMISSION_DENIED = 8,
  CODE_UNAUTHENTICATED = 9,
  CODE_RESOURCE_EXHAUSTED = 10,
  CODE_FAILED_PRECONDITION = 11,
  CODE_ABORTED = 12,
  CODE_OUT_OF_RANGE = 13,
  CODE_UNIMPLEMENTED = 14,
  CODE_INTERNAL = 15,
  CODE_UNAVAILABLE = 16,
  CODE_DATA_LOSS = 17,
  Code_INT_MIN_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::min(),
  Code_INT_MAX_SENTINEL_DO_NOT_USE_ = std::numeric_limits<int32_t>::max()
};
bool Code_IsValid(int value);
constexpr Code Code_MIN = CODE_UNSPECIFIED;
constexpr Code Code_MAX = CODE_DATA_LOSS;
constexpr int Code_ARRAYSIZE = Code_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Code_descriptor();
template<typename T>
inline const std::string& Code_Name(T enum_t_value) {
  static_assert(::std::is_same<T, Code>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function Code_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    Code_descriptor(), enum_t_value);
}
inline bool Code_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, Code* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<Code>(
    Code_descriptor(), name, value);
}
// ===================================================================


// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace errors

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::errors::v1::Code> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::errors::v1::Code>() {
  return ::errors::v1::Code_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferror_5fcodes_2eproto
