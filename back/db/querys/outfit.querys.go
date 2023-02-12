package querys

// GetOutfits
const GetOutfits = `SELECT 
id,
top_id,
pants_id,
shoes_id
FROM outfit
ORDER BY id;`

// GetOutfit
const GetOutfit = `SELECT 
id,
top_id,
pants_id,
shoes_id
FROM outfit
WHERE id = $1;`

// DeleteOutfit
const DeleteOutfit = `DELETE FROM outfit WHERE id = $1;`

// CreateOutfit
const CreateOutfit = `INSERT INTO outfit (
	top_id,
	pants_id,
	shoes_id
) VALUES (
  $1, $2, $3
)
RETURNING id;`

//UpdateOutfit

const UpdateOutfitQuery = `UPDATE outfit (
SET top_id = $2, 
pants_id = $3,
shoes_id = $4
WHERE id = $1
RETURNING id;`
