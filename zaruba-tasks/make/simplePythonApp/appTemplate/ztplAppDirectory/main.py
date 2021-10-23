from http.server import BaseHTTPRequestHandler, HTTPServer  
import os  

  
class RequestHandler(BaseHTTPRequestHandler):  

  def do_GET(self):  
    self.send_response(200, 'Hello world 🐍')


def run():  
  
  # get http_port
  http_port = int(os.getenv('APP_HTTP_PORT', '3000'))
  server_address = ('127.0.0.1', http_port)  
  httpd = HTTPServer(server_address, RequestHandler)  

  # serve
  print('Serve HTTP on port %d'.format(http_port))  
  httpd.serve_forever()  


if __name__ == '__main__':  
  run()  