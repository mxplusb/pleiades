// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: kvstore/v1/kv.proto

package com.github.mxplusb.pleiades.api.kvstore.v1;

/**
 * Protobuf type {@code kvstore.v1.KeyValueDescriptor}
 */
public final class KeyValueDescriptor extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:kvstore.v1.KeyValueDescriptor)
    KeyValueDescriptorOrBuilder {
private static final long serialVersionUID = 0L;
  // Use KeyValueDescriptor.newBuilder() to construct.
  private KeyValueDescriptor(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private KeyValueDescriptor() {
    versions_ = emptyIntList();
    currentKey_ = com.google.protobuf.ByteString.EMPTY;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new KeyValueDescriptor();
  }

  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_KeyValueDescriptor_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_KeyValueDescriptor_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.class, com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.Builder.class);
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

  public static final int CURRENT_KEY_FIELD_NUMBER = 2;
  private com.google.protobuf.ByteString currentKey_ = com.google.protobuf.ByteString.EMPTY;
  /**
   * <code>bytes current_key = 2 [json_name = "currentKey"];</code>
   * @return The currentKey.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString getCurrentKey() {
    return currentKey_;
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
    getSerializedSize();
    if (getVersionsList().size() > 0) {
      output.writeUInt32NoTag(10);
      output.writeUInt32NoTag(versionsMemoizedSerializedSize);
    }
    for (int i = 0; i < versions_.size(); i++) {
      output.writeUInt32NoTag(versions_.getInt(i));
    }
    if (!currentKey_.isEmpty()) {
      output.writeBytes(2, currentKey_);
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
    if (!currentKey_.isEmpty()) {
      size += com.google.protobuf.CodedOutputStream
        .computeBytesSize(2, currentKey_);
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
    if (!(obj instanceof com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor)) {
      return super.equals(obj);
    }
    com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor other = (com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor) obj;

    if (!getVersionsList()
        .equals(other.getVersionsList())) return false;
    if (!getCurrentKey()
        .equals(other.getCurrentKey())) return false;
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
    hash = (37 * hash) + CURRENT_KEY_FIELD_NUMBER;
    hash = (53 * hash) + getCurrentKey().hashCode();
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor parseFrom(
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
  public static Builder newBuilder(com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor prototype) {
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
   * Protobuf type {@code kvstore.v1.KeyValueDescriptor}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:kvstore.v1.KeyValueDescriptor)
      com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptorOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_KeyValueDescriptor_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_KeyValueDescriptor_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.class, com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.Builder.class);
    }

    // Construct using com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.newBuilder()
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
      currentKey_ = com.google.protobuf.ByteString.EMPTY;
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KvProto.internal_static_kvstore_v1_KeyValueDescriptor_descriptor;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor getDefaultInstanceForType() {
      return com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.getDefaultInstance();
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor build() {
      com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor buildPartial() {
      com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor result = new com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor(this);
      buildPartialRepeatedFields(result);
      if (bitField0_ != 0) { buildPartial0(result); }
      onBuilt();
      return result;
    }

    private void buildPartialRepeatedFields(com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor result) {
      if (((bitField0_ & 0x00000001) != 0)) {
        versions_.makeImmutable();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.versions_ = versions_;
    }

    private void buildPartial0(com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor result) {
      int from_bitField0_ = bitField0_;
      if (((from_bitField0_ & 0x00000002) != 0)) {
        result.currentKey_ = currentKey_;
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
      if (other instanceof com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor) {
        return mergeFrom((com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor other) {
      if (other == com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor.getDefaultInstance()) return this;
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
      if (other.getCurrentKey() != com.google.protobuf.ByteString.EMPTY) {
        setCurrentKey(other.getCurrentKey());
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
            case 18: {
              currentKey_ = input.readBytes();
              bitField0_ |= 0x00000002;
              break;
            } // case 18
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

    private com.google.protobuf.ByteString currentKey_ = com.google.protobuf.ByteString.EMPTY;
    /**
     * <code>bytes current_key = 2 [json_name = "currentKey"];</code>
     * @return The currentKey.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString getCurrentKey() {
      return currentKey_;
    }
    /**
     * <code>bytes current_key = 2 [json_name = "currentKey"];</code>
     * @param value The currentKey to set.
     * @return This builder for chaining.
     */
    public Builder setCurrentKey(com.google.protobuf.ByteString value) {
      if (value == null) { throw new NullPointerException(); }
      currentKey_ = value;
      bitField0_ |= 0x00000002;
      onChanged();
      return this;
    }
    /**
     * <code>bytes current_key = 2 [json_name = "currentKey"];</code>
     * @return This builder for chaining.
     */
    public Builder clearCurrentKey() {
      bitField0_ = (bitField0_ & ~0x00000002);
      currentKey_ = getDefaultInstance().getCurrentKey();
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


    // @@protoc_insertion_point(builder_scope:kvstore.v1.KeyValueDescriptor)
  }

  // @@protoc_insertion_point(class_scope:kvstore.v1.KeyValueDescriptor)
  private static final com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor();
  }

  public static com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<KeyValueDescriptor>
      PARSER = new com.google.protobuf.AbstractParser<KeyValueDescriptor>() {
    @java.lang.Override
    public KeyValueDescriptor parsePartialFrom(
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

  public static com.google.protobuf.Parser<KeyValueDescriptor> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<KeyValueDescriptor> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public com.github.mxplusb.pleiades.api.kvstore.v1.KeyValueDescriptor getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

