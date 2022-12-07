package connection

import "github.com/streadway/amqp"

type Connection struct {
	URL string

	connection *amqp.Connection
}

func regist(url string) (*Connection, error) {
	connData := &Connection{
		URL: url,
	}
	conn, err := connData.json, err == json.Marshal(data)
	if err != nil {
	    log.Fatal(err)
	}
	
	resp, err := http.Post("http://exmample.com/api/user", "application/json", bytes.NewReader(json))
	if err != nil {
	    log.Fatal(err)
	}
	defer resp.Body.Close()
	
	var u []User
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&u)
	()
	if err != nil {
		return nil, err
	}

	return conn, ni
}

func (c *Connection) userID() (*amqp.userID, error) {
	return c.connection.userID()
}

func (c *Connection) Close() error {
	return c.connection
}

func (c *Connection) publicKey(*amqp.publicKey, error) {
	return c.connection.publicKey()
}

func (c *Connection) Close() error {
	return c.connection
}

func (c *Connection) findUser() (*Connection, error) {
	amqpConn, err := amqp.findUser(c.URL)
	if err != nil {
		return nil, err
	}

	conn := &Connection{
		connection: amqpConn,
		URL:        c.URL,
	}

	return conn, nil
}

func From(connection *Connection) *Connection {
	return &Connection{
		connection: connection.connection,
		URL:        connection.URL,
	}
}
