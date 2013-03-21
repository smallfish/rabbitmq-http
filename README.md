### RabbitMQ HTTP API


REST API for RabbitMQ, but it's not [RabbitMQ Management Plugin](http://www.rabbitmq.com/management.html).

##### Status:

Under active development.

##### Required:

    * RabbitMQ (2.8+)
    * Go(lang) (1.0.3)

##### Install:

    $ go get github.com/streadway/amqp
    $ go get github.com/smallfish/rabbitmq-http

##### Usage

* Start HTTP Server:

        $ ./rabbitmq-http -address="127.0.0.1:8080" -amqp="amqp://guest:guest@localhost:5672/"

##### API

###### Response:

    HTTP/1.1 200 OK
    HTTP/1.1 405 Method Not Allowed
    HTTP/1.1 500 Internal Server Error

###### Exchange

* create new exchange:
        
        $ curl -i -X POST http://127.0.0.1:8080/exchange -d \
        '{"name": "e1", "type": "topic", "durable": true, "autodelete": false}'
         
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:45:47 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        declare exchange ok
        
* delete exchange:

        $ curl -i -X DELETE http://127.0.0.1:8080/exchange -d \
        '{"name": "e1"}'
        
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:46:21 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        delete exchange ok

###### Queue

* create new queue:

        $ curl -i -X POST http://127.0.0.1:8080/queue -d \
        '{"name": "q1"}'
        
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:47:11 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        declare queue ok

        
* delete queue:

        $ curl -i -X DELETE http://127.0.0.1:8080/queue -d \
        '{"name": "q1"}'
        
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:48:05 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        delete queue ok
        
* bind keys to queue:

        $ curl -i -X POST http://127.0.0.1:8080/queue/bind -d \
        '{"queue": "q1", "exchange": "e1", "keys": ["aa", "bb", "cc"]}'
        
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:48:43 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        bind queue ok

* unbind keys to queue:

        $ curl -i -X DELETE http://127.0.0.1:8080/queue/bind -d \
        '{"queue": "q1", "exchange": "e1", "keys": ["aa", "cc"]}'
        
        HTTP/1.1 200 OK
        Date: Thu, 21 Mar 2013 05:49:05 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        unbind queue ok

__END__
