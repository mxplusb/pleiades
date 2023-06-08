// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_shard.proto

package com.github.mxplusb.pleiades.api.raft.v1;

public interface ShardStateOrBuilder extends
    // @@protoc_insertion_point(interface_extends:raft.v1.ShardState)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>.google.protobuf.Timestamp last_updated = 1 [json_name = "lastUpdated"];</code>
   * @return Whether the lastUpdated field is set.
   */
  boolean hasLastUpdated();
  /**
   * <code>.google.protobuf.Timestamp last_updated = 1 [json_name = "lastUpdated"];</code>
   * @return The lastUpdated.
   */
  com.google.protobuf.Timestamp getLastUpdated();
  /**
   * <code>.google.protobuf.Timestamp last_updated = 1 [json_name = "lastUpdated"];</code>
   */
  com.google.protobuf.TimestampOrBuilder getLastUpdatedOrBuilder();

  /**
   * <code>uint64 shard_id = 2 [json_name = "shardId"];</code>
   * @return The shardId.
   */
  long getShardId();

  /**
   * <code>uint64 config_change_id = 3 [json_name = "configChangeId"];</code>
   * @return The configChangeId.
   */
  long getConfigChangeId();

  /**
   * <code>map&lt;uint64, string&gt; replicas = 4 [json_name = "replicas"];</code>
   */
  int getReplicasCount();
  /**
   * <code>map&lt;uint64, string&gt; replicas = 4 [json_name = "replicas"];</code>
   */
  boolean containsReplicas(
      long key);
  /**
   * Use {@link #getReplicasMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.Long, java.lang.String>
  getReplicas();
  /**
   * <code>map&lt;uint64, string&gt; replicas = 4 [json_name = "replicas"];</code>
   */
  java.util.Map<java.lang.Long, java.lang.String>
  getReplicasMap();
  /**
   * <code>map&lt;uint64, string&gt; replicas = 4 [json_name = "replicas"];</code>
   */
  /* nullable */
java.lang.String getReplicasOrDefault(
      long key,
      /* nullable */
java.lang.String defaultValue);
  /**
   * <code>map&lt;uint64, string&gt; replicas = 4 [json_name = "replicas"];</code>
   */
  java.lang.String getReplicasOrThrow(
      long key);

  /**
   * <code>map&lt;uint64, string&gt; observers = 5 [json_name = "observers"];</code>
   */
  int getObserversCount();
  /**
   * <code>map&lt;uint64, string&gt; observers = 5 [json_name = "observers"];</code>
   */
  boolean containsObservers(
      long key);
  /**
   * Use {@link #getObserversMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.Long, java.lang.String>
  getObservers();
  /**
   * <code>map&lt;uint64, string&gt; observers = 5 [json_name = "observers"];</code>
   */
  java.util.Map<java.lang.Long, java.lang.String>
  getObserversMap();
  /**
   * <code>map&lt;uint64, string&gt; observers = 5 [json_name = "observers"];</code>
   */
  /* nullable */
java.lang.String getObserversOrDefault(
      long key,
      /* nullable */
java.lang.String defaultValue);
  /**
   * <code>map&lt;uint64, string&gt; observers = 5 [json_name = "observers"];</code>
   */
  java.lang.String getObserversOrThrow(
      long key);

  /**
   * <code>map&lt;uint64, string&gt; witnesses = 6 [json_name = "witnesses"];</code>
   */
  int getWitnessesCount();
  /**
   * <code>map&lt;uint64, string&gt; witnesses = 6 [json_name = "witnesses"];</code>
   */
  boolean containsWitnesses(
      long key);
  /**
   * Use {@link #getWitnessesMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.Long, java.lang.String>
  getWitnesses();
  /**
   * <code>map&lt;uint64, string&gt; witnesses = 6 [json_name = "witnesses"];</code>
   */
  java.util.Map<java.lang.Long, java.lang.String>
  getWitnessesMap();
  /**
   * <code>map&lt;uint64, string&gt; witnesses = 6 [json_name = "witnesses"];</code>
   */
  /* nullable */
java.lang.String getWitnessesOrDefault(
      long key,
      /* nullable */
java.lang.String defaultValue);
  /**
   * <code>map&lt;uint64, string&gt; witnesses = 6 [json_name = "witnesses"];</code>
   */
  java.lang.String getWitnessesOrThrow(
      long key);

  /**
   * <code>map&lt;uint64, string&gt; removed = 7 [json_name = "removed"];</code>
   */
  int getRemovedCount();
  /**
   * <code>map&lt;uint64, string&gt; removed = 7 [json_name = "removed"];</code>
   */
  boolean containsRemoved(
      long key);
  /**
   * Use {@link #getRemovedMap()} instead.
   */
  @java.lang.Deprecated
  java.util.Map<java.lang.Long, java.lang.String>
  getRemoved();
  /**
   * <code>map&lt;uint64, string&gt; removed = 7 [json_name = "removed"];</code>
   */
  java.util.Map<java.lang.Long, java.lang.String>
  getRemovedMap();
  /**
   * <code>map&lt;uint64, string&gt; removed = 7 [json_name = "removed"];</code>
   */
  /* nullable */
java.lang.String getRemovedOrDefault(
      long key,
      /* nullable */
java.lang.String defaultValue);
  /**
   * <code>map&lt;uint64, string&gt; removed = 7 [json_name = "removed"];</code>
   */
  java.lang.String getRemovedOrThrow(
      long key);

  /**
   * <code>.raft.v1.StateMachineType type = 8 [json_name = "type"];</code>
   * @return The enum numeric value on the wire for type.
   */
  int getTypeValue();
  /**
   * <code>.raft.v1.StateMachineType type = 8 [json_name = "type"];</code>
   * @return The type.
   */
  com.github.mxplusb.pleiades.api.raft.v1.StateMachineType getType();
}
