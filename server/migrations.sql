begin;

create table songs (
  id SERIAL PRIMARY KEY,
  name TEXT,
);

create table if not exists resources (
   id serial primary key,
   link text not null,
   label text not null,
   description text,
   song_id int not null,
   constraint fk_song
     foreign key (song_id)
       references songs(id)
       on delete cascade
 );

commit;
