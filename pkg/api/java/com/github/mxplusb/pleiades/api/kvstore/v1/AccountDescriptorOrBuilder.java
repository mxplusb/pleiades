// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

public interface AccountDescriptorOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kvstore.v1.AccountDescriptor)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint64 account_id = 1 [json_name = "accountId"];</code>
   * @return The accountId.
   */
  long getAccountId();

  /**
   * <code>string owner = 2 [json_name = "owner"];</code>
   * @return The owner.
   */
  java.lang.String getOwner();
  /**
   * <code>string owner = 2 [json_name = "owner"];</code>
   * @return The bytes for owner.
   */
  com.google.protobuf.ByteString
      getOwnerBytes();

  /**
   * <code>.google.protobuf.Timestamp created = 3 [json_name = "created"];</code>
   * @return Whether the created field is set.
   */
  boolean hasCreated();
  /**
   * <code>.google.protobuf.Timestamp created = 3 [json_name = "created"];</code>
   * @return The created.
   */
  com.google.protobuf.Timestamp getCreated();
  /**
   * <code>.google.protobuf.Timestamp created = 3 [json_name = "created"];</code>
   */
  com.google.protobuf.TimestampOrBuilder getCreatedOrBuilder();

  /**
   * <code>.google.protobuf.Timestamp last_updated = 4 [json_name = "lastUpdated"];</code>
   * @return Whether the lastUpdated field is set.
   */
  boolean hasLastUpdated();
  /**
   * <code>.google.protobuf.Timestamp last_updated = 4 [json_name = "lastUpdated"];</code>
   * @return The lastUpdated.
   */
  com.google.protobuf.Timestamp getLastUpdated();
  /**
   * <code>.google.protobuf.Timestamp last_updated = 4 [json_name = "lastUpdated"];</code>
   */
  com.google.protobuf.TimestampOrBuilder getLastUpdatedOrBuilder();

  /**
   * <code>uint64 bucket_count = 5 [json_name = "bucketCount"];</code>
   * @return The bucketCount.
   */
  long getBucketCount();

  /**
   * <code>repeated string buckets = 6 [json_name = "buckets"];</code>
   * @return A list containing the buckets.
   */
  java.util.List<java.lang.String>
      getBucketsList();
  /**
   * <code>repeated string buckets = 6 [json_name = "buckets"];</code>
   * @return The count of buckets.
   */
  int getBucketsCount();
  /**
   * <code>repeated string buckets = 6 [json_name = "buckets"];</code>
   * @param index The index of the element to return.
   * @return The buckets at the given index.
   */
  java.lang.String getBuckets(int index);
  /**
   * <code>repeated string buckets = 6 [json_name = "buckets"];</code>
   * @param index The index of the value to return.
   * @return The bytes of the buckets at the given index.
   */
  com.google.protobuf.ByteString
      getBucketsBytes(int index);
}
