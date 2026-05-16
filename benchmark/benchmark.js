import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  vus: 50,
  duration: '30s',

  thresholds: {
    http_req_duration: ['p(95)<200', 'p(99)<400'], // latency SLAs
    http_req_failed: ['rate<0.01'],                // <1% errors
  },
};

export default function () {
  const payload = JSON.stringify({
    id: 1,
    user_id: 101,
    amount: 500
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const res = http.post('http://localhost:8080/payment', payload, params);

  check(res, {
    'status is 200': (r) => r.status === 200,
  });

  sleep(0.1); // prevents unrealistic hammering
}