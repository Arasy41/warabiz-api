package repository

const (
	qCreateCategory = `
	INSERT INTO category (category_name, created_by)
	VALUES (?,?)
	RETURNING id
	`
	qCheckDuplicateCategory = `
	SELECT EXISTS(SELECT 1 FROM category WHERE deleted_at IS NULL AND category_name = ?)
	`
	qGetCategoryByID = `
	SELECT id, category_name, created_by, created_at, updated_by, updated_at
	FROM category
	WHERE deleted_at IS NULL AND id = ?
	`
	qCountCategory = `
	SELECT COUNT(*) AS total
	FROM category
	WHERE deleted_at IS NULL
	`
	qGetAllCategory = `
	SELECT id, category_name
	FROM category
	WHERE deleted_at IS NULL
	`
	qUpdateCategory = `
	UPDATE category
	SET category_name = ?,
		updated_by = ?,
		updated_at = timezone('utc', now())
	WHERE id = ?
	`

	qDeleteCategory = `
	UPDATE category
	SET deleted_at = timezone('utc', now())
	WHERE id = ?
	`

	qGetCategoryDetailByID = `
	SELECT c.id AS category_id, c.category_name, w.id AS waralaba_id, w.category_id, w.waralaba_name, w.prize, w.contact,  w.brochure_link, w.since, w.outlet_total, w.license_duration
	FROM category c
	INNER JOIN waralaba w ON w.category_id = c.id
	WHERE c.id = ?
	`
)
