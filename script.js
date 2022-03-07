import http from 'k6/http';

export let options = {
    insecureSkipTLSVerify: true,
    vus: 10,
    iterations: 1000000,
};

let count = 0

export default function () {
    const url = 'http://localhost:80/vote';
    const payload = JSON.stringify({
        key: `${count++}`,
        candidateId: 'any-candidate-id',
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}