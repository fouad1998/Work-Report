package gitlab

func (g *Gitlab) SetToken(token string) error {
	if err := g.config.Read(); err != nil {
		return err
	}

	g.config.GitlabToken = token

	return g.config.Save()
}
