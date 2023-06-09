/*
 * Copyright (c) 2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file raft/v1/raft_event.proto (package raft.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {proto3, Timestamp} from "@bufbuild/protobuf";

/**
 * @generated from enum raft.v1.EventType
 */
export const EventType = proto3.makeEnum(
  "raft.v1.EventType",
  [
    {no: 0, name: "EVENT_TYPE_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "EVENT_TYPE_LOG_ENTRY", localName: "LOG_ENTRY"},
    {no: 2, name: "EVENT_TYPE_SNAPSHOT", localName: "SNAPSHOT"},
    {no: 3, name: "EVENT_TYPE_CONNECTION", localName: "CONNECTION"},
    {no: 4, name: "EVENT_TYPE_HOST", localName: "HOST"},
    {no: 5, name: "EVENT_TYPE_NODE", localName: "NODE"},
    {no: 6, name: "EVENT_TYPE_RAFT", localName: "RAFT"},
  ],
);

/**
 * @generated from enum raft.v1.Event
 */
export const Event = proto3.makeEnum(
  "raft.v1.Event",
  [
    {no: 0, name: "EVENT_UNSPECIFIED", localName: "UNSPECIFIED"},
    {no: 1, name: "EVENT_CONNECTION_ESTABLISHED", localName: "CONNECTION_ESTABLISHED"},
    {no: 2, name: "EVENT_CONNECTION_FAILED", localName: "CONNECTION_FAILED"},
    {no: 3, name: "EVENT_LOG_COMPACTED", localName: "LOG_COMPACTED"},
    {no: 4, name: "EVENT_LOGDB_COMPACTED", localName: "LOGDB_COMPACTED"},
    {no: 5, name: "EVENT_MEMBERSHIP_CHANGED", localName: "MEMBERSHIP_CHANGED"},
    {no: 6, name: "EVENT_NODE_HOST_SHUTTING_DOWN", localName: "NODE_HOST_SHUTTING_DOWN"},
    {no: 7, name: "EVENT_NODE_READY", localName: "NODE_READY"},
    {no: 8, name: "EVENT_NODE_UNLOADED", localName: "NODE_UNLOADED"},
    {no: 9, name: "EVENT_SEND_SNAPSHOT_ABORTED", localName: "SEND_SNAPSHOT_ABORTED"},
    {no: 10, name: "EVENT_SEND_SNAPSHOT_COMPLETED", localName: "SEND_SNAPSHOT_COMPLETED"},
    {no: 11, name: "EVENT_SEND_SNAPSHOT_STARTED", localName: "SEND_SNAPSHOT_STARTED"},
    {no: 12, name: "EVENT_SNAPSHOT_COMPACTED", localName: "SNAPSHOT_COMPACTED"},
    {no: 13, name: "EVENT_SNAPSHOT_CREATED", localName: "SNAPSHOT_CREATED"},
    {no: 14, name: "EVENT_SNAPSHOT_RECEIVED", localName: "SNAPSHOT_RECEIVED"},
    {no: 15, name: "EVENT_SNAPSHOT_RECOVERED", localName: "SNAPSHOT_RECOVERED"},
    {no: 16, name: "EVENT_LEADER_UPDATED", localName: "LEADER_UPDATED"},
  ],
);

/**
 * @generated from message raft.v1.RaftEvent
 */
export const RaftEvent = proto3.makeMessageType(
  "raft.v1.RaftEvent",
  () => [
    { no: 1, name: "typ", kind: "enum", T: proto3.getEnumType(EventType) },
    { no: 2, name: "action", kind: "enum", T: proto3.getEnumType(Event) },
    { no: 3, name: "timestamp", kind: "message", T: Timestamp },
    { no: 4, name: "log_entry", kind: "message", T: RaftLogEntryEvent, oneof: "event" },
    { no: 5, name: "snapshot", kind: "message", T: RaftSnapshotEvent, oneof: "event" },
    { no: 6, name: "connection", kind: "message", T: RaftConnectionEvent, oneof: "event" },
    { no: 7, name: "node", kind: "message", T: RaftNodeEvent, oneof: "event" },
    { no: 8, name: "host_shutdown", kind: "message", T: RaftHostShutdown, oneof: "event" },
    { no: 9, name: "leader_update", kind: "message", T: RaftLeaderInfo, oneof: "event" },
  ],
);

/**
 * @generated from message raft.v1.RaftLeaderInfo
 */
export const RaftLeaderInfo = proto3.makeMessageType(
  "raft.v1.RaftLeaderInfo",
  () => [
    { no: 1, name: "shard_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "replica_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "term", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "leader_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message raft.v1.RaftLogEntryEvent
 */
export const RaftLogEntryEvent = proto3.makeMessageType(
  "raft.v1.RaftLogEntryEvent",
  () => [
    { no: 1, name: "shard_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "replica_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "index", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message raft.v1.RaftSnapshotEvent
 */
export const RaftSnapshotEvent = proto3.makeMessageType(
  "raft.v1.RaftSnapshotEvent",
  () => [
    { no: 1, name: "shard_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "replica_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "from_index", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 4, name: "to_index", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message raft.v1.RaftConnectionEvent
 */
export const RaftConnectionEvent = proto3.makeMessageType(
  "raft.v1.RaftConnectionEvent",
  () => [
    { no: 1, name: "address", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "is_snapshot", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ],
);

/**
 * @generated from message raft.v1.RaftNodeEvent
 */
export const RaftNodeEvent = proto3.makeMessageType(
  "raft.v1.RaftNodeEvent",
  () => [
    { no: 1, name: "shard_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 2, name: "replica_id", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ],
);

/**
 * @generated from message raft.v1.RaftHostShutdown
 */
export const RaftHostShutdown = proto3.makeMessageType(
  "raft.v1.RaftHostShutdown",
  [],
);

