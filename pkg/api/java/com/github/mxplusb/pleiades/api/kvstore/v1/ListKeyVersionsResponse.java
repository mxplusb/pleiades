// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

/**
 * Protobuf type {@code kvstore.v1.ListKeyVersionsResponse}
 */
public final class ListKeyVersionsResponse extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:kvstore.v1.ListKeyVersionsResponse)
    ListKeyVersionsResponseOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ListKeyVersionsResponse.newBuilder() to construct.
  private ListKeyVersionsResponse(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ListKeyVersionsResponse() {
    versions_ = emptyIntList();
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ListKeyVersionsResponse();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_ListKeyVersionsResponse_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_ListKeyVersionsResponse_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.class, com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.Builder.class);
  }

  public static final int VERSIONS_FIELD_NUMBER = 1;
  @SuppressWarnings("serial")
  private com.google.protobuf.Internal.IntList versions_;
  /**
   * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
   * @return A list containing the versions.
   */
  @java.lang.Override
  public java.util.List<java.lang.Integer>
      getVersionsList() {
    return versions_;
  }
  /**
   * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
   * @return The count of versions.
   */
  public int getVersionsCount() {
    return versions_.size();
  }
  /**
   * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
   * @param index The index of the element to return.
   * @return The versions at the given index.
   */
  public int getVersions(int index) {
    return versions_.getInt(index);
  }
  private int versionsMemoizedSerializedSize = -1;

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
    getSerializedSize();
    if (getVersionsList().size() > 0) {
      output.writeUInt32NoTag(10);
      output.writeUInt32NoTag(versionsMemoizedSerializedSize);
    }
    for (int i = 0; i < versions_.size(); i++) {
      output.writeUInt32NoTag(versions_.getInt(i));
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    {
      int dataSize = 0;
      for (int i = 0; i < versions_.size(); i++) {
        dataSize += com.google.protobuf.CodedOutputStream
          .computeUInt32SizeNoTag(versions_.getInt(i));
      }
      size += dataSize;
      if (!getVersionsList().isEmpty()) {
        size += 1;
        size += com.google.protobuf.CodedOutputStream
            .computeInt32SizeNoTag(dataSize);
      }
      versionsMemoizedSerializedSize = dataSize;
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
    if (!(obj instanceof com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse)) {
      return super.equals(obj);
    }
    com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse other = (com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse) obj;

    if (!getVersionsList()
        .equals(other.getVersionsList())) return false;
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
    if (getVersionsCount() > 0) {
      hash = (37 * hash) + VERSIONS_FIELD_NUMBER;
      hash = (53 * hash) + getVersionsList().hashCode();
    }
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse parseFrom(
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
  public static Builder newBuilder(com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse prototype) {
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
   * Protobuf type {@code kvstore.v1.ListKeyVersionsResponse}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:kvstore.v1.ListKeyVersionsResponse)
      com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponseOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_ListKeyVersionsResponse_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_ListKeyVersionsResponse_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.class, com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.Builder.class);
    }

    // Construct using com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.newBuilder()
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
      versions_ = emptyIntList();
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_ListKeyVersionsResponse_descriptor;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse getDefaultInstanceForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse build() {
      com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse buildPartial() {
      com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse result = new com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse(this);
      buildPartialRepeatedFields(result);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartialRepeatedFields(com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse result) {
      if (((bitField0_ & 0x00000001) != 0)) {
        versions_.makeImmutable();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.versions_ = versions_;
    }

    private void buildPartial0(com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse result) {
      int from_bitField0_ = bitField0_;
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
      if (other instanceof com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse) {
        return mergeFrom((com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse other) {
      if (other == com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse.getDefaultInstance()) return this;
      if (!other.versions_.isEmpty()) {
        if (versions_.isEmpty()) {
          versions_ = other.versions_;
          bitField0_ = (bitField0_ & ~0x00000001);
        } else {
          ensureVersionsIsMutable();
          versions_.addAll(other.versions_);
        }
        onChanged();
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
              int v = input.readUInt32();
              ensureVersionsIsMutable();
              versions_.addInt(v);
              break;
            } // case 8
            case 10: {
              int length = input.readRawVarint32();
              int limit = input.pushLimit(length);
              ensureVersionsIsMutable();
              while (input.getBytesUntilLimit() > 0) {
                versions_.addInt(input.readUInt32());
              }
              input.popLimit(limit);
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

    private com.google.protobuf.Internal.IntList versions_ = emptyIntList();
    private void ensureVersionsIsMutable() {
      if (!((bitField0_ & 0x00000001) != 0)) {
        versions_ = mutableCopy(versions_);
        bitField0_ |= 0x00000001;
      }
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @return A list containing the versions.
     */
    public java.util.List<java.lang.Integer>
        getVersionsList() {
      return ((bitField0_ & 0x00000001) != 0) ?
               java.util.Collections.unmodifiableList(versions_) : versions_;
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @return The count of versions.
     */
    public int getVersionsCount() {
      return versions_.size();
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @param index The index of the element to return.
     * @return The versions at the given index.
     */
    public int getVersions(int index) {
      return versions_.getInt(index);
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @param index The index to set the value at.
     * @param value The versions to set.
     * @return This builder for chaining.
     */
    public Builder setVersions(
        int index, int value) {

      ensureVersionsIsMutable();
      versions_.setInt(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @param value The versions to add.
     * @return This builder for chaining.
     */
    public Builder addVersions(int value) {

      ensureVersionsIsMutable();
      versions_.addInt(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @param values The versions to add.
     * @return This builder for chaining.
     */
    public Builder addAllVersions(
        java.lang.Iterable<? extends java.lang.Integer> values) {
      ensureVersionsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, versions_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated uint32 versions = 1 [json_name = "versions"];</code>
     * @return This builder for chaining.
     */
    public Builder clearVersions() {
      versions_ = emptyIntList();
      bitField0_ = (bitField0_ & ~0x00000001);
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


    // @@protoc_insertion_point(builder_scope:kvstore.v1.ListKeyVersionsResponse)
  }

  // @@protoc_insertion_point(class_scope:kvstore.v1.ListKeyVersionsResponse)
  private static final com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse();
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ListKeyVersionsResponse>
      PARSER = new com.google.protobuf.AbstractParser<ListKeyVersionsResponse>() {
    @java.lang.Override
    public ListKeyVersionsResponse parsePartialFrom(
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

  public static com.google.protobuf.Parser<ListKeyVersionsResponse> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ListKeyVersionsResponse> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.mxplusb.pleiades.api.kvstore.v1.ListKeyVersionsResponse getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

