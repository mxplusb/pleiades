// @generated by protoc-gen-connect-web v0.2.1 with parameter "target=js+dts"
// @generated from file kvstore/v1/service.proto (package kvstore.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import {CreateAccountRequest, CreateAccountResponse, CreateBucketRequest, CreateBucketResponse, DeleteAccountRequest, DeleteAccountResponse, DeleteBucketRequest, DeleteBucketResponse, DeleteKeyRequest, DeleteKeyResponse, GetKeyRequest, GetKeyResponse, PutKeyRequest, PutKeyResponse} from "./kv_pb.js";
import {MethodKind} from "@bufbuild/protobuf";
import {CloseTransactionRequest, CloseTransactionResponse, CommitRequest, CommitResponse, NewTransactionRequest, NewTransactionResponse} from "./transactions_pb.js";

/**
 * @generated from service kvstore.v1.KvStoreService
 */
export declare const KvStoreService: {
  readonly typeName: "kvstore.v1.KvStoreService",
  readonly methods: {
    /**
     * @generated from rpc kvstore.v1.KvStoreService.CreateAccount
     */
    readonly createAccount: {
      readonly name: "CreateAccount",
      readonly I: typeof CreateAccountRequest,
      readonly O: typeof CreateAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.DeleteAccount
     */
    readonly deleteAccount: {
      readonly name: "DeleteAccount",
      readonly I: typeof DeleteAccountRequest,
      readonly O: typeof DeleteAccountResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.CreateBucket
     */
    readonly createBucket: {
      readonly name: "CreateBucket",
      readonly I: typeof CreateBucketRequest,
      readonly O: typeof CreateBucketResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.DeleteBucket
     */
    readonly deleteBucket: {
      readonly name: "DeleteBucket",
      readonly I: typeof DeleteBucketRequest,
      readonly O: typeof DeleteBucketResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.GetKey
     */
    readonly getKey: {
      readonly name: "GetKey",
      readonly I: typeof GetKeyRequest,
      readonly O: typeof GetKeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.PutKey
     */
    readonly putKey: {
      readonly name: "PutKey",
      readonly I: typeof PutKeyRequest,
      readonly O: typeof PutKeyResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.KvStoreService.DeleteKey
     */
    readonly deleteKey: {
      readonly name: "DeleteKey",
      readonly I: typeof DeleteKeyRequest,
      readonly O: typeof DeleteKeyResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};

/**
 * @generated from service kvstore.v1.TransactionsService
 */
export declare const TransactionsService: {
  readonly typeName: "kvstore.v1.TransactionsService",
  readonly methods: {
    /**
     * @generated from rpc kvstore.v1.TransactionsService.NewTransaction
     */
    readonly newTransaction: {
      readonly name: "NewTransaction",
      readonly I: typeof NewTransactionRequest,
      readonly O: typeof NewTransactionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.TransactionsService.CloseTransaction
     */
    readonly closeTransaction: {
      readonly name: "CloseTransaction",
      readonly I: typeof CloseTransactionRequest,
      readonly O: typeof CloseTransactionResponse,
      readonly kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc kvstore.v1.TransactionsService.Commit
     */
    readonly commit: {
      readonly name: "Commit",
      readonly I: typeof CommitRequest,
      readonly O: typeof CommitResponse,
      readonly kind: MethodKind.Unary,
    },
  }
};
