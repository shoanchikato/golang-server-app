package setup

import d "app/pkg/di"

func CleanUp(app d.DI) {
	defer app.DB.Close()
}
