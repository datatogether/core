
-- name: insert-primers
insert into primers values
  ('5b1031f4-38a8-40b3-be91-c324bf686a87','2017-01-01 00:00:01','2017-01-01 00:00:01', 'Environmental Protection Agency', 'The mission of the Environmental Protection Agency is to protect human health and the environment through the development and enforcement of regulations. The EPA is responsible for administering a number of laws that span various sectors, such as agriculture, transportation, utilities, construction, and oil and gas. In the budget for FY 2017, the agency lays out goals to better support communities and address climate change following the President’s Climate Action Plan. Additionally, the agency aims to improve community water infrastructure, chemical plant safety, and collaborative partnerships among federal, state, and tribal levels.',false),
	('d9deff9d-15e8-43f1-9d00-51160c0bffbe','2017-01-01 00:00:01','2017-01-01 00:00:01', 'US Census', 'US Census Bureau',false);
-- name: delete-primers
delete from primers;

--name: insert-subprimers
insert into subprimers values
  ('326fcfa0-d3e6-4b2d-8f95-e77220e16109', 'www.epa.gov', '2017-01-01 00:00:01', '2017-01-01 00:00:01', '5b1031f4-38a8-40b3-be91-c324bf686a87',true,43200000,null,null,null),
  ('440d9779-406c-4015-8f2d-404b04ead3a2', 'www.census.gov', '2017-01-01 00:00:01', '2017-01-01 00:00:01', 'd9deff9d-15e8-43f1-9d00-51160c0bffbe',true,43200000,null,null,null);
--name: delete-subprimers
delete from subprimers;

-- name: insert-urls
insert into urls values
	-- url,created,updated,last_head,last_get,status,content_type,content_sniff,content_length,title,id,headers_took,download_took,headers,meta,hash
	('http://www.epa.gov', '2017-01-01 00:00:01', '2017-01-01 00:00:01', '2017-01-01 00:00:01', null, 200, 'text/html; charset=utf-8', 'text/html;', -1, 'United States Environmental Protection Agency, US EPA', 'cee7bbd4-2bf9-4b83-b2c8-be6aeb70e771',0,0, '["X-Content-Type-Options","nosniff","Expires","Fri, 24 Feb 2017 21:53:45 GMT","Date","Fri, 24 Feb 2017 21:53:45 GMT","Etag","W/\"7f53-549471782bb42\"","X-Ua-Compatible","IE=Edge,chrome=1","X-Cached-By","Boost","Content-Type","text/html; charset=utf-8","Vary","Accept-Encoding","Accept-Ranges","bytes","Cache-Control","no-cache, no-store, must-revalidate, post-check=0, pre-check=0","Server","Apache","Connection","keep-alive","Strict-Transport-Security","max-age=31536000; preload;"]', null, '1220459219b10032cc86dcdbc0f83aea15a9d3e1119e7b5170beaee233008ea2c2de'),
  ('https://www.census.gov/nometa.pdf','2017-03-15 17:36:40', '2017-03-21 22:25:21','2017-03-21 22:25:20.88', '2017-03-21 22:25:20.88', 200,'text/html','application/pdf; charset=utf-8' ,164010,'North American Industry Classification System (NAICS) Main Page � U.S. Census Bureau','4c5fc7b8-1397-4d34-980b-1d01247f9ee4',0,0 ,'["Date","Tue, 21 Mar 2017 22:25:20 GMT","Accept-Ranges","bytes","Content-Type","text/html","Strict-Transport-Security","max-age=31536000","Vary","Accept-Encoding"]',null,'1220af06510193276b5fd9ad2fc55dcc004ada557d9259ca3505478bfef0b12ed988'),
  ('https://www.census.gov/topics/economy/classification-codes.html','2017-03-15 17:36:40', '2017-03-21 22:25:21','2017-03-21 22:25:20.88','2017-03-21 22:25:20.88384',200,'text/html','text/plain; charset=utf-8' ,164010,'North American Industry Classification System (NAICS) Main Page � U.S. Census Bureau','4c5fc7b8-1397-4d34-980b-1d01247f9ee4',0,0 ,'["Date","Tue, 21 Mar 2017 22:25:20 GMT","Accept-Ranges","bytes","Content-Type","text/html","Strict-Transport-Security","max-age=31536000","Vary","Accept-Encoding"]',null,'12207b06510193276b5fd9ad2fc55dcc004ada557d9259ca3505478bfef0b16ed977');
-- name: delete-urls
delete from urls;

-- name: insert-links
-- insert into links values
-- ('2017-01-01 00:00:02','2017-01-01 00:00:02','http://www.epa.gov','http://www.epa.gov');
-- name: delete-links
delete from links;

-- name: insert-snapshots
-- insert into snapshots values
-- 	();
-- name: delete-snapshots
delete from snapshots;

-- name: insert-metadata
insert into metadata values
  -- hash,time_stamp,key_id,subject,prev,meta,deleted
  ('1220b499c5e883f2a3d47fe96c51779d05d5758c53c24caabc447db137b220c1688a', '2017-03-15 17:48:33', 'a5d6f8d8cbb15c0f60159b7fd58b1c555dbde076388ffce96c53f04b91ed377a', '12207b06510193276b5fd9ad2fc55dcc004ada557d9259ca3505478bfef0b16ed977', '', '{"description":"","title":"NAICS Classification Codes"}', false);
-- name: delete-metadata
delete from metadata;

-- name: insert-collections
insert into collections values
  ('76dd07ac-54cb-4f9d-b0a6-88d3d55c0d9d', '2017-01-01 00:00:01', '2017-01-01 00:00:01', 'bernidette', 'a collection of urls', null, null);
-- name: delete-collections
delete from collections;

-- name: insert-archive_requests
-- insert into archive_requests values
--  ('8b14f3d6-882f-4dd5-92f8-abaac220864f','2017-01-01 00:00:01','http://www.apple.com','');
-- name: delete-archive_requests
delete from archive_requests;