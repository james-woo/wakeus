/**
 * @fileoverview gRPC-Web generated client stub for rpc
 * @enhanceable
 * @public
 */

/* eslint-disable */
// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.rpc = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rpc.HardwareCommandClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rpc.HardwareCommandPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.BasicRequest,
 *   !proto.rpc.BasicResponse>}
 */
const methodInfo_HardwareCommand_Basic = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.BasicResponse,
  /** @param {!proto.rpc.BasicRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.BasicResponse.deserializeBinary
);


/**
 * @param {!proto.rpc.BasicRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.BasicResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.BasicResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.HardwareCommandClient.prototype.basic =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.HardwareCommand/Basic',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Basic,
      callback);
};


/**
 * @param {!proto.rpc.BasicRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.BasicResponse>}
 *     A native promise that resolves to the response
 */
proto.rpc.HardwareCommandPromiseClient.prototype.basic =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.HardwareCommand/Basic',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Basic);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.FadeRequest,
 *   !proto.rpc.FadeResponse>}
 */
const methodInfo_HardwareCommand_Fade = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.FadeResponse,
  /** @param {!proto.rpc.FadeRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.FadeResponse.deserializeBinary
);


/**
 * @param {!proto.rpc.FadeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.FadeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.FadeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.HardwareCommandClient.prototype.fade =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.HardwareCommand/Fade',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Fade,
      callback);
};


/**
 * @param {!proto.rpc.FadeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.FadeResponse>}
 *     A native promise that resolves to the response
 */
proto.rpc.HardwareCommandPromiseClient.prototype.fade =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.HardwareCommand/Fade',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Fade);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.ClearRequest,
 *   !proto.rpc.ClearResponse>}
 */
const methodInfo_HardwareCommand_Clear = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.ClearResponse,
  /** @param {!proto.rpc.ClearRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.ClearResponse.deserializeBinary
);


/**
 * @param {!proto.rpc.ClearRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.ClearResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.ClearResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.HardwareCommandClient.prototype.clear =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.HardwareCommand/Clear',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Clear,
      callback);
};


/**
 * @param {!proto.rpc.ClearRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.ClearResponse>}
 *     A native promise that resolves to the response
 */
proto.rpc.HardwareCommandPromiseClient.prototype.clear =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.HardwareCommand/Clear',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Clear);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.TestRequest,
 *   !proto.rpc.TestResponse>}
 */
const methodInfo_HardwareCommand_Test = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.TestResponse,
  /** @param {!proto.rpc.TestRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.TestResponse.deserializeBinary
);


/**
 * @param {!proto.rpc.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.TestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.TestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.HardwareCommandClient.prototype.test =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.HardwareCommand/Test',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Test,
      callback);
};


/**
 * @param {!proto.rpc.TestRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.TestResponse>}
 *     A native promise that resolves to the response
 */
proto.rpc.HardwareCommandPromiseClient.prototype.test =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.HardwareCommand/Test',
      request,
      metadata || {},
      methodInfo_HardwareCommand_Test);
};


module.exports = proto.rpc;

