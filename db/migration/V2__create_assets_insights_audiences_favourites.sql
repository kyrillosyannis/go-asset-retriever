create table if not exists assets (
    id bigserial primary key,
    description varchar(255)
);

alter table charts add column asset_id bigint unique references assets(id) on delete cascade;

create table if not exists insights (
    id bigserial primary key,
    asset_id bigint unique references assets(id) on delete cascade, -- One-to-One Relationship
    text varchar(255)
);

create table if not exists audiences (
    id bigserial primary key,
    asset_id bigint unique references assets(id) on delete cascade, -- One-to-One Relationship
    gender varchar(10),
    country_of_birth varchar(100),
    age_group varchar(50),
    daily_hours_on_social_media smallint,
    last_month_number_of_purchases int
);

create table if not exists favourites (
    id bigserial primary key,
    user_id bigint,
    asset_id bigint references assets(id) on delete cascade,
    unique(user_id, asset_id) -- Prevent duplicate favorites
);
