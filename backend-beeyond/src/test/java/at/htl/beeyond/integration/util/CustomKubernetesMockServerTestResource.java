package at.htl.beeyond.integration.util;

import io.quarkus.test.kubernetes.client.KubernetesServerTestResource;

public class CustomKubernetesMockServerTestResource extends KubernetesServerTestResource {

    @Override
    protected void configureServer() {
        server.expect().post().withPath("/api/v1/namespaces")
                .andReturn(201, null)
                .always();
    }
}
