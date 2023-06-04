-- +migrate Up
create table if not exists "commpics" (
		"cid" integer references comms(cid),
		"name" text
	  );

-- +migrate Down
drop table "commpics";
