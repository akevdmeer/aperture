// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: aperture/flowcontrol/check/v1/check.proto

package com.fluxninja.generated.aperture.flowcontrol.check.v1;

public interface ClassifierInfoOrBuilder extends
    // @@protoc_insertion_point(interface_extends:aperture.flowcontrol.check.v1.ClassifierInfo)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string policy_name = 1 [json_name = "policyName"];</code>
   * @return The policyName.
   */
  java.lang.String getPolicyName();
  /**
   * <code>string policy_name = 1 [json_name = "policyName"];</code>
   * @return The bytes for policyName.
   */
  com.google.protobuf.ByteString
      getPolicyNameBytes();

  /**
   * <code>string policy_hash = 2 [json_name = "policyHash"];</code>
   * @return The policyHash.
   */
  java.lang.String getPolicyHash();
  /**
   * <code>string policy_hash = 2 [json_name = "policyHash"];</code>
   * @return The bytes for policyHash.
   */
  com.google.protobuf.ByteString
      getPolicyHashBytes();

  /**
   * <code>int64 classifier_index = 3 [json_name = "classifierIndex"];</code>
   * @return The classifierIndex.
   */
  long getClassifierIndex();

  /**
   * <code>.aperture.flowcontrol.check.v1.ClassifierInfo.Error error = 5 [json_name = "error"];</code>
   * @return The enum numeric value on the wire for error.
   */
  int getErrorValue();
  /**
   * <code>.aperture.flowcontrol.check.v1.ClassifierInfo.Error error = 5 [json_name = "error"];</code>
   * @return The error.
   */
  com.fluxninja.generated.aperture.flowcontrol.check.v1.ClassifierInfo.Error getError();
}
