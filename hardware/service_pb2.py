# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: service.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='service.proto',
  package='rpc',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rservice.proto\x12\x03rpc\"(\n\x05\x43olor\x12\t\n\x01r\x18\x01 \x01(\x05\x12\t\n\x01g\x18\x02 \x01(\x05\x12\t\n\x01\x62\x18\x03 \x01(\x05\"<\n\x0c\x42\x61sicRequest\x12\x19\n\x05\x63olor\x18\x01 \x01(\x0b\x32\n.rpc.Color\x12\x11\n\tintensity\x18\x02 \x01(\x02\"\x1f\n\rBasicResponse\x12\x0e\n\x06result\x18\x01 \x01(\x08\"\x8b\x01\n\x0b\x46\x61\x64\x65Request\x12\x1e\n\nstartColor\x18\x01 \x01(\x0b\x32\n.rpc.Color\x12\x1c\n\x08\x65ndColor\x18\x02 \x01(\x0b\x32\n.rpc.Color\x12\x16\n\x0estartIntensity\x18\x03 \x01(\x02\x12\x14\n\x0c\x65ndIntensity\x18\x04 \x01(\x02\x12\x10\n\x08\x64uration\x18\x05 \x01(\x05\"\x10\n\x0eRainbowRequest\"!\n\x0fRainbowResponse\x12\x0e\n\x06result\x18\x01 \x01(\x08\"\x1e\n\x0c\x46\x61\x64\x65Response\x12\x0e\n\x06result\x18\x01 \x01(\x08\"\r\n\x0bTestRequest\"\x1e\n\x0cTestResponse\x12\x0e\n\x06result\x18\x01 \x01(\x08\"\x0e\n\x0c\x43learRequest\"\x1f\n\rClearResponse\x12\x0e\n\x06result\x18\x01 \x01(\x08\x32\x81\x02\n\x0fHardwareCommand\x12.\n\x05\x42\x61sic\x12\x11.rpc.BasicRequest\x1a\x12.rpc.BasicResponse\x12+\n\x04\x46\x61\x64\x65\x12\x10.rpc.FadeRequest\x1a\x11.rpc.FadeResponse\x12\x34\n\x07Rainbow\x12\x13.rpc.RainbowRequest\x1a\x14.rpc.RainbowResponse\x12.\n\x05\x43lear\x12\x11.rpc.ClearRequest\x1a\x12.rpc.ClearResponse\x12+\n\x04Test\x12\x10.rpc.TestRequest\x1a\x11.rpc.TestResponseb\x06proto3')
)




_COLOR = _descriptor.Descriptor(
  name='Color',
  full_name='rpc.Color',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='r', full_name='rpc.Color.r', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='g', full_name='rpc.Color.g', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='b', full_name='rpc.Color.b', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=22,
  serialized_end=62,
)


_BASICREQUEST = _descriptor.Descriptor(
  name='BasicRequest',
  full_name='rpc.BasicRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='color', full_name='rpc.BasicRequest.color', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='intensity', full_name='rpc.BasicRequest.intensity', index=1,
      number=2, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=64,
  serialized_end=124,
)


_BASICRESPONSE = _descriptor.Descriptor(
  name='BasicResponse',
  full_name='rpc.BasicResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='rpc.BasicResponse.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=126,
  serialized_end=157,
)


_FADEREQUEST = _descriptor.Descriptor(
  name='FadeRequest',
  full_name='rpc.FadeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='startColor', full_name='rpc.FadeRequest.startColor', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='endColor', full_name='rpc.FadeRequest.endColor', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='startIntensity', full_name='rpc.FadeRequest.startIntensity', index=2,
      number=3, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='endIntensity', full_name='rpc.FadeRequest.endIntensity', index=3,
      number=4, type=2, cpp_type=6, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='duration', full_name='rpc.FadeRequest.duration', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=160,
  serialized_end=299,
)


_RAINBOWREQUEST = _descriptor.Descriptor(
  name='RainbowRequest',
  full_name='rpc.RainbowRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=301,
  serialized_end=317,
)


_RAINBOWRESPONSE = _descriptor.Descriptor(
  name='RainbowResponse',
  full_name='rpc.RainbowResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='rpc.RainbowResponse.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=319,
  serialized_end=352,
)


_FADERESPONSE = _descriptor.Descriptor(
  name='FadeResponse',
  full_name='rpc.FadeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='rpc.FadeResponse.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=354,
  serialized_end=384,
)


_TESTREQUEST = _descriptor.Descriptor(
  name='TestRequest',
  full_name='rpc.TestRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=386,
  serialized_end=399,
)


_TESTRESPONSE = _descriptor.Descriptor(
  name='TestResponse',
  full_name='rpc.TestResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='rpc.TestResponse.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=401,
  serialized_end=431,
)


_CLEARREQUEST = _descriptor.Descriptor(
  name='ClearRequest',
  full_name='rpc.ClearRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=433,
  serialized_end=447,
)


_CLEARRESPONSE = _descriptor.Descriptor(
  name='ClearResponse',
  full_name='rpc.ClearResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='rpc.ClearResponse.result', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=449,
  serialized_end=480,
)

