import http from "k6/http";

export let options = {
  vus: 500,
  duration: "10s",
};

export default function () {
  http.get("http://host.docker.internal:8090/movie/getallmovie");
}
