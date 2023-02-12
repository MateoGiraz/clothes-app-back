package querys

// GetClothes
const GetClothes = `SELECT 
id,
is_available,
name,
description,
color,
category
FROM clothing
ORDER BY id;`

// GetClothing
const GetClothing = `SELECT 
id,
is_available,
name,
description,
color,
category
FROM clothing
WHERE id = $1;`

// DeleteClothing
const DeleteClothing = `DELETE FROM clothing WHERE id = $1;`

// CreateClothing
const CreateClothing = `INSERT INTO clothing (
  is_available, 
  name,
  description,
  color,
  category
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id;`

//UpdateClothing

const UpdateClothingQuery = `UPDATE clothing (
SET is_availabe = $2, 
WHERE id = $1
RETURNING id;`
