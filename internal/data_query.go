package internal

type Queries struct {
	Insert string
	Check  string
}

var query = Queries{
	Insert: `INSERT INTO application (name, type, bank, opsyears, ssm, prevgateway, prodtype, storetype, inventory, reference, socmedia, litigation, score)
                 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
	Check: `SELECT EXISTS (SELECT 1 FROM application WHERE name = $1)`,
}
