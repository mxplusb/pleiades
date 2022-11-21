// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_shard.proto

package io.a13s.api.raft.v1;

public interface StartReplicaRequestOrBuilder extends
    // @@protoc_insertion_point(interface_extends:raft.v1.StartReplicaRequest)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
   * @return The shardId.
   */
  long getShardId();

  /**
   * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
   * @return The replicaId.
   */
  long getReplicaId();

  /**
   * <code>.raft.v1.StateMachineType type = 3 [json_name = "type"];</code>
   * @return The enum numeric value on the wire for type.
   */
  int getTypeValue();
  /**
   * <code>.raft.v1.StateMachineType type = 3 [json_name = "type"];</code>
   * @return The type.
   */
  io.a13s.api.raft.v1.StateMachineType getType();

  /**
   * <code>bool restart = 4 [json_name = "restart"];</code>
   * @return The restart.
   */
  boolean getRestart();
}
