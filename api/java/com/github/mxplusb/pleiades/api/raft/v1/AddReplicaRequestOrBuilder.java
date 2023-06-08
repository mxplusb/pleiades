// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_shard.proto

package com.github.mxplusb.pleiades.api.raft.v1;

public interface AddReplicaRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:raft.v1.AddReplicaRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * replica_id is a non-zero value used to identify a node within a Raft cluster.
   * </pre>
   *
   * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
   * @return The shardId.
   */
  long getShardId();

  /**
   * <pre>
   * shard_id is the unique value used to identify a Raft cluster.
   * </pre>
   *
   * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
   * @return The replicaId.
   */
  long getReplicaId();

  /**
   * <code>string hostname = 4 [json_name = "hostname"];</code>
   * @return The hostname.
   */
  java.lang.String getHostname();
  /**
   * <code>string hostname = 4 [json_name = "hostname"];</code>
   * @return The bytes for hostname.
   */
  com.google.protobuf.ByteString
      getHostnameBytes();

  /**
   * <code>int64 timeout = 5 [json_name = "timeout"];</code>
   * @return The timeout.
   */
  long getTimeout();
}
