// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

public interface GetKeyRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kvstore.v1.GetKeyRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint64 account_id = 1 [json_name = "accountId"];</code>
   * @return The accountId.
   */
  long getAccountId();

  /**
   * <code>string bucket_name = 2 [json_name = "bucketName"];</code>
   * @return The bucketName.
   */
  java.lang.String getBucketName();
  /**
   * <code>string bucket_name = 2 [json_name = "bucketName"];</code>
   * @return The bytes for bucketName.
   */
  com.google.protobuf.ByteString
      getBucketNameBytes();

  /**
   * <code>bytes key = 3 [json_name = "key"];</code>
   * @return The key.
   */
  com.google.protobuf.ByteString getKey();

  /**
   * <code>optional uint32 version = 4 [json_name = "version"];</code>
   * @return Whether the version field is set.
   */
  boolean hasVersion();
  /**
   * <code>optional uint32 version = 4 [json_name = "version"];</code>
   * @return The version.
   */
  int getVersion();
}