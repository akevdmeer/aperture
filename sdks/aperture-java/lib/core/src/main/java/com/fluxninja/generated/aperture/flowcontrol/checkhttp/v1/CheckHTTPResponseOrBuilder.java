// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: aperture/flowcontrol/checkhttp/v1/checkhttp.proto

package com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1;

public interface CheckHTTPResponseOrBuilder extends
    // @@protoc_insertion_point(interface_extends:aperture.flowcontrol.checkhttp.v1.CheckHTTPResponse)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * Status `OK` allows the request. Any other status indicates the request should be denied
   * </pre>
   *
   * <code>.google.rpc.Status status = 1 [json_name = "status"];</code>
   * @return Whether the status field is set.
   */
  boolean hasStatus();
  /**
   * <pre>
   * Status `OK` allows the request. Any other status indicates the request should be denied
   * </pre>
   *
   * <code>.google.rpc.Status status = 1 [json_name = "status"];</code>
   * @return The status.
   */
  com.fluxninja.generated.google.rpc.Status getStatus();
  /**
   * <pre>
   * Status `OK` allows the request. Any other status indicates the request should be denied
   * </pre>
   *
   * <code>.google.rpc.Status status = 1 [json_name = "status"];</code>
   */
  com.fluxninja.generated.google.rpc.StatusOrBuilder getStatusOrBuilder();

  /**
   * <pre>
   * Supplies http attributes for a denied response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.DeniedHttpResponse denied_response = 2 [json_name = "deniedResponse"];</code>
   * @return Whether the deniedResponse field is set.
   */
  boolean hasDeniedResponse();
  /**
   * <pre>
   * Supplies http attributes for a denied response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.DeniedHttpResponse denied_response = 2 [json_name = "deniedResponse"];</code>
   * @return The deniedResponse.
   */
  com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.DeniedHttpResponse getDeniedResponse();
  /**
   * <pre>
   * Supplies http attributes for a denied response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.DeniedHttpResponse denied_response = 2 [json_name = "deniedResponse"];</code>
   */
  com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.DeniedHttpResponseOrBuilder getDeniedResponseOrBuilder();

  /**
   * <pre>
   * Supplies http attributes for an ok response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.OkHttpResponse ok_response = 3 [json_name = "okResponse"];</code>
   * @return Whether the okResponse field is set.
   */
  boolean hasOkResponse();
  /**
   * <pre>
   * Supplies http attributes for an ok response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.OkHttpResponse ok_response = 3 [json_name = "okResponse"];</code>
   * @return The okResponse.
   */
  com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.OkHttpResponse getOkResponse();
  /**
   * <pre>
   * Supplies http attributes for an ok response.
   * </pre>
   *
   * <code>.aperture.flowcontrol.checkhttp.v1.OkHttpResponse ok_response = 3 [json_name = "okResponse"];</code>
   */
  com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.OkHttpResponseOrBuilder getOkResponseOrBuilder();

  /**
   * <pre>
   * Optional response metadata that will be emitted as dynamic metadata to be consumed by the next
   * filter.
   * </pre>
   *
   * <code>.google.protobuf.Struct dynamic_metadata = 4 [json_name = "dynamicMetadata"];</code>
   * @return Whether the dynamicMetadata field is set.
   */
  boolean hasDynamicMetadata();
  /**
   * <pre>
   * Optional response metadata that will be emitted as dynamic metadata to be consumed by the next
   * filter.
   * </pre>
   *
   * <code>.google.protobuf.Struct dynamic_metadata = 4 [json_name = "dynamicMetadata"];</code>
   * @return The dynamicMetadata.
   */
  com.google.protobuf.Struct getDynamicMetadata();
  /**
   * <pre>
   * Optional response metadata that will be emitted as dynamic metadata to be consumed by the next
   * filter.
   * </pre>
   *
   * <code>.google.protobuf.Struct dynamic_metadata = 4 [json_name = "dynamicMetadata"];</code>
   */
  com.google.protobuf.StructOrBuilder getDynamicMetadataOrBuilder();

  /**
   * <pre>
   * When true, dry_run informs the receiver to allow all traffic regardless of Agent's decision.
   * </pre>
   *
   * <code>bool dry_run = 5 [json_name = "dryRun"];</code>
   * @return The dryRun.
   */
  boolean getDryRun();

  com.fluxninja.generated.aperture.flowcontrol.checkhttp.v1.CheckHTTPResponse.HttpResponseCase getHttpResponseCase();
}
