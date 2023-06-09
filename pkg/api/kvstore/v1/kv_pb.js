/*
 * Copyright (c) 2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file kvstore/v1/kv.proto (package kvstore.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {proto3, Timestamp} from "@bufbuild/protobuf";
import {Error} from "../../errors/v1/errors_pb.js";
import {Transaction} from "./transactions_pb.js";

/**
 * @generated from enum kvstore.v1.KeyOperationType
 */
export const KeyOperationType = proto3.makeEnum(
  "kvstore.v1.KeyOperationType",
  [
    {no: 0, name: "KEY_OPERATION_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "KEY_OPERATION_TYPE_GET", localName: "GET"},
    {no: 2, name: "KEY_OPERATION_TYPE_PUT", localName: "PUT"},
    {no: 3, name: "KEY_OPERATION_TYPE_DELETE", localName: "DELETE"},
  ],
);

/**
 * @generated from message kvstore.v1.KVStoreWrapper
 */
export const KVStoreWrapper = proto3.makeMessageType(
  "kvstore.v1.KVStoreWrapper",
  () => [
    { no: 1, name: "account", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "bucket", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "typ", kind: "enum", T: proto3.getEnumType(KVStoreWrapper_RequestType) },
    { no: 4, name: "create_account_request", kind: "message", T: CreateAccountRequest, oneof: "payload" },
    { no: 5, name: "create_account_reply", kind: "message", T: CreateAccountResponse, oneof: "payload" },
    { no: 6, name: "delete_account_request", kind: "message", T: DeleteAccountRequest, oneof: "payload" },
    { no: 7, name: "delete_account_reply", kind: "message", T: DeleteAccountResponse, oneof: "payload" },
    { no: 8, name: "get_account_descriptor_request", kind: "message", T: GetAccountDescriptorRequest, oneof: "payload" },
    { no: 9, name: "get_account_descriptor_reply", kind: "message", T: GetAccountDescriptorResponse, oneof: "payload" },
    { no: 10, name: "create_bucket_request", kind: "message", T: CreateBucketRequest, oneof: "payload" },
    { no: 11, name: "create_bucket_reply", kind: "message", T: CreateBucketResponse, oneof: "payload" },
    { no: 12, name: "delete_bucket_request", kind: "message", T: DeleteBucketRequest, oneof: "payload" },
    { no: 13, name: "delete_bucket_reply", kind: "message", T: DeleteBucketResponse, oneof: "payload" },
    { no: 14, name: "get_key_request", kind: "message", T: GetKeyRequest, oneof: "payload" },
    { no: 15, name: "get_key_reply", kind: "message", T: GetKeyResponse, oneof: "payload" },
    { no: 16, name: "put_key_request", kind: "message", T: PutKeyRequest, oneof: "payload" },
    { no: 17, name: "put_key_reply", kind: "message", T: PutKeyResponse, oneof: "payload" },
    { no: 18, name: "delete_key_request", kind: "message", T: DeleteKeyRequest, oneof: "payload" },
    { no: 19, name: "delete_key_reply", kind: "message", T: DeleteKeyResponse, oneof: "payload" },
    { no: 20, name: "get_bucket_descriptor_request", kind: "message", T: GetBucketDescriptorRequest, oneof: "payload" },
    { no: 21, name: "get_bucket_descriptor_reply", kind: "message", T: GetBucketDescriptorResponse, oneof: "payload" },
    { no: 22, name: "error", kind: "message", T: Error, oneof: "payload" },
  ],
);

/**
 * @generated from enum kvstore.v1.KVStoreWrapper.RequestType
 */
export const KVStoreWrapper_RequestType = proto3.makeEnum(
  "kvstore.v1.KVStoreWrapper.RequestType",
  [
    {no: 0, name: "REQUEST_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "REQUEST_TYPE_CREATE_ACCOUNT_REQUEST", localName: "CREATE_ACCOUNT_REQUEST"},
    {no: 2, name: "REQUEST_TYPE_CREATE_ACCOUNT_REPLY", localName: "CREATE_ACCOUNT_REPLY"},
    {no: 3, name: "REQUEST_TYPE_DELETE_ACCOUNT_REQUEST", localName: "DELETE_ACCOUNT_REQUEST"},
    {no: 4, name: "REQUEST_TYPE_DELETE_ACCOUNT_REPLY", localName: "DELETE_ACCOUNT_REPLY"},
    {no: 5, name: "REQUEST_TYPE_GET_ACCOUNT_DESCRIPTOR_REQUEST", localName: "GET_ACCOUNT_DESCRIPTOR_REQUEST"},
    {no: 6, name: "REQUEST_TYPE_GET_ACCOUNT_DESCRIPTOR_REPLY", localName: "GET_ACCOUNT_DESCRIPTOR_REPLY"},
    {no: 7, name: "REQUEST_TYPE_CREATE_BUCKET_REQUEST", localName: "CREATE_BUCKET_REQUEST"},
    {no: 8, name: "REQUEST_TYPE_CREATE_BUCKET_REPLY", localName: "CREATE_BUCKET_REPLY"},
    {no: 9, name: "REQUEST_TYPE_DELETE_BUCKET_REQUEST", localName: "DELETE_BUCKET_REQUEST"},
    {no: 10, name: "REQUEST_TYPE_DELETE_BUCKET_REPLY", localName: "DELETE_BUCKET_REPLY"},
    {no: 11, name: "REQUEST_TYPE_GET_KEY_REQUEST", localName: "GET_KEY_REQUEST"},
    {no: 12, name: "REQUEST_TYPE_GET_KEY_REPLY", localName: "GET_KEY_REPLY"},
    {no: 13, name: "REQUEST_TYPE_PUT_KEY_REQUEST", localName: "PUT_KEY_REQUEST"},
    {no: 14, name: "REQUEST_TYPE_PUT_KEY_REPLY", localName: "PUT_KEY_REPLY"},
    {no: 15, name: "REQUEST_TYPE_DELETE_KEY_REQUEST", localName: "DELETE_KEY_REQUEST"},
    {no: 16, name: "REQUEST_TYPE_DELETE_KEY_REPLY", localName: "DELETE_KEY_REPLY"},
    {no: 17, name: "REQUEST_TYPE_GET_BUCKET_DESCRIPTOR_REQUEST", localName: "GET_BUCKET_DESCRIPTOR_REQUEST"},
    {no: 18, name: "REQUEST_TYPE_GET_BUCKET_DESCRIPTOR_REPLY", localName: "GET_BUCKET_DESCRIPTOR_REPLY"},
    {no: 19, name: "REQUEST_TYPE_RECOVERABLE_ERROR", localName: "RECOVERABLE_ERROR"},
  ],
);

/**
 * @generated from message kvstore.v1.CreateAccountRequest
 */
export const CreateAccountRequest = proto3.makeMessageType(
  "kvstore.v1.CreateAccountRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "owner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.CreateAccountResponse
 */
export const CreateAccountResponse = proto3.makeMessageType(
  "kvstore.v1.CreateAccountResponse",
  () => [
    { no: 1, name: "account_descriptor", kind: "message", T: AccountDescriptor },
    { no: 2, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteAccountRequest
 */
export const DeleteAccountRequest = proto3.makeMessageType(
  "kvstore.v1.DeleteAccountRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "owner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteAccountResponse
 */
export const DeleteAccountResponse = proto3.makeMessageType(
  "kvstore.v1.DeleteAccountResponse",
  () => [
    { no: 1, name: "ok", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.GetAccountDescriptorRequest
 */
export const GetAccountDescriptorRequest = proto3.makeMessageType(
  "kvstore.v1.GetAccountDescriptorRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message kvstore.v1.GetAccountDescriptorResponse
 */
export const GetAccountDescriptorResponse = proto3.makeMessageType(
  "kvstore.v1.GetAccountDescriptorResponse",
  () => [
    { no: 1, name: "account_descriptor", kind: "message", T: AccountDescriptor },
  ],
);

/**
 * @generated from message kvstore.v1.AccountDescriptor
 */
export const AccountDescriptor = proto3.makeMessageType(
  "kvstore.v1.AccountDescriptor",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "owner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "created", kind: "message", T: Timestamp },
    { no: 4, name: "last_updated", kind: "message", T: Timestamp },
    { no: 5, name: "bucket_count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 6, name: "buckets", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
  ],
);

/**
 * @generated from message kvstore.v1.CreateBucketRequest
 */
export const CreateBucketRequest = proto3.makeMessageType(
  "kvstore.v1.CreateBucketRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "owner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.CreateBucketResponse
 */
export const CreateBucketResponse = proto3.makeMessageType(
  "kvstore.v1.CreateBucketResponse",
  () => [
    { no: 1, name: "bucket_descriptor", kind: "message", T: BucketDescriptor },
    { no: 2, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteBucketRequest
 */
export const DeleteBucketRequest = proto3.makeMessageType(
  "kvstore.v1.DeleteBucketRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteBucketResponse
 */
export const DeleteBucketResponse = proto3.makeMessageType(
  "kvstore.v1.DeleteBucketResponse",
  () => [
    { no: 1, name: "ok", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.BucketDescriptor
 */
export const BucketDescriptor = proto3.makeMessageType(
  "kvstore.v1.BucketDescriptor",
  () => [
    { no: 1, name: "owner", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "size", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "key_count", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "created", kind: "message", T: Timestamp },
    { no: 5, name: "last_updated", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message kvstore.v1.GetBucketDescriptorRequest
 */
export const GetBucketDescriptorRequest = proto3.makeMessageType(
  "kvstore.v1.GetBucketDescriptorRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "bucket_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message kvstore.v1.GetBucketDescriptorResponse
 */
export const GetBucketDescriptorResponse = proto3.makeMessageType(
  "kvstore.v1.GetBucketDescriptorResponse",
  () => [
    { no: 1, name: "bucket_descriptor", kind: "message", T: BucketDescriptor },
  ],
);

/**
 * @generated from message kvstore.v1.GetKeyRequest
 */
export const GetKeyRequest = proto3.makeMessageType(
  "kvstore.v1.GetKeyRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "bucket_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "version", kind: "scalar", T: 13 /* ScalarType.UINT32 */, opt: true },
  ],
);

/**
 * @generated from message kvstore.v1.GetKeyResponse
 */
export const GetKeyResponse = proto3.makeMessageType(
  "kvstore.v1.GetKeyResponse",
  () => [
    { no: 1, name: "key_value_pair", kind: "message", T: KeyValue },
  ],
);

/**
 * @generated from message kvstore.v1.PutKeyRequest
 */
export const PutKeyRequest = proto3.makeMessageType(
  "kvstore.v1.PutKeyRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "bucket_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key_value_pair", kind: "message", T: KeyValue },
    { no: 4, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.PutKeyResponse
 */
export const PutKeyResponse = proto3.makeMessageType(
  "kvstore.v1.PutKeyResponse",
  () => [
    { no: 1, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteKeyRequest
 */
export const DeleteKeyRequest = proto3.makeMessageType(
  "kvstore.v1.DeleteKeyRequest",
  () => [
    { no: 1, name: "account_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "bucket_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 4, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.DeleteKeyResponse
 */
export const DeleteKeyResponse = proto3.makeMessageType(
  "kvstore.v1.DeleteKeyResponse",
  () => [
    { no: 1, name: "ok", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 3, name: "transaction", kind: "message", T: Transaction },
  ],
);

/**
 * @generated from message kvstore.v1.KeyValueDescriptor
 */
export const KeyValueDescriptor = proto3.makeMessageType(
  "kvstore.v1.KeyValueDescriptor",
  () => [
    { no: 1, name: "versions", kind: "scalar", T: 13 /* ScalarType.UINT32 */, repeated: true },
    { no: 2, name: "current_key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ],
);

/**
 * @generated from message kvstore.v1.ListKeyVersionsRequest
 */
export const ListKeyVersionsRequest = proto3.makeMessageType(
  "kvstore.v1.ListKeyVersionsRequest",
  () => [
    { no: 1, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
  ],
);

/**
 * @generated from message kvstore.v1.ListKeyVersionsResponse
 */
export const ListKeyVersionsResponse = proto3.makeMessageType(
  "kvstore.v1.ListKeyVersionsResponse",
  () => [
    { no: 1, name: "versions", kind: "scalar", T: 13 /* ScalarType.UINT32 */, repeated: true },
  ],
);

/**
 * @generated from message kvstore.v1.KeyValue
 */
export const KeyValue = proto3.makeMessageType(
  "kvstore.v1.KeyValue",
  () => [
    { no: 1, name: "key", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 2, name: "create_revision", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "mod_revision", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "version", kind: "scalar", T: 13 /* ScalarType.UINT32 */ },
    { no: 5, name: "value", kind: "scalar", T: 12 /* ScalarType.BYTES */ },
    { no: 6, name: "lease", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ],
);

/**
 * @generated from message kvstore.v1.Event
 */
export const Event = proto3.makeMessageType(
  "kvstore.v1.Event",
  () => [
    { no: 1, name: "type", kind: "enum", T: proto3.getEnumType(KeyOperationType) },
    { no: 2, name: "kv", kind: "message", T: KeyValue },
    { no: 3, name: "prev_kv", kind: "message", T: KeyValue, opt: true },
  ],
);

