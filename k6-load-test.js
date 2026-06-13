import http from "k6/http";
import { check, sleep } from "k6";

export const options = {
  vus: 20,
  duration: "30s",
};

export default function () {
  const clientId = `client-${__VU}`;

  const res = http.get("http://localhost:8080/api/data", {
    headers: {
      "X-Client-ID": clientId,
    },
  });

  check(res, {
    "gateway returned expected response": (r) => r.status === 200 || r.status === 429,
  });

  sleep(1);
}