/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/spf13/pflag"
)

var (
	username    = pflag.String("username", "", "database username")
	password    = pflag.String("password", "", "database password")
	cosmosdbURL = pflag.String("cosmos-db-url", "", "URL to cosmos db")
)

func main() {
	// connect to the cluster
	pflag.Parse()
	fmt.Println(*cosmosdbURL)
	cluster := gocql.NewCluster(*cosmosdbURL)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: *username,
		Password: *password,
	}
	cluster.Port = 10350
	cluster.Keyspace = "mla"
	cluster.Consistency = gocql.LocalQuorum
	cluster.SslOpts = &gocql.SslOptions{
		EnableHostVerification: false,
	}
	cluster.DisableInitialHostLookup = true
	cluster.NumConns = 1
	go func() {
		session, err := cluster.CreateSession()
		if err != nil {
			fmt.Println(err)
			return
		}

		defer session.Close()

		var keyspaceName string
		for i := 0; i < 10000; i++ {
			if err := session.Query("SELECT keyspace_name FROM system_schema.keyspaces LIMIT 1").Scan(&keyspaceName); err != nil {
				log.Fatal(err)
			}
			fmt.Println("select system_schema.keyspaces")
			fmt.Println("keyspace:", i, keyspaceName)
			time.Sleep(time.Second)
		}
	}()

	fmt.Scanln()
	fmt.Println("done")
}
