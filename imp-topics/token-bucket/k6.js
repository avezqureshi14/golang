import http from 'k6/http';
import { sleep } from 'k6';

export let options = {
  vus: 5,          // 20 virtual users
  duration: '3s',  // run for 10 seconds
};

export default function () {
  let res = http.get('http://localhost:8080/api');

  console.log(`Status: ${res.status}`);

  sleep(0.1); // each user waits 100ms before next request
}