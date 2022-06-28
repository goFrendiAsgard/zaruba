from typing import Any, Mapping
import http.client
import json

def send_request(conn: http.client.HTTPConnection, method: str, url: str, payload: str, headers: Mapping[str, str]):
    conn.request(method, url, payload, headers)
    res = conn.getresponse()
    data = res.read()
    raw_response = data.decode('utf-8')
    print('raw response', raw_response)
    return raw_response

def login(conn: http.client.HTTPConnection, username: str, password: str):
    payload = 'grant_type=&username={}&password={}&scope=&client_id=&client_secret='.format(username, password)
    headers = {
        'accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded',
    }
    raw_response = send_request(conn, 'POST', '/api/v1/create-oauth-access-token/', payload, headers)
    decoded_response = json.loads(raw_response)
    print('decoded response', decoded_response)
    return decoded_response

def send_json_request(conn: http.client.HTTPConnection, method: str, url: str, payload_object: Mapping[str, Any], access_token: str = '', headers: Mapping[str, str] = {}):
    if not 'Content-Type' in headers:
        headers['Content-Type'] = 'application/json'
    if access_token != '':
        headers['Authorization'] = 'Bearer {}'.format(access_token)
    payload = json.dumps(payload_object)
    raw_response = send_request(conn, method, url, payload, headers)
    decoded_response = json.loads(raw_response)
    print('decoded response', decoded_response)
    return decoded_response

conn = http.client.HTTPConnection("localhost", 3000)

print('🧪 login')
response = login(conn, 'root', 'Alch3mist')
assert 'access_token' in response
assert response['token_type'] == 'bearer'
access_token = response['access_token']

print('🧪 create book')
response = send_json_request(conn, 'POST', '/api/v1/books/', {
    'title': 'Homo Sapiens',
    'author': 'Yuval Noah Harari',
    'synopsis': 'A brief history of mankind',
}, access_token)
assert 'id' in response
assert response['title'] == 'Homo Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'
book_id = response['id']

print('🧪 get book')
response = send_json_request(conn, 'GET', '/api/v1/books/{}'.format(book_id), {}, access_token)
assert response['id'] == book_id
assert response['title'] == 'Homo Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('🧪 update book')
response = send_json_request(conn, 'PUT', '/api/v1/books/{}'.format(book_id), {
    'title': 'Sapiens',
    'author': 'Yuval Noah Harari',
    'synopsis': 'A brief history of mankind',
}, access_token)
assert response['id'] == book_id
assert response['title'] == 'Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('🧪 get books')
response = send_json_request(conn, 'GET', '/api/v1/books/', {}, access_token)
assert response['count'] == 1
assert len(response['rows']) == 1

print('🧪 delete book')
response = send_json_request(conn, 'DELETE', '/api/v1/books/{}'.format(book_id), {}, access_token)
assert response['id'] == book_id
assert response['title'] == 'Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('🧪 get books after delete')
response = send_json_request(conn, 'GET', '/api/v1/books/', {}, access_token)
assert response['count'] == 0
assert len(response['rows']) == 0