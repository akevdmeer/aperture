local commonConfig = import '../../common/config-defaults.libsonnet';

/**
* @schema (driver_criteria.enabled: bool) Enables the driver.
* @schema (driver_criteria.threshold: float64) Threshold for the driver.
*/
local driver_criteria_defaults = {
  enabled: '__REQUIRED_FIELD__',
  threshold: '__REQUIRED_FIELD__',
};

/**
* @schema (overload_confirmation_driver.pod_cpu: driver_criteria) The driver for using CPU usage as overload confirmation.
* @schema (overload_confirmation_driver.pod_memory: driver_criteria) The driver for using CPU usage as overload confirmation.
*/
local overload_confirmation_driver_defaults = {
  pod_cpu: {},
  pod_memory: {},
};

/**
* @schema (kubelet_overload_confirmations.infra_context: aperture.spec.v1.KubernetesObjectSelector) Kubernetes selector for scraping metrics.
* @schema (kubelet_overload_confirmations.criteria: overload_confirmation_driver) Criteria for overload confirmation.
*/
local kubelet_overload_confirmations_defaults = {
  infra_context: '__REQUIRED_FIELD__',
  criteria: '__REQUIRED_FIELD__',
};

local service_protection_core_defaults = {
  overload_confirmations: [],

  adaptive_load_scheduler: {
    load_scheduler: {
      selectors: commonConfig.selectors_defaults,
    },
    gradient: {
      slope: -1,
      min_gradient: 0.1,
      max_gradient: 1.0,
    },
    alerter: {
      alert_name: 'Load Throttling Event',
    },
    max_load_multiplier: 2.0,
    load_multiplier_linear_increment: 0.025,
  },

  dry_run: false,

  kubelet_overload_confirmations: {},
};

commonConfig {
  /**
  * @param (policy.evaluation_interval: string) The interval between successive evaluations of the Circuit.
  * @param (policy.service_protection_core.overload_confirmations: []overload_confirmation) List of overload confirmation criteria. Load scheduler can throttle flows when all of the specified overload confirmation criteria are met.
  * @schema (overload_confirmation.query_string: string) The Prometheus query to be run. Must return a scalar or a vector with a single element.
  * @schema (overload_confirmation.threshold: float64) The threshold for the overload confirmation criteria.
  * @schema (overload_confirmation.operator: string) The operator for the overload confirmation criteria. oneof: `gt | lt | gte | lte | eq | neq`
  * @param (policy.service_protection_core.adaptive_load_scheduler: aperture.spec.v1.AdaptiveLoadSchedulerParameters) Parameters for Adaptive Load Scheduler.
  * @param (policy.service_protection_core.dry_run: bool) Default configuration for setting dry run mode on Load Scheduler. In dry run mode, the Load Scheduler acts as a passthrough and does not throttle flows. This config can be updated at runtime without restarting the policy.
  * @param (policy.service_protection_core.kubelet_overload_confirmations: kubelet_overload_confirmations) Overload confirmation signals from kubelet.
  */
  policy+: {
    evaluation_interval: '10s',
    service_protection_core: service_protection_core_defaults,
  },

  dashboard+: {
    title: 'Aperture Service Protection',
    variant_name: 'Service Protection',
  },

  kubelet_overload_confirmations: kubelet_overload_confirmations_defaults,
  overload_confirmation_driver: overload_confirmation_driver_defaults,
  driver_criteria: driver_criteria_defaults,
}
