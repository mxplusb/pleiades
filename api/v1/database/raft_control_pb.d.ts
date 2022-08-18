// package: database
// file: api/v1/database/raft_control.proto

import * as jspb from "google-protobuf";
import * as api_v1_database_kv_pb from "../../../api/v1/database/kv_pb";
import * as api_v1_errors_pb from "../../../api/v1/errors_pb";

export class RaftControlPayload extends jspb.Message {
  hasGetleaderidrequest(): boolean;
  clearGetleaderidrequest(): void;
  getGetleaderidrequest(): GetLeaderIDRequest | undefined;
  setGetleaderidrequest(value?: GetLeaderIDRequest): void;

  hasGetleaderidresponse(): boolean;
  clearGetleaderidresponse(): void;
  getGetleaderidresponse(): GetLeaderIDResponse | undefined;
  setGetleaderidresponse(value?: GetLeaderIDResponse): void;

  hasIdrequest(): boolean;
  clearIdrequest(): void;
  getIdrequest(): IdRequest | undefined;
  setIdrequest(value?: IdRequest): void;

  hasIdresponse(): boolean;
  clearIdresponse(): void;
  getIdresponse(): IdResponse | undefined;
  setIdresponse(value?: IdResponse): void;

  hasIndexstate(): boolean;
  clearIndexstate(): void;
  getIndexstate(): IndexState | undefined;
  setIndexstate(value?: IndexState): void;

  hasModifynoderequest(): boolean;
  clearModifynoderequest(): void;
  getModifynoderequest(): ModifyNodeRequest | undefined;
  setModifynoderequest(value?: ModifyNodeRequest): void;

  hasReadindexrequest(): boolean;
  clearReadindexrequest(): void;
  getReadindexrequest(): ReadIndexRequest | undefined;
  setReadindexrequest(value?: ReadIndexRequest): void;

  hasReadlocalnoderequest(): boolean;
  clearReadlocalnoderequest(): void;
  getReadlocalnoderequest(): ReadLocalNodeRequest | undefined;
  setReadlocalnoderequest(value?: ReadLocalNodeRequest): void;

  hasRequestleadertransferresponse(): boolean;
  clearRequestleadertransferresponse(): void;
  getRequestleadertransferresponse(): RequestLeaderTransferResponse | undefined;
  setRequestleadertransferresponse(value?: RequestLeaderTransferResponse): void;

  hasRequestsnapshotrequest(): boolean;
  clearRequestsnapshotrequest(): void;
  getRequestsnapshotrequest(): RequestSnapshotRequest | undefined;
  setRequestsnapshotrequest(value?: RequestSnapshotRequest): void;

  hasSnapshotoption(): boolean;
  clearSnapshotoption(): void;
  getSnapshotoption(): SnapshotOption | undefined;
  setSnapshotoption(value?: SnapshotOption): void;

  hasStopnoderesponse(): boolean;
  clearStopnoderesponse(): void;
  getStopnoderesponse(): StopNodeResponse | undefined;
  setStopnoderesponse(value?: StopNodeResponse): void;

  hasStoprequest(): boolean;
  clearStoprequest(): void;
  getStoprequest(): StopRequest | undefined;
  setStoprequest(value?: StopRequest): void;

  hasStopresponse(): boolean;
  clearStopresponse(): void;
  getStopresponse(): StopResponse | undefined;
  setStopresponse(value?: StopResponse): void;

  hasSysopstate(): boolean;
  clearSysopstate(): void;
  getSysopstate(): SysOpState | undefined;
  setSysopstate(value?: SysOpState): void;

  hasError(): boolean;
  clearError(): void;
  getError(): api_v1_errors_pb.DBError | undefined;
  setError(value?: api_v1_errors_pb.DBError): void;

  getMethod(): RaftControlPayload.MethodNameMap[keyof RaftControlPayload.MethodNameMap];
  setMethod(value: RaftControlPayload.MethodNameMap[keyof RaftControlPayload.MethodNameMap]): void;

