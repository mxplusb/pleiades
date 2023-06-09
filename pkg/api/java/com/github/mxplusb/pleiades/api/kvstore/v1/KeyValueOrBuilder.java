// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

public interface KeyValueOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kvstore.v1.KeyValue)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * key is the key in bytes. An empty key is not allowed.
   * </pre>
   *
   * <code>bytes key = 1 [json_name = "key"];</code>
   * @return The key.
   */
  com.google.protobuf.ByteString getKey();

  /**
   * <pre>
   * create_revision is the revision of last creation on this key.
   * </pre>
   *
   * <code>int64 create_revision = 2 [json_name = "createRevision"];</code>
   * @return The createRevision.
   */
  long getCreateRevision();

  /**
   * <pre>
   * mod_revision is the revision of last modification on this key.
   * </pre>
   *
   * <code>int64 mod_revision = 3 [json_name = "modRevision"];</code>
   * @return The modRevision.
   */
  long getModRevision();

  /**
   * <pre>
   * version is the version of the key. A deletion resets
   * the version to zero and any modification of the key
   * increases its version.
   * </pre>
   *
   * <code>uint32 version = 4 [json_name = "version"];</code>
   * @return The version.
   */
  int getVersion();

  /**
   * <pre>
   * value is the value held by the key, in bytes.
   * </pre>
   *
   * <code>bytes value = 5 [json_name = "value"];</code>
   * @return The value.
   */
  com.google.protobuf.ByteString getValue();

  /**
   * <pre>
   * lease is the ID of the lease that attached to key.
   * When the attached lease expires, the key will be deleted.
   * If lease is 0, then no lease is attached to the key.
   * </pre>
   *
   * <code>int64 lease = 6 [json_name = "lease"];</code>
   * @return The lease.
   */
  long getLease();
}
