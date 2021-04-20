package at.htl.beeyond.integration.util;

import com.github.tomakehurst.wiremock.WireMockServer;
import io.quarkus.test.common.QuarkusTestResourceLifecycleManager;

import java.util.Collections;
import java.util.Map;

import static com.github.tomakehurst.wiremock.client.WireMock.*;

public class WiremockValidation implements QuarkusTestResourceLifecycleManager {

    private WireMockServer wireMockServer;

    @Override
    public Map<String, String> start() {
        this.wireMockServer = new WireMockServer(8082);
        this.wireMockServer.stubFor(post(urlEqualTo("/api/validate")).willReturn(aResponse().withStatus(200)));
        this.wireMockServer.start();

        return Collections.singletonMap("at.htl.beeyond.service.ValidationRestClient/mp-rest/url", wireMockServer.baseUrl());
    }

    @Override
    public void stop() {
        if (null != wireMockServer) {
            wireMockServer.stop();
        }
    }
}
