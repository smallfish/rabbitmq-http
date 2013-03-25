### RabbitMQ HTTP API


REST HTTP API for RabbitMQ, it's not [RabbitMQ Management Plugin](http://www.rabbitmq.com/management.html).

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

##### HTTP Response

    200 OK
    405 Method Not Allowed
    500 Internal Server Error

##### API List

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

###### Message

* publish new message:

        $ curl -i -X POST "http://127.0.0.1:8080/publish" -d \
        '{"exchange": "e1", "key": "bb", "deliverymode": 1, "priority": 99, "body": "hahaha"}'

        HTTP/1.1 200 OK
        Date: Mon, 25 Mar 2013 11:56:22 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        publish message ok

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

* consume queue:

        $ curl -i -X GET "http://127.0.0.1:8080/queue?name=q1"

        HTTP/1.1 200 OK
        Date: Fri, 22 Mar 2013 04:11:59 GMT
        Transfer-Encoding: chunked
        Content-Type: text/plain; charset=utf-8

        <DATA>\n
        <DATA>\n
        ...

##### Copyright and License

This module is licensed under the BSD license.

Copyright (C) 2013, by Chen "smallfish" Xiaoyu (陈小玉) <smallfish.xy@gmail.com>.

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
