package archive

const QSubprimerUndescribedContent = `
select
  url, created, updated, last_head, last_get, status, content_type, content_sniff, content_length, 
  title, id, headers_took, download_took, headers, meta, hash 
from 
  urls 
where 
  url ilike $1
  and content_sniff != 'text/html; charset=utf-8'
  and last_get is not null
  -- confirm is not empty hash
  and hash != '1220e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
  and not exists (select null from metadata where urls.hash = metadata.subject) 
limit $2 offset $3;`

const QSubprimerDescribedContent = `
select
  url, created, updated, last_head, last_get, status, content_type, content_sniff, content_length, 
  title, id, headers_took, download_took, headers, meta, hash 
from 
  urls 
where 
  url ilike $1
  and content_sniff != 'text/html; charset=utf-8'
  and last_get is not null
  -- confirm is not empty hash
  and hash != '1220e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855'
  and exists (select null from metadata where urls.hash = metadata.subject) 
limit $2 offset $3;`
