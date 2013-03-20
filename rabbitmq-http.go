// Copyright (C) 2013 Chen "smallfish" Xiaoyu (陈小玉)

// TODO:
//  * Logging error
//  * Check unmarshal/connect error
//  * Defer rabbit close
//  * Format response, result as JSON

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	address = flag.String("address", "127.0.0.1:8080", "bind host:port")
	amqpUri = flag.String("amqp", "amqp://guest:guest@127.0.0.1:5672/", "amqp uri")
)

func init() {
	flag.Parse()
}

// Entity for HTTP Request Body: Exchange/Queue/QueueBind JSON Input
type ExchangeEntity struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	NoWait     bool   `json:"nowait"`
}

type QueueEntity struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

type QueueBindEntity struct {
	Queue    string   `json:"queue"`
	Exchange string   `json:"exchange"`
	NoWait   bool     `json:"nowait"`
	Keys     []string `json:"keys"` // bind/routing keys
}

// RabbitMQ Operate Wrapper
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func (r *RabbitMQ) Connect() (err error) {
	r.conn, err = amqp.Dial(*amqpUri)
	if err != nil {
		// TODO: log error
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

func (r *RabbitMQ) DeclareExchange(name, typ string, durable, autodelete, nowait bool) (err error) {
	err = r.channel.ExchangeDeclare(name, typ, durable, autodelete, false, nowait, nil)
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

func (r *RabbitMQ) DeleteExchange(name string) (err error) {
	err = r.channel.ExchangeDelete(name, false, false)
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

func (r *RabbitMQ) DeclareQueue(name string, durable, autodelete, exclusive, nowait bool) (err error) {
	_, err = r.channel.QueueDeclare(name, durable, autodelete, exclusive, nowait, nil)
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

func (r *RabbitMQ) DeleteQueue(name string) (err error) {
	// TODO: other property wrapper
	_, err = r.channel.QueueDelete(name, false, false, false)
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

func (r *RabbitMQ) BindQueue(queue, exchange string, keys []string, nowait bool) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueBind(queue, key, exchange, nowait, nil); err != nil {
			// TODO: log error
			return err
		}
	}
	return nil
}

func (r *RabbitMQ) UnBindQueue(queue, exchange string, keys []string) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueUnbind(queue, key, exchange, nil); err != nil {
			return err
		}
	}
	return nil
}

func (r *RabbitMQ) Close() (err error) {
	err = r.conn.Close()
	if err != nil {
		// TODO: log error
		return err
	}
	return nil
}

// QueueHandler
func QueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "read body error: %s", err)
			return
		}
		entity := new(QueueEntity)
		json.Unmarshal(body, entity) // TODO: check error
		rabbit := new(RabbitMQ)
		rabbit.Connect() // TODO: check error
		if r.Method == "POST" {
			err = rabbit.DeclareQueue(entity.Name, entity.Durable, entity.AutoDelete, entity.Exclusive, entity.NoWait)
			if err != nil {
				fmt.Fprintf(w, "declare queue error: %s", err)
			} else {
				fmt.Fprintf(w, "declare queue ok")
			}
		} else if r.Method == "DELETE" {
			err = rabbit.DeleteQueue(entity.Name)
			if err != nil {
				fmt.Fprintf(w, "delete queue error: %s", err)
			} else {
				fmt.Fprintf(w, "delete queue ok")
			}
		}
		rabbit.Close()
	} else {
		fmt.Fprintf(w, "method %s not allow", r.Method)
	}
}

func QueueBindHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "read body error: %s", err)
			return
		}
		entity := new(QueueBindEntity)
		json.Unmarshal(body, entity) // TODO: check error
		rabbit := new(RabbitMQ)
		rabbit.Connect()
		if r.Method == "POST" {
			err = rabbit.BindQueue(entity.Queue, entity.Exchange, entity.Keys, entity.NoWait)
			if err != nil {
				fmt.Fprintf(w, "bind queue error: %s", err)
			} else {
				fmt.Fprintf(w, "bind queue ok")
			}
		} else if r.Method == "DELETE" {
			err = rabbit.UnBindQueue(entity.Queue, entity.Exchange, entity.Keys)
			if err != nil {
				fmt.Fprintf(w, "unbind queue error: %s", err)
			} else {
				fmt.Fprintf(w, "unbind queue ok")
			}
		}
		rabbit.Close()
	} else {
		fmt.Fprintf(w, "method %s not allow", r.Method)
	}
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "DELETE" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "read body error: %s", err)
			return
		}
		entity := new(ExchangeEntity)
		json.Unmarshal(body, entity)
		rabbit := new(RabbitMQ)
		rabbit.Connect()
		if r.Method == "POST" {
			err = rabbit.DeclareExchange(entity.Name, entity.Type, entity.Durable, entity.AutoDelete, entity.NoWait)
			if err != nil {
				fmt.Fprintf(w, "declare exchange error: %s", err)
			} else {
				fmt.Fprintf(w, "declare exchange ok")
			}
		} else if r.Method == "DELETE" {
			err = rabbit.DeleteExchange(entity.Name)
			if err != nil {
				fmt.Fprintf(w, "delete exchange error: %s", err)
			} else {
				fmt.Fprintf(w, "delete exchange ok")
			}
		}
		rabbit.Close()
	} else {
		fmt.Fprintf(w, "method %s not allow", r.Method)
	}
}

func main() {
	// Register HTTP Handlers
	http.HandleFunc("/exchange", ExchangeHandler)
	http.HandleFunc("/queue/bind", QueueBindHandler)
	http.HandleFunc("/queue", QueueHandler)

	// Start HTTP Server
	err := http.ListenAndServe(*address, nil)
	if err != nil {
		log.Fatal(err)
	}
}
