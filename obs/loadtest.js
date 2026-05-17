import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  vus: 50,          // number of virtual users
  duration: '30s',  // test duration
};

export default function () {
  http.get('http://localhost:8080/api/hello');
  sleep(0.1);
}