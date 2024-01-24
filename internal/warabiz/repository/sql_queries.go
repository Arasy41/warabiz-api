package repository

const (
	qCreateWarabiz = `
	INSERT INTO waralaba (id, category_id, waralaba_name, prize, contact, brochure_link, since, outlet_total, license_duration, created_by)
	VALUES (?,?,?,?,?,?,?,?,?,?)
	RETURNING id
	`
	qCheckDuplicateWarabiz = `
	SELECT EXISTS(SELECT 1 FROM waralaba WHERE deleted_at IS NULL AND waralaba_name = ?)
	`
	qGetWarabizByID = `
	SELECT w.id, w.category_id, c.category_name, w.waralaba_name, w.prize, w.contact,  w.brochure_link, w.since, w.outlet_total, w.license_duration, w.created_by, w.created_at, w.updated_by, w.updated_at
	FROM waralaba w
	INNER JOIN category c ON c.id = w.category_id
	WHERE w.deleted_at IS NULL AND w.id = ?
	`
	qCountWarabiz = `
	SELECT COUNT(*) AS total
	FROM waralaba
	WHERE deleted_at IS NULL
	`
	qGetAllWarabiz = `
	SELECT id, category_id, waralaba_name, prize
	FROM waralaba
	WHERE deleted_at IS NULL
	`
	qUpdateWarabiz = `
	UPDATE waralaba
	SET slug = ?,
		title = ?,
		content = ?,
		excerpt = ?,
		description = ?,
		thumbnail_url = ?,
		author_id = ?,
		publisher_id = ?,
		updated_by = ?,
		updated_at = timezone('utc', now())
	WHERE id = ?
	`
	qDeleteWarabiz = `
	UPDATE waralaba
	SET deleted_at = timezone('utc', now())
	WHERE id = ?
	`
)
