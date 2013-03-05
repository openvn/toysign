package toysign

func (h *handler) initSubRoutes() {
	h._subRoutes = []route{
		route{"login", Login},
		route{"login2", Login2},
	}
}
