// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: errors/v1/errors.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferrors_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferrors_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3021000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3021009 < PROTOBUF_MIN_PROTOC_VERSION
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
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "errors/v1/error_codes.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_errors_2fv1_2ferrors_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_errors_2fv1_2ferrors_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_errors_2fv1_2ferrors_2eproto;
namespace errors {
namespace v1 {
class Error;
struct ErrorDefaultTypeInternal;
extern ErrorDefaultTypeInternal _Error_default_instance_;
}  // namespace v1
}  // namespace errors
PROTOBUF_NAMESPACE_OPEN
template<> ::errors::v1::Error* Arena::CreateMaybeMessage<::errors::v1::Error>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace errors {
namespace v1 {

// ===================================================================

class Error final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:errors.v1.Error) */ {
 public:
  inline Error() : Error(nullptr) {}
  ~Error() override;
  explicit PROTOBUF_CONSTEXPR Error(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Error(const Error& from);
  Error(Error&& from) noexcept
    : Error() {
    *this = ::std::move(from);
  }

  inline Error& operator=(const Error& from) {
    CopyFrom(from);
    return *this;
  }
  inline Error& operator=(Error&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Error& default_instance() {
    return *internal_default_instance();
  }
  static inline const Error* internal_default_instance() {
    return reinterpret_cast<const Error*>(
               &_Error_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Error& a, Error& b) {
    a.Swap(&b);
  }
  inline void Swap(Error* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Error* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Error* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Error>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Error& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom( const Error& from) {
    Error::MergeImpl(*this, from);
  }
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message& to_msg, const ::PROTOBUF_NAMESPACE_ID::Message& from_msg);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::PROTOBUF_NAMESPACE_ID::Arena* arena, bool is_message_owned);
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Error* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "errors.v1.Error";
  }
  protected:
  explicit Error(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kMessageFieldNumber = 2,
    kCodeFieldNumber = 1,
  };
  // string message = 2 [json_name = "message"];
  void clear_message();
  const std::string& message() const;
  template <typename ArgT0 = const std::string&, typename... ArgT>
  void set_message(ArgT0&& arg0, ArgT... args);
  std::string* mutable_message();
  PROTOBUF_NODISCARD std::string* release_message();
  void set_allocated_message(std::string* message);
  private:
  const std::string& _internal_message() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_message(const std::string& value);
  std::string* _internal_mutable_message();
  public:

  // .errors.v1.Code code = 1 [json_name = "code"];
  void clear_code();
  ::errors::v1::Code code() const;
  void set_code(::errors::v1::Code value);
  private:
  ::errors::v1::Code _internal_code() const;
  void _internal_set_code(::errors::v1::Code value);
  public:

  // @@protoc_insertion_point(class_scope:errors.v1.Error)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  struct Impl_ {
    ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr message_;
    int code_;
    mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_errors_2fv1_2ferrors_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Error

// .errors.v1.Code code = 1 [json_name = "code"];
inline void Error::clear_code() {
  _impl_.code_ = 0;
}
inline ::errors::v1::Code Error::_internal_code() const {
  return static_cast< ::errors::v1::Code >(_impl_.code_);
}
inline ::errors::v1::Code Error::code() const {
  // @@protoc_insertion_point(field_get:errors.v1.Error.code)
  return _internal_code();
}
inline void Error::_internal_set_code(::errors::v1::Code value) {
  
  _impl_.code_ = value;
}
inline void Error::set_code(::errors::v1::Code value) {
  _internal_set_code(value);
  // @@protoc_insertion_point(field_set:errors.v1.Error.code)
}

// string message = 2 [json_name = "message"];
inline void Error::clear_message() {
  _impl_.message_.ClearToEmpty();
}
inline const std::string& Error::message() const {
  // @@protoc_insertion_point(field_get:errors.v1.Error.message)
  return _internal_message();
}
template <typename ArgT0, typename... ArgT>
inline PROTOBUF_ALWAYS_INLINE
void Error::set_message(ArgT0&& arg0, ArgT... args) {
 
 _impl_.message_.Set(static_cast<ArgT0 &&>(arg0), args..., GetArenaForAllocation());
  // @@protoc_insertion_point(field_set:errors.v1.Error.message)
}
inline std::string* Error::mutable_message() {
  std::string* _s = _internal_mutable_message();
  // @@protoc_insertion_point(field_mutable:errors.v1.Error.message)
  return _s;
}
inline const std::string& Error::_internal_message() const {
  return _impl_.message_.Get();
}
inline void Error::_internal_set_message(const std::string& value) {
  
  _impl_.message_.Set(value, GetArenaForAllocation());
}
inline std::string* Error::_internal_mutable_message() {
  
  return _impl_.message_.Mutable(GetArenaForAllocation());
}
inline std::string* Error::release_message() {
  // @@protoc_insertion_point(field_release:errors.v1.Error.message)
  return _impl_.message_.Release();
}
inline void Error::set_allocated_message(std::string* message) {
  if (message != nullptr) {
    
  } else {
    
  }
  _impl_.message_.SetAllocated(message, GetArenaForAllocation());
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (_impl_.message_.IsDefault()) {
    _impl_.message_.Set("", GetArenaForAllocation());
  }
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:errors.v1.Error.message)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace errors

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_errors_2fv1_2ferrors_2eproto
