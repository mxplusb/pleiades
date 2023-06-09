// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_host.proto

package com.github.mxplusb.pleiades.api.raft.v1;

/**
 * Protobuf type {@code raft.v1.GetHostConfigResponse}
 */
public final class GetHostConfigResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:raft.v1.GetHostConfigResponse)
    GetHostConfigResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use GetHostConfigResponse.newBuilder() to construct.
  private GetHostConfigResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GetHostConfigResponse() {
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new GetHostConfigResponse();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.mxplusb.pleiades.api.raft.v1.RaftHostProto.internal_static_raft_v1_GetHostConfigResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.mxplusb.pleiades.api.raft.v1.RaftHostProto.internal_static_raft_v1_GetHostConfigResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.class, com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.Builder.class);
  }

  public static final int CONFIG_FIELD_NUMBER = 1;
  private com.github.mxplusb.pleiades.api.raft.v1.HostConfig config_;
  /**
   * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
   * @return Whether the config field is set.
   */
  @java.lang.Override
  public boolean hasConfig() {
    return config_ != null;
  }
  /**
   * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
   * @return The config.
   */
  @java.lang.Override
  public com.github.mxplusb.pleiades.api.raft.v1.HostConfig getConfig() {
    return config_ == null ? com.github.mxplusb.pleiades.api.raft.v1.HostConfig.getDefaultInstance() : config_;
  }
  /**
   * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
   */
  @java.lang.Override
  public com.github.mxplusb.pleiades.api.raft.v1.HostConfigOrBuilder getConfigOrBuilder() {
    return config_ == null ? com.github.mxplusb.pleiades.api.raft.v1.HostConfig.getDefaultInstance() : config_;
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
    if (config_ != null) {
      output.writeMessage(1, getConfig());
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (config_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(1, getConfig());
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
    if (!(obj instanceof com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse)) {
      return super.equals(obj);
    }
    com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse other = (com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse) obj;

    if (hasConfig() != other.hasConfig()) return false;
    if (hasConfig()) {
      if (!getConfig()
          .equals(other.getConfig())) return false;
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
    if (hasConfig()) {
      hash = (37 * hash) + CONFIG_FIELD_NUMBER;
      hash = (53 * hash) + getConfig().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse parseFrom(
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
  public static Builder newBuilder(com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse prototype) {
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
   * Protobuf type {@code raft.v1.GetHostConfigResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:raft.v1.GetHostConfigResponse)
      com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftHostProto.internal_static_raft_v1_GetHostConfigResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftHostProto.internal_static_raft_v1_GetHostConfigResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.class, com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.Builder.class);
    }

    // Construct using com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.newBuilder()
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
      config_ = null;
      if (configBuilder_ != null) {
        configBuilder_.dispose();
        configBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.mxplusb.pleiades.api.raft.v1.RaftHostProto.internal_static_raft_v1_GetHostConfigResponse_descriptor;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse getDefaultInstanceForType() {
      return com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse build() {
      com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse buildPartial() {
      com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse result = new com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse(this);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartial0(com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000001) != 0)) {
        result.config_ = configBuilder_ == null
            ? config_
            : configBuilder_.build();
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
      if (other instanceof com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse) {
        return mergeFrom((com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse other) {
      if (other == com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse.getDefaultInstance()) return this;
      if (other.hasConfig()) {
        mergeConfig(other.getConfig());
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
                  getConfigFieldBuilder().getBuilder(),
                  extensionRegistry);
              bitField0_ |= 0x00000001;
              break;
            } // case 10
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

    private com.github.mxplusb.pleiades.api.raft.v1.HostConfig config_;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.github.mxplusb.pleiades.api.raft.v1.HostConfig, com.github.mxplusb.pleiades.api.raft.v1.HostConfig.Builder, com.github.mxplusb.pleiades.api.raft.v1.HostConfigOrBuilder> configBuilder_;
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     * @return Whether the config field is set.
     */
    public boolean hasConfig() {
      return ((bitField0_ & 0x00000001) != 0);
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     * @return The config.
     */
    public com.github.mxplusb.pleiades.api.raft.v1.HostConfig getConfig() {
      if (configBuilder_ == null) {
        return config_ == null ? com.github.mxplusb.pleiades.api.raft.v1.HostConfig.getDefaultInstance() : config_;
      } else {
        return configBuilder_.getMessage();
      }
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public Builder setConfig(com.github.mxplusb.pleiades.api.raft.v1.HostConfig value) {
      if (configBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        config_ = value;
      } else {
        configBuilder_.setMessage(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public Builder setConfig(
        com.github.mxplusb.pleiades.api.raft.v1.HostConfig.Builder builderForValue) {
      if (configBuilder_ == null) {
        config_ = builderForValue.build();
      } else {
        configBuilder_.setMessage(builderForValue.build());
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public Builder mergeConfig(com.github.mxplusb.pleiades.api.raft.v1.HostConfig value) {
      if (configBuilder_ == null) {
        if (((bitField0_ & 0x00000001) != 0) &&
          config_ != null &&
          config_ != com.github.mxplusb.pleiades.api.raft.v1.HostConfig.getDefaultInstance()) {
          getConfigBuilder().mergeFrom(value);
        } else {
          config_ = value;
        }
      } else {
        configBuilder_.mergeFrom(value);
      }
      bitField0_ |= 0x00000001;
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public Builder clearConfig() {
      bitField0_ = (bitField0_ & ~0x00000001);
      config_ = null;
      if (configBuilder_ != null) {
        configBuilder_.dispose();
        configBuilder_ = null;
      }
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public com.github.mxplusb.pleiades.api.raft.v1.HostConfig.Builder getConfigBuilder() {
      bitField0_ |= 0x00000001;
      onChanged();
      return getConfigFieldBuilder().getBuilder();
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    public com.github.mxplusb.pleiades.api.raft.v1.HostConfigOrBuilder getConfigOrBuilder() {
      if (configBuilder_ != null) {
        return configBuilder_.getMessageOrBuilder();
      } else {
        return config_ == null ?
            com.github.mxplusb.pleiades.api.raft.v1.HostConfig.getDefaultInstance() : config_;
      }
    }
    /**
     * <code>.raft.v1.HostConfig config = 1 [json_name = "config"];</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.github.mxplusb.pleiades.api.raft.v1.HostConfig, com.github.mxplusb.pleiades.api.raft.v1.HostConfig.Builder, com.github.mxplusb.pleiades.api.raft.v1.HostConfigOrBuilder> 
        getConfigFieldBuilder() {
      if (configBuilder_ == null) {
        configBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.github.mxplusb.pleiades.api.raft.v1.HostConfig, com.github.mxplusb.pleiades.api.raft.v1.HostConfig.Builder, com.github.mxplusb.pleiades.api.raft.v1.HostConfigOrBuilder>(
                getConfig(),
                getParentForChildren(),
                isClean());
        config_ = null;
      }
      return configBuilder_;
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


    // @@protoc_insertion_point(builder_scope:raft.v1.GetHostConfigResponse)
  }

  // @@protoc_insertion_point(class_scope:raft.v1.GetHostConfigResponse)
  private static final com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse();
  }

  public static com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GetHostConfigResponse>
      PARSER = new com.google.protobuf.AbstractParser<GetHostConfigResponse>() {
    @java.lang.Override
    public GetHostConfigResponse parsePartialFrom(
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

  public static com.google.protobuf.Parser<GetHostConfigResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GetHostConfigResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.mxplusb.pleiades.api.raft.v1.GetHostConfigResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

