// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/service.proto

#include "kvstore/v1/service.pb.h"

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
namespace kvstore {
namespace v1 {
}  // namespace v1
}  // namespace kvstore
static constexpr const ::_pb::EnumDescriptor**
    file_level_enum_descriptors_kvstore_2fv1_2fservice_2eproto = nullptr;
static constexpr const ::_pb::ServiceDescriptor**
    file_level_service_descriptors_kvstore_2fv1_2fservice_2eproto = nullptr;
const ::uint32_t TableStruct_kvstore_2fv1_2fservice_2eproto::offsets[1] = {};
static constexpr ::_pbi::MigrationSchema* schemas = nullptr;
static constexpr ::_pb::Message* const* file_default_instances = nullptr;
const char descriptor_table_protodef_kvstore_2fv1_2fservice_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
    "\n\030kvstore/v1/service.proto\022\nkvstore.v1\032\023"
    "kvstore/v1/kv.proto\032\035kvstore/v1/transact"
    "ions.proto2\277\004\n\016KvStoreService\022V\n\rCreateA"
    "ccount\022 .kvstore.v1.CreateAccountRequest"
    "\032!.kvstore.v1.CreateAccountResponse\"\000\022V\n"
    "\rDeleteAccount\022 .kvstore.v1.DeleteAccoun"
    "tRequest\032!.kvstore.v1.DeleteAccountRespo"
    "nse\"\000\022S\n\014CreateBucket\022\037.kvstore.v1.Creat"
    "eBucketRequest\032 .kvstore.v1.CreateBucket"
    "Response\"\000\022S\n\014DeleteBucket\022\037.kvstore.v1."
    "DeleteBucketRequest\032 .kvstore.v1.DeleteB"
    "ucketResponse\"\000\022D\n\006GetKey\022\031.kvstore.v1.G"
    "etKeyRequest\032\032.kvstore.v1.GetKeyResponse"
    "\"\003\220\002\001\022A\n\006PutKey\022\031.kvstore.v1.PutKeyReque"
    "st\032\032.kvstore.v1.PutKeyResponse\"\000\022J\n\tDele"
    "teKey\022\034.kvstore.v1.DeleteKeyRequest\032\035.kv"
    "store.v1.DeleteKeyResponse\"\0002\216\002\n\023Transac"
    "tionsService\022W\n\016NewTransaction\022!.kvstore"
    ".v1.NewTransactionRequest\032\".kvstore.v1.N"
    "ewTransactionResponse\022]\n\020CloseTransactio"
    "n\022#.kvstore.v1.CloseTransactionRequest\032$"
    ".kvstore.v1.CloseTransactionResponse\022\?\n\006"
    "Commit\022\031.kvstore.v1.CommitRequest\032\032.kvst"
    "ore.v1.CommitResponseB\275\001\n*com.github.mxp"
    "lusb.pleiades.api.kvstore.v1B\014ServicePro"
    "toP\001Z8github.com/mxplusb/pleiades/pkg/ap"
    "i/kvstore/v1;kvstorev1\242\002\003KXX\252\002\nKvstore.V"
    "1\312\002\nKvstore\\V1\342\002\026Kvstore\\V1\\GPBMetadata\352"
    "\002\013Kvstore::V1b\006proto3"
};
static const ::_pbi::DescriptorTable* const descriptor_table_kvstore_2fv1_2fservice_2eproto_deps[2] =
    {
        &::descriptor_table_kvstore_2fv1_2fkv_2eproto,
        &::descriptor_table_kvstore_2fv1_2ftransactions_2eproto,
};
static ::absl::once_flag descriptor_table_kvstore_2fv1_2fservice_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_kvstore_2fv1_2fservice_2eproto = {
    false,
    false,
    1141,
    descriptor_table_protodef_kvstore_2fv1_2fservice_2eproto,
    "kvstore/v1/service.proto",
    &descriptor_table_kvstore_2fv1_2fservice_2eproto_once,
    descriptor_table_kvstore_2fv1_2fservice_2eproto_deps,
    2,
    0,
    schemas,
    file_default_instances,
    TableStruct_kvstore_2fv1_2fservice_2eproto::offsets,
    nullptr,
    file_level_enum_descriptors_kvstore_2fv1_2fservice_2eproto,
    file_level_service_descriptors_kvstore_2fv1_2fservice_2eproto,
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
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_kvstore_2fv1_2fservice_2eproto_getter() {
  return &descriptor_table_kvstore_2fv1_2fservice_2eproto;
}
// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2
static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_kvstore_2fv1_2fservice_2eproto(&descriptor_table_kvstore_2fv1_2fservice_2eproto);
namespace kvstore {
namespace v1 {
// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace kvstore
PROTOBUF_NAMESPACE_OPEN
PROTOBUF_NAMESPACE_CLOSE
// @@protoc_insertion_point(global_scope)
#include "google/protobuf/port_undef.inc"