_BASICREQUEST.fields_by_name['color'].message_type = _COLOR
_FADEREQUEST.fields_by_name['startColor'].message_type = _COLOR
_FADEREQUEST.fields_by_name['endColor'].message_type = _COLOR
DESCRIPTOR.message_types_by_name['Color'] = _COLOR
DESCRIPTOR.message_types_by_name['BasicRequest'] = _BASICREQUEST
DESCRIPTOR.message_types_by_name['BasicResponse'] = _BASICRESPONSE
DESCRIPTOR.message_types_by_name['FadeRequest'] = _FADEREQUEST
DESCRIPTOR.message_types_by_name['RainbowRequest'] = _RAINBOWREQUEST
DESCRIPTOR.message_types_by_name['RainbowResponse'] = _RAINBOWRESPONSE
DESCRIPTOR.message_types_by_name['FadeResponse'] = _FADERESPONSE
DESCRIPTOR.message_types_by_name['TestRequest'] = _TESTREQUEST
DESCRIPTOR.message_types_by_name['TestResponse'] = _TESTRESPONSE
DESCRIPTOR.message_types_by_name['ClearRequest'] = _CLEARREQUEST
DESCRIPTOR.message_types_by_name['ClearResponse'] = _CLEARRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Color = _reflection.GeneratedProtocolMessageType('Color', (_message.Message,), dict(
  DESCRIPTOR = _COLOR,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.Color)
  ))
_sym_db.RegisterMessage(Color)

BasicRequest = _reflection.GeneratedProtocolMessageType('BasicRequest', (_message.Message,), dict(
  DESCRIPTOR = _BASICREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.BasicRequest)
  ))
_sym_db.RegisterMessage(BasicRequest)

BasicResponse = _reflection.GeneratedProtocolMessageType('BasicResponse', (_message.Message,), dict(
  DESCRIPTOR = _BASICRESPONSE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.BasicResponse)
  ))
_sym_db.RegisterMessage(BasicResponse)

FadeRequest = _reflection.GeneratedProtocolMessageType('FadeRequest', (_message.Message,), dict(
  DESCRIPTOR = _FADEREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.FadeRequest)
  ))
_sym_db.RegisterMessage(FadeRequest)

RainbowRequest = _reflection.GeneratedProtocolMessageType('RainbowRequest', (_message.Message,), dict(
  DESCRIPTOR = _RAINBOWREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.RainbowRequest)
  ))
_sym_db.RegisterMessage(RainbowRequest)

RainbowResponse = _reflection.GeneratedProtocolMessageType('RainbowResponse', (_message.Message,), dict(
  DESCRIPTOR = _RAINBOWRESPONSE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.RainbowResponse)
  ))
_sym_db.RegisterMessage(RainbowResponse)

FadeResponse = _reflection.GeneratedProtocolMessageType('FadeResponse', (_message.Message,), dict(
  DESCRIPTOR = _FADERESPONSE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.FadeResponse)
  ))
_sym_db.RegisterMessage(FadeResponse)

TestRequest = _reflection.GeneratedProtocolMessageType('TestRequest', (_message.Message,), dict(
  DESCRIPTOR = _TESTREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.TestRequest)
  ))
_sym_db.RegisterMessage(TestRequest)

TestResponse = _reflection.GeneratedProtocolMessageType('TestResponse', (_message.Message,), dict(
  DESCRIPTOR = _TESTRESPONSE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.TestResponse)
  ))
_sym_db.RegisterMessage(TestResponse)

ClearRequest = _reflection.GeneratedProtocolMessageType('ClearRequest', (_message.Message,), dict(
  DESCRIPTOR = _CLEARREQUEST,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.ClearRequest)
  ))
_sym_db.RegisterMessage(ClearRequest)

ClearResponse = _reflection.GeneratedProtocolMessageType('ClearResponse', (_message.Message,), dict(
  DESCRIPTOR = _CLEARRESPONSE,
  __module__ = 'service_pb2'
  # @@protoc_insertion_point(class_scope:rpc.ClearResponse)
  ))
_sym_db.RegisterMessage(ClearResponse)



_HARDWARECOMMAND = _descriptor.ServiceDescriptor(
  name='HardwareCommand',
  full_name='rpc.HardwareCommand',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=483,
  serialized_end=740,
  methods=[
  _descriptor.MethodDescriptor(
    name='Basic',
    full_name='rpc.HardwareCommand.Basic',
    index=0,
    containing_service=None,
    input_type=_BASICREQUEST,
    output_type=_BASICRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Fade',
    full_name='rpc.HardwareCommand.Fade',
    index=1,
    containing_service=None,
    input_type=_FADEREQUEST,
    output_type=_FADERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Rainbow',
    full_name='rpc.HardwareCommand.Rainbow',
    index=2,
    containing_service=None,
    input_type=_RAINBOWREQUEST,
    output_type=_RAINBOWRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Clear',
    full_name='rpc.HardwareCommand.Clear',
    index=3,
    containing_service=None,
    input_type=_CLEARREQUEST,
    output_type=_CLEARRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Test',
    full_name='rpc.HardwareCommand.Test',
    index=4,
    containing_service=None,
    input_type=_TESTREQUEST,
    output_type=_TESTRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_HARDWARECOMMAND)

DESCRIPTOR.services_by_name['HardwareCommand'] = _HARDWARECOMMAND

# @@protoc_insertion_point(module_scope)
