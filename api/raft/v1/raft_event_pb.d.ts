// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file raft/v1/raft_event.proto (package raft.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from enum raft.v1.EventType
 */
export declare enum EventType {
  /**
   * @generated from enum value: EVENT_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: EVENT_TYPE_LOG_ENTRY = 1;
   */
  LOG_ENTRY = 1,

  /**
   * @generated from enum value: EVENT_TYPE_SNAPSHOT = 2;
   */
  SNAPSHOT = 2,

  /**
   * @generated from enum value: EVENT_TYPE_CONNECTION = 3;
   */
  CONNECTION = 3,

  /**
   * @generated from enum value: EVENT_TYPE_HOST = 4;
   */
  HOST = 4,

  /**
   * @generated from enum value: EVENT_TYPE_NODE = 5;
   */
  NODE = 5,

  /**
   * @generated from enum value: EVENT_TYPE_RAFT = 6;
   */
  RAFT = 6,
}

/**
 * @generated from enum raft.v1.Event
 */
export declare enum Event {
  /**
   * @generated from enum value: EVENT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: EVENT_CONNECTION_ESTABLISHED = 1;
   */
  CONNECTION_ESTABLISHED = 1,

  /**
   * @generated from enum value: EVENT_CONNECTION_FAILED = 2;
   */
  CONNECTION_FAILED = 2,

  /**
   * @generated from enum value: EVENT_LOG_COMPACTED = 3;
   */
  LOG_COMPACTED = 3,

  /**
   * @generated from enum value: EVENT_LOGDB_COMPACTED = 4;
   */
  LOGDB_COMPACTED = 4,

  /**
   * @generated from enum value: EVENT_MEMBERSHIP_CHANGED = 5;
   */
  MEMBERSHIP_CHANGED = 5,

  /**
   * @generated from enum value: EVENT_NODE_HOST_SHUTTING_DOWN = 6;
   */
  NODE_HOST_SHUTTING_DOWN = 6,

  /**
   * @generated from enum value: EVENT_NODE_READY = 7;
   */
  NODE_READY = 7,

  /**
   * @generated from enum value: EVENT_NODE_UNLOADED = 8;
   */
  NODE_UNLOADED = 8,

  /**
   * @generated from enum value: EVENT_SEND_SNAPSHOT_ABORTED = 9;
   */
  SEND_SNAPSHOT_ABORTED = 9,

  /**
   * @generated from enum value: EVENT_SEND_SNAPSHOT_COMPLETED = 10;
   */
  SEND_SNAPSHOT_COMPLETED = 10,

  /**
   * @generated from enum value: EVENT_SEND_SNAPSHOT_STARTED = 11;
   */
  SEND_SNAPSHOT_STARTED = 11,

  /**
   * @generated from enum value: EVENT_SNAPSHOT_COMPACTED = 12;
   */
  SNAPSHOT_COMPACTED = 12,

  /**
   * @generated from enum value: EVENT_SNAPSHOT_CREATED = 13;
   */
  SNAPSHOT_CREATED = 13,

  /**
   * @generated from enum value: EVENT_SNAPSHOT_RECEIVED = 14;
   */
  SNAPSHOT_RECEIVED = 14,

  /**
   * @generated from enum value: EVENT_SNAPSHOT_RECOVERED = 15;
   */
  SNAPSHOT_RECOVERED = 15,

  /**
   * @generated from enum value: EVENT_LEADER_UPDATED = 16;
   */
  LEADER_UPDATED = 16,
}

/**
 * @generated from message raft.v1.RaftEvent
 */
export declare class RaftEvent extends Message<RaftEvent> {
  /**
   * @generated from field: raft.v1.EventType typ = 1;
   */
  typ: EventType;

  /**
   * @generated from field: raft.v1.Event action = 2;
   */
  action: Event;

  /**
   * @generated from field: google.protobuf.Timestamp timestamp = 3;
   */
  timestamp?: Timestamp;

  /**
   * @generated from oneof raft.v1.RaftEvent.event
   */
  event: {
    /**
     * @generated from field: raft.v1.RaftLogEntryEvent log_entry = 4;
     */
    value: RaftLogEntryEvent;
    case: "logEntry";
  } | {
    /**
     * @generated from field: raft.v1.RaftSnapshotEvent snapshot = 5;
     */
    value: RaftSnapshotEvent;
    case: "snapshot";
  } | {
    /**
     * @generated from field: raft.v1.RaftConnectionEvent connection = 6;
     */
    value: RaftConnectionEvent;
    case: "connection";
  } | {
    /**
     * @generated from field: raft.v1.RaftNodeEvent node = 7;
     */
    value: RaftNodeEvent;
    case: "node";
  } | {
    /**
     * @generated from field: raft.v1.RaftHostShutdown host_shutdown = 8;
     */
    value: RaftHostShutdown;
    case: "hostShutdown";
  } | {
    /**
     * @generated from field: raft.v1.RaftLeaderInfo leader_update = 9;
     */
    value: RaftLeaderInfo;
    case: "leaderUpdate";
  } | { case: undefined; value?: undefined };

