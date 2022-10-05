// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: raft/v1/raft_shard.proto

package io.a13s.api.raft.v1;

/**
 * Protobuf type {@code raft.v1.ShardStateEvent}
 */
public final class ShardStateEvent extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:raft.v1.ShardStateEvent)
    ShardStateEventOrBuilder {
private static final long serialVersionUID = 0L;
  // Use ShardStateEvent.newBuilder() to construct.
  private ShardStateEvent(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private ShardStateEvent() {
    cmd_ = 0;
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new ShardStateEvent();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  private ShardStateEvent(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    if (extensionRegistry == null) {
      throw new java.lang.NullPointerException();
    }
    com.google.protobuf.UnknownFieldSet.Builder unknownFields =
        com.google.protobuf.UnknownFieldSet.newBuilder();
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          case 8: {
            int rawValue = input.readEnum();

            cmd_ = rawValue;
            break;
          }
          case 18: {
            io.a13s.api.raft.v1.ShardState.Builder subBuilder = null;
            if (event_ != null) {
              subBuilder = event_.toBuilder();
            }
            event_ = input.readMessage(io.a13s.api.raft.v1.ShardState.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(event_);
              event_ = subBuilder.buildPartial();
            }

            break;
          }
          default: {
            if (!parseUnknownField(
                input, unknownFields, extensionRegistry, tag)) {
              done = true;
            }
            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (com.google.protobuf.UninitializedMessageException e) {
      throw e.asInvalidProtocolBufferException().setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      this.unknownFields = unknownFields.build();
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.a13s.api.raft.v1.RaftShardProto.internal_static_raft_v1_ShardStateEvent_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.a13s.api.raft.v1.RaftShardProto.internal_static_raft_v1_ShardStateEvent_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.a13s.api.raft.v1.ShardStateEvent.class, io.a13s.api.raft.v1.ShardStateEvent.Builder.class);
  }

  /**
   * Protobuf enum {@code raft.v1.ShardStateEvent.CmdType}
   */
  public enum CmdType
      implements com.google.protobuf.ProtocolMessageEnum {
    /**
     * <code>CMD_TYPE_UNSPECIFIED = 0;</code>
     */
    CMD_TYPE_UNSPECIFIED(0),
    /**
     * <code>CMD_TYPE_PUT = 1;</code>
     */
    CMD_TYPE_PUT(1),
    /**
     * <code>CMD_TYPE_DELETE = 2;</code>
     */
    CMD_TYPE_DELETE(2),
    UNRECOGNIZED(-1),
    ;

    /**
     * <code>CMD_TYPE_UNSPECIFIED = 0;</code>
     */
    public static final int CMD_TYPE_UNSPECIFIED_VALUE = 0;
    /**
     * <code>CMD_TYPE_PUT = 1;</code>
     */
    public static final int CMD_TYPE_PUT_VALUE = 1;
    /**
     * <code>CMD_TYPE_DELETE = 2;</code>
     */
    public static final int CMD_TYPE_DELETE_VALUE = 2;


    public final int getNumber() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalArgumentException(
            "Can't get the number of an unknown enum value.");
      }
      return value;
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     * @deprecated Use {@link #forNumber(int)} instead.
     */
    @java.lang.Deprecated
    public static CmdType valueOf(int value) {
      return forNumber(value);
    }

    /**
     * @param value The numeric wire value of the corresponding enum entry.
     * @return The enum associated with the given numeric wire value.
     */
    public static CmdType forNumber(int value) {
      switch (value) {
        case 0: return CMD_TYPE_UNSPECIFIED;
        case 1: return CMD_TYPE_PUT;
        case 2: return CMD_TYPE_DELETE;
        default: return null;
      }
    }

    public static com.google.protobuf.Internal.EnumLiteMap<CmdType>
        internalGetValueMap() {
      return internalValueMap;
    }
    private static final com.google.protobuf.Internal.EnumLiteMap<
        CmdType> internalValueMap =
          new com.google.protobuf.Internal.EnumLiteMap<CmdType>() {
            public CmdType findValueByNumber(int number) {
              return CmdType.forNumber(number);
            }
          };

    public final com.google.protobuf.Descriptors.EnumValueDescriptor
        getValueDescriptor() {
      if (this == UNRECOGNIZED) {
        throw new java.lang.IllegalStateException(
            "Can't get the descriptor of an unrecognized enum value.");
      }
      return getDescriptor().getValues().get(ordinal());
    }
    public final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptorForType() {
      return getDescriptor();
    }
    public static final com.google.protobuf.Descriptors.EnumDescriptor
        getDescriptor() {
      return io.a13s.api.raft.v1.ShardStateEvent.getDescriptor().getEnumTypes().get(0);
    }

    private static final CmdType[] VALUES = values();

    public static CmdType valueOf(
        com.google.protobuf.Descriptors.EnumValueDescriptor desc) {
      if (desc.getType() != getDescriptor()) {
        throw new java.lang.IllegalArgumentException(
          "EnumValueDescriptor is not for this type.");
      }
      if (desc.getIndex() == -1) {
        return UNRECOGNIZED;
      }
      return VALUES[desc.getIndex()];
    }

    private final int value;

    private CmdType(int value) {
      this.value = value;
    }

    // @@protoc_insertion_point(enum_scope:raft.v1.ShardStateEvent.CmdType)
  }

  public static final int CMD_FIELD_NUMBER = 1;
  private int cmd_;
  /**
   * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
   * @return The enum numeric value on the wire for cmd.
   */
  @java.lang.Override public int getCmdValue() {
    return cmd_;
  }
  /**
   * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
   * @return The cmd.
   */
  @java.lang.Override public io.a13s.api.raft.v1.ShardStateEvent.CmdType getCmd() {
    @SuppressWarnings("deprecation")
    io.a13s.api.raft.v1.ShardStateEvent.CmdType result = io.a13s.api.raft.v1.ShardStateEvent.CmdType.valueOf(cmd_);
    return result == null ? io.a13s.api.raft.v1.ShardStateEvent.CmdType.UNRECOGNIZED : result;
  }

  public static final int EVENT_FIELD_NUMBER = 2;
  private io.a13s.api.raft.v1.ShardState event_;
  /**
   * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
   * @return Whether the event field is set.
   */
  @java.lang.Override
  public boolean hasEvent() {
    return event_ != null;
  }
  /**
   * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
   * @return The event.
   */
  @java.lang.Override
  public io.a13s.api.raft.v1.ShardState getEvent() {
    return event_ == null ? io.a13s.api.raft.v1.ShardState.getDefaultInstance() : event_;
  }
  /**
   * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
   */
  @java.lang.Override
  public io.a13s.api.raft.v1.ShardStateOrBuilder getEventOrBuilder() {
    return getEvent();
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
    if (cmd_ != io.a13s.api.raft.v1.ShardStateEvent.CmdType.CMD_TYPE_UNSPECIFIED.getNumber()) {
      output.writeEnum(1, cmd_);
    }
    if (event_ != null) {
      output.writeMessage(2, getEvent());
    }
    unknownFields.writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (cmd_ != io.a13s.api.raft.v1.ShardStateEvent.CmdType.CMD_TYPE_UNSPECIFIED.getNumber()) {
      size += com.google.protobuf.CodedOutputStream
        .computeEnumSize(1, cmd_);
    }
    if (event_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getEvent());
    }
    size += unknownFields.getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.a13s.api.raft.v1.ShardStateEvent)) {
      return super.equals(obj);
    }
    io.a13s.api.raft.v1.ShardStateEvent other = (io.a13s.api.raft.v1.ShardStateEvent) obj;

    if (cmd_ != other.cmd_) return false;
    if (hasEvent() != other.hasEvent()) return false;
    if (hasEvent()) {
      if (!getEvent()
          .equals(other.getEvent())) return false;
    }
    if (!unknownFields.equals(other.unknownFields)) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + CMD_FIELD_NUMBER;
    hash = (53 * hash) + cmd_;
    if (hasEvent()) {
      hash = (37 * hash) + EVENT_FIELD_NUMBER;
      hash = (53 * hash) + getEvent().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.a13s.api.raft.v1.ShardStateEvent parseFrom(
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
  public static Builder newBuilder(io.a13s.api.raft.v1.ShardStateEvent prototype) {
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
   * Protobuf type {@code raft.v1.ShardStateEvent}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:raft.v1.ShardStateEvent)
      io.a13s.api.raft.v1.ShardStateEventOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.a13s.api.raft.v1.RaftShardProto.internal_static_raft_v1_ShardStateEvent_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.a13s.api.raft.v1.RaftShardProto.internal_static_raft_v1_ShardStateEvent_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.a13s.api.raft.v1.ShardStateEvent.class, io.a13s.api.raft.v1.ShardStateEvent.Builder.class);
    }

    // Construct using io.a13s.api.raft.v1.ShardStateEvent.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      cmd_ = 0;

      if (eventBuilder_ == null) {
        event_ = null;
      } else {
        event_ = null;
        eventBuilder_ = null;
      }
      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.a13s.api.raft.v1.RaftShardProto.internal_static_raft_v1_ShardStateEvent_descriptor;
    }

    @java.lang.Override
    public io.a13s.api.raft.v1.ShardStateEvent getDefaultInstanceForType() {
      return io.a13s.api.raft.v1.ShardStateEvent.getDefaultInstance();
    }

    @java.lang.Override
    public io.a13s.api.raft.v1.ShardStateEvent build() {
      io.a13s.api.raft.v1.ShardStateEvent result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.a13s.api.raft.v1.ShardStateEvent buildPartial() {
      io.a13s.api.raft.v1.ShardStateEvent result = new io.a13s.api.raft.v1.ShardStateEvent(this);
      result.cmd_ = cmd_;
      if (eventBuilder_ == null) {
        result.event_ = event_;
      } else {
        result.event_ = eventBuilder_.build();
      }
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
      if (other instanceof io.a13s.api.raft.v1.ShardStateEvent) {
        return mergeFrom((io.a13s.api.raft.v1.ShardStateEvent)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.a13s.api.raft.v1.ShardStateEvent other) {
      if (other == io.a13s.api.raft.v1.ShardStateEvent.getDefaultInstance()) return this;
      if (other.cmd_ != 0) {
        setCmdValue(other.getCmdValue());
      }
      if (other.hasEvent()) {
        mergeEvent(other.getEvent());
      }
      this.mergeUnknownFields(other.unknownFields);
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
      io.a13s.api.raft.v1.ShardStateEvent parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (io.a13s.api.raft.v1.ShardStateEvent) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }

    private int cmd_ = 0;
    /**
     * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
     * @return The enum numeric value on the wire for cmd.
     */
    @java.lang.Override public int getCmdValue() {
      return cmd_;
    }
    /**
     * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
     * @param value The enum numeric value on the wire for cmd to set.
     * @return This builder for chaining.
     */
    public Builder setCmdValue(int value) {
      
      cmd_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
     * @return The cmd.
     */
    @java.lang.Override
    public io.a13s.api.raft.v1.ShardStateEvent.CmdType getCmd() {
      @SuppressWarnings("deprecation")
      io.a13s.api.raft.v1.ShardStateEvent.CmdType result = io.a13s.api.raft.v1.ShardStateEvent.CmdType.valueOf(cmd_);
      return result == null ? io.a13s.api.raft.v1.ShardStateEvent.CmdType.UNRECOGNIZED : result;
    }
    /**
     * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
     * @param value The cmd to set.
     * @return This builder for chaining.
     */
    public Builder setCmd(io.a13s.api.raft.v1.ShardStateEvent.CmdType value) {
      if (value == null) {
        throw new NullPointerException();
      }
      
      cmd_ = value.getNumber();
      onChanged();
      return this;
    }
    /**
     * <code>.raft.v1.ShardStateEvent.CmdType cmd = 1 [json_name = "cmd"];</code>
     * @return This builder for chaining.
     */
    public Builder clearCmd() {
      
      cmd_ = 0;
      onChanged();
      return this;
    }

    private io.a13s.api.raft.v1.ShardState event_;
    private com.google.protobuf.SingleFieldBuilderV3<
        io.a13s.api.raft.v1.ShardState, io.a13s.api.raft.v1.ShardState.Builder, io.a13s.api.raft.v1.ShardStateOrBuilder> eventBuilder_;
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     * @return Whether the event field is set.
     */
    public boolean hasEvent() {
      return eventBuilder_ != null || event_ != null;
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     * @return The event.
     */
    public io.a13s.api.raft.v1.ShardState getEvent() {
      if (eventBuilder_ == null) {
        return event_ == null ? io.a13s.api.raft.v1.ShardState.getDefaultInstance() : event_;
      } else {
        return eventBuilder_.getMessage();
      }
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public Builder setEvent(io.a13s.api.raft.v1.ShardState value) {
      if (eventBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        event_ = value;
        onChanged();
      } else {
        eventBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public Builder setEvent(
        io.a13s.api.raft.v1.ShardState.Builder builderForValue) {
      if (eventBuilder_ == null) {
        event_ = builderForValue.build();
        onChanged();
      } else {
        eventBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public Builder mergeEvent(io.a13s.api.raft.v1.ShardState value) {
      if (eventBuilder_ == null) {
        if (event_ != null) {
          event_ =
            io.a13s.api.raft.v1.ShardState.newBuilder(event_).mergeFrom(value).buildPartial();
        } else {
          event_ = value;
        }
        onChanged();
      } else {
        eventBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public Builder clearEvent() {
      if (eventBuilder_ == null) {
        event_ = null;
        onChanged();
      } else {
        event_ = null;
        eventBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public io.a13s.api.raft.v1.ShardState.Builder getEventBuilder() {
      
      onChanged();
      return getEventFieldBuilder().getBuilder();
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    public io.a13s.api.raft.v1.ShardStateOrBuilder getEventOrBuilder() {
      if (eventBuilder_ != null) {
        return eventBuilder_.getMessageOrBuilder();
      } else {
        return event_ == null ?
            io.a13s.api.raft.v1.ShardState.getDefaultInstance() : event_;
      }
    }
    /**
     * <code>.raft.v1.ShardState event = 2 [json_name = "event"];</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        io.a13s.api.raft.v1.ShardState, io.a13s.api.raft.v1.ShardState.Builder, io.a13s.api.raft.v1.ShardStateOrBuilder> 
        getEventFieldBuilder() {
      if (eventBuilder_ == null) {
        eventBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            io.a13s.api.raft.v1.ShardState, io.a13s.api.raft.v1.ShardState.Builder, io.a13s.api.raft.v1.ShardStateOrBuilder>(
                getEvent(),
                getParentForChildren(),
                isClean());
        event_ = null;
      }
      return eventBuilder_;
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


    // @@protoc_insertion_point(builder_scope:raft.v1.ShardStateEvent)
  }

  // @@protoc_insertion_point(class_scope:raft.v1.ShardStateEvent)
  private static final io.a13s.api.raft.v1.ShardStateEvent DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.a13s.api.raft.v1.ShardStateEvent();
  }

  public static io.a13s.api.raft.v1.ShardStateEvent getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<ShardStateEvent>
      PARSER = new com.google.protobuf.AbstractParser<ShardStateEvent>() {
    @java.lang.Override
    public ShardStateEvent parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return new ShardStateEvent(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<ShardStateEvent> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<ShardStateEvent> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.a13s.api.raft.v1.ShardStateEvent getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