  getTypesCase(): RaftControlPayload.TypesCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RaftControlPayload.AsObject;
  static toObject(includeInstance: boolean, msg: RaftControlPayload): RaftControlPayload.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RaftControlPayload, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RaftControlPayload;
  static deserializeBinaryFromReader(message: RaftControlPayload, reader: jspb.BinaryReader): RaftControlPayload;
}

export namespace RaftControlPayload {
  export type AsObject = {
    getleaderidrequest?: GetLeaderIDRequest.AsObject,
    getleaderidresponse?: GetLeaderIDResponse.AsObject,
    idrequest?: IdRequest.AsObject,
    idresponse?: IdResponse.AsObject,
    indexstate?: IndexState.AsObject,
    modifynoderequest?: ModifyNodeRequest.AsObject,
    readindexrequest?: ReadIndexRequest.AsObject,
    readlocalnoderequest?: ReadLocalNodeRequest.AsObject,
    requestleadertransferresponse?: RequestLeaderTransferResponse.AsObject,
    requestsnapshotrequest?: RequestSnapshotRequest.AsObject,
    snapshotoption?: SnapshotOption.AsObject,
    stopnoderesponse?: StopNodeResponse.AsObject,
    stoprequest?: StopRequest.AsObject,
    stopresponse?: StopResponse.AsObject,
    sysopstate?: SysOpState.AsObject,
    error?: api_v1_errors_pb.DBError.AsObject,
    method: RaftControlPayload.MethodNameMap[keyof RaftControlPayload.MethodNameMap],
  }

  export interface MethodNameMap {
    ADD_NODE: 0;
    ADD_OBSERVER: 1;
    ADD_WITNESS: 2;
    GET_ID: 3;
    GET_LEADER_ID: 4;
    READ_INDEX: 5;
    READ_LOCAL_NODE: 6;
    REQUEST_COMPACTION: 7;
    REQUEST_DELETE_NODE: 8;
    REQUEST_LEADER_TRANSFER: 9;
    REQUEST_SNAPSHOT: 10;
    STOP: 11;
    STOP_NODE: 12;
  }

  export const MethodName: MethodNameMap;

  export enum TypesCase {
    TYPES_NOT_SET = 0,
    GETLEADERIDREQUEST = 1,
    GETLEADERIDRESPONSE = 2,
    IDREQUEST = 3,
    IDRESPONSE = 4,
    INDEXSTATE = 5,
    MODIFYNODEREQUEST = 6,
    READINDEXREQUEST = 7,
    READLOCALNODEREQUEST = 8,
    REQUESTLEADERTRANSFERRESPONSE = 9,
    REQUESTSNAPSHOTREQUEST = 10,
    SNAPSHOTOPTION = 12,
    STOPNODERESPONSE = 13,
    STOPREQUEST = 14,
    STOPRESPONSE = 15,
    SYSOPSTATE = 16,
    ERROR = 18,
  }
}

export class StopNodeResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StopNodeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StopNodeResponse): StopNodeResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StopNodeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StopNodeResponse;
  static deserializeBinaryFromReader(message: StopNodeResponse, reader: jspb.BinaryReader): StopNodeResponse;
}

export namespace StopNodeResponse {
  export type AsObject = {
  }
}

export class StopRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StopRequest.AsObject;
  static toObject(includeInstance: boolean, msg: StopRequest): StopRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StopRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StopRequest;
  static deserializeBinaryFromReader(message: StopRequest, reader: jspb.BinaryReader): StopRequest;
}

export namespace StopRequest {
  export type AsObject = {
  }
}

export class StopResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StopResponse.AsObject;
  static toObject(includeInstance: boolean, msg: StopResponse): StopResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StopResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StopResponse;
  static deserializeBinaryFromReader(message: StopResponse, reader: jspb.BinaryReader): StopResponse;
}

export namespace StopResponse {
  export type AsObject = {
  }
}

export class RequestSnapshotRequest extends jspb.Message {
  getClusterid(): number;
  setClusterid(value: number): void;

  hasOptions(): boolean;
  clearOptions(): void;
  getOptions(): SnapshotOption | undefined;
  setOptions(value?: SnapshotOption): void;

