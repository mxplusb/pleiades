// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

public interface GetBucketDescriptorRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kvstore.v1.GetBucketDescriptorRequest)
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
}
