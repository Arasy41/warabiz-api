package repository

const (
	qCreateWaraCareer = `
	INSERT INTO wara_career (career_title, description, address, image_url, created_by)
	VALUES (?,?)
	RETURNING id
	`
	qCheckDuplicateWaraCareer = `
	SELECT EXISTS(SELECT 1 FROM wara_career WHERE deleted_at IS NULL AND career_title = ?)
	`
	qGetWaraCareerByID = `
	SELECT id, career_title, description, address, image_url, created_by, created_at, updated_by, updated_at
	FROM wara_career
	WHERE deleted_at IS NULL AND id = ?
	`
	qCountWaraCareer = `
	SELECT COUNT(*) AS total
	FROM wara_career
	WHERE deleted_at IS NULL
	`
	qGetAllWaraCareer = `
	SELECT id, career_title, description, address, image_url
	FROM wara_career
	WHERE deleted_at IS NULL
	`
	qUpdateWaraCareer = `
	UPDATE wara_career
	SET career_title = ?,
	 	description = ?, 
		address = ?, 
		image_url = ?,
		updated_by = ?,
		updated_at = timezone('utc', now())
	WHERE id = ?
	`

	qDeleteWaraCareer = `
	UPDATE wara_career
	SET deleted_at = timezone('utc', now())
	WHERE id = ?
	`
)
