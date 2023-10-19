-- name: CreateAccount :execresult
INSERT INTO accounts (
    owner,
    balance, 
    currency
) VALUES (
  ?,?,?
  );

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = ? LIMIT 1
FOR UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT ?
OFFSET ?;

-- name: UpdateAccount :exec
update accounts 
set balance = ?
where id = ?;

-- name: AddAccountBalance :exec
update accounts 
set balance = balance + sqlc.arg(amount)
where id = sqlc.arg(id);

-- name: DeleteAccount :exec
delete from accounts 
where id = ?;
