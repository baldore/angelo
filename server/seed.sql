begin;

insert into songs (name, labels)
values
  ('billets doux', '[]'),
  ('what is this thing called love', '[]'),
  ('swing 42', '[{ "name": "julian concert" }]'),
  ('rhythm futur', '[{ "name": "julian concert" }]');

commit;
