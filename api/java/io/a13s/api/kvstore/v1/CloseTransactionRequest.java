// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/transactions.proto

package io.a13s.api.kvstore.v1;

/**
 * Protobuf type {@code kvstore.v1.CloseTransactionRequest}
 */
public final class CloseTransactionRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:kvstore.v1.CloseTransactionRequest)
    CloseTransactionRequestOrBuilder {
private static final long serialVersionUID = 0L;
  // Use CloseTransactionRequest.newBuilder() to construct.
  private CloseTransactionRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private CloseTransactionRequest() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new CloseTransactionRequest();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.a13s.api.kvstore.v1.TransactionsProto.internal_static_kvstore_v1_CloseTransactionRequest_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.a13s.api.kvstore.v1.TransactionsProto.internal_static_kvstore_v1_CloseTransactionRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.a13s.api.kvstore.v1.CloseTransactionRequest.class, io.a13s.api.kvstore.v1.CloseTransactionRequest.Builder.class);
  }

  public static final int TRANSACTION_FIELD_NUMBER = 1;
  private io.a13s.api.kvstore.v1.Transaction transaction_;
  /**
   * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
   * @return Whether the transaction field is set.
   */
  @java.lang.Override
  public boolean hasTransaction() {
    return transaction_ != null;
  }
  /**
   * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
   * @return The transaction.
   */
  @java.lang.Override
  public io.a13s.api.kvstore.v1.Transaction getTransaction() {
    return transaction_ == null ? io.a13s.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
  }
  /**
   * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
   */
  @java.lang.Override
  public io.a13s.api.kvstore.v1.TransactionOrBuilder getTransactionOrBuilder() {
    return getTransaction();
  }

  public static final int TIMEOUT_FIELD_NUMBER = 2;
  private long timeout_;
  /**
   * <code>int64 timeout = 2 [json_name = "timeout"];</code>
   * @return The timeout.
   */
  @java.lang.Override
  public long getTimeout() {
    return timeout_;
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
    if (transaction_ != null) {
      output.writeMessage(1, getTransaction());
    }
    if (timeout_ != 0L) {
      output.writeInt64(2, timeout_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (transaction_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getTransaction());
    }
    if (timeout_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(2, timeout_);
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
    if (!(obj instanceof io.a13s.api.kvstore.v1.CloseTransactionRequest)) {
      return super.equals(obj);
    }
    io.a13s.api.kvstore.v1.CloseTransactionRequest other = (io.a13s.api.kvstore.v1.CloseTransactionRequest) obj;

    if (hasTransaction() != other.hasTransaction()) return false;
    if (hasTransaction()) {
      if (!getTransaction()
          .equals(other.getTransaction())) return false;
    }
    if (getTimeout()
        != other.getTimeout()) return false;
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
    if (hasTransaction()) {
      hash = (37 * hash) + TRANSACTION_FIELD_NUMBER;
      hash = (53 * hash) + getTransaction().hashCode();
    }
    hash = (37 * hash) + TIMEOUT_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getTimeout());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.a13s.api.kvstore.v1.CloseTransactionRequest parseFrom(
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
  public static Builder newBuilder(io.a13s.api.kvstore.v1.CloseTransactionRequest prototype) {
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
   * Protobuf type {@code kvstore.v1.CloseTransactionRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:kvstore.v1.CloseTransactionRequest)
      io.a13s.api.kvstore.v1.CloseTransactionRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.a13s.api.kvstore.v1.TransactionsProto.internal_static_kvstore_v1_CloseTransactionRequest_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.a13s.api.kvstore.v1.TransactionsProto.internal_static_kvstore_v1_CloseTransactionRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.a13s.api.kvstore.v1.CloseTransactionRequest.class, io.a13s.api.kvstore.v1.CloseTransactionRequest.Builder.class);
    }

    // Construct using io.a13s.api.kvstore.v1.CloseTransactionRequest.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      if (transactionBuilder_ == null) {
        transaction_ = null;
      } else {
        transaction_ = null;
        transactionBuilder_ = null;
      }
      timeout_ = 0L;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.a13s.api.kvstore.v1.TransactionsProto.internal_static_kvstore_v1_CloseTransactionRequest_descriptor;
    }

    @java.lang.Override
    public io.a13s.api.kvstore.v1.CloseTransactionRequest getDefaultInstanceForType() {
      return io.a13s.api.kvstore.v1.CloseTransactionRequest.getDefaultInstance();
    }

    @java.lang.Override
    public io.a13s.api.kvstore.v1.CloseTransactionRequest build() {
      io.a13s.api.kvstore.v1.CloseTransactionRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.a13s.api.kvstore.v1.CloseTransactionRequest buildPartial() {
      io.a13s.api.kvstore.v1.CloseTransactionRequest result = new io.a13s.api.kvstore.v1.CloseTransactionRequest(this);
      if (transactionBuilder_ == null) {
        result.transaction_ = transaction_;
      } else {
        result.transaction_ = transactionBuilder_.build();
      }
      result.timeout_ = timeout_;
      onBuilt();
      return result;
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
      if (other instanceof io.a13s.api.kvstore.v1.CloseTransactionRequest) {
        return mergeFrom((io.a13s.api.kvstore.v1.CloseTransactionRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.a13s.api.kvstore.v1.CloseTransactionRequest other) {
      if (other == io.a13s.api.kvstore.v1.CloseTransactionRequest.getDefaultInstance()) return this;
      if (other.hasTransaction()) {
        mergeTransaction(other.getTransaction());
      }
      if (other.getTimeout() != 0L) {
        setTimeout(other.getTimeout());
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
            case 10: {
              input.readMessage(
                  getTransactionFieldBuilder().getBuilder(),
                  extensionRegistry);

              break;
            } // case 10
            case 16: {
              timeout_ = input.readInt64();

              break;
            } // case 16
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

    private io.a13s.api.kvstore.v1.Transaction transaction_;
    private com.google.protobuf.SingleFieldBuilderV3<
        io.a13s.api.kvstore.v1.Transaction, io.a13s.api.kvstore.v1.Transaction.Builder, io.a13s.api.kvstore.v1.TransactionOrBuilder> transactionBuilder_;
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     * @return Whether the transaction field is set.
     */
    public boolean hasTransaction() {
      return transactionBuilder_ != null || transaction_ != null;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     * @return The transaction.
     */
    public io.a13s.api.kvstore.v1.Transaction getTransaction() {
      if (transactionBuilder_ == null) {
        return transaction_ == null ? io.a13s.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
      } else {
        return transactionBuilder_.getMessage();
      }
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public Builder setTransaction(io.a13s.api.kvstore.v1.Transaction value) {
      if (transactionBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        transaction_ = value;
        onChanged();
      } else {
        transactionBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public Builder setTransaction(
        io.a13s.api.kvstore.v1.Transaction.Builder builderForValue) {
      if (transactionBuilder_ == null) {
        transaction_ = builderForValue.build();
        onChanged();
      } else {
        transactionBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public Builder mergeTransaction(io.a13s.api.kvstore.v1.Transaction value) {
      if (transactionBuilder_ == null) {
        if (transaction_ != null) {
          transaction_ =
            io.a13s.api.kvstore.v1.Transaction.newBuilder(transaction_).mergeFrom(value).buildPartial();
        } else {
          transaction_ = value;
        }
        onChanged();
      } else {
        transactionBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public Builder clearTransaction() {
      if (transactionBuilder_ == null) {
        transaction_ = null;
        onChanged();
      } else {
        transaction_ = null;
        transactionBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public io.a13s.api.kvstore.v1.Transaction.Builder getTransactionBuilder() {
      
      onChanged();
      return getTransactionFieldBuilder().getBuilder();
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    public io.a13s.api.kvstore.v1.TransactionOrBuilder getTransactionOrBuilder() {
      if (transactionBuilder_ != null) {
        return transactionBuilder_.getMessageOrBuilder();
      } else {
        return transaction_ == null ?
            io.a13s.api.kvstore.v1.Transaction.getDefaultInstance() : transaction_;
      }
    }
    /**
     * <code>.kvstore.v1.Transaction transaction = 1 [json_name = "transaction"];</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        io.a13s.api.kvstore.v1.Transaction, io.a13s.api.kvstore.v1.Transaction.Builder, io.a13s.api.kvstore.v1.TransactionOrBuilder> 
        getTransactionFieldBuilder() {
      if (transactionBuilder_ == null) {
        transactionBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            io.a13s.api.kvstore.v1.Transaction, io.a13s.api.kvstore.v1.Transaction.Builder, io.a13s.api.kvstore.v1.TransactionOrBuilder>(
                getTransaction(),
                getParentForChildren(),
                isClean());
        transaction_ = null;
      }
      return transactionBuilder_;
    }

    private long timeout_ ;
    /**
     * <code>int64 timeout = 2 [json_name = "timeout"];</code>
     * @return The timeout.
     */
    @java.lang.Override
    public long getTimeout() {
      return timeout_;
    }
    /**
     * <code>int64 timeout = 2 [json_name = "timeout"];</code>
     * @param value The timeout to set.
     * @return This builder for chaining.
     */
    public Builder setTimeout(long value) {
      
      timeout_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 timeout = 2 [json_name = "timeout"];</code>
     * @return This builder for chaining.
     */
    public Builder clearTimeout() {
      
      timeout_ = 0L;
      onChanged();
      return this;
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


    // @@protoc_insertion_point(builder_scope:kvstore.v1.CloseTransactionRequest)
  }

  // @@protoc_insertion_point(class_scope:kvstore.v1.CloseTransactionRequest)
  private static final io.a13s.api.kvstore.v1.CloseTransactionRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.a13s.api.kvstore.v1.CloseTransactionRequest();
  }

  public static io.a13s.api.kvstore.v1.CloseTransactionRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<CloseTransactionRequest>
      PARSER = new com.google.protobuf.AbstractParser<CloseTransactionRequest>() {
    @java.lang.Override
    public CloseTransactionRequest parsePartialFrom(
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

  public static com.google.protobuf.Parser<CloseTransactionRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<CloseTransactionRequest> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.a13s.api.kvstore.v1.CloseTransactionRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

