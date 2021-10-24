from http.server import BaseHTTPRequestHandler, HTTPServer  
import os  

  
class RequestHandler(BaseHTTPRequestHandler):  

  def do_GET(self):  
    self.send_response(200)
    self.send_header('Content-type', 'text/html; charset=utf-8')
    self.end_headers()
    self.wfile.write(bytes('Hello world 🐍', 'UTF-8'))


def run():  
  # get http_port
  http_port = int(os.getenv('APP_HTTP_PORT', '3000'))
  server_address = ('', http_port)  
  httpd = HTTPServer(server_address, RequestHandler)  
  # serve
  print('Serve HTTP on port {}'.format(http_port))  
  httpd.serve_forever()  


if __name__ == '__main__':  
  run()  