  constructor(data?: PartialMessage<RaftEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftEvent;

  static equals(a: RaftEvent | PlainMessage<RaftEvent> | undefined, b: RaftEvent | PlainMessage<RaftEvent> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftLeaderInfo
 */
export declare class RaftLeaderInfo extends Message<RaftLeaderInfo> {
  /**
   * @generated from field: uint64 shard_id = 1;
   */
  shardId: bigint;

  /**
   * @generated from field: uint64 replica_id = 2;
   */
  replicaId: bigint;

  /**
   * @generated from field: uint64 term = 3;
   */
  term: bigint;

  /**
   * @generated from field: uint64 leader_id = 4;
   */
  leaderId: bigint;

  constructor(data?: PartialMessage<RaftLeaderInfo>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftLeaderInfo";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftLeaderInfo;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftLeaderInfo;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftLeaderInfo;

  static equals(a: RaftLeaderInfo | PlainMessage<RaftLeaderInfo> | undefined, b: RaftLeaderInfo | PlainMessage<RaftLeaderInfo> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftLogEntryEvent
 */
export declare class RaftLogEntryEvent extends Message<RaftLogEntryEvent> {
  /**
   * @generated from field: uint64 shard_id = 1;
   */
  shardId: bigint;

  /**
   * @generated from field: uint64 replica_id = 2;
   */
  replicaId: bigint;

  /**
   * @generated from field: uint64 index = 3;
   */
  index: bigint;

  constructor(data?: PartialMessage<RaftLogEntryEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftLogEntryEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftLogEntryEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftLogEntryEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftLogEntryEvent;

  static equals(a: RaftLogEntryEvent | PlainMessage<RaftLogEntryEvent> | undefined, b: RaftLogEntryEvent | PlainMessage<RaftLogEntryEvent> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftSnapshotEvent
 */
export declare class RaftSnapshotEvent extends Message<RaftSnapshotEvent> {
  /**
   * @generated from field: uint64 shard_id = 1;
   */
  shardId: bigint;

  /**
   * @generated from field: uint64 replica_id = 2;
   */
  replicaId: bigint;

  /**
   * @generated from field: uint64 from_index = 3;
   */
  fromIndex: bigint;

  /**
   * @generated from field: uint64 to_index = 4;
   */
  toIndex: bigint;

  constructor(data?: PartialMessage<RaftSnapshotEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftSnapshotEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftSnapshotEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftSnapshotEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftSnapshotEvent;

  static equals(a: RaftSnapshotEvent | PlainMessage<RaftSnapshotEvent> | undefined, b: RaftSnapshotEvent | PlainMessage<RaftSnapshotEvent> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftConnectionEvent
 */
export declare class RaftConnectionEvent extends Message<RaftConnectionEvent> {
  /**
   * @generated from field: string address = 1;
   */
  address: string;

  /**
   * @generated from field: bool is_snapshot = 2;
   */
  isSnapshot: boolean;

  constructor(data?: PartialMessage<RaftConnectionEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftConnectionEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftConnectionEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftConnectionEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftConnectionEvent;

  static equals(a: RaftConnectionEvent | PlainMessage<RaftConnectionEvent> | undefined, b: RaftConnectionEvent | PlainMessage<RaftConnectionEvent> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftNodeEvent
 */
export declare class RaftNodeEvent extends Message<RaftNodeEvent> {
  /**
   * @generated from field: uint64 shard_id = 1;
   */
  shardId: bigint;

  /**
   * @generated from field: uint64 replica_id = 2;
   */
  replicaId: bigint;

  constructor(data?: PartialMessage<RaftNodeEvent>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftNodeEvent";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftNodeEvent;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftNodeEvent;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftNodeEvent;

  static equals(a: RaftNodeEvent | PlainMessage<RaftNodeEvent> | undefined, b: RaftNodeEvent | PlainMessage<RaftNodeEvent> | undefined): boolean;
}

/**
 * @generated from message raft.v1.RaftHostShutdown
 */
export declare class RaftHostShutdown extends Message<RaftHostShutdown> {
  constructor(data?: PartialMessage<RaftHostShutdown>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "raft.v1.RaftHostShutdown";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RaftHostShutdown;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RaftHostShutdown;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RaftHostShutdown;

  static equals(a: RaftHostShutdown | PlainMessage<RaftHostShutdown> | undefined, b: RaftHostShutdown | PlainMessage<RaftHostShutdown> | undefined): boolean;
}

