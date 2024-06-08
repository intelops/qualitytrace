
import { check } from "k6";
import { textSummary } from "https://jslib.k6.io/k6-summary/0.0.2/index.js";
import { Http, Tracetest } from "k6/x/qualitytrace";
import { sleep } from "k6";

export const options = {
  stages: [
    { duration: "5m", target: 30 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<500"],
  },
};

const qualitytrace = Tracetest({
  serverUrl: "http://localhost:11633",
});
const testId = "kc_MgKoVR";
const http = new Http();
const url = "http://localhost:8081/pokemon?take=5";

export default function () {
  const params = {
    qualitytrace: {
      testId,
    },
    headers: {
      "Content-Type": "application/json",
    },
  };

  const response = http.get(url, params);

  check(response, {
    "is status 200": (r) => r.status === 200,
  });

  sleep(1);
}

// enable this to return a non-zero status code if a qualitytrace test fails
export function teardown() {
  qualitytrace.validateResult();
}

export function handleSummary(data) {
  // combine the default summary with the qualitytrace summary
  const qualitytraceSummary = qualitytrace.summary();
  const defaultSummary = textSummary(data);
  const summary = `
    ${defaultSummary}
    ${qualitytraceSummary}
  `;

  return {
    stdout: summary,
    "qualitytrace.json": qualitytrace.json(),
  };
}
