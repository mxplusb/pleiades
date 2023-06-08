// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

public interface KVStoreWrapperOrBuilder extends
    // @@protoc_insertion_point(interface_extends:kvstore.v1.KVStoreWrapper)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>uint64 account = 1 [json_name = "account"];</code>
   * @return The account.
   */
  long getAccount();

  /**
   * <code>string bucket = 2 [json_name = "bucket"];</code>
   * @return The bucket.
   */
  java.lang.String getBucket();
  /**
   * <code>string bucket = 2 [json_name = "bucket"];</code>
   * @return The bytes for bucket.
   */
  com.google.protobuf.ByteString
      getBucketBytes();

  /**
   * <code>.kvstore.v1.KVStoreWrapper.RequestType typ = 3 [json_name = "typ"];</code>
   * @return The enum numeric value on the wire for typ.
   */
  int getTypValue();
  /**
   * <code>.kvstore.v1.KVStoreWrapper.RequestType typ = 3 [json_name = "typ"];</code>
   * @return The typ.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.KVStoreWrapper.RequestType getTyp();

  /**
   * <code>.kvstore.v1.CreateAccountRequest create_account_request = 4 [json_name = "createAccountRequest"];</code>
   * @return Whether the createAccountRequest field is set.
   */
  boolean hasCreateAccountRequest();
  /**
   * <code>.kvstore.v1.CreateAccountRequest create_account_request = 4 [json_name = "createAccountRequest"];</code>
   * @return The createAccountRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateAccountRequest getCreateAccountRequest();
  /**
   * <code>.kvstore.v1.CreateAccountRequest create_account_request = 4 [json_name = "createAccountRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateAccountRequestOrBuilder getCreateAccountRequestOrBuilder();

  /**
   * <code>.kvstore.v1.CreateAccountResponse create_account_reply = 5 [json_name = "createAccountReply"];</code>
   * @return Whether the createAccountReply field is set.
   */
  boolean hasCreateAccountReply();
  /**
   * <code>.kvstore.v1.CreateAccountResponse create_account_reply = 5 [json_name = "createAccountReply"];</code>
   * @return The createAccountReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateAccountResponse getCreateAccountReply();
  /**
   * <code>.kvstore.v1.CreateAccountResponse create_account_reply = 5 [json_name = "createAccountReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateAccountResponseOrBuilder getCreateAccountReplyOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteAccountRequest delete_account_request = 6 [json_name = "deleteAccountRequest"];</code>
   * @return Whether the deleteAccountRequest field is set.
   */
  boolean hasDeleteAccountRequest();
  /**
   * <code>.kvstore.v1.DeleteAccountRequest delete_account_request = 6 [json_name = "deleteAccountRequest"];</code>
   * @return The deleteAccountRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteAccountRequest getDeleteAccountRequest();
  /**
   * <code>.kvstore.v1.DeleteAccountRequest delete_account_request = 6 [json_name = "deleteAccountRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteAccountRequestOrBuilder getDeleteAccountRequestOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteAccountResponse delete_account_reply = 7 [json_name = "deleteAccountReply"];</code>
   * @return Whether the deleteAccountReply field is set.
   */
  boolean hasDeleteAccountReply();
  /**
   * <code>.kvstore.v1.DeleteAccountResponse delete_account_reply = 7 [json_name = "deleteAccountReply"];</code>
   * @return The deleteAccountReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteAccountResponse getDeleteAccountReply();
  /**
   * <code>.kvstore.v1.DeleteAccountResponse delete_account_reply = 7 [json_name = "deleteAccountReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteAccountResponseOrBuilder getDeleteAccountReplyOrBuilder();

  /**
   * <code>.kvstore.v1.GetAccountDescriptorRequest get_account_descriptor_request = 8 [json_name = "getAccountDescriptorRequest"];</code>
   * @return Whether the getAccountDescriptorRequest field is set.
   */
  boolean hasGetAccountDescriptorRequest();
  /**
   * <code>.kvstore.v1.GetAccountDescriptorRequest get_account_descriptor_request = 8 [json_name = "getAccountDescriptorRequest"];</code>
   * @return The getAccountDescriptorRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetAccountDescriptorRequest getGetAccountDescriptorRequest();
  /**
   * <code>.kvstore.v1.GetAccountDescriptorRequest get_account_descriptor_request = 8 [json_name = "getAccountDescriptorRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetAccountDescriptorRequestOrBuilder getGetAccountDescriptorRequestOrBuilder();

  /**
   * <code>.kvstore.v1.GetAccountDescriptorResponse get_account_descriptor_reply = 9 [json_name = "getAccountDescriptorReply"];</code>
   * @return Whether the getAccountDescriptorReply field is set.
   */
  boolean hasGetAccountDescriptorReply();
  /**
   * <code>.kvstore.v1.GetAccountDescriptorResponse get_account_descriptor_reply = 9 [json_name = "getAccountDescriptorReply"];</code>
   * @return The getAccountDescriptorReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetAccountDescriptorResponse getGetAccountDescriptorReply();
  /**
   * <code>.kvstore.v1.GetAccountDescriptorResponse get_account_descriptor_reply = 9 [json_name = "getAccountDescriptorReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetAccountDescriptorResponseOrBuilder getGetAccountDescriptorReplyOrBuilder();

  /**
   * <code>.kvstore.v1.CreateBucketRequest create_bucket_request = 10 [json_name = "createBucketRequest"];</code>
   * @return Whether the createBucketRequest field is set.
   */
  boolean hasCreateBucketRequest();
  /**
   * <code>.kvstore.v1.CreateBucketRequest create_bucket_request = 10 [json_name = "createBucketRequest"];</code>
   * @return The createBucketRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateBucketRequest getCreateBucketRequest();
  /**
   * <code>.kvstore.v1.CreateBucketRequest create_bucket_request = 10 [json_name = "createBucketRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateBucketRequestOrBuilder getCreateBucketRequestOrBuilder();

  /**
   * <code>.kvstore.v1.CreateBucketResponse create_bucket_reply = 11 [json_name = "createBucketReply"];</code>
   * @return Whether the createBucketReply field is set.
   */
  boolean hasCreateBucketReply();
  /**
   * <code>.kvstore.v1.CreateBucketResponse create_bucket_reply = 11 [json_name = "createBucketReply"];</code>
   * @return The createBucketReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateBucketResponse getCreateBucketReply();
  /**
   * <code>.kvstore.v1.CreateBucketResponse create_bucket_reply = 11 [json_name = "createBucketReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.CreateBucketResponseOrBuilder getCreateBucketReplyOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteBucketRequest delete_bucket_request = 12 [json_name = "deleteBucketRequest"];</code>
   * @return Whether the deleteBucketRequest field is set.
   */
  boolean hasDeleteBucketRequest();
  /**
   * <code>.kvstore.v1.DeleteBucketRequest delete_bucket_request = 12 [json_name = "deleteBucketRequest"];</code>
   * @return The deleteBucketRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteBucketRequest getDeleteBucketRequest();
  /**
   * <code>.kvstore.v1.DeleteBucketRequest delete_bucket_request = 12 [json_name = "deleteBucketRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteBucketRequestOrBuilder getDeleteBucketRequestOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteBucketResponse delete_bucket_reply = 13 [json_name = "deleteBucketReply"];</code>
   * @return Whether the deleteBucketReply field is set.
   */
  boolean hasDeleteBucketReply();
  /**
   * <code>.kvstore.v1.DeleteBucketResponse delete_bucket_reply = 13 [json_name = "deleteBucketReply"];</code>
   * @return The deleteBucketReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteBucketResponse getDeleteBucketReply();
  /**
   * <code>.kvstore.v1.DeleteBucketResponse delete_bucket_reply = 13 [json_name = "deleteBucketReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteBucketResponseOrBuilder getDeleteBucketReplyOrBuilder();

  /**
   * <code>.kvstore.v1.GetKeyRequest get_key_request = 14 [json_name = "getKeyRequest"];</code>
   * @return Whether the getKeyRequest field is set.
   */
  boolean hasGetKeyRequest();
  /**
   * <code>.kvstore.v1.GetKeyRequest get_key_request = 14 [json_name = "getKeyRequest"];</code>
   * @return The getKeyRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetKeyRequest getGetKeyRequest();
  /**
   * <code>.kvstore.v1.GetKeyRequest get_key_request = 14 [json_name = "getKeyRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetKeyRequestOrBuilder getGetKeyRequestOrBuilder();

  /**
   * <code>.kvstore.v1.GetKeyResponse get_key_reply = 15 [json_name = "getKeyReply"];</code>
   * @return Whether the getKeyReply field is set.
   */
  boolean hasGetKeyReply();
  /**
   * <code>.kvstore.v1.GetKeyResponse get_key_reply = 15 [json_name = "getKeyReply"];</code>
   * @return The getKeyReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetKeyResponse getGetKeyReply();
  /**
   * <code>.kvstore.v1.GetKeyResponse get_key_reply = 15 [json_name = "getKeyReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetKeyResponseOrBuilder getGetKeyReplyOrBuilder();

  /**
   * <code>.kvstore.v1.PutKeyRequest put_key_request = 16 [json_name = "putKeyRequest"];</code>
   * @return Whether the putKeyRequest field is set.
   */
  boolean hasPutKeyRequest();
  /**
   * <code>.kvstore.v1.PutKeyRequest put_key_request = 16 [json_name = "putKeyRequest"];</code>
   * @return The putKeyRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.PutKeyRequest getPutKeyRequest();
  /**
   * <code>.kvstore.v1.PutKeyRequest put_key_request = 16 [json_name = "putKeyRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.PutKeyRequestOrBuilder getPutKeyRequestOrBuilder();

  /**
   * <code>.kvstore.v1.PutKeyResponse put_key_reply = 17 [json_name = "putKeyReply"];</code>
   * @return Whether the putKeyReply field is set.
   */
  boolean hasPutKeyReply();
  /**
   * <code>.kvstore.v1.PutKeyResponse put_key_reply = 17 [json_name = "putKeyReply"];</code>
   * @return The putKeyReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.PutKeyResponse getPutKeyReply();
  /**
   * <code>.kvstore.v1.PutKeyResponse put_key_reply = 17 [json_name = "putKeyReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.PutKeyResponseOrBuilder getPutKeyReplyOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteKeyRequest delete_key_request = 18 [json_name = "deleteKeyRequest"];</code>
   * @return Whether the deleteKeyRequest field is set.
   */
  boolean hasDeleteKeyRequest();
  /**
   * <code>.kvstore.v1.DeleteKeyRequest delete_key_request = 18 [json_name = "deleteKeyRequest"];</code>
   * @return The deleteKeyRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyRequest getDeleteKeyRequest();
  /**
   * <code>.kvstore.v1.DeleteKeyRequest delete_key_request = 18 [json_name = "deleteKeyRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyRequestOrBuilder getDeleteKeyRequestOrBuilder();

  /**
   * <code>.kvstore.v1.DeleteKeyResponse delete_key_reply = 19 [json_name = "deleteKeyReply"];</code>
   * @return Whether the deleteKeyReply field is set.
   */
  boolean hasDeleteKeyReply();
  /**
   * <code>.kvstore.v1.DeleteKeyResponse delete_key_reply = 19 [json_name = "deleteKeyReply"];</code>
   * @return The deleteKeyReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse getDeleteKeyReply();
  /**
   * <code>.kvstore.v1.DeleteKeyResponse delete_key_reply = 19 [json_name = "deleteKeyReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponseOrBuilder getDeleteKeyReplyOrBuilder();

  /**
   * <code>.kvstore.v1.GetBucketDescriptorRequest get_bucket_descriptor_request = 20 [json_name = "getBucketDescriptorRequest"];</code>
   * @return Whether the getBucketDescriptorRequest field is set.
   */
  boolean hasGetBucketDescriptorRequest();
  /**
   * <code>.kvstore.v1.GetBucketDescriptorRequest get_bucket_descriptor_request = 20 [json_name = "getBucketDescriptorRequest"];</code>
   * @return The getBucketDescriptorRequest.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetBucketDescriptorRequest getGetBucketDescriptorRequest();
  /**
   * <code>.kvstore.v1.GetBucketDescriptorRequest get_bucket_descriptor_request = 20 [json_name = "getBucketDescriptorRequest"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetBucketDescriptorRequestOrBuilder getGetBucketDescriptorRequestOrBuilder();

  /**
   * <code>.kvstore.v1.GetBucketDescriptorResponse get_bucket_descriptor_reply = 21 [json_name = "getBucketDescriptorReply"];</code>
   * @return Whether the getBucketDescriptorReply field is set.
   */
  boolean hasGetBucketDescriptorReply();
  /**
   * <code>.kvstore.v1.GetBucketDescriptorResponse get_bucket_descriptor_reply = 21 [json_name = "getBucketDescriptorReply"];</code>
   * @return The getBucketDescriptorReply.
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetBucketDescriptorResponse getGetBucketDescriptorReply();
  /**
   * <code>.kvstore.v1.GetBucketDescriptorResponse get_bucket_descriptor_reply = 21 [json_name = "getBucketDescriptorReply"];</code>
   */
  com.github.mxplusb.pleiades.api.kvstore.v1.GetBucketDescriptorResponseOrBuilder getGetBucketDescriptorReplyOrBuilder();

  /**
   * <code>.errors.v1.Error error = 22 [json_name = "error"];</code>
   * @return Whether the error field is set.
   */
  boolean hasError();
  /**
   * <code>.errors.v1.Error error = 22 [json_name = "error"];</code>
   * @return The error.
   */
  com.github.mxplusb.pleiades.api.errors.v1.Error getError();
  /**
   * <code>.errors.v1.Error error = 22 [json_name = "error"];</code>
   */
  com.github.mxplusb.pleiades.api.errors.v1.ErrorOrBuilder getErrorOrBuilder();

  com.github.mxplusb.pleiades.api.kvstore.v1.KVStoreWrapper.PayloadCase getPayloadCase();
}