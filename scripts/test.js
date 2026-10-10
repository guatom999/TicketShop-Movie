// import http from "k6/http";

// export let options = {
//   vus: 500,
//   duration: "10s",
// };

// export default function () {
//   http.get("http://host.docker.internal:8090/movie/getallmovie");
// }

import http from "k6/http";
import { sleep } from "k6";

export let options = {
  vus: 500,
  duration: "5s",
};

export default function () {
  // Make a GET request to the target URL
  http.get("http://localhost:8080/web/index.html");

  // Sleep for 1 second to simulate real-world usage
  // sleep(0.1);
}
