// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

/**
 * Protobuf type {@code kvstore.v1.DeleteKeyResponse}
 */
public final class DeleteKeyResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:kvstore.v1.DeleteKeyResponse)
    DeleteKeyResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use DeleteKeyResponse.newBuilder() to construct.
  private DeleteKeyResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private DeleteKeyResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new DeleteKeyResponse();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_DeleteKeyResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_DeleteKeyResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.class, com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.Builder.class);
  }

  public static final int OK_FIELD_NUMBER = 1;
  private boolean ok_ = false;
  /**
   * <code>bool ok = 1 [json_name = "ok"];</code>
   * @return The ok.
   */
  @java.lang.Override
  public boolean getOk() {
    return ok_;
  }

  public static final int TRANSACTION_FIELD_NUMBER = 3;
  private com.github.mxplusb.pleiades.api.kvstore.v1.Transaction transaction_;
  /**
   * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
   * @return Whether the transaction field is set.
   */
  @java.lang.Override
  public boolean hasTransaction() {
    return transaction_ != null;
  }
  /**
   * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
   * @return The transaction.
   */
  @java.lang.Override
  public com.github.mxplusb.pleiades.api.kvstore.v1.Transaction getTransaction() {
    return transaction_ == null ? com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
  }
  /**
   * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
   */
  @java.lang.Override
  public com.github.mxplusb.pleiades.api.kvstore.v1.TransactionOrBuilder getTransactionOrBuilder() {
    return transaction_ == null ? com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (ok_ != false) {
      output.writeBool(1, ok_);
    }
    if (transaction_ != null) {
      output.writeMessage(3, getTransaction());
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (ok_ != false) {
      size += com.google.protobuf.CodedOutputStream
        .computeBoolSize(1, ok_);
    }
    if (transaction_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(3, getTransaction());
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse)) {
      return super.equals(obj);
    }
    com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse other = (com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse) obj;

    if (getOk()
        != other.getOk()) return false;
    if (hasTransaction() != other.hasTransaction()) return false;
    if (hasTransaction()) {
      if (!getTransaction()
          .equals(other.getTransaction())) return false;
    }
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + OK_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
        getOk());
    if (hasTransaction()) {
      hash = (37 * hash) + TRANSACTION_FIELD_NUMBER;
      hash = (53 * hash) + getTransaction().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code kvstore.v1.DeleteKeyResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:kvstore.v1.DeleteKeyResponse)
      com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_DeleteKeyResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_DeleteKeyResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.class, com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.Builder.class);
    }

    // Construct using com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      bitField0_ = 0;
      ok_ = false;
      transaction_ = null;
      if (transactionBuilder_ != null) {
        transactionBuilder_.dispose();
        transactionBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_DeleteKeyResponse_descriptor;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse getDefaultInstanceForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse build() {
      com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse buildPartial() {
      com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse result = new com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.ok_ = ok_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.transaction_ = transactionBuilder_ == null
            ? transaction_
            : transactionBuilder_.build();
      }
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse) {
        return mergeFrom((com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse other) {
      if (other == com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse.getDefaultInstance()) return this;
      if (other.getOk() != false) {
        setOk(other.getOk());
      }
      if (other.hasTransaction()) {
        mergeTransaction(other.getTransaction());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {
              ok_ = input.readBool();
              bitField0_ |= 0x00000001;
              break;
            } // case 8
            case 26: {
              input.readMessage(
                  getTransactionFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000002;
              break;
            } // case 26
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }
    private int bitField0_;

    private boolean ok_ ;
    /**
     * <code>bool ok = 1 [json_name = "ok"];</code>
     * @return The ok.
     */
    @java.lang.Override
    public boolean getOk() {
      return ok_;
    }
    /**
     * <code>bool ok = 1 [json_name = "ok"];</code>
     * @param value The ok to set.
     * @return This builder for chaining.
     */
    public Builder setOk(boolean value) {

      ok_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>bool ok = 1 [json_name = "ok"];</code>
     * @return This builder for chaining.
     */
    public Builder clearOk() {
      bitField0_ = (bitField0_ & ~0x00000001);
      ok_ = false;
      onChanged();
      return this;
    }

    private com.github.mxplusb.pleiades.api.kvstore.v1.Transaction transaction_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.github.mxplusb.pleiades.api.kvstore.v1.Transaction, com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.Builder, com.github.mxplusb.pleiades.api.kvstore.v1.TransactionOrBuilder> transactionBuilder_;
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     * @return Whether the transaction field is set.
     */
    public boolean hasTransaction() {
      return ((bitField0_ & 0x00000002) != 0);
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     * @return The transaction.
     */
    public com.github.mxplusb.pleiades.api.kvstore.v1.Transaction getTransaction() {
      if (transactionBuilder_ == null) {
        return transaction_ == null ? com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
      } else {
        return transactionBuilder_.getMessage();
      }
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public Builder setTransaction(com.github.mxplusb.pleiades.api.kvstore.v1.Transaction value) {
      if (transactionBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        transaction_ = value;
      } else {
        transactionBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public Builder setTransaction(
        com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.Builder builderForValue) {
      if (transactionBuilder_ == null) {
        transaction_ = builderForValue.build();
      } else {
        transactionBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public Builder mergeTransaction(com.github.mxplusb.pleiades.api.kvstore.v1.Transaction value) {
      if (transactionBuilder_ == null) {
        if (((bitField0_ & 0x00000002) != 0) &&
          transaction_ != null &&
          transaction_ != com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.getDefaultInstance()) {
          getTransactionBuilder().mergeFrom(value);
        } else {
          transaction_ = value;
        }
      } else {
        transactionBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public Builder clearTransaction() {
      bitField0_ = (bitField0_ & ~0x00000002);
      transaction_ = null;
      if (transactionBuilder_ != null) {
        transactionBuilder_.dispose();
        transactionBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.Builder getTransactionBuilder() {
      bitField0_ |= 0x00000002;
      onChanged();
      return getTransactionFieldBuilder().getBuilder();
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    public com.github.mxplusb.pleiades.api.kvstore.v1.TransactionOrBuilder getTransactionOrBuilder() {
      if (transactionBuilder_ != null) {
        return transactionBuilder_.getMessageOrBuilder();
      } else {
        return transaction_ == null ?
            com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
      }
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 3 [json_name = "transaction"];</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.github.mxplusb.pleiades.api.kvstore.v1.Transaction, com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.Builder, com.github.mxplusb.pleiades.api.kvstore.v1.TransactionOrBuilder> 
        getTransactionFieldBuilder() {
      if (transactionBuilder_ == null) {
        transactionBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.github.mxplusb.pleiades.api.kvstore.v1.Transaction, com.github.mxplusb.pleiades.api.kvstore.v1.Transaction.Builder, com.github.mxplusb.pleiades.api.kvstore.v1.TransactionOrBuilder>(
                getTransaction(),
                getParentForChildren(),
                isClean());
        transaction_ = null;
      }
      return transactionBuilder_;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:kvstore.v1.DeleteKeyResponse)
  }

  // @@protoc_insertion_point(class_scope:kvstore.v1.DeleteKeyResponse)
  private static final com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse();
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<DeleteKeyResponse>
      PARSER = new com.google.protobuf.AbstractParser<DeleteKeyResponse>() {
    @java.lang.Override
    public DeleteKeyResponse parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<DeleteKeyResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<DeleteKeyResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.mxplusb.pleiades.api.kvstore.v1.DeleteKeyResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

