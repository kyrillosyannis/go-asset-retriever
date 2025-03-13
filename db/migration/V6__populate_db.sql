insert into assets (description, type) values ('Chart One', 'Chart'), ('Audience One', 'Audience');

insert into charts (title, x_title, y_title, x_coord, y_coord, asset_id) values ('Chart One', 'X title', 'Y title', 4, 5, 1);

insert into audiences (asset_id, gender, country_of_birth, age_group, daily_hours_on_social_media, last_month_number_of_purchases)
values (2, 'M', 'GR', '30-40', 2, 2);

insert into favourites (user_id, asset_id) values (1, 1), (1, 2);
