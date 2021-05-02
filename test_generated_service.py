import http.client
import json

def send_request(conn, method, url, payload_object, headers=None):
    if headers is None:
        headers = {'Content-Type': 'application/json'}
    payload = json.dumps(payload_object)
    conn.request(method, url, payload, headers)
    res = conn.getresponse()
    data = res.read()
    return json.loads(data.decode("utf-8"))

conn = http.client.HTTPConnection("localhost", 3000)

print('create book')
response = send_request(conn, 'POST', '/book/', {
    'title': 'Homo Sapiens',
    'author': 'Yuval Noah Harari',
    'synopsis': 'A brief history of mankind',
})
print(response)
assert 'id' in response
assert response['title'] == 'Homo Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'
book_id = response['id']

print('get book')
response = send_request(conn, 'GET', '/book/{}'.format(book_id), {})
print(response)
assert response['id'] == book_id
assert response['title'] == 'Homo Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('update book')
response = send_request(conn, 'PUT', '/book/{}'.format(book_id), {
    'title': 'Sapiens',
    'author': 'Yuval Noah Harari',
    'synopsis': 'A brief history of mankind',
})
print(response)
assert response['id'] == book_id
assert response['title'] == 'Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('get books')
response = send_request(conn, 'GET', '/book/', {})
print(response)
assert len(response) == 1

print('delete books')
response = send_request(conn, 'DELETE', '/book/{}'.format(book_id), {})
print(response)
assert response['id'] == book_id
assert response['title'] == 'Sapiens'
assert response['author'] == 'Yuval Noah Harari'
assert response['synopsis'] == 'A brief history of mankind'

print('get books after delete')
response = send_request(conn, 'GET', '/book/', {})
print(response)
assert len(response) == 0