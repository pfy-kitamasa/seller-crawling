/*
 domain
 */
-- name: GetDomainByID :one
SELECT * FROM domain WHERE key = $1 LIMIT 1;

-- name: UpdateDomainLastedVisitURL :exec
UPDATE "domain" SET
    latest_visit_url = $2
,   updated_at = now()
WHERE key = $1;

/*
 seller
 */
-- name: CreateSeller :one
INSERT INTO "seller" ("url", "key", "name", "tel", "prefecture", "address", "longitude", "latitude", "disabled", "meta", "opened_at") VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8, false, NULL, $9)
RETURNING id;

-- name: UpdateSellerDisabled :exec
UPDATE "seller" SET
    disabled = true
,   updated_at = now()
WHERE id = $1;

-- name: UpdateSellerExported :exec
UPDATE "seller" SET
    exported = true
,   exported_at = now()
WHERE id = $1;
