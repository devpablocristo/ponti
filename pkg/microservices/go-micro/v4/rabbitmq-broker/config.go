package pkgrabbit

import (
	"fmt"
)

// config estructura que implementa la interfaz Config para RabbitMQ.
type config struct {
	name string
	host string

	user         string
	password     string
	vhost        string
	queue        string
	exchange     string
	exchangeType string
	routingKey   string

	port      int
	autoAck   bool
	exclusive bool
	noLocal   bool
	noWait    bool
}

// newConfig crea una nueva configuración para RabbitMQ con opciones adicionales.
func newConfig(name, host, user, password, vhost, queue, exchange, exchangeType, routingKey string, port int, autoAck, exclusive, noLocal, noWait bool) Config {
	return &config{
		name:         name,
		host:         host,
		port:         port,
		user:         user,
		password:     password,
		vhost:        vhost,
		queue:        queue,
		exchange:     exchange,
		exchangeType: exchangeType,
		routingKey:   routingKey,
		autoAck:      autoAck,
		exclusive:    exclusive,
		noLocal:      noLocal,
		noWait:       noWait,
	}
}

// Getters y Setters

func (c *config) GetName() string     { return c.name }
func (c *config) SetName(name string) { c.name = name }

func (c *config) GetHost() string     { return c.host }
func (c *config) SetHost(host string) { c.host = host }

func (c *config) GetPort() int     { return c.port }
func (c *config) SetPort(port int) { c.port = port }

func (c *config) GetUser() string     { return c.user }
func (c *config) SetUser(user string) { c.user = user }

func (c *config) GetPassword() string         { return c.password }
func (c *config) SetPassword(password string) { c.password = password }

func (c *config) GetVHost() string      { return c.vhost }
func (c *config) SetVHost(vhost string) { c.vhost = vhost }

func (c *config) GetQueue() string      { return c.queue }
func (c *config) SetQueue(queue string) { c.queue = queue }

func (c *config) GetExchange() string         { return c.exchange }
func (c *config) SetExchange(exchange string) { c.exchange = exchange }

func (c *config) GetExchangeType() string             { return c.exchangeType }
func (c *config) SetExchangeType(exchangeType string) { c.exchangeType = exchangeType }

func (c *config) GetRoutingKey() string           { return c.routingKey }
func (c *config) SetRoutingKey(routingKey string) { c.routingKey = routingKey }

func (c *config) GetAutoAck() bool        { return c.autoAck }
func (c *config) SetAutoAck(autoAck bool) { c.autoAck = autoAck }

func (c *config) GetExclusive() bool          { return c.exclusive }
func (c *config) SetExclusive(exclusive bool) { c.exclusive = exclusive }

func (c *config) GetNoLocal() bool        { return c.noLocal }
func (c *config) SetNoLocal(noLocal bool) { c.noLocal = noLocal }

func (c *config) GetNoWait() bool       { return c.noWait }
func (c *config) SetNoWait(noWait bool) { c.noWait = noWait }

// GetAddress genera la URL de conexión para RabbitMQ.
func (c *config) GetAddress() string {
	vhost := c.vhost
	if vhost == "" {
		vhost = "/"
	}
	return fmt.Sprintf("amqp://%s:%s@%s:%d%s", c.user, c.password, c.host, c.port, vhost)
}

// Validate verifica que todos los parámetros de configuración sean válidos.
func (c *config) Validate() error {
	if c.name == "" {
		return fmt.Errorf("rabbitmq service name is not configured")
	}
	if c.host == "" {
		return fmt.Errorf("rabbitmq host is not configured")
	}
	if c.port == 0 {
		return fmt.Errorf("rabbitmq port is not configured")
	}
	if c.user == "" {
		return fmt.Errorf("rabbitmq user is not configured")
	}
	if c.password == "" {
		return fmt.Errorf("rabbitmq password is not configured")
	}
	if c.vhost == "" {
		return fmt.Errorf("rabbitmq vhost is not configured")
	}
	if c.queue == "" {
		return fmt.Errorf("rabbitmq queue is not configured")
	}
	if c.exchange == "" {
		return fmt.Errorf("rabbitmq exchange is not configured")
	}
	if c.exchangeType == "" {
		return fmt.Errorf("rabbitmq exchange type is not configured")
	}
	if c.routingKey == "" {
		return fmt.Errorf("rabbitmq routing key is not configured")
	}
	if !isValidExchangeType(c.exchangeType) {
		return fmt.Errorf("invalid exchange type: %s", c.exchangeType)
	}
	return nil
}

// isValidExchangeType valida si el tipo de exchange es válido.
func isValidExchangeType(exchangeType string) bool {
	validTypes := map[string]bool{
		"direct":  true,
		"fanout":  true,
		"topic":   true,
		"headers": true,
	}
	return validTypes[exchangeType]
}
