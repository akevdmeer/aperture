package com.fluxninja.example.controller;

import com.fluxninja.aperture.servlet.javax.ApertureFilter;
import com.fluxninja.example.filter.ApertureFeatureFilter;
import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.core.env.Environment;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AppController {
    // /super endpoint is protected by imported, ready-to-use Aperture Filter
    @RequestMapping(value = "/super", method = RequestMethod.GET)
    public String hello() {
        return "Hello World";
    }

    // /super2 endpoint is protected by a Filter created using Aperture SDK feature flow
    @RequestMapping(value = "/super2", method = RequestMethod.GET)
    public String hello2() {
        return "Hello World 2";
    }

    @RequestMapping(value = "/health", method = RequestMethod.GET)
    public String health() {
        return "Healthy";
    }

    @RequestMapping(value = "/connected", method = RequestMethod.GET)
    public String connected() {
        return "";
    }

    // Register imported Aperture Filter to apply to /super endpoint
    @Bean
    public FilterRegistrationBean<ApertureFilter> apertureFilter(Environment env) {
        FilterRegistrationBean<ApertureFilter> registrationBean = new FilterRegistrationBean<>();

        registrationBean.setFilter(new ApertureFilter());
        registrationBean.addUrlPatterns("/super");

        String agentHost = env.getProperty("FN_AGENT_HOST");
        registrationBean.addInitParameter("agent_host", agentHost);
        String agentPort = env.getProperty("FN_AGENT_PORT");
        registrationBean.addInitParameter("agent_port", agentPort);
        String failOpen = env.getProperty("FN_ENABLE_FAIL_OPEN");
        registrationBean.addInitParameter("enable_fail_open", failOpen);
        String controlPointName = env.getProperty("FN_CONTROL_POINT_NAME");
        registrationBean.addInitParameter("control_point_name", controlPointName);
        String grpcTimeoutMs = env.getProperty("FN_GRPC_TIMEOUT_MS");
        registrationBean.addInitParameter("timeout_ms", grpcTimeoutMs);
        String insecureGrpc = env.getProperty("FN_INSECURE_GRPC");
        registrationBean.addInitParameter("insecure_grpc", insecureGrpc);
        String rootCertificateFile = env.getProperty("FN_ROOT_CERTIFICATE_FILE");
        registrationBean.addInitParameter("root_certificate_file", rootCertificateFile);

        return registrationBean;
    }

    // Register locally defined Aperture Feature Filter to apply to /super2 endpoint
    @Bean
    public FilterRegistrationBean<ApertureFeatureFilter> apertureFeatureFilter(Environment env) {
        FilterRegistrationBean<ApertureFeatureFilter> registrationBean =
                new FilterRegistrationBean<>();

        registrationBean.setFilter(new ApertureFeatureFilter());
        registrationBean.addUrlPatterns("/super2");

        String agentHost = env.getProperty("FN_AGENT_HOST");
        registrationBean.addInitParameter("agent_host", agentHost);
        String agentPort = env.getProperty("FN_AGENT_PORT");
        registrationBean.addInitParameter("agent_port", agentPort);
        String failOpen = env.getProperty("FN_ENABLE_FAIL_OPEN");
        registrationBean.addInitParameter("enable_fail_open", failOpen);
        String insecureGrpc = env.getProperty("FN_INSECURE_GRPC");
        registrationBean.addInitParameter("insecure_grpc", insecureGrpc);
        String rootCertificateFile = env.getProperty("FN_ROOT_CERTIFICATE_FILE");
        registrationBean.addInitParameter("root_certificate_file", rootCertificateFile);

        return registrationBean;
    }
}