  getTimeout(): number;
  setTimeout(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RequestSnapshotRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RequestSnapshotRequest): RequestSnapshotRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RequestSnapshotRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RequestSnapshotRequest;
  static deserializeBinaryFromReader(message: RequestSnapshotRequest, reader: jspb.BinaryReader): RequestSnapshotRequest;
}

export namespace RequestSnapshotRequest {
  export type AsObject = {
    clusterid: number,
    options?: SnapshotOption.AsObject,
    timeout: number,
  }
}

export class SnapshotOption extends jspb.Message {
  getCompactionoverhead(): number;
  setCompactionoverhead(value: number): void;

  getExportpath(): string;
  setExportpath(value: string): void;

  getExported(): boolean;
  setExported(value: boolean): void;

  getOverridecompactionoverhead(): boolean;
  setOverridecompactionoverhead(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SnapshotOption.AsObject;
  static toObject(includeInstance: boolean, msg: SnapshotOption): SnapshotOption.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SnapshotOption, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SnapshotOption;
  static deserializeBinaryFromReader(message: SnapshotOption, reader: jspb.BinaryReader): SnapshotOption;
}

export namespace SnapshotOption {
  export type AsObject = {
    compactionoverhead: number,
    exportpath: string,
    exported: boolean,
    overridecompactionoverhead: boolean,
  }
}

export class RequestLeaderTransferResponse extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RequestLeaderTransferResponse.AsObject;
  static toObject(includeInstance: boolean, msg: RequestLeaderTransferResponse): RequestLeaderTransferResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RequestLeaderTransferResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RequestLeaderTransferResponse;
  static deserializeBinaryFromReader(message: RequestLeaderTransferResponse, reader: jspb.BinaryReader): RequestLeaderTransferResponse;
}

export namespace RequestLeaderTransferResponse {
  export type AsObject = {
  }
}

export class SysOpState extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SysOpState.AsObject;
  static toObject(includeInstance: boolean, msg: SysOpState): SysOpState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SysOpState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SysOpState;
  static deserializeBinaryFromReader(message: SysOpState, reader: jspb.BinaryReader): SysOpState;
}

export namespace SysOpState {
  export type AsObject = {
  }
}

export class ModifyNodeRequest extends jspb.Message {
  hasClusterid(): boolean;
  clearClusterid(): void;
  getClusterid(): number;
  setClusterid(value: number): void;

  hasNodeid(): boolean;
  clearNodeid(): void;
  getNodeid(): number;
  setNodeid(value: number): void;

  hasTarget(): boolean;
  clearTarget(): void;
  getTarget(): string;
  setTarget(value: string): void;

  hasConfigchangeindex(): boolean;
  clearConfigchangeindex(): void;
  getConfigchangeindex(): number;
  setConfigchangeindex(value: number): void;

  hasTimeout(): boolean;
  clearTimeout(): void;
  getTimeout(): number;
  setTimeout(value: number): void;

  hasSynchronous(): boolean;
  clearSynchronous(): void;
  getSynchronous(): boolean;
  setSynchronous(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ModifyNodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ModifyNodeRequest): ModifyNodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ModifyNodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ModifyNodeRequest;
  static deserializeBinaryFromReader(message: ModifyNodeRequest, reader: jspb.BinaryReader): ModifyNodeRequest;
}

export namespace ModifyNodeRequest {
  export type AsObject = {
    clusterid: number,
    nodeid: number,
    target: string,
    configchangeindex: number,
    timeout: number,
    synchronous: boolean,
  }
}

export class ReadLocalNodeRequest extends jspb.Message {
  hasQuery(): boolean;
  clearQuery(): void;
  getQuery(): api_v1_database_kv_pb.KeyValue | undefined;
  setQuery(value?: api_v1_database_kv_pb.KeyValue): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadLocalNodeRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadLocalNodeRequest): ReadLocalNodeRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadLocalNodeRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadLocalNodeRequest;
  static deserializeBinaryFromReader(message: ReadLocalNodeRequest, reader: jspb.BinaryReader): ReadLocalNodeRequest;
}

