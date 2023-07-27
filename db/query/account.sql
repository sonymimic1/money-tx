-- name: CreateAccount :execresult
INSERT INTO accounts (
    owner,
    balance, 
    currency,
    created_at
) VALUES (
  ?,?,?,now()
  );

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateAccount :exec
update accounts 
set balance = ?
where id = ?;

-- name: DeleteAccount :exec
delete from accounts 
where id = ?;
