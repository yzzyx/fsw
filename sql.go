package main

import "github.com/olebedev/config"

type SqlCollector struct {
	driver     string
	datasource string
	query      string
	interval   int
}

func (sc *SqlCollector) SetConfig(cfg *config.Config) (err error) {

	if sc.driver, err = cfg.String("driver"); err != nil {
		return err
	}

	if sc.datasource, err = cfg.String("datasource"); err != nil {
		return err
	}

	if sc.query, err = cfg.String("query"); err != nil {
		return err
	}

	sc.interval = cfg.UInt("interval", 3600)
	return nil
}
