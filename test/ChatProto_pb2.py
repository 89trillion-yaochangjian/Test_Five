# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ChatProto.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='ChatProto.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\016../model;model',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0f\x43hatProto.proto\"\x9d\x01\n\x0b\x43hatRequest\x12\x10\n\x08userName\x18\x01 \x01(\t\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x03 \x01(\t\x12,\n\x08userList\x18\x04 \x03(\x0b\x32\x1a.ChatRequest.UserListEntry\x1a/\n\rUserListEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\x42\x10Z\x0e../model;modelb\x06proto3'
)




_CHATREQUEST_USERLISTENTRY = _descriptor.Descriptor(
  name='UserListEntry',
  full_name='ChatRequest.UserListEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='ChatRequest.UserListEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='value', full_name='ChatRequest.UserListEntry.value', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=b'8\001',
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=130,
  serialized_end=177,
)

_CHATREQUEST = _descriptor.Descriptor(
  name='ChatRequest',
  full_name='ChatRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='userName', full_name='ChatRequest.userName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='type', full_name='ChatRequest.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='content', full_name='ChatRequest.content', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='userList', full_name='ChatRequest.userList', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[_CHATREQUEST_USERLISTENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=177,
)

_CHATREQUEST_USERLISTENTRY.containing_type = _CHATREQUEST
_CHATREQUEST.fields_by_name['userList'].message_type = _CHATREQUEST_USERLISTENTRY
DESCRIPTOR.message_types_by_name['ChatRequest'] = _CHATREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ChatRequest = _reflection.GeneratedProtocolMessageType('ChatRequest', (_message.Message,), {

  'UserListEntry' : _reflection.GeneratedProtocolMessageType('UserListEntry', (_message.Message,), {
    'DESCRIPTOR' : _CHATREQUEST_USERLISTENTRY,
    '__module__' : 'ChatProto_pb2'
    # @@protoc_insertion_point(class_scope:ChatRequest.UserListEntry)
    })
  ,
  'DESCRIPTOR' : _CHATREQUEST,
  '__module__' : 'ChatProto_pb2'
  # @@protoc_insertion_point(class_scope:ChatRequest)
  })
_sym_db.RegisterMessage(ChatRequest)
_sym_db.RegisterMessage(ChatRequest.UserListEntry)


DESCRIPTOR._options = None
_CHATREQUEST_USERLISTENTRY._options = None
# @@protoc_insertion_point(module_scope)
