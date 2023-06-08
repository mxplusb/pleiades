/*
 * Copyright (c) 2023 Sienna Lloyd
 *
 * Licensed under the PolyForm Internal Use License 1.0.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License here:
 *  https://github.com/mxplusb/pleiades/blob/mainline/LICENSE
 */

// @generated by protoc-gen-es v0.1.1 with parameter "target=js+dts"
// @generated from file errors/v1/errors.proto (package errors.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";
import type {Code} from "./error_codes_pb.js";

/**
 * The `Status` type defines a logical error model that is suitable for
 * different programming environments, including REST APIs and RPC APIs.
 *
 * @generated from message errors.v1.Error
 */
export declare class Error extends Message<Error> {
  /**
   * A simple error code that can be easily handled by the client. The
   * actual error code is defined by `google.rpc.Code`.
   *
   * @generated from field: errors.v1.Code code = 1;
   */
  code: Code;

  /**
   * A developer-facing human-readable error message in English. It should
   * both explain the error and offer an actionable resolution to it.
   *
   * @generated from field: string message = 2;
   */
  message: string;

  constructor(data?: PartialMessage<Error>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "errors.v1.Error";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Error;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Error;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Error;

  static equals(a: Error | PlainMessage<Error> | undefined, b: Error | PlainMessage<Error> | undefined): boolean;
}

