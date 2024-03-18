-- name: SearchMobileNumber :one
select * from mobile_number where mobile_number = $1 limit 1;