export namespace ReadLocalNodeRequest {
  export type AsObject = {
    query?: api_v1_database_kv_pb.KeyValue.AsObject,
  }
}

export class ReadIndexRequest extends jspb.Message {
  getClusterid(): number;
  setClusterid(value: number): void;

  getTimeout(): number;
  setTimeout(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ReadIndexRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ReadIndexRequest): ReadIndexRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ReadIndexRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ReadIndexRequest;
  static deserializeBinaryFromReader(message: ReadIndexRequest, reader: jspb.BinaryReader): ReadIndexRequest;
}

export namespace ReadIndexRequest {
  export type AsObject = {
    clusterid: number,
    timeout: number,
  }
}

export class IndexState extends jspb.Message {
  getStatus(): IndexState.ResultCodeMap[keyof IndexState.ResultCodeMap];
  setStatus(value: IndexState.ResultCodeMap[keyof IndexState.ResultCodeMap]): void;

  getSnapshotindex(): number;
  setSnapshotindex(value: number): void;

  hasResults(): boolean;
  clearResults(): void;
  getResults(): Result | undefined;
  setResults(value?: Result): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IndexState.AsObject;
  static toObject(includeInstance: boolean, msg: IndexState): IndexState.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IndexState, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IndexState;
  static deserializeBinaryFromReader(message: IndexState, reader: jspb.BinaryReader): IndexState;
}

export namespace IndexState {
  export type AsObject = {
    status: IndexState.ResultCodeMap[keyof IndexState.ResultCodeMap],
    snapshotindex: number,
    results?: Result.AsObject,
  }

  export interface ResultCodeMap {
    TIMEOUT: 0;
    COMPLETED: 1;
    TERMINATED: 2;
    REJECTED: 3;
    DROPPED: 4;
    ABORTED: 5;
    COMMITTED: 6;
  }

  export const ResultCode: ResultCodeMap;
}

export class Result extends jspb.Message {
  getValue(): number;
  setValue(value: number): void;

  getData(): Uint8Array | string;
  getData_asU8(): Uint8Array;
  getData_asB64(): string;
  setData(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Result.AsObject;
  static toObject(includeInstance: boolean, msg: Result): Result.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Result, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Result;
  static deserializeBinaryFromReader(message: Result, reader: jspb.BinaryReader): Result;
}

export namespace Result {
  export type AsObject = {
    value: number,
    data: Uint8Array | string,
  }
}

export class IdRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IdRequest.AsObject;
  static toObject(includeInstance: boolean, msg: IdRequest): IdRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IdRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IdRequest;
  static deserializeBinaryFromReader(message: IdRequest, reader: jspb.BinaryReader): IdRequest;
}

export namespace IdRequest {
  export type AsObject = {
  }
}

export class IdResponse extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IdResponse.AsObject;
  static toObject(includeInstance: boolean, msg: IdResponse): IdResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IdResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IdResponse;
  static deserializeBinaryFromReader(message: IdResponse, reader: jspb.BinaryReader): IdResponse;
}

export namespace IdResponse {
  export type AsObject = {
    id: string,
  }
}

export class GetLeaderIDRequest extends jspb.Message {
  getClusterid(): number;
  setClusterid(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLeaderIDRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetLeaderIDRequest): GetLeaderIDRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLeaderIDRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLeaderIDRequest;
  static deserializeBinaryFromReader(message: GetLeaderIDRequest, reader: jspb.BinaryReader): GetLeaderIDRequest;
}

export namespace GetLeaderIDRequest {
  export type AsObject = {
    clusterid: number,
  }
}

export class GetLeaderIDResponse extends jspb.Message {
  getLeaderid(): number;
  setLeaderid(value: number): void;

  getIsleader(): boolean;
  setIsleader(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetLeaderIDResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetLeaderIDResponse): GetLeaderIDResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetLeaderIDResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetLeaderIDResponse;
  static deserializeBinaryFromReader(message: GetLeaderIDResponse, reader: jspb.BinaryReader): GetLeaderIDResponse;
}

export namespace GetLeaderIDResponse {
  export type AsObject = {
    leaderid: number,
    isleader: boolean,
  }
}

