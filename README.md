### RabbitMQ HTTP API

REST API for RabbitMQ, but it's not [RabbitMQ Management Plugin](http://www.rabbitmq.com/management.html).

#### Required:

    * RabbitMQ (2.8+)
    * Go(lang) (1.0.3)

#### Install:

    $ go get github.com/streadway/amqp
    $ go get github.com/smallfish/rabbitmq-http

#### Usage

* Start HTTP Server:

    $ ./rabbitmq-http -bind="127.0.0.1:8080" -amqp="amqp://guest:guest@localhost:5672/"

#### API

#####Exchange

* create new exchange:
        
        $ curl -X POST http://127.0.0.1:8080/exchange -d \
         '{"name": "e1", "type": "topic", "durable": true, "auto_delete": false}'
        
* delete exchange:
        
        $ curl -X DELETE http://127.0.0.1:8080/exchange -d '{"name": "e1"}'

#####Queue

* create new queue:

        $ curl -X POST http://127.0.0.1:8080/queue -d '{"name": "q1"}'
        
* delete queue:

        $ curl -X DELETE http://127.0.0.1:8080/queue -d '{"name": "q1"}'

__END__
