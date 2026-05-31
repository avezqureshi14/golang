import http from 'k6/http';

export const options = {
  vus: 50,           // 50 concurrent users
  duration: '30s',   // run for 30 seconds
};

export default function () {
  http.get('http://localhost:8080/');
}