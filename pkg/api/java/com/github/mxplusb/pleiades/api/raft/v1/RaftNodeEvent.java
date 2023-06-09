// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_event.proto

package com.github.mxplusb.pleiades.api.raft.v1;

/**
 * Protobuf type {@code raft.v1.RaftNodeEvent}
 */
public final class RaftNodeEvent extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:raft.v1.RaftNodeEvent)
    RaftNodeEventOrBuilder {
private static final long serialVersionUID = 0L;
  // Use RaftNodeEvent.newBuilder() to construct.
  private RaftNodeEvent(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private RaftNodeEvent() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new RaftNodeEvent();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.mxplusb.pleiades.api.raft.v1.RaftEventProto.internal_static_raft_v1_RaftNodeEvent_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.mxplusb.pleiades.api.raft.v1.RaftEventProto.internal_static_raft_v1_RaftNodeEvent_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.class, com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.Builder.class);
  }

  public static final int SHARD_ID_FIELD_NUMBER = 1;
  private long shardId_ = 0L;
  /**
   * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
   * @return The shardId.
   */
  @java.lang.Override
  public long getShardId() {
    return shardId_;
  }

  public static final int REPLICA_ID_FIELD_NUMBER = 2;
  private long replicaId_ = 0L;
  /**
   * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
   * @return The replicaId.
   */
  @java.lang.Override
  public long getReplicaId() {
    return replicaId_;
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
    if (shardId_ != 0L) {
      output.writeUInt64(1, shardId_);
    }
    if (replicaId_ != 0L) {
      output.writeUInt64(2, replicaId_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (shardId_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt64Size(1, shardId_);
    }
    if (replicaId_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeUInt64Size(2, replicaId_);
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
    if (!(obj instanceof com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent)) {
      return super.equals(obj);
    }
    com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent other = (com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent) obj;

    if (getShardId()
        != other.getShardId()) return false;
    if (getReplicaId()
        != other.getReplicaId()) return false;
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
    hash = (37 * hash) + SHARD_ID_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getShardId());
    hash = (37 * hash) + REPLICA_ID_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getReplicaId());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent parseFrom(
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
  public static Builder newBuilder(com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent prototype) {
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
   * Protobuf type {@code raft.v1.RaftNodeEvent}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:raft.v1.RaftNodeEvent)
      com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEventOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftEventProto.internal_static_raft_v1_RaftNodeEvent_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftEventProto.internal_static_raft_v1_RaftNodeEvent_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.class, com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.Builder.class);
    }

    // Construct using com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.newBuilder()
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
      shardId_ = 0L;
      replicaId_ = 0L;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftEventProto.internal_static_raft_v1_RaftNodeEvent_descriptor;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent getDefaultInstanceForType() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent build() {
      com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent buildPartial() {
      com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent result = new com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.shardId_ = shardId_;
      }
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.replicaId_ = replicaId_;
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
      if (other instanceof com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent) {
        return mergeFrom((com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent other) {
      if (other == com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent.getDefaultInstance()) return this;
      if (other.getShardId() != 0L) {
        setShardId(other.getShardId());
      }
      if (other.getReplicaId() != 0L) {
        setReplicaId(other.getReplicaId());
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
              shardId_ = input.readUInt64();
              bitField0_ |= 0x00000001;
              break;
            } // case 8
            case 16: {
              replicaId_ = input.readUInt64();
              bitField0_ |= 0x00000002;
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
    private int bitField0_;

    private long shardId_ ;
    /**
     * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
     * @return The shardId.
     */
    @java.lang.Override
    public long getShardId() {
      return shardId_;
    }
    /**
     * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
     * @param value The shardId to set.
     * @return This builder for chaining.
     */
    public Builder setShardId(long value) {

      shardId_ = value;
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>uint64 shard_id = 1 [json_name = "shardId"];</code>
     * @return This builder for chaining.
     */
    public Builder clearShardId() {
      bitField0_ = (bitField0_ & ~0x00000001);
      shardId_ = 0L;
      onChanged();
      return this;
    }

    private long replicaId_ ;
    /**
     * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
     * @return The replicaId.
     */
    @java.lang.Override
    public long getReplicaId() {
      return replicaId_;
    }
    /**
     * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
     * @param value The replicaId to set.
     * @return This builder for chaining.
     */
    public Builder setReplicaId(long value) {

      replicaId_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>uint64 replica_id = 2 [json_name = "replicaId"];</code>
     * @return This builder for chaining.
     */
    public Builder clearReplicaId() {
      bitField0_ = (bitField0_ & ~0x00000002);
      replicaId_ = 0L;
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


    // @@protoc_insertion_point(builder_scope:raft.v1.RaftNodeEvent)
  }

  // @@protoc_insertion_point(class_scope:raft.v1.RaftNodeEvent)
  private static final com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent();
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<RaftNodeEvent>
      PARSER = new com.google.protobuf.AbstractParser<RaftNodeEvent>() {
    @java.lang.Override
    public RaftNodeEvent parsePartialFrom(
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

  public static com.google.protobuf.Parser<RaftNodeEvent> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<RaftNodeEvent> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.mxplusb.pleiades.api.raft.v1.RaftNodeEvent getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

