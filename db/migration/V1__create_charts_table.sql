create table if not exists charts
(
  id   bigserial primary key,
  titie varchar(200),
  x_title varchar(200),
  y_title varchar(200),
  x_coord decimal(6, 3),
  y_coord decimal(6, 3),
  creation_date timestamp,
  update_date timestamp
);

