import http from 'k6/http';
import {check, sleep} from 'k6';

export default function () {
    let response = http.get('http://localhost:8080/v1/satellite/positions');

    check(response, {
        'is status 200': (r) => r.status === 200,
    });

    sleep(1);
}
