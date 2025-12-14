package setup

import d "app/internal/di"

func CleanUp(app d.DI) {
	defer app.DB.Close()
}
