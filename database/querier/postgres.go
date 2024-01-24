package querier

import "fmt"

type Postgres struct{}

func (p *Postgres) CreateDatabase(name, templateName string) string {
	if templateName != "" {
		return fmt.Sprintf(`CREATE DATABASE "%s" TEMPLATE "%s"`, name, templateName)

	}
	return fmt.Sprintf(`CREATE DATABASE "%s"`, name)
}

func (p *Postgres) DeleteDatabase(name string) string {
	q := `DROP DATABASE "%s"`
	return fmt.Sprintf(q, name)
}

func (p *Postgres) DisconnectFomDatabase(name string) string {
	q := `SELECT pg_terminate_backend(pg_stat_activity.pid)
            FROM pg_stat_activity
            WHERE pg_stat_activity.datname = '%s'
            AND pg_stat_activity.pid <> pg_backend_pid()`
	return fmt.Sprintf(q, name)

}
