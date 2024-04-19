# Database Connection Pooling Test

## Run MariaDB with Docker
```shell
docker-compose up -d
```

### Application configuration

Database connection information is in the .env file.
```dotenv
DB_HOST=127.0.0.1
DB_PORT=3310
DB_USER=sslcom
DB_PASSWORD=123123
DB_NAME=sslcom
DB_MAX_IDLE_CONN=150
DB_MAX_OPEN_CONN=480
DB_MAX_IDLE_LIFE_TIME=10
DB_MAX_CONN_LIFE_TIME=60
SERVER_PORT=7001
```

### Run Application
```shell
go run main.go
```

### The load test is started by opening the SSLcom.jmx file with JMeter

Jmeter Settings;

 * Number of Threads
 * Ramp-up Period
 * Loop Count

By changing the values, different tests can be performed according to the maximum connection number.

After running the JMeter load test, a connection to the database can be established and active connections can be checked with the following SQL.

```sql
SELECT count(*) FROM information_schema.processlist WHERE DB = 'sslcom'
```