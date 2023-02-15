import http from 'k6/http';
import {
    sleep
} from 'k6';
export const options = {
    vus: 10000,
    duration: '30s',
};
export default function () {
    const url = 'http://localhost:8080/ping';
    const payload = JSON.stringify({
        email: 'aaa',
        password: 'bbb',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
            "Authorization": "123"
        },
    };

    http.post(url, payload, params);
    sleep(1)